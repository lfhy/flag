package flag

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// FlagVar 定义一个支持多来源设置的标志变量
type FlagVar struct {
	// Value 存储实际值的指针，确保可以通过反射修改其内容
	Value interface{}
	// Name 是用于命令行参数解析的名字，例如 "--port"
	Name string
	// Env 是对应的环境变量名，例如 "APP_PORT"
	Env string
	// ConfigSection 表示在配置文件中的节标题，例如 "[server]"
	ConfigSection string
	// ConfigKey 表示在配置文件中的键名，例如 "port"
	ConfigKey string
	// Usage 提供帮助文档描述此参数的作用
	Usage string
	// Hidden 控制该参数是否出现在 --help 输出中
	Hidden bool
	// DefaultValue 设置默认值，当所有来源都没有指定时使用
	DefaultValue interface{}
}

func anyToBool(value any) bool {
	switch data := value.(type) {
	case bool:
		return data
	case int:
		return data != 0
	case int64:
		return data != 0
	case uint:
		return data != 0
	case uint64:
		return data != 0
	case string:
		return strings.ToLower(data) == "true"
	}
	return false
}

func anyToString(value any) string {
	switch data := value.(type) {
	case string:
		return data
	case int:
		return fmt.Sprintf("%d", data)
	case int64:
		return fmt.Sprintf("%d", data)
	case uint:
		return fmt.Sprintf("%d", data)
	case uint64:
		return fmt.Sprintf("%d", data)
	case bool:
		return fmt.Sprintf("%t", data)
	}
	return ""
}

func anyToInt(value any) int {
	switch data := value.(type) {
	case int:
		return data
	case int64:
		return int(data)
	case uint:
		return int(data)
	case uint64:
		return int(data)
	case string:
		i, _ := strconv.Atoi(data)
		return i
	}
	return 0
}

func anyToInt64(value any) int64 {
	switch data := value.(type) {
	case int:
		return int64(data)
	case int64:
		return data
	case uint:
		return int64(data)
	case uint64:
		return int64(data)
	case string:
		i, _ := strconv.ParseInt(data, 10, 64)
		return i
	}
	return 0
}

func anyToUint(value any) uint {
	switch data := value.(type) {
	case int:
		return uint(data)
	case int64:
		return uint(data)
	case uint:
		return data
	case uint64:
		return uint(data)
	case string:
		i, _ := strconv.ParseUint(data, 10, 64)
		return uint(i)
	}
	return 0
}

func anyToUint64(value any) uint64 {
	switch data := value.(type) {
	case int:
		return uint64(data)
	case int64:
		return uint64(data)
	case uint:
		return uint64(data)
	case uint64:
		return data
	case string:
		i, _ := strconv.ParseUint(data, 10, 64)
		return i
	}
	return 0
}

func anyToFloat64(value any) float64 {
	switch data := value.(type) {
	case float64:
		return data
	case int:
		return float64(data)
	case string:
		f, _ := strconv.ParseFloat(data, 64)
		return f
	}
	return 0
}

func anyToTimeDuration(value any) time.Duration {
	switch data := value.(type) {
	case time.Duration:
		return data
	case string:
		d, _ := time.ParseDuration(data)
		return d
	}
	return 0
}

func (f *FlagSet) Var(opt *FlagVar) {
	if opt.Name == "" && opt.ConfigSection != "" && opt.ConfigKey != "" {
		opt.Name = opt.ConfigSection + "-" + opt.ConfigKey
	}
	if opt.ConfigKey == "" && opt.ConfigSection == "" {
		info := strings.Split(opt.Name, "-")
		if len(info) >= 2 {
			opt.ConfigSection = info[0]
			opt.ConfigKey = strings.Join(info[0:], "-")
		} else {
			opt.ConfigSection = "config"
			opt.ConfigKey = opt.Name
		}
	}
	if opt.ConfigKey == "" && opt.ConfigSection != "" {
		opt.ConfigKey = opt.Name
	}
	value, err := newReflectValue(opt.Value, opt.DefaultValue)
	if err != nil {
		f.handleError(err)
		return
	}
	f.VarFlag(value, opt.Name, opt.ConfigSection, opt.ConfigKey, opt.Env, opt.Hidden, opt.Usage)
}

func (f *FlagSet) handleError(err error) {
	switch f.errorHandling {
	case ContinueOnError:
		fmt.Fprintln(f.out(), err)
	case ExitOnError:
		fmt.Fprintln(f.out(), err)
		os.Exit(2)
	case PanicOnError:
		panic(err)
	default:
		panic(err)
	}
}
