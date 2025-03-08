package flag

import "time"

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
	f.Var(newBoolValue(value, p), name, usage)
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
	f.Var(newStringValue(value, p), name, usage)
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
	f.Var(newFloat64Value(value, p), name, usage)
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
	f.Var(newDurationValue(value, p), name, usage)
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
	f.Var(newIntValue(value, p), name, usage)
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
	f.Var(newInt64Value(value, p), name, usage)
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
	f.Var(newUintValue(value, p), name, usage)
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
	f.Var(newUint64Value(value, p), name, usage)
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
