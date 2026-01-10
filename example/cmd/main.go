package main

import (
	"fmt"

	"github.com/lfhy/flag"
)

type Cmd1 struct {
}

func (Cmd1) Name() string {
	return "cmd1"
}

func (Cmd1) Init(args ...string) error {
	return nil
}

func (Cmd1) Run(args ...string) error {
	fmt.Println("运行CMD1")
	return nil
}

func (Cmd1) Help() string {
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
	flag.RegisterCommand(Cmd1{})
	flag.RegisterCommand(Cmd2{})
	err := flag.Run()
	if err != nil {
		fmt.Println(err)
	}
}
