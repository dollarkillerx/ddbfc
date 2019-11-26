/**
 * @Author: DollarKillerX
 * @Description: scan.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:23 2019/11/26
 */
package cmd

import (
	"ddbf/model"
	"errors"
	"github.com/urfave/cli"
	"log"
	"strings"
	"time"
)

var (
	d400 = errors.New("Incorrect domain name entered")
)

// 绑定相关参数 并初始化系统
func ScanIc(ctx *cli.Context) error {
	if ctx.IsSet("domain") {
		domain := ctx.String("domain")
		// domain validate
		if domain == "" {
			return d400
		}
		split := strings.Split(domain, ".")
		if len(split) < 2 {
			return d400
		}
		model.BaseModel.Domain = domain
	}

	if ctx.IsSet("timeout") {
		timeoutInt := ctx.Int("timeout")
		if timeoutInt <= 0 {
			return d400
		}
		model.BaseModel.TimeOut = time.Duration(timeoutInt) * time.Millisecond
	}

	if ctx.IsSet("tryName") {
		tryName := ctx.Int("tryName")
		if tryName <= 0 {
			return d400
		}
		model.BaseModel.TryNum = tryName
	}

	if ctx.IsSet("max") {
		maxConcurrent := ctx.Int("max")
		if maxConcurrent <= 0 {
			return d400
		}
		log.Println(maxConcurrent)
		model.BaseModel.Max = maxConcurrent
		log.Println(model.BaseModel.Max)
	}

	// 启动初始化程序
	EngineInit()

	return nil
}
