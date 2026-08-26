package flag_test

import (
	"fmt"
	"io"
	"os"
	"reflect"
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

// TestParseAllowsInterspersedArguments 验证位置参数前后均可出现选项，且非布尔选项仍能消费以负号开头的值。
func TestParseAllowsInterspersedArguments(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	var detail bool
	var count int
	f.BoolVar(&detail, "detail", false, "输出详情")
	f.IntVar(&count, "count", 0, "计数")

	if err := f.Parse([]string{"/mnt/osca", "-count", "-1", "middle", "-detail"}); err != nil {
		t.Fatalf("交错参数解析失败: %v", err)
	}
	if !detail || count != -1 {
		t.Fatalf("选项值错误: detail=%t count=%d", detail, count)
	}
	if want := []string{"/mnt/osca", "middle"}; !reflect.DeepEqual(f.Args(), want) {
		t.Fatalf("位置参数 = %v，want %v", f.Args(), want)
	}
}

// TestParseStandardStopsAtFirstPositional 验证顶层解析在子命令名处停止，不会抢先解析子命令选项。
func TestParseStandardStopsAtFirstPositional(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	var debug bool
	f.BoolVar(&debug, "debug", false, "全局调试")

	if err := f.ParseStandard([]string{"-debug", "mount-stat", "-detail", "/mnt/osca"}); err != nil {
		t.Fatalf("标准参数解析失败: %v", err)
	}
	if !debug {
		t.Fatal("全局 -debug 未解析")
	}
	if want := []string{"mount-stat", "-detail", "/mnt/osca"}; !reflect.DeepEqual(f.Args(), want) {
		t.Fatalf("标准位置参数 = %v，want %v", f.Args(), want)
	}
}

// TestParseKeepsBooleanPathAndDoubleDashLiteral 验证布尔选项不会吞掉路径，且 -- 后的 - 参数保持位置参数语义。
func TestParseKeepsBooleanPathAndDoubleDashLiteral(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	var detail bool
	f.BoolVar(&detail, "detail", false, "输出详情")

	if err := f.Parse([]string{"-detail", "/mnt/osca", "--", "-detail", "literal"}); err != nil {
		t.Fatalf("布尔参数解析失败: %v", err)
	}
	if !detail {
		t.Fatal("-detail 未设置为 true")
	}
	if want := []string{"/mnt/osca", "-detail", "literal"}; !reflect.DeepEqual(f.Args(), want) {
		t.Fatalf("位置参数 = %v，want %v", f.Args(), want)
	}
}

// TestParseSupportsSeparatedBooleanLiteral 验证兼容历史上的 -detail false 写法，同时拒绝无效的显式布尔值。
func TestParseSupportsSeparatedBooleanLiteral(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	var detail bool
	f.BoolVar(&detail, "detail", true, "输出详情")

	if err := f.Parse([]string{"target", "-detail", "false"}); err != nil {
		t.Fatalf("分离布尔值解析失败: %v", err)
	}
	if detail {
		t.Fatal("-detail false 未设置为 false")
	}
	if want := []string{"target"}; !reflect.DeepEqual(f.Args(), want) {
		t.Fatalf("位置参数 = %v，want %v", f.Args(), want)
	}

	invalid := flag.NewFlagSet("invalid", flag.ContinueOnError)
	invalid.Bool("detail", false, "输出详情")
	if err := invalid.Parse([]string{"-detail=invalid"}); err == nil {
		t.Fatal("无效的显式布尔值未返回错误")
	}
}

// TestParseReadsConfigAfterPositionals 验证 -c 即使位于位置参数之后也会按正常选项解析并加载配置。
func TestParseReadsConfigAfterPositionals(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	var host string
	f.StringConfigVar(&host, "host", "", "host", "default", "服务地址")

	cfgPath := t.TempDir() + "/app.toml"
	if err := os.WriteFile(cfgPath, []byte(`host = "from-config"`), 0644); err != nil {
		t.Fatalf("写入配置文件失败: %v", err)
	}
	if err := f.Parse([]string{"target", "-c", cfgPath}); err != nil {
		t.Fatalf("交错配置参数解析失败: %v", err)
	}
	if host != "from-config" {
		t.Fatalf("位置参数后的配置未生效: host=%q", host)
	}
	if want := []string{"target"}; !reflect.DeepEqual(f.Args(), want) {
		t.Fatalf("位置参数 = %v，want %v", f.Args(), want)
	}
}

// TestParseDoesNotReadConfigAfterDoubleDash 验证 -- 后的 -c 只作为位置参数，不能触发配置文件读取。
func TestParseDoesNotReadConfigAfterDoubleDash(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	var host string
	f.StringConfigVar(&host, "host", "", "host", "default", "服务地址")

	cfgPath := t.TempDir() + "/app.toml"
	if err := os.WriteFile(cfgPath, []byte(`host = "from-config"`), 0644); err != nil {
		t.Fatalf("写入配置文件失败: %v", err)
	}
	if err := f.Parse([]string{"target", "--", "-c", cfgPath}); err != nil {
		t.Fatalf("双横线参数解析失败: %v", err)
	}
	if host != "default" {
		t.Fatalf("-- 后的配置参数意外生效: host=%q", host)
	}
	if want := []string{"target", "-c", cfgPath}; !reflect.DeepEqual(f.Args(), want) {
		t.Fatalf("位置参数 = %v，want %v", f.Args(), want)
	}
}

func TestStringSliceAndIntSliceValues(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)

	var tags flag.StringSlice
	var ids flag.IntSlice

	f.StringSliceVar(&tags, "tags", nil, "标签")
	f.IntSliceVar(&ids, "ids", nil, "ID列表")

	if err := f.Parse([]string{"-tags=go,linux", "-tags=cli", "-ids=1,2", "-ids=3"}); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}

	if want := (flag.StringSlice{"go", "linux", "cli"}); !reflect.DeepEqual(tags, want) {
		t.Fatalf("StringSlice 解析失败，got=%v want=%v", tags, want)
	}
	if got := tags.String(); got != "go,linux,cli" {
		t.Fatalf("StringSlice String 失败，got=%q", got)
	}

	if want := (flag.IntSlice{1, 2, 3}); !reflect.DeepEqual(ids, want) {
		t.Fatalf("IntSlice 解析失败，got=%v want=%v", ids, want)
	}
	if got := ids.String(); got != "1,2,3" {
		t.Fatalf("IntSlice String 失败，got=%q", got)
	}
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

