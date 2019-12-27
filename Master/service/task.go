/**
 * @Author: DollarKillerX
 * @Description: task 相关服务
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:45 2019/12/27
 */
package service

import (
	"ddbf/Master/definition"
	"ddbf/Master/shared"
	"ddbf/Master/utils"
	"github.com/dollarkillerx/easyutils"
	"github.com/gin-gonic/gin"
	"github.com/saracen/fastzip"
	"io"
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

}
