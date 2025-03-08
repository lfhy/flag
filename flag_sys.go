package flag

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 定义系统标志
var sysflag = ArgsFlag{NewFlagSet(os.Args[0], ExitOnError)}

// 解析命令行参数
func Parse() {
	sysflag.f.Parse(os.Args[1:])
}

// Bool类型定义
// 定义一个布尔类型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func BoolFullVar(p *bool, name string, title, key, env string, value bool, usage string) {
	sysflag.f.FullVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个布尔类型的标志，并设置其名称、标题、键、默认值和用法
func BoolConfigVar(p *bool, name string, title, key string, value bool, usage string) {
	sysflag.f.FullVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个布尔类型的标志，并设置其名称、环境变量、默认值和用法
func BoolEnvVar(p *bool, name string, env string, value bool, usage string) {
	sysflag.f.FullVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个布尔类型的标志，并设置其名称、默认值和用法
func BoolVar(p *bool, name string, value bool, usage string) {
	sysflag.f.Var(newBoolValue(value, p), name, usage)
}

// 定义一个布尔类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func BoolFull(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个布尔类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func BoolConfig(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个布尔类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func BoolEnv(name, env string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个布尔类型的标志，并设置其名称、默认值和用法，并返回指针
func Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolVar(p, name, value, usage)
	return p
}

// String类型定义
// 定义一个字符串类型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func StringFullVar(p *string, name string, title, key, env string, value string, usage string) {
	sysflag.f.FullVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个字符串类型的标志，并设置其名称、标题、键、默认值和用法
func StringConfigVar(p *string, name string, title, key string, value string, usage string) {
	sysflag.f.FullVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个字符串类型的标志，并设置其名称、环境变量、默认值和用法
func StringEnvVar(p *string, name string, env string, value string, usage string) {
	sysflag.f.FullVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个字符串类型的标志，并设置其名称、默认值和用法
func StringVar(p *string, name string, value string, usage string) {
	sysflag.f.Var(newStringValue(value, p), name, usage)
}

// 定义一个字符串类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func StringFull(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	sysflag.StringFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个字符串类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func StringConfig(name, title, key string, value string, usage string) *string {
	p := new(string)
	sysflag.StringConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个字符串类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func StringEnv(name, env string, value string, usage string) *string {
	p := new(string)
	sysflag.StringEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个字符串类型的标志，并设置其名称、默认值和用法，并返回指针
func String(name string, value string, usage string) *string {
	p := new(string)
	sysflag.StringVar(p, name, value, usage)
	return p
}

// 浮点类型定义
// 定义一个浮点类型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func Float64FullVar(p *float64, name string, title, key, env string, value float64, usage string) {
	sysflag.f.FullVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// 定义一个浮点类型的标志，并设置其名称、标题、键、默认值和用法
func Float64ConfigVar(p *float64, name string, title, key string, value float64, usage string) {
	sysflag.f.FullVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// 定义一个浮点类型的标志，并设置其名称、环境变量、默认值和用法
func Float64EnvVar(p *float64, name string, env string, value float64, usage string) {
	sysflag.f.FullVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// 定义一个浮点类型的标志，并设置其名称、默认值和用法
func Float64Var(p *float64, name string, value float64, usage string) {
	sysflag.f.Var(newFloat64Value(value, p), name, usage)
}

// 定义一个浮点类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func Float64Full(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个浮点类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func Float64Config(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64ConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个浮点类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func Float64Env(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64EnvVar(p, name, env, value, usage)
	return p
}

// 定义一个浮点类型的标志，并设置其名称、默认值和用法，并返回指针
func Float64(name string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64Var(p, name, value, usage)
	return p
}

// 延迟定义
// 定义一个延迟类型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func DurationFullVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	sysflag.f.FullVar(newDurationValue(value, p), name, title, key, env, usage)
}

// 定义一个延迟类型的标志，并设置其名称、标题、键、默认值和用法
func DurationConfigVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	sysflag.f.FullVar(newDurationValue(value, p), name, title, key, "", usage)
}

// 定义一个延迟类型的标志，并设置其名称、环境变量、默认值和用法
func DurationEnvVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	sysflag.f.FullVar(newDurationValue(value, p), name, "", "", env, usage)
}

// 定义一个延迟类型的标志，并设置其名称、默认值和用法
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	sysflag.f.Var(newDurationValue(value, p), name, usage)
}

// 定义一个延迟类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func DurationFull(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个延迟类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func DurationConfig(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个延迟类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func DurationEnv(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个延迟类型的标志，并设置其名称、默认值和用法，并返回指针
func Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationVar(p, name, value, usage)
	return p
}

// Int类型定义
// 定义一个整型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func IntFullVar(p *int, name string, title, key, env string, value int, usage string) {
	sysflag.f.FullVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个整型的标志，并设置其名称、标题、键、默认值和用法
