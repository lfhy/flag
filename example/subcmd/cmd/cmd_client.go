package cmd

import (
	"fmt"

	"github.com/lfhy/flag"
)

// ClientCmd 结构体，包含一个 FlagSet 和一个 Addr 字段
type ClientCmd struct {
	*flag.FlagSet
	Addr string
}

// 初始化函数，注册 ClientCmd 命令
func init() {
	RegisterCommand(&ClientCmd{})
}

// 返回命令的简短描述
func (c *ClientCmd) UsageShort() string {
	return "client"
}

// 打印命令的详细描述
func (c *ClientCmd) Usage() {
	fmt.Println(c.Name(), ":", c.UsageShort())
	fmt.Println()
	c.FlagSet.PrintDefaults()
}

// 返回命令的名称
func (*ClientCmd) Name() []string { return []string{"client"} }

// 初始化命令，设置 FlagSet 和 Addr 字段
func (c *ClientCmd) Init(arg ...string) error {
	c.FlagSet = flag.NewFlagSet(c.Name()[0], flag.ContinueOnError)
	c.StringVar(&c.Addr, "s", "127.0.0.1:8080", "server")
	return c.FlagSet.Parse(arg)
}

// 运行命令，打印 Addr 字段
func (s *ClientCmd) Run() error {
	fmt.Println("Server:", s.Addr)
	return nil
}
