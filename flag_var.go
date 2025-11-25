package flag

import (
	"fmt"
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
	var value Value
	switch data := opt.Value.(type) {
	case bool:
		value = newBoolValue(anyToBool(opt.DefaultValue), &data)
	case *bool:
		value = newBoolValue(anyToBool(opt.DefaultValue), data)
	case string:
		value = newStringValue(anyToString(opt.DefaultValue), &data)
	case *string:
		value = newStringValue(anyToString(opt.DefaultValue), data)
	case int:
		value = newIntValue(anyToInt(opt.DefaultValue), &data)
	case *int:
		value = newIntValue(anyToInt(opt.DefaultValue), data)
	case int64:
		value = newInt64Value(anyToInt64(opt.DefaultValue), &data)
	case *int64:
		value = newInt64Value(anyToInt64(opt.DefaultValue), data)
	case uint:
		value = newUintValue(anyToUint(opt.DefaultValue), &data)
	case *uint:
		value = newUintValue(anyToUint(opt.DefaultValue), data)
	case uint64:
		value = newUint64Value(anyToUint64(opt.DefaultValue), &data)
	case *uint64:
		value = newUint64Value(anyToUint64(opt.DefaultValue), data)
	case float64:
		value = newFloat64Value(anyToFloat64(opt.DefaultValue), &data)
	case *float64:
		value = newFloat64Value(anyToFloat64(opt.DefaultValue), data)
	case time.Duration:
		value = newDurationValue(anyToTimeDuration(opt.DefaultValue), &data)
	case *time.Duration:
		value = newDurationValue(anyToTimeDuration(opt.DefaultValue), data)
	}
	f.VarFlag(value, opt.Name, opt.ConfigSection, opt.ConfigKey, opt.Env, opt.Hidden, opt.Usage)
}

// Bool类型定义
// 定义一个Bool类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) BoolFullVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个Bool类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) BoolConfigVar(p *bool, name string, title, key string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个Bool类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) BoolEnvVar(p *bool, name string, env string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个Bool类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, "", "", "", usage)
}

// 定义一个Bool类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) BoolFull(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) BoolConfig(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) BoolEnv(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVar(p, name, value, usage)
	return p
}

// String类型定义
// 定义一个String类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) StringFullVar(p *string, name string, title, key, env string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个String类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) StringConfigVar(p *string, name string, title, key string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个String类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) StringEnvVar(p *string, name string, env string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个String类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, "", "", "", usage)
}

// 定义一个String类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) StringFull(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.StringFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) StringConfig(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.StringConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) StringEnv(name, env string, value string, usage string) *string {
	p := new(string)
	f.StringEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) String(name string, value string, usage string) *string {
	p := new(string)
	f.StringVar(p, name, value, usage)
	return p
}

// 浮点类型定义
// 定义一个Float64类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) Float64FullVar(p *float64, name string, title, key, env string, value float64, usage string) {
	f.FullVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// 定义一个Float64类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) Float64ConfigVar(p *float64, name string, title, key string, value float64, usage string) {
	f.FullVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// 定义一个Float64类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) Float64EnvVar(p *float64, name string, env string, value float64, usage string) {
	f.FullVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// 定义一个Float64类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string) {
	f.FullVar(newFloat64Value(value, p), name, "", "", "", usage)
}

// 定义一个Float64类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) Float64Full(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Float64类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Float64Config(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64ConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Float64类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) Float64Env(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64EnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Float64类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Float64(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64Var(p, name, value, usage)
	return p
}

// 延迟定义
// 定义一个Duration类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) DurationFullVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	f.FullVar(newDurationValue(value, p), name, title, key, env, usage)
}

// 定义一个Duration类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) DurationConfigVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	f.FullVar(newDurationValue(value, p), name, title, key, "", usage)
}

// 定义一个Duration类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) DurationEnvVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	f.FullVar(newDurationValue(value, p), name, "", "", env, usage)
}

// 定义一个Duration类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.FullVar(newDurationValue(value, p), name, "", "", "", usage)
}

// 定义一个Duration类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) DurationFull(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Duration类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) DurationConfig(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Duration类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) DurationEnv(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Duration类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, value, usage)
	return p
}

// Int类型定义
// 定义一个Int类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) IntFullVar(p *int, name string, title, key, env string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个Int类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) IntConfigVar(p *int, name string, title, key string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个Int类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) IntEnvVar(p *int, name string, env string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个Int类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) IntVar(p *int, name string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, "", "", "", usage)
}

