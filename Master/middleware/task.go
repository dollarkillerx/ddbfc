/**
 * @Author: DollarKillerX
 * @Description: task 相关中间件
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:47 2019/12/27
 */
package middleware

import (
	"ddbf/Master/definition"
	"ddbf/Master/shared"
	"ddbf/Master/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

// 流量限制中间件
func TaskLimit(ctx *gin.Context) {
	// 如果不是空的就说明当前任务还在执行
	if len(shared.LimitChannel) != 0 {
		utils.RespStandard(ctx, definition.RespBusy)
		ctx.Abort()
		return
	}
	// 插入一个任务
	shared.LimitChannel <- true
}

// 参数校验中间件
func TaskValidate(ctx *gin.Context) {
	// 检测domain
	domain := ctx.PostForm("domain")
	if domain == "" {
		utils.RespStandard(ctx, definition.Resp400)
		<-shared.LimitChannel
		ctx.Abort()
		return
	}

	// 检测文件是否存在
	header, e := ctx.FormFile("file")
	if e != nil {
		utils.RespStandard(ctx, definition.Resp400)
		<-shared.LimitChannel
		ctx.Abort()
		return
	}

	// 检测文件格式是否正确  文件大小是否正确
	if strings.Index(header.Filename, "zip") == -1 || header.Size == 0 {
		utils.RespStandard(ctx, definition.Resp400)
		<-shared.LimitChannel
		ctx.Abort()
		return
	}
	ctx.Set("domain", domain)
}
