/**
 * @Author: DollarKillerX
 * @Description: cmd基础命令
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:21 2019/11/26
 */
package cmd

import "github.com/urfave/cli"

var Scan = cli.Command{
	Name:        "scan",
	Usage:       "cmd ...",
	Description: "cmd ...",
	Action:      scan,
	Flags: []cli.Flag{
		stringFlag("domain, d", "", "domain"),
		intFlag("timeout, t", 3, "Single DNS query timeout Millisecond"),
		intFlag("tryNum, r", 3, "Number of attempts"),
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