// 定义一个Int类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) IntFull(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.IntFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) IntConfig(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.IntConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) IntEnv(name, env string, value int, usage string) *int {
	p := new(int)
	f.IntEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Int(name string, value int, usage string) *int {
	p := new(int)
	f.IntVar(p, name, value, usage)
	return p
}

// Int64类型定义
// 定义一个Int64类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) Int64FullVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个Int64类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) Int64ConfigVar(p *int64, name string, title, key string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个Int64类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) Int64EnvVar(p *int64, name string, env string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个Int64类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, "", "", "", usage)
}

// 定义一个Int64类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) Int64Full(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Int64Config(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) Int64Env(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64Var(p, name, value, usage)
	return p
}

// Uint类型定义
// 定义一个Uint类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) UintFullVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.FullVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个Uint类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) UintConfigVar(p *uint, name string, title, key string, value uint, usage string) {
	f.FullVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个Uint类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) UintEnvVar(p *uint, name string, env string, value uint, usage string) {
	f.FullVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个Uint类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string) {
	f.FullVar(newUintValue(value, p), name, "", "", "", usage)
}

// 定义一个Uint类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) UintFull(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintFullVar(p, name, title, key, env, value, usage)
	return p
}

func (f *FlagSet) UintConfig(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	f.UintConfigVar(p, name, title, key, value, usage)
	return p
}

func (f *FlagSet) UintEnv(name, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintEnvVar(p, name, env, value, usage)
	return p
}

func (f *FlagSet) Uint(name string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVar(p, name, value, usage)
	return p
}

// Uint64类型定义
func (f *FlagSet) Uint64FullVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	f.FullVar(newUint64Value(value, p), name, title, key, env, usage)
}

// Uint64ConfigVar 函数用于将一个 uint64 类型的变量与一个配置文件中的变量进行绑定
func (f *FlagSet) Uint64ConfigVar(p *uint64, name string, title, key string, value uint64, usage string) {
	// 调用 FullVar 函数，将 uint64 类型的变量与配置文件中的变量进行绑定
	f.FullVar(newUint64Value(value, p), name, title, key, "", usage)
}

// Uint64EnvVar 函数用于将一个 uint64 类型的变量与一个环境变量进行绑定
func (f *FlagSet) Uint64EnvVar(p *uint64, name string, env string, value uint64, usage string) {
	// 调用 FullVar 函数，将 uint64 类型的变量与环境变量进行绑定
	f.FullVar(newUint64Value(value, p), name, "", "", env, usage)
}

// Uint64Var 函数用于将一个 uint64 类型的变量与一个命令行参数进行绑定
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
	// 调用 Var 函数，将 uint64 类型的变量与命令行参数进行绑定
	f.FullVar(newUint64Value(value, p), name, "", "", "", usage)
}

// Uint64Full 函数用于将一个 uint64 类型的变量与一个配置文件、环境变量和命令行参数进行绑定
func (f *FlagSet) Uint64Full(name, title, key, env string, value uint64, usage string) *uint64 {
	// 创建一个 uint64 类型的变量
	p := new(uint64)
	// 调用 Uint64FullVar 函数，将 uint64 类型的变量与配置文件、环境变量和命令行参数进行绑定
	f.Uint64FullVar(p, name, title, key, env, value, usage)
	// 返回 uint64 类型的变量
	return p
}

// Uint64Config 函数用于将一个 uint64 类型的变量与一个配置文件和命令行参数进行绑定
func (f *FlagSet) Uint64Config(name, title, key string, value uint64, usage string) *uint64 {
	// 创建一个 uint64 类型的变量
	p := new(uint64)
	// 调用 Uint64ConfigVar 函数，将 uint64 类型的变量与配置文件和命令行参数进行绑定
	f.Uint64ConfigVar(p, name, title, key, value, usage)
	// 返回 uint64 类型的变量
	return p
}

// Uint64Env 函数用于将一个 uint64 类型的变量与一个环境变量和命令行参数进行绑定
func (f *FlagSet) Uint64Env(name, env string, value uint64, usage string) *uint64 {
	// 创建一个 uint64 类型的变量
	p := new(uint64)
	// 调用 Uint64EnvVar 函数，将 uint64 类型的变量与环境变量和命令行参数进行绑定
	f.Uint64EnvVar(p, name, env, value, usage)
	// 返回 uint64 类型的变量
	return p
}

// Uint64函数用于设置一个uint64类型的flag，并返回一个指向该flag的指针
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64 {
	// 创建一个uint64类型的指针
	p := new(uint64)
	// 调用Uint64Var函数，将指针、flag名称、默认值和用法作为参数传入
	f.Uint64Var(p, name, value, usage)
	// 返回指向该flag的指针
	return p
}
