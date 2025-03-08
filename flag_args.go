package flag

import (
	"fmt"
	"io"
	"os"
	"time"
)

type ArgsFlag struct {
	f *FlagSet
}

// 创建一个新的ArgsFlag实例
func NewArgsFlag() *ArgsFlag {
	// 使用os.Args[0]作为FlagSet的名称，ExitOnError表示在遇到错误时退出程序
	return &ArgsFlag{NewFlagSet(os.Args[0], ExitOnError)}
}

// 解析ArgsFlag结构体的Parse方法
func (f *ArgsFlag) Parse() {
	// 调用f.f的Parse方法，传入os.Args[1:]作为参数
	f.f.Parse(os.Args[1:])
}

// Bool类型定义
// BoolFullVar 函数用于创建一个布尔类型的全变量标志
func (f *ArgsFlag) BoolFullVar(p *bool, name string, title, key, env string, value bool, usage string) {
	// 调用 FullVar 函数，传入一个 newBoolValue 函数返回的值，以及标志的名称、标题、键、环境变量、默认值和用法
	f.f.FullVar(newBoolValue(value, p), name, title, key, env, usage)
}

// BoolConfigVar 函数用于将一个布尔类型的配置变量添加到 ArgsFlag 中
func (f *ArgsFlag) BoolConfigVar(p *bool, name string, title, key string, value bool, usage string) {
	// 调用 FullVar 函数，将布尔类型的配置变量添加到 ArgsFlag 中
	f.f.FullVar(newBoolValue(value, p), name, title, key, "", usage)
}

// BoolEnvVar 函数用于设置一个布尔类型的环境变量
func (f *ArgsFlag) BoolEnvVar(p *bool, name string, env string, value bool, usage string) {
	// 调用 FullVar 函数，传入一个 newBoolValue 函数返回的值，以及 name、env、value、usage 参数
	f.f.FullVar(newBoolValue(value, p), name, "", "", env, usage)
}

// BoolVar函数用于将一个bool类型的变量与一个flag绑定
func (f *ArgsFlag) BoolVar(p *bool, name string, value bool, usage string) {
	// 调用ArgsFlag的Var方法，将一个bool类型的变量与一个flag绑定
	f.f.Var(newBoolValue(value, p), name, usage)
}

// BoolFull 函数用于创建一个布尔类型的标志，并返回一个指向该标志的指针
func (f *ArgsFlag) BoolFull(name, title, key, env string, value bool, usage string) *bool {
	// 创建一个布尔类型的变量
	p := new(bool)
	// 调用 BoolFullVar 函数，将变量 p 传递给该函数
	f.BoolFullVar(p, name, title, key, env, value, usage)
	// 返回指向该变量的指针
	return p
}

// BoolConfig 函数用于创建一个布尔类型的配置项
func (f *ArgsFlag) BoolConfig(name, title, key string, value bool, usage string) *bool {
	// 创建一个布尔类型的指针
	p := new(bool)
	// 调用 BoolConfigVar 函数，将指针、名称、标题、键、值和用法作为参数传入
	f.BoolConfigVar(p, name, title, key, value, usage)
	// 返回指针
	return p
}

// BoolEnv 函数用于设置一个布尔类型的环境变量
func (f *ArgsFlag) BoolEnv(name, env string, value bool, usage string) *bool {
	// 创建一个布尔类型的指针
	p := new(bool)
	// 调用 BoolEnvVar 函数，设置环境变量
	f.BoolEnvVar(p, name, env, value, usage)
	// 返回布尔类型的指针
	return p
}

// Bool函数用于创建一个布尔类型的flag，并返回一个指向该flag的指针
func (f *ArgsFlag) Bool(name string, value bool, usage string) *bool {
	// 创建一个指向bool类型的指针
	p := new(bool)
	// 调用BoolVar函数，将指针、名称、默认值和用法作为参数传入
	f.BoolVar(p, name, value, usage)
	// 返回指向该flag的指针
	return p
}

// String类型定义
// StringFullVar 函数用于创建一个带有全名变量的字符串标志
func (f *ArgsFlag) StringFullVar(p *string, name string, title, key, env string, value string, usage string) {
	// 调用 FullVar 函数，传入一个带有初始值的字符串值，以及标志的名称、标题、键、环境变量和用法
	f.f.FullVar(newStringValue(value, p), name, title, key, env, usage)
}