// TestRunLeavesSubcommandFlagsForCommand 验证顶层 Run 在命令名处停止，将后续选项完整交给子命令 FlagSet。
func TestRunLeavesSubcommandFlagsForCommand(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	command := &helpCmd{}
	f.RegisterCommand(command)

	if err := f.Run("cmd1", "-p", "value"); err != nil {
		t.Fatalf("运行子命令失败: %v", err)
	}
	if command.param != "value" {
		t.Fatalf("子命令参数 = %q，want value", command.param)
	}
}

type customInt int
type customString string
type customBool bool

type aliasCmd struct {
	ran bool
}

func (*aliasCmd) Name() string {
	return "list"
}

func (c *aliasCmd) Run(args ...string) error {
	c.ran = true
	return nil
}

func (*aliasCmd) Help() string {
	return "列出数据"
}

func TestVarSupportsDefinedTypes(t *testing.T) {
	args := []string{"-port=7", "-name=demo", "-debug"}
	f := flag.NewFlagSet("test", flag.PanicOnError)

	var port customInt
	var name customString
	var debug customBool

	f.Var(&flag.FlagVar{
		Value: &port,
		Name:  "port",
	})
	f.Var(&flag.FlagVar{
		Value: &name,
		Name:  "name",
	})
	f.Var(&flag.FlagVar{
		Value: &debug,
		Name:  "debug",
	})

	if err := f.Parse(args); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}

	if port != customInt(7) {
		t.Fatalf("自定义 int 类型解析失败，got=%v", port)
	}
	if name != customString("demo") {
		t.Fatalf("自定义 string 类型解析失败，got=%v", name)
	}
	if debug != customBool(true) {
		t.Fatalf("自定义 bool 类型解析失败，got=%v", debug)
	}
}