func IntConfigVar(p *int, name string, title, key string, value int, usage string) {
	sysflag.f.FullVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个整型的标志，并设置其名称、环境变量、默认值和用法
func IntEnvVar(p *int, name string, env string, value int, usage string) {
	sysflag.f.FullVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个整型的标志，并设置其名称、默认值和用法
func IntVar(p *int, name string, value int, usage string) {
	sysflag.f.Var(newIntValue(value, p), name, usage)
}

// 定义一个整型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func IntFull(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	sysflag.IntFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个整型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func IntConfig(name, title, key string, value int, usage string) *int {
	p := new(int)
	sysflag.IntConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个整型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func IntEnv(name, env string, value int, usage string) *int {
	p := new(int)
	sysflag.IntEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个整型的标志，并设置其名称、默认值和用法，并返回指针
func Int(name string, value int, usage string) *int {
	p := new(int)
	sysflag.IntVar(p, name, value, usage)
	return p
}

// Int64类型定义
// 定义一个64位整型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func Int64FullVar(p *int64, name string, title, key, env string, value int64, usage string) {
	sysflag.f.FullVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个64位整型的标志，并设置其名称、标题、键、默认值和用法
func Int64ConfigVar(p *int64, name string, title, key string, value int64, usage string) {
	sysflag.f.FullVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个64位整型的标志，并设置其名称、环境变量、默认值和用法
func Int64EnvVar(p *int64, name string, env string, value int64, usage string) {
	sysflag.f.FullVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个64位整型的标志，并设置其名称、默认值和用法
func Int64Var(p *int64, name string, value int64, usage string) {
	sysflag.f.Var(newInt64Value(value, p), name, usage)
}

// 定义一个64位整型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func Int64Full(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个64位整型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func Int64Config(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64ConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个64位整型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func Int64Env(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64EnvVar(p, name, env, value, usage)
	return p
}

// 定义一个64位整型的标志，并设置其名称、默认值和用法，并返回指针
func Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64Var(p, name, value, usage)
	return p
}

// Uint类型定义
// 定义一个无符号整型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func UintFullVar(p *uint, name string, title, key, env string, value uint, usage string) {
	sysflag.f.FullVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个无符号整型的标志，并设置其名称、标题、键、默认值和用法
func UintConfigVar(p *uint, name string, title, key string, value uint, usage string) {
	sysflag.f.FullVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个无符号整型的标志，并设置其名称、环境变量、默认值和用法
func UintEnvVar(p *uint, name string, env string, value uint, usage string) {
	sysflag.f.FullVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个无符号整型的标志，并设置其名称、默认值和用法
func UintVar(p *uint, name string, value uint, usage string) {
	sysflag.f.Var(newUintValue(value, p), name, usage)
}

// 定义一个无符号整型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func UintFull(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个无符号整型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func UintConfig(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个无符号整型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func UintEnv(name, env string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个无符号整型的标志，并设置其名称、默认值和用法，并返回指针
func Uint(name string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintVar(p, name, value, usage)
	return p
}

// Uint64类型定义
// 定义一个64位无符号整型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func Uint64FullVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	sysflag.f.FullVar(newUint64Value(value, p), name, title, key, env, usage)
}

// 定义64位无符号整型的标志，并设置其名称、标题、键、默认值和用法
func Uint64ConfigVar(p *uint64, name string, title, key string, value uint64, usage string) {
	sysflag.f.FullVar(newUint64Value(value, p), name, title, key, "", usage)
}

// 定义一个64位无符号整型的标志，并设置其名称、环境变量、默认值和用法
func Uint64EnvVar(p *uint64, name string, env string, value uint64, usage string) {
	sysflag.f.FullVar(newUint64Value(value, p), name, "", "", env, usage)
}

// Uint64Var函数用于将一个uint64类型的变量注册到命令行参数中
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	// 将uint64类型的变量注册到命令行参数中
	sysflag.f.Var(newUint64Value(value, p), name, usage)
}

// Uint64Full函数用于将一个uint64类型的变量注册到命令行参数中，并设置参数的名称、标题、键和值
func Uint64Full(name, title, key, env string, value uint64, usage string) *uint64 {
	// 创建一个uint64类型的变量
	p := new(uint64)
	// 将uint64类型的变量注册到命令行参数中，并设置参数的名称、标题、键和值
	sysflag.Uint64FullVar(p, name, title, key, env, value, usage)
	return p
}

// Uint64Config函数用于将一个uint64类型的变量注册到命令行参数中，并设置参数的名称、标题、键和值
func Uint64Config(name, title, key string, value uint64, usage string) *uint64 {
	// 创建一个uint64类型的变量
	p := new(uint64)
	// 将uint64类型的变量注册到命令行参数中，并设置参数的名称、标题、键和值
	sysflag.Uint64ConfigVar(p, name, title, key, value, usage)
	return p
}

// Uint64Env函数用于将一个uint64类型的变量注册到命令行参数中，并设置参数的名称、环境变量和值
func Uint64Env(name, env string, value uint64, usage string) *uint64 {
	// 创建一个uint64类型的变量
	p := new(uint64)
	// 将uint64类型的变量注册到命令行参数中，并设置参数的名称、环境变量和值
	sysflag.Uint64EnvVar(p, name, env, value, usage)
	return p
}

// Uint64函数用于将一个uint64类型的变量注册到命令行参数中，并设置参数的名称和值
func Uint64(name string, value uint64, usage string) *uint64 {
	// 创建一个uint64类型的变量
	p := new(uint64)
	// 将uint64类型的变量注册到命令行参数中，并设置参数的名称和值
	sysflag.Uint64Var(p, name, value, usage)
	return p
}

// Args函数用于获取命令行参数
func Args() []string { return sysflag.f.args }

// ArgsLen函数用于获取命令行参数的长度
func ArgsLen() int { return len(sysflag.f.args) }

// Usage函数用于打印命令行参数的使用方法
func Usage() {
	// 打印命令行参数的使用方法
	Greenf(fmt.Sprintf("%s的使用方法", os.Args[0]), "\n\n")
	sysflag.f.PrintDefaults()
}

// PrintAll函数用于打印所有的命令行参数
func PrintAll() { sysflag.f.PrintAll() }

// SetOutput函数用于设置命令行参数的输出
func SetOutput(output io.Writer) { sysflag.f.SetOutput(output) }