// StringConfigVar 函数用于设置字符串类型的配置变量
func (f *ArgsFlag) StringConfigVar(p *string, name string, title, key string, value string, usage string) {
	// 调用 FullVar 函数，传入 stringValue 类型的参数，设置字符串类型的配置变量
	f.f.FullVar(newStringValue(value, p), name, title, key, "", usage)
}

// StringEnvVar 函数用于设置一个带有环境变量的字符串标志
func (f *ArgsFlag) StringEnvVar(p *string, name string, env string, value string, usage string) {
	// 调用 FullVar 函数，传入参数为：带有环境变量的字符串值，标志名称，空字符串，空字符串，环境变量名称，标志描述
	f.f.FullVar(newStringValue(value, p), name, "", "", env, usage)
}

// StringVar函数用于将一个字符串类型的参数绑定到ArgsFlag结构体中
func (f *ArgsFlag) StringVar(p *string, name string, value string, usage string) {
	// 调用ArgsFlag结构体的Var方法，将一个StringValue类型的参数绑定到ArgsFlag结构体中
	f.f.Var(newStringValue(value, p), name, usage)
}

// StringFull 函数用于创建一个 ArgsFlag 类型的指针，并返回一个字符串类型的指针
func (f *ArgsFlag) StringFull(name, title, key, env string, value string, usage string) *string {
	// 创建一个字符串类型的指针
	p := new(string)
	// 调用 StringFullVar 函数，将参数传递给该函数
	f.StringFullVar(p, name, title, key, env, value, usage)
	// 返回字符串类型的指针
	return p
}

// StringConfig 函数用于创建一个字符串类型的配置项
func (f *ArgsFlag) StringConfig(name, title, key string, value string, usage string) *string {
	// 创建一个字符串类型的指针
	p := new(string)
	// 调用 StringConfigVar 函数，将指针、名称、标题、键、值和用法作为参数传入
	f.StringConfigVar(p, name, title, key, value, usage)
	// 返回指针
	return p
}

// StringEnv 函数用于设置一个字符串类型的命令行参数，并从环境变量中获取该参数的值
func (f *ArgsFlag) StringEnv(name, env string, value string, usage string) *string {
	// 创建一个指向字符串的指针
	p := new(string)
	// 调用 StringEnvVar 函数，将指针、参数名、环境变量名、默认值和参数说明作为参数传入
	f.StringEnvVar(p, name, env, value, usage)
	// 返回指向字符串的指针
	return p
}

// String方法用于将一个字符串类型的参数添加到ArgsFlag中，并返回该参数的指针
func (f *ArgsFlag) String(name string, value string, usage string) *string {
	// 创建一个字符串类型的指针
	p := new(string)
	// 调用StringVar方法将参数添加到ArgsFlag中
	f.StringVar(p, name, value, usage)
	// 返回该参数的指针
	return p
}

// 浮点类型定义
// Float64FullVar 函数用于设置一个浮点型变量的全名、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) Float64FullVar(p *float64, name string, title, key, env string, value float64, usage string) {
	// 调用 FullVar 函数，传入一个 Float64Value 类型的参数，该参数包含默认值和指针
	f.f.FullVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// Float64ConfigVar 函数用于将一个 float64 类型的变量注册到 ArgsFlag 中