func TestVarUnsupportedTypeContinueOnError(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)

	output := captureStdout(t, func() {
		f.Var(&flag.FlagVar{
			Value: &struct{}{},
			Name:  "bad",
		})
	})

	if !strings.Contains(output, "unsupported flag value type") {
		t.Fatalf("应输出不支持类型的错误，实际输出: %q", output)
	}
	if got := f.Lookup("bad"); got != nil {
		t.Fatalf("ContinueOnError 不应注册不支持的 flag")
	}
}

func TestVarUnsupportedTypePanicOnError(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("PanicOnError 应触发 panic")
		}
	}()

	f.Var(&flag.FlagVar{
		Value: &struct{}{},
		Name:  "bad",
	})
}

func TestFlagAliasParsesCanonicalFlag(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	user := f.String("user", "", "用户名")
	f.Alias("user", "u")

	if err := f.Parse([]string{"-u", "tom"}); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if *user != "tom" {
		t.Fatalf("别名解析失败，got=%q", *user)
	}
	if got := f.Lookup("u"); got == nil || got.Name != "user" {
		t.Fatalf("Lookup 应通过别名返回主 flag")
	}
}

func TestFlagAliasSetUsesCanonicalFlag(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	user := f.String("user", "", "用户名")
	f.Alias("user", "u")

	if err := f.Set("u", "tom"); err != nil {
		t.Fatalf("Set 返回错误: %v", err)
	}
	if *user != "tom" {
		t.Fatalf("Set 别名失败，got=%q", *user)
	}
}

func TestCommandAliasRunsSameCommand(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	cmd := &aliasCmd{}
	f.RegisterCommand(cmd)
	f.AliasCmd("list", "ls")

	if err := f.RunCmd("ls"); err != nil {
		t.Fatalf("RunCmd 返回错误: %v", err)
	}
	if !cmd.ran {
		t.Fatal("命令别名没有执行到原命令")
	}
}

func TestFlagAliasContinueOnErrorDoesNotPanic(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)

	output := captureStdout(t, func() {
		f.Alias("missing", "m")
	})

	if !strings.Contains(output, "参数不存在，无法设置别名: missing") {
		t.Fatalf("应输出参数别名错误，实际输出: %q", output)
	}
	if got := f.Lookup("m"); got != nil {
		t.Fatalf("ContinueOnError 不应注册无效别名")
	}
}

func TestCommandAliasContinueOnErrorDoesNotPanic(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)

	output := captureStdout(t, func() {
		f.AliasCmd("missing", "m")
	})

	if !strings.Contains(output, "命令不存在，无法设置别名: missing") {
		t.Fatalf("应输出命令别名错误，实际输出: %q", output)
	}
	if f.LookupCmd("m") {
		t.Fatalf("ContinueOnError 不应注册无效命令别名")
	}
}

func TestPrintDefaultsShowsAliases(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	f.String("user", "", "用户名")
	f.Alias("user", "u")
	cmd := &aliasCmd{}
	f.RegisterCommand(cmd)
	f.AliasCmd("list", "ls")

	output := captureStdout(t, func() {
		f.PrintDefaults()
	})

	if !strings.Contains(output, "-user, -u") {
		t.Fatalf("帮助信息未输出参数别名:\n%s", output)
	}
	if !strings.Contains(output, "list, ls") {
		t.Fatalf("帮助信息未输出命令别名:\n%s", output)
	}
}

func TestStringsVarDefault(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	tags := f.Strings("tags", "a,b,c", "标签")
	if got := len(*tags); got != 3 {
		t.Fatalf("默认值切分错误，got=%d", got)
	}
	if (*tags)[0] != "a" || (*tags)[1] != "b" || (*tags)[2] != "c" {
		t.Fatalf("默认值解析错误，got=%v", *tags)
	}
}

