package flag

import (
	"fmt"
	"os"
)

type ArgsFlag struct {
	*FlagSet
}

// 创建一个新的ArgsFlag实例
func NewArgsFlag() *ArgsFlag {
	// 使用os.Args[0]作为FlagSet的名称，ExitOnError表示在遇到错误时退出程序
	return &ArgsFlag{NewFlagSet(os.Args[0], ExitOnError)}
}

// 解析ArgsFlag结构体的Parse方法
func (f *ArgsFlag) Parse() error {
	// 调用f的Parse方法，传入os.Args[1:]作为参数
	return f.FlagSet.Parse(os.Args[1:])
}

// Usage函数用于打印使用方法
func (f *ArgsFlag) Usage() {
	// 打印程序名称
	Greenf(fmt.Sprintf("%s的使用方法", os.Args[0]), "\n\n")
	// 打印默认值
	f.PrintDefaults()
}

func RegisterCommand(cmd Cmd) {
	GetDefaultFlagSet().RegisterCommand(cmd)
}

func Run() error {
	return GetDefaultFlagSet().Run(os.Args[1:]...)
}

func LookupCmd(cmd string) bool {
	return GetDefaultFlagSet().LookupCmd(cmd)
}

func RunCmd(cmd string, args ...string) error {
	return GetDefaultFlagSet().RunCmd(cmd, args...)
}
