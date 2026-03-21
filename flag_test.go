package flag_test

import (
	"fmt"
	"io"
	"os"
	"strings"
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

type helpCmd struct {
	*flag.FlagSet
	param string
}

func (*helpCmd) Name() string {
	return "cmd1"
}

func (c *helpCmd) Init(args ...string) error {
	c.FlagSet = flag.NewFlagSet("cmd1", flag.ContinueOnError)
	c.StringVar(&c.param, "p", "", "cmd1参数")
	return c.Parse(args)
}

func (*helpCmd) Run(args ...string) error {
	return nil
}

func (*helpCmd) Help() string {
	return "运行CMD1"
}

func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("创建 stdout 管道失败: %v", err)
	}

	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	fn()

	if err := w.Close(); err != nil {
		t.Fatalf("关闭 stdout 写端失败: %v", err)
	}

	output, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("读取 stdout 失败: %v", err)
	}
	if err := r.Close(); err != nil {
		t.Fatalf("关闭 stdout 读端失败: %v", err)
	}
	return string(output)
}

func TestRunCmdHelpPrintedOnce(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	f.RegisterCommand(&helpCmd{})

	var err error
	output := captureStdout(t, func() {
		err = f.RunCmd("cmd1", "-h")
	})

	if err != nil {
		t.Fatalf("帮助请求不应再作为错误返回: %v", err)
	}
	if count := strings.Count(output, "cmd1参数"); count != 1 {
		t.Fatalf("帮助信息应只打印一次，实际打印了 %d 次:\n%s", count, output)
	}
	if strings.Contains(output, "运行CMD1") {
		t.Fatalf("帮助请求不应重复打印 Help 文本:\n%s", output)
	}
}
