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

// Parse 解析顶层进程参数，并在首个位置参数（通常是子命令）处停止，避免把子命令选项当作全局选项。
func (f *ArgsFlag) Parse() error {
	return f.FlagSet.ParseStandard(os.Args[1:])
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

func Alias(name, alias string) {
	GetDefaultFlagSet().Alias(name, alias)
}

func AliasCmd(name, alias string) {
	GetDefaultFlagSet().AliasCmd(name, alias)
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