func (f *ArgsFlag) Float64ConfigVar(p *float64, name string, title, key string, value float64, usage string) {
	// 调用 FullVar 函数，将 float64 类型的变量注册到 ArgsFlag 中
	f.f.FullVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// Float64EnvVar 函数用于设置一个浮点型环境变量
func (f *ArgsFlag) Float64EnvVar(p *float64, name string, env string, value float64, usage string) {
	// 调用 FullVar 函数，传入一个 newFloat64Value 函数返回的值，name，env，value，usage
	f.f.FullVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// Float64Var函数用于将一个float64类型的参数绑定到ArgsFlag结构体中
func (f *ArgsFlag) Float64Var(p *float64, name string, value float64, usage string) {
	// 调用ArgsFlag结构体的Var方法，将newFloat64Value函数返回的Float64Value类型的参数绑定到ArgsFlag结构体中
	f.f.Var(newFloat64Value(value, p), name, usage)
}

// Float64Full函数用于创建一个float64类型的flag，并返回该flag的指针
func (f *ArgsFlag) Float64Full(name, title, key, env string, value float64, usage string) *float64 {
	// 创建一个float64类型的指针
	p := new(float64)
	// 调用Float64FullVar函数，将指针、name、title、key、env、value、usage作为参数传入
	f.Float64FullVar(p, name, title, key, env, value, usage)
	// 返回指针
	return p
}

// Float64Config 函数用于创建一个 float64 类型的配置项
func (f *ArgsFlag) Float64Config(name, title, key string, value float64, usage string) *float64 {
	// 创建一个 float64 类型的指针
	p := new(float64)
	// 调用 Float64ConfigVar 函数，将指针、名称、标题、键、值和用法作为参数传入
	f.Float64ConfigVar(p, name, title, key, value, usage)
	// 返回指针
	return p
}

// Float64Env 函数用于从环境变量中获取一个 float64 类型的值，并将其赋值给指针 p
func (f *ArgsFlag) Float64Env(name, env string, value float64, usage string) *float64 {
	// 创建一个 float64 类型的指针 p
	p := new(float64)
	// 调用 Float64EnvVar 函数，将指针 p、name、env、value 和 usage 作为参数传入
	f.Float64EnvVar(p, name, env, value, usage)
	// 返回指针 p
	return p
}

// Float64函数用于设置一个float64类型的参数，并返回该参数的指针
func (f *ArgsFlag) Float64(name string, value float64, usage string) *float64 {
	// 创建一个float64类型的指针
	p := new(float64)
	// 调用Float64Var函数，将指针、参数名、默认值和参数说明作为参数传入
	f.Float64Var(p, name, value, usage)
	// 返回该参数的指针
	return p
}

// 延迟定义
// DurationFullVar 函数用于设置一个带有全变量的时间类型标志
func (f *ArgsFlag) DurationFullVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	// 调用 FullVar 函数，传入一个 newDurationValue 函数返回的 DurationValue 类型的参数
	f.f.FullVar(newDurationValue(value, p), name, title, key, env, usage)
}

// DurationConfigVar 函数用于设置一个 time.Duration 类型的配置变量
func (f *ArgsFlag) DurationConfigVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	// 调用 FullVar 函数，传入一个 newDurationValue 函数返回的值，以及 name、title、key、value 和 usage 参数
	f.f.FullVar(newDurationValue(value, p), name, title, key, "", usage)
}

// DurationEnvVar 函数用于设置一个带有环境变量的时间类型参数
func (f *ArgsFlag) DurationEnvVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	// 调用 FullVar 函数，传入 newDurationValue 函数生成的值，参数 name，env，usage
	f.f.FullVar(newDurationValue(value, p), name, "", "", env, usage)
}

// DurationVar函数：将time.Duration类型的值绑定到指定的flag上
func (f *ArgsFlag) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	// 创建一个新的DurationValue类型的值，并将其绑定到指定的flag上
	f.f.Var(newDurationValue(value, p), name, usage)
}

// DurationFull函数：将time.Duration类型的值绑定到指定的flag上，并设置title、key、env等参数
func (f *ArgsFlag) DurationFull(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	// 创建一个新的time.Duration类型的指针
	p := new(time.Duration)
	// 调用DurationFullVar函数，将time.Duration类型的值绑定到指定的flag上，并设置title、key、env等参数
	// 返回time.Duration类型的指针
	f.DurationFullVar(p, name, title, key, env, value, usage)
	return p
}

// DurationConfig函数：将time.Duration类型的值绑定到指定的flag上，并设置key、env等参数
func (f *ArgsFlag) DurationConfig(name, title, key string, value time.Duration, usage string) *time.Duration {
	// 创建一个新的time.Duration类型的指针
	p := new(time.Duration)
	// 调用DurationConfigVar函数，将time.Duration类型的值绑定到指定的flag上，并设置key、env等参数
	f.DurationConfigVar(p, name, title, key, value, usage)
	// 返回time.Duration类型的指针
	return p
}

// DurationEnv函数：将time.Duration类型的值绑定到指定的flag上，并设置env参数
func (f *ArgsFlag) DurationEnv(name, env string, value time.Duration, usage string) *time.Duration {
	// 创建一个新的time.Duration类型的指针
	p := new(time.Duration)
	// 调用DurationEnvVar函数，将time.Duration类型的值绑定到指定的flag上，并设置env参数
	f.DurationEnvVar(p, name, env, value, usage)
	// 返回time.Duration类型的指针
	return p
}

