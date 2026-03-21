package main

import (
	"fmt"

	"github.com/lfhy/flag"
)

var Global string

func init() {
	flag.StringVar(&Global, "g", "", "全局参数")
}

type Cmd1 struct {
	*flag.FlagSet
	Cmd1Params string
}

func (*Cmd1) Name() string {
	return "cmd1"
}

func (s *Cmd1) Init(args ...string) error {
	s.FlagSet = flag.NewFlagSet("cmd1", flag.ContinueOnError)
	s.Var(&flag.FlagVar{
		Value: &s.Cmd1Params,
		Name:  "p",
		Usage: "cmd1的参数",
	})
	return s.Parse(args)
}

func (*Cmd1) Run(args ...string) error {
	fmt.Println("运行CMD1")
	return nil
}

func (*Cmd1) Help() string {
	return "运行CMD1"
}

type Cmd2 struct {
}

func (Cmd2) Name() string {
	return "cmd2"
}

func (Cmd2) Init(args ...string) error {
	return nil
}

func (Cmd2) Run(args ...string) error {
	fmt.Println("运行CMD2")
	return nil
}

func main() {
	flag.RegisterCommand(&Cmd1{})
	flag.RegisterCommand(Cmd2{})
	flag.Parse()
	err := flag.Run()
	if err != nil {
		fmt.Println(err)
	}
}
