/**
 * @Author: DollarKillerX
 * @Description: main程序入口
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:17 2019/11/26
 */
package main

import (
	"ddbf/cmd"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// 进行性能分析
	//go scp()

	app := cli.NewApp()
	app.Name = "DdbFC"
	app.Author = "DollarKiller"
	app.Email = "adapawang@gmail.com"
	app.Version = "0.1"
	app.Usage = "Distributed DNS brute force cracking CLi"

	app.Flags = append(app.Flags, cmd.Scan.Flags...)

	app.Action = cmd.ScanIc
	err := app.Run(os.Args)

	if err != nil {
		log.Fatalln(err)
	}

}

//func scp() {
//	cpuf, e := os.Create("cpu_profile")
//	if e != nil {
//		log.Fatalln(e)
//	}
//	pprof.StartCPUProfile(cpuf)
//
//	time.Sleep(10 * time.Second)
//	defer pprof.StopCPUProfile()
//
//	log.Println("关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||关闭分析|||")
//}