// Duration函数：将time.Duration类型的值绑定到指定的flag上
func (f *ArgsFlag) Duration(name string, value time.Duration, usage string) *time.Duration {
	// 创建一个新的time.Duration类型的指针
	p := new(time.Duration)
	// 调用DurationVar函数，将time.Duration类型的值绑定到指定的flag上
	f.DurationVar(p, name, value, usage)
	// 返回time.Duration类型的指针
	return p
}

// Int类型定义
// IntFullVar 函数用于创建一个带有全名、标题、键、环境变量、默认值和用法的整型标志
func (f *ArgsFlag) IntFullVar(p *int, name string, title, key, env string, value int, usage string) {
	// 调用 FullVar 函数，传入一个 newIntValue 函数创建的整型值，以及标志的全名、标题、键、环境变量和用法
	f.f.FullVar(newIntValue(value, p), name, title, key, env, usage)
}

// IntConfigVar 函数用于将一个整型变量与一个配置文件变量关联
func (f *ArgsFlag) IntConfigVar(p *int, name string, title, key string, value int, usage string) {
	// 调用 FullVar 函数，将整型变量与配置文件变量关联
	f.f.FullVar(newIntValue(value, p), name, title, key, "", usage)
}

// IntEnvVar 函数用于将一个整型变量与一个环境变量关联
func (f *ArgsFlag) IntEnvVar(p *int, name string, env string, value int, usage string) {
	// 调用 FullVar 函数，将整型变量与环境变量关联
	f.f.FullVar(newIntValue(value, p), name, "", "", env, usage)
}

// IntVar 函数用于将一个整型变量与一个命令行参数关联
func (f *ArgsFlag) IntVar(p *int, name string, value int, usage string) {
	// 调用 Var 函数，将整型变量与命令行参数关联
	f.f.Var(newIntValue(value, p), name, usage)
}

// IntFull 函数用于将一个整型变量与一个配置文件变量、一个命令行参数关联
func (f *ArgsFlag) IntFull(name, title, key, env string, value int, usage string) *int {
	// 创建一个新的整型变量
	p := new(int)
	// 调用 IntFullVar 函数，将整型变量与配置文件变量、命令行参数关联
	f.IntFullVar(p, name, title, key, env, value, usage)
	// 返回整型变量
	return p
}

// IntConfig 函数用于将一个整型变量与一个配置文件变量关联
func (f *ArgsFlag) IntConfig(name, title, key string, value int, usage string) *int {
	// 创建一个新的整型变量
	p := new(int)
	// 调用 IntConfigVar 函数，将整型变量与配置文件变量关联
	f.IntConfigVar(p, name, title, key, value, usage)
	// 返回整型变量
	return p
}

// IntEnv 函数用于将一个整型变量与一个环境变量关联
func (f *ArgsFlag) IntEnv(name, env string, value int, usage string) *int {
	// 创建一个新的整型变量
	p := new(int)
	// 调用 IntEnvVar 函数，将整型变量与环境变量关联
	f.IntEnvVar(p, name, env, value, usage)
	// 返回整型变量
	return p
}

// Int 函数用于将一个整型变量与一个命令行参数关联
func (f *ArgsFlag) Int(name string, value int, usage string) *int {
	p := new(int)
	f.IntVar(p, name, value, usage)
	return p
}

