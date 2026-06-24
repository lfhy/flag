package flag

import (
	"fmt"
	"io"
	"os"
)

// 定义系统标志
var sysflag = ArgsFlag{NewFlagSet(os.Args[0], ExitOnError)}

func GetDefaultFlagSet() *FlagSet {
	return sysflag.FlagSet
}

func SetErrorHandling(h ErrorHandling) {
	sysflag.SetErrorHandling(h)
}

// 解析命令行参数
func Parse() error {
	return sysflag.Parse()
}

func GetConfig() *Config {
	return sysflag.FlagSet.GetConfig()
}

func Var(opt *FlagVar) {
	sysflag.Var(opt)
}

// Args函数用于获取命令行参数
func Args() []string { return sysflag.args }

// ArgsLen函数用于获取命令行参数的长度
func ArgsLen() int { return len(sysflag.args) }

// Usage函数用于打印命令行参数的使用方法
func Usage() {
	Greenf(fmt.Sprintf("%s的使用方法", os.Args[0]), "\n\n")
	sysflag.PrintDefaults()
}

// PrintAll函数用于打印所有的命令行参数
func PrintAll() { sysflag.PrintAll() }

// SetOutput函数用于设置命令行参数的输出
func SetOutput(output io.Writer) { sysflag.SetOutput(output) }

// 删除参数
func DelArg(arg string) {
	sysflag.DelArg(arg)
}
