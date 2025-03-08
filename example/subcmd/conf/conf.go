package conf

import (
	"fmt"
	"os"

	"github.com/lfhy/flag"
)

var (
	Debug bool
	Help  bool
)

// 初始化函数
func init() {
	// 设置错误处理方式为继续执行
	flag.SetErrorHandling(flag.ContinueOnError)
	// 将Debug变量与命令行参数debug绑定
	flag.BoolVar(&Debug, "debug", false, "debug")
	// 解析命令行参数
	err := flag.Parse()
	// 如果解析出错
	if err != nil {
		// 如果错误是帮助信息
		if err == flag.ErrHelp {
			// 设置Help变量为true
			Help = true
		} else {
			// 否则打印错误信息并退出程序
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
}
