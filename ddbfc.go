/**
 * @Author: DollarKillerX
 * @Description: main程序入口
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:17 2019/11/26
 */
package main

import (
	"ddbf/cmd"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "DdbFC"
	app.Author = "DollarKiller"
	app.Email = "adapawang@gmail.com"
	app.Version = "0.1"
	app.Usage = "Distributed DNS brute force cracking CLi"

	app.Commands = []cli.Command{cmd.Scan}
	app.Flags = append(app.Flags, cmd.Scan.Flags...)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

}