func TestStringsVarCommaSplit(t *testing.T) {
	args := []string{"-tags=x,y,z"}
	f := flag.NewFlagSet("test", flag.PanicOnError)
	tags := f.Strings("tags", "", "标签")
	if err := f.Parse(args); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if len(*tags) != 3 || (*tags)[0] != "x" || (*tags)[1] != "y" || (*tags)[2] != "z" {
		t.Fatalf("逗号分隔解析错误，got=%v", *tags)
	}
}

func TestStringsVarAppend(t *testing.T) {
	args := []string{"-tags=a", "-tags=b", "-tags=c,d"}
	f := flag.NewFlagSet("test", flag.PanicOnError)
	tags := f.Strings("tags", "", "标签")
	if err := f.Parse(args); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if len(*tags) != 4 {
		t.Fatalf("多次传参累加失败，got=%v", *tags)
	}
	if (*tags)[0] != "a" || (*tags)[1] != "b" || (*tags)[2] != "c" || (*tags)[3] != "d" {
		t.Fatalf("多次传参累加内容错误，got=%v", *tags)
	}
}

func TestStringsHidden(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	tags := f.StringsHidden("tags", "x,y", "标签")
	if len(*tags) != 2 || (*tags)[0] != "x" || (*tags)[1] != "y" {
		t.Fatalf("Hidden 默认值解析错误，got=%v", *tags)
	}
	if fl := f.Lookup("tags"); fl == nil || !fl.Hidden {
		t.Fatalf("Hidden 标志未正确注册")
	}
}

// 验证 Var() 通用入口也支持 []string（反射路径）
func TestVarStringSliceDefault(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	var tags []string
	f.Var(&flag.FlagVar{
		Value:        &tags,
		Name:         "tags",
		DefaultValue: "go,linux",
	})
	if len(tags) != 2 || tags[0] != "go" || tags[1] != "linux" {
		t.Fatalf("Var 默认值解析错误，got=%v", tags)
	}
}

func TestVarStringSliceAppend(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	var tags []string
	f.Var(&flag.FlagVar{
		Value: &tags,
		Name:  "tags",
	})
	if err := f.Parse([]string{"-tags=a,b", "-tags=c"}); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if len(tags) != 3 || tags[0] != "a" || tags[1] != "b" || tags[2] != "c" {
		t.Fatalf("Var 多次传参累加错误，got=%v", tags)
	}
}

func TestVarStringSliceRawDefault(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	var tags []string
	f.Var(&flag.FlagVar{
		Value:        &tags,
		Name:         "tags",
		DefaultValue: []string{"x", "y", "z"},
	})
	if len(tags) != 3 || tags[0] != "x" || tags[1] != "y" || tags[2] != "z" {
		t.Fatalf("Var 直接 []string 默认值解析错误，got=%v", tags)
	}
}

// ============================== 数值切片测试 ==============================

func TestIntsVarDefault(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	ids := f.Ints("ids", "1,2,3", "ID列表")
	if len(*ids) != 3 || (*ids)[0] != 1 || (*ids)[1] != 2 || (*ids)[2] != 3 {
		t.Fatalf("Ints 默认值解析错误，got=%v", *ids)
	}
}

func TestIntsVarCommaSplit(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	ids := f.Ints("ids", "", "ID列表")
	if err := f.Parse([]string{"-ids=4,5,6"}); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if len(*ids) != 3 || (*ids)[0] != 4 || (*ids)[1] != 5 || (*ids)[2] != 6 {
		t.Fatalf("Ints 逗号解析错误，got=%v", *ids)
	}
}

func TestIntsVarAppend(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	ids := f.Ints("ids", "", "ID列表")
	if err := f.Parse([]string{"-ids=1", "-ids=2,3"}); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if len(*ids) != 3 || (*ids)[0] != 1 || (*ids)[1] != 2 || (*ids)[2] != 3 {
		t.Fatalf("Ints 多次累加错误，got=%v", *ids)
	}
}

