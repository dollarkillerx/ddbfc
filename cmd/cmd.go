/**
 * @Author: DollarKillerX
 * @Description: cmd基础命令
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:21 2019/11/26
 */
package cmd

import (
	"ddbf/model"
	"github.com/urfave/cli"
)

var Scan = cli.Command{
	Name:        "scan",
	Usage:       "start to crack weak password",
	Description: "start to crack weak password",
	Action:      ScanIc,
	Flags: []cli.Flag{
		stringFlag("domain, d", "", "domain"),
		//intFlag("timeout, t", 400, "Single DNS query timeout Millisecond"),
		//intFlag("tryNum, r", 3, "Number of attempts"),
		//intFlag("max, m", 200, "Maximum number of concurrency"),
		cli.IntFlag{
			Name:        "max, m",
			Value:       200,
			Usage:       "Maximum number of concurrency",
			Destination: &model.BaseModel.Max,
		},
		cli.IntFlag{
			Name:        "tryNum, r",
			Value:       3,
			Usage:       "Number of attempts",
			Destination: &model.BaseModel.TryNum,
		},
		cli.IntFlag{
			Name:        "timeout, t",
			Value:       300,
			Usage:       "Single DNS query timeout Millisecond",
			Destination: &model.BaseModel.TimeOut,
		},
	},
}

func stringFlag(name, value, usage string) cli.StringFlag {
	return cli.StringFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func boolFlag(name, usage string) cli.BoolFlag {
	return cli.BoolFlag{
		Name:  name,
		Usage: usage,
	}
}

func intFlag(name string, value int, usage string) cli.IntFlag {
	return cli.IntFlag{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}
