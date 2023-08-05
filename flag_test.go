package flag_test

import (
	"fmt"
	"testing"

	"github.com/lfhy/flag"
)

// 测试解析
func TestVar(t *testing.T) {
	args := []string{"-int=1", "-string=2", "-float=3"}
	f := flag.NewFlagSet("test", flag.PanicOnError)
	i := f.Int("int", 0, "测试Int")
	s := f.String("string", "0", "测试String")
	fl := f.Float64("float", 0, "测试Float")
	f.Parse(args)
	if *i != 1 {
		t.Fatal("Int解析错误")
	}
	if *s != "2" {
		t.Fatal("String解析错误")
	}
	if *fl != 3 {
		t.Fatal("Float解析错误")
	}
	fmt.Println("测试通过")

}

// 测试打印表格
func TestPrintTable(t *testing.T) {
	args := []string{"-int=1", "-string=2", "-float=3"}
	f := flag.NewFlagSet("test", flag.PanicOnError)
	f.Int("int", 0, "测试Int")
	f.String("string", "0", "测试String")
	f.Float64("float", 0, "测试Float")
	f.Parse(args)
	f.PrintAll()
}
