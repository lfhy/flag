package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lfhy/flag/example/subcmd/cmd"
	"github.com/lfhy/flag/example/subcmd/conf"
)

func main() {
	// 如果Help为true，则列出所有子命令
	if conf.Help {
		cmd.ListCommand()
		os.Exit(0)
	}
	// 运行子命令
	if len(flag.Args()) > 0 {
		defer os.Exit(0)
		// 执行命令
		if conf.Debug {
			fmt.Println("Run:", flag.Args())
		}
		if err := cmd.RunCmd(flag.Args()[0], flag.Args()[1:]...); err != nil {
			if conf.Debug {
				fmt.Println(err)
			}
		}
	} else {
		// 不存在则默认启用Server端
		cmd.RunCmd("server", os.Args[1:]...)
	}
}
