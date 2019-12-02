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
)

var (
	d400 = errors.New("Incorrect domain name entered")
)

// 绑定相关参数 并初始化系统
func ScanIc(ctx *cli.Context) error {
	if model.BaseModel.Domain == "" {
		log.Fatalln("-h View help")
	}
	// 启动初始化程序
	EngineInit()

	return nil
}
