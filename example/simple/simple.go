package main

import (
	"fmt"

	"github.com/lfhy/flag"
)

// 定义全局变量
var (
	Integer = 1
	String  = "hello"
	Float   = 1.1
	Bool    = true
)

// 初始化函数，用于解析命令行参数
func init() {
	// 定义命令行参数
	flag.IntVar(&Integer, "i", 2, "integer")
	flag.StringVar(&String, "s", "string", "string")
	flag.Float64Var(&Float, "f", 2.2, "float")
	flag.BoolVar(&Bool, "b", false, "bool")
	// 解析命令行参数
	flag.Parse()
}

func main() {
	// 打印命令行参数的值
	fmt.Printf("Integer: %v\n", Integer)
	fmt.Printf("String: %v\n", String)
	fmt.Printf("Float: %v\n", Float)
	fmt.Printf("Bool: %v\n", Bool)
}
