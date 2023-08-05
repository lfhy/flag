package flag

import (
	"fmt"
	"reflect"
)

type ErrorHandling int

const (
	ContinueOnError ErrorHandling = iota // 忽略错误
	ExitOnError                          // 错误退出
	PanicOnError                         // 崩溃退出
)

type Value interface {
	String() string
	Set(string) error
}

type Flag struct {
	Name        string // 命令行名称
	EnvName     string // 环境变量名称
	ConfigTitle string // 配置文件标题
	ConfigKey   string // 配置文件关键字
	Usage       string // 帮助信息
	Value       Value  // 参数值
	DefValue    string // 默认值
	Hiden       bool   //是否隐藏参数
}

// 默认使用方法
func defaultUsage(f *FlagSet) {
	if f.name == "" {
		fmt.Fprintf(f.out(), "使用方法:\n")
	} else {
		fmt.Fprint(f.out(), Green(fmt.Sprintf("%s的使用方法:\n\n", f.name)))
	}
	f.PrintDefaults()
}

func isZeroValue(flag *Flag, value string) bool {
	typ := reflect.TypeOf(flag.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Ptr {
		z = reflect.New(typ.Elem())
	} else {
		z = reflect.Zero(typ)
	}
	if value == z.Interface().(Value).String() {
		return true
	}

	switch value {
	case "false", "", "0":
		return true
	}
	return false
}

func UnquoteUsage(flag *Flag) (name string, usage string) {
	// Look for a back-quoted name, but avoid the strings package.
	usage = flag.Usage
	for i := 0; i < len(usage); i++ {
		if usage[i] == '`' {
			for j := i + 1; j < len(usage); j++ {
				if usage[j] == '`' {
					name = usage[i+1 : j]
					usage = usage[:i] + name + usage[j+1:]
					return name, usage
				}
			}
			break // Only one back quote; use type name.
		}
	}
	// No explicit name, so use type if we can find one.
	name = "value"
	switch flag.Value.(type) {
	case boolFlag:
		name = ""
	case *durationValue:
		name = "duration"
	case *float64Value:
		name = "float"
	case *intValue, *int64Value:
		name = "int"
	case *stringValue:
		name = "string"
	case *uintValue, *uint64Value:
		name = "uint"
	}
	return
}
