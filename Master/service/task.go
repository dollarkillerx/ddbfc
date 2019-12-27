/**
 * @Author: DollarKillerX
 * @Description: task 相关服务
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:45 2019/12/27
 */
package service

import (
	"bufio"
	"ddbf/Master/definition"
	"ddbf/Master/shared"
	"ddbf/Master/utils"
	"ddbf/pb/pb_work"
	"github.com/dollarkillerx/easyutils"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/gin-gonic/gin"
	"github.com/saracen/fastzip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// 注册任务
func Task(ctx *gin.Context) {
	header, _ := ctx.FormFile("file")

	open, e := header.Open()
	if e != nil {
		utils.RespStandard(ctx, definition.RespFileErr)
		<-shared.LimitChannel
		ctx.Abort()
		return
	}
	defer open.Close()

	filename := easyutils.SuperRand()
	filePath := filepath.Join(definition.ZIPFILE, filename+".zip")
	file, e := os.Create(filePath)
	if e != nil {
		utils.RespStandard(ctx, definition.RespFileErr)
		<-shared.LimitChannel
		ctx.Abort()
		return
	}
	defer file.Close()
	// 进行持久化
	_, e = io.Copy(file, open)
	if e != nil {
		utils.RespStandard(ctx, definition.RespFileErr)
		<-shared.LimitChannel
		ctx.Abort()
		return
	}

	// 基本文件处理完毕

	// # 以下是文件的解压
	// 文件解压目录
	unzipPath := filepath.Join(definition.UNZIPFILE, filename)
	e = easyutils.DirPing(unzipPath)
	if e != nil {
		utils.RespStandard(ctx, definition.RespFileErr)
		<-shared.LimitChannel
		ctx.Abort()
		return
	}
	extractor, e := fastzip.NewExtractor(filePath, unzipPath)
	if e != nil {
		utils.RespStandard(ctx, definition.RespFileErr)
		<-shared.LimitChannel
		ctx.Abort()
		return
	}
	defer extractor.Close()
	if e := extractor.Extract(); e != nil {
		utils.RespStandard(ctx, definition.RespFileErr)
		<-shared.LimitChannel
		ctx.Abort()
		return
	}

	domain, _ := ctx.Get("domain") // 测试域名
	go subcontract(filename, domain.(string), unzipPath)

	// 基本的校验已经完结
	utils.RespStandard(ctx, &definition.ResponseStandard{
		HttpCode: 200,
		Code:     200,
		Msg:      filename,
	})

	// 解压完毕 删除上传完的zip文件
	if e := os.Remove(filePath); e != nil {
		log.Println(e)
	}

}

// 分发任务
func subcontract(key, domain, unzipPath string) {
	dirs, e := ioutil.ReadDir(unzipPath)
	if e != nil {
		clog.PrintWa(e)
		<-shared.LimitChannel
		return
	}

	task := pb_work.Request{TaskId: key, TaskItem: []string{}}
	taskNum := 0 // 初始化任务个数
	log.Printf("进入分包阶段 taskId: %s domain: %s  INFOS：%s\n", key, domain, dirs)
	taskStatistics := definition.Task{Id: key, Num: 0}
	shared.TaskNum[key] = &taskStatistics // 初始化
	for _, v := range dirs {
		if !v.IsDir() {
			dirName := filepath.Join(unzipPath, v.Name())
			file, e := os.Open(dirName)
			if e != nil {
				clog.PrintWa(e)
				continue
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				task.TaskItem = append(task.TaskItem, scanner.Text()+"."+domain) // 拼接域名

				if len(task.TaskItem) == 120000 { // 单个任务包的大小1.2W
					taskNum++
					shared.TaskNum[key].Num = taskNum

					shared.TaskPool <- task // 发送给任务队列
					// 重置task item
					task.TaskItem = make([]string, 0)
				}
			}
		}
	}
	// 如果任务大小没有到120000就会进入这里
	if len(task.TaskItem) > 0 {
		log.Printf("当前任务 %s  以及全部发送完毕!!! 200 OK", key)
		shared.TaskNum[key].Over = true

		taskNum++
		shared.TaskNum[key].Num = taskNum
		shared.TaskPool <- task // 发送给任务队列
	}

	// 删除字典目录
	if e := os.RemoveAll(unzipPath); e != nil {
		clog.PrintWa(e)
	}
}

// 获取报告
func ReportGet(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		utils.RespStandard(ctx, definition.Resp400)
		return
	}
	// 打开文件
	file, e := os.Open(reportGetFile(id))
	if e != nil {
		utils.RespStandard(ctx, definition.Resp404)
		return
	}
	defer file.Close()
	ctx.Header("content-type", "application/json")
	_, e = io.Copy(ctx.Writer, file)
	if e != nil {
		utils.RespStandard(ctx, definition.Resp404)
	}
}

func reportGetFile(id string) string {
	return filepath.Join(definition.OUTFILE, id+".txt")
}