// Int64类型定义
// 定义一个Int64类型的FullVar方法，用于设置Int64类型的变量
func (f *ArgsFlag) Int64FullVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.f.FullVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个Int64类型的ConfigVar方法，用于设置Int64类型的变量
func (f *ArgsFlag) Int64ConfigVar(p *int64, name string, title, key string, value int64, usage string) {
	f.f.FullVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个Int64类型的EnvVar方法，用于设置Int64类型的变量
func (f *ArgsFlag) Int64EnvVar(p *int64, name string, env string, value int64, usage string) {
	f.f.FullVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个Int64类型的Var方法，用于设置Int64类型的变量
func (f *ArgsFlag) Int64Var(p *int64, name string, value int64, usage string) {
	f.f.Var(newInt64Value(value, p), name, usage)
}

// 定义一个Int64类型的Full方法，用于设置Int64类型的变量
func (f *ArgsFlag) Int64Full(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int64类型的Config方法，用于设置Int64类型的变量
func (f *ArgsFlag) Int64Config(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int64类型的Env方法，用于设置Int64类型的变量
func (f *ArgsFlag) Int64Env(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Int64类型的Int64方法，用于设置Int64类型的变量
func (f *ArgsFlag) Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64Var(p, name, value, usage)
	return p
}

// Uint类型定义
// 定义一个Uint类型的FullVar方法，用于设置Uint类型的变量
func (f *ArgsFlag) UintFullVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.f.FullVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个Uint类型的ConfigVar方法，用于设置Uint类型的变量
func (f *ArgsFlag) UintConfigVar(p *uint, name string, title, key string, value uint, usage string) {
	f.f.FullVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个Uint类型的EnvVar方法，用于设置Uint类型的变量
func (f *ArgsFlag) UintEnvVar(p *uint, name string, env string, value uint, usage string) {
	f.f.FullVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个Uint类型的Var方法，用于设置Uint类型的变量
func (f *ArgsFlag) UintVar(p *uint, name string, value uint, usage string) {
	f.f.Var(newUintValue(value, p), name, usage)
}

// 定义一个Uint类型的Full方法，用于设置Uint类型的变量
func (f *ArgsFlag) UintFull(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Uint类型的Config方法，用于设置Uint类型的变量
func (f *ArgsFlag) UintConfig(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	f.UintConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Uint类型的Env方法，用于设置Uint类型的变量
func (f *ArgsFlag) UintEnv(name, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Uint类型的Uint方法，用于设置Uint类型的变量
func (f *ArgsFlag) Uint(name string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVar(p, name, value, usage)
	return p
}

// Uint64类型定义
// 定义一个Uint64类型的FullVar方法，用于设置Uint64类型的变量
func (f *ArgsFlag) Uint64FullVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	f.f.FullVar(newUint64Value(value, p), name, title, key, env, usage)
}

// 定义一个Uint64类型的ConfigVar方法，用于设置Uint64类型的变量
func (f *ArgsFlag) Uint64ConfigVar(p *uint64, name string, title, key string, value uint64, usage string) {
	f.f.FullVar(newUint64Value(value, p), name, title, key, "", usage)
}

// 定义一个Uint64类型的EnvVar方法，用于设置Uint64类型的变量
func (f *ArgsFlag) Uint64EnvVar(p *uint64, name string, env string, value uint64, usage string) {
	f.f.FullVar(newUint64Value(value, p), name, "", "", env, usage)
}

// 定义一个Uint64类型的Var方法，用于设置Uint64类型的变量
func (f *ArgsFlag) Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.f.Var(newUint64Value(value, p), name, usage)
}

// 定义一个Uint64类型的Full方法，用于设置Uint64类型的变量
func (f *ArgsFlag) Uint64Full(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Uint64类型的Config方法，用于设置Uint64类型的变量
func (f *ArgsFlag) Uint64Config(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64ConfigVar(p, name, title, key, value, usage)
	return p
}

// Uint64Env函数用于设置一个uint64类型的环境变量
func (f *ArgsFlag) Uint64Env(name, env string, value uint64, usage string) *uint64 {
	// 创建一个uint64类型的指针
	p := new(uint64)
	// 调用Uint64EnvVar函数设置环境变量
	f.Uint64EnvVar(p, name, env, value, usage)
	// 返回指针
	return p
}

// Uint64函数用于设置一个uint64类型的参数
func (f *ArgsFlag) Uint64(name string, value uint64, usage string) *uint64 {
	// 创建一个uint64类型的指针
	p := new(uint64)
	// 调用Uint64Var函数设置参数
	f.Uint64Var(p, name, value, usage)
	// 返回指针
	return p
}

// Args函数用于获取所有的参数
func (f *ArgsFlag) Args() []string { return f.f.args }

// ArgsLen函数用于获取参数的长度
func (f *ArgsFlag) ArgsLen() int { return len(f.f.args) }

// Usage函数用于打印使用方法
func (f *ArgsFlag) Usage() {
	// 打印程序名称
	Greenf(fmt.Sprintf("%s的使用方法", os.Args[0]), "\n\n")
	// 打印默认值
	f.f.PrintDefaults()
}

// PrintAll函数用于打印所有的参数
func (f *ArgsFlag) PrintAll() { f.f.PrintAll() }

// SetOutput函数用于设置输出
func (f *ArgsFlag) SetOutput(output io.Writer) { f.f.SetOutput(output) }
