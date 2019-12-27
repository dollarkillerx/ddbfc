/**
 * @Author: DollarKillerX
 * @Description: resp.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:44 2019/12/27
 */
package utils

import (
	"ddbf/Master/definition"
	"github.com/gin-gonic/gin"
)

// 通用返回
func RespStandard(ctx *gin.Context, resp *definition.ResponseStandard) {
	ctx.JSON(resp.HttpCode, resp)
}
