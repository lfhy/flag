package cmd

import (
	"fmt"

	"github.com/lfhy/flag"
)

// ServerCmd 结构体，包含一个 FlagSet 和一个 Port 字段
type ServerCmd struct {
	*flag.FlagSet
	Port string
}

// 初始化函数，注册 ServerCmd 命令
func init() {
	RegisterCommand(&ServerCmd{})
}

// 返回命令的简短描述
func (c *ServerCmd) UsageShort() string {
	return "server"
}

// 打印命令的详细描述
func (c *ServerCmd) Usage() {
	fmt.Println(c.Name(), ":", c.UsageShort())
	fmt.Println()
	c.FlagSet.PrintDefaults()
}

// 返回命令的名称
func (*ServerCmd) Name() []string { return []string{"server"} }

// 初始化命令，设置 FlagSet 和 Port 字段的值
func (c *ServerCmd) Init(arg ...string) error {
	c.FlagSet = flag.NewFlagSet(c.Name()[0], flag.ContinueOnError)
	c.StringVar(&c.Port, "p", "8080", "port")
	return c.FlagSet.Parse(arg)
}

// 运行命令，打印 Port 字段的值
func (s *ServerCmd) Run() error {
	fmt.Println("Port:", s.Port)
	return nil
}