func TestInt64sVarAppend(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	ids := f.Int64s("ids", "", "ID列表")
	if err := f.Parse([]string{"-ids=100", "-ids=200"}); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if len(*ids) != 2 || (*ids)[0] != 100 || (*ids)[1] != 200 {
		t.Fatalf("Int64s 累加错误，got=%v", *ids)
	}
}

func TestUintsVarAppend(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	ids := f.Uints("ids", "", "ID列表")
	if err := f.Parse([]string{"-ids=7,8", "-ids=9"}); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if len(*ids) != 3 || (*ids)[0] != 7 || (*ids)[1] != 8 || (*ids)[2] != 9 {
		t.Fatalf("Uints 累加错误，got=%v", *ids)
	}
}

func TestUint64sVarAppend(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	ids := f.Uint64s("ids", "", "ID列表")
	if err := f.Parse([]string{"-ids=1,2", "-ids=3"}); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if len(*ids) != 3 || (*ids)[0] != 1 || (*ids)[1] != 2 || (*ids)[2] != 3 {
		t.Fatalf("Uint64s 累加错误，got=%v", *ids)
	}
}

func TestIntsHidden(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	ids := f.IntsHidden("ids", "1,2", "ID列表")
	if len(*ids) != 2 || (*ids)[0] != 1 || (*ids)[1] != 2 {
		t.Fatalf("Ints Hidden 默认值错误，got=%v", *ids)
	}
	if fl := f.Lookup("ids"); fl == nil || !fl.Hidden {
		t.Fatalf("Ints Hidden 未正确注册")
	}
}

// Var() 反射路径支持数值切片
func TestVarIntSliceAppend(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	var ids []int
	f.Var(&flag.FlagVar{Value: &ids, Name: "ids"})
	if err := f.Parse([]string{"-ids=1,2", "-ids=3"}); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}
	if len(ids) != 3 || ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
		t.Fatalf("Var []int 累加错误，got=%v", ids)
	}
}

func TestVarUint64SliceDefault(t *testing.T) {
	f := flag.NewFlagSet("test", flag.PanicOnError)
	var ids []uint64
	f.Var(&flag.FlagVar{Value: &ids, Name: "ids", DefaultValue: "5,10"})
	if len(ids) != 2 || ids[0] != 5 || ids[1] != 10 {
		t.Fatalf("Var []uint64 默认值错误，got=%v", ids)
	}
}

// 自定义配置文件参数名
func TestCustomConfigFlagName(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	f.SetConfigFlagName("conf")

	// 确认 ConfigFlagName() 反映自定义值
	if got := f.ConfigFlagName(); got != "conf" {
		t.Fatalf("ConfigFlagName() 返回 %q，期望 conf", got)
	}

	var host string
	f.StringVar(&host, "host", "default", "host")

	// 临时写一个配置文件
	tmpDir := t.TempDir()
	cfgPath := tmpDir + "/app.toml"
	if err := os.WriteFile(cfgPath, []byte(`host = "from-config"`), 0644); err != nil {
		t.Fatal(err)
	}

	// 用自定义参数名 -conf 指定配置文件（不应报错）
	if err := f.Parse([]string{"-conf", cfgPath}); err != nil {
		t.Fatalf("Parse 返回错误: %v", err)
	}

	// 自动注册的隐藏参数名应是 conf 而非 c
	if fl := f.Lookup("conf"); fl == nil {
		t.Fatalf("未注册自定义配置参数名 conf")
	}
	if fl := f.Lookup("c"); fl != nil {
		t.Fatalf("不应再注册默认的 c 参数")
	}

	// 配置文件参数本身被正确解析（actual 里有值）
	if cf := f.Lookup("conf"); cf == nil || cf.Value.String() != cfgPath {
		t.Fatalf("配置参数 conf 未正确取到值，got=%v", cf)
	}
}

// 未设置时回落到 DefaultConfigFlagName
func TestDefaultConfigFlagNameFallback(t *testing.T) {
	f := flag.NewFlagSet("test", flag.ContinueOnError)
	if got := f.ConfigFlagName(); got != "c" {
		t.Fatalf("未设置时应回落到 c，got=%q", got)
	}
}
