/**
 * @Author: DollarKillerX
 * @Description: Registered.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:31 2019/12/27
 */
package router

import (
	"ddbf/Master/middleware"
	"ddbf/Master/service"
	"github.com/gin-gonic/gin"
)

func Registered(app *gin.Engine) {
	app.POST("/task", middleware.TaskLimit, middleware.TaskValidate, service.Task) // 上传后返回任务的唯一ID
	app.GET("/report/:id", service.ReportGet)                                      // 更具id获取任务报告
}
