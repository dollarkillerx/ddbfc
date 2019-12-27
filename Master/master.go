/**
 * @Author: DollarKillerX
 * @Description: master.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:30 2019/12/27
 */
package main

import (
	"ddbf/Master/discovery"
	"ddbf/Master/router"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalln("参数规则 ./main api服务地址 rpc地址")
	}
	apiHost := os.Args[1]
	rpcHost := os.Args[2]

	app := gin.New()

	app.Use(gin.Recovery()) // 注册防止崩溃中间件
	router.Registered(app)  // 注册路由

	go discovery.RunServer(rpcHost) // 注册 rpc服务器

	if err := app.Run(apiHost); err != nil {
		log.Fatalln(err)
	}
}
