package flag

import "time"

// Bool类型定义
// 定义一个Bool类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) BoolFullHiddenVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) BoolConfigHiddenVar(p *bool, name string, title, key string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个Bool类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) BoolEnvHiddenVar(p *bool, name string, env string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个Bool类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) BoolHiddenVar(p *bool, name string, value bool, usage string) {
	f.HiddenVar(newBoolValue(value, p), name, usage)
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) BoolFullHidden(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) BoolConfigHidden(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) BoolEnvHidden(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) BoolHidden(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolHiddenVar(p, name, value, usage)
	return p
}

// String类型定义
// 定义一个String类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) StringFullHiddenVar(p *string, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个String类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) StringConfigHiddenVar(p *string, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个String类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) StringEnvHiddenVar(p *string, name string, env string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个String类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) StringHiddenVar(p *string, name string, value string, usage string) {
	f.HiddenVar(newStringValue(value, p), name, usage)
}

// 定义一个String类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) StringFullHidden(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.StringFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) StringConfigHidden(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.StringConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) StringEnvHidden(name, env string, value string, usage string) *string {
	p := new(string)
	f.StringEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) StringHidden(name string, value string, usage string) *string {
	p := new(string)
	f.StringHiddenVar(p, name, value, usage)
	return p
}

// 浮点类型定义
// 定义一个Float64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) Float64FullHiddenVar(p *float64, name string, title, key, env string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) Float64ConfigHiddenVar(p *float64, name string, title, key string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// 定义一个Float64类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) Float64EnvHiddenVar(p *float64, name string, env string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// 定义一个Float64类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) Float64HiddenVar(p *float64, name string, value float64, usage string) {
	f.HiddenVar(newFloat64Value(value, p), name, usage)
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Float64FullHidden(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Float64ConfigHidden(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Float64EnvHidden(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Float64Hidden(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64HiddenVar(p, name, value, usage)
	return p
}

// 延迟定义
// 定义一个time.Duration类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) DurationFullHiddenVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, title, key, env, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) DurationConfigHiddenVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, title, key, "", usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) DurationEnvHiddenVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, "", "", env, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) DurationHiddenVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.HiddenVar(newDurationValue(value, p), name, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) DurationFullHidden(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) DurationConfigHidden(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) DurationEnvHidden(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) DurationHidden(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationHiddenVar(p, name, value, usage)
	return p
}

// Int类型定义
// 定义一个Int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) IntFullHiddenVar(p *int, name string, title, key, env string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个Int类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) IntConfigHiddenVar(p *int, name string, title, key string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个Int类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) IntEnvHiddenVar(p *int, name string, env string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个Int类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) IntHiddenVar(p *int, name string, value int, usage string) {
	f.HiddenVar(newIntValue(value, p), name, usage)
}

// 定义一个Int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) IntFullHidden(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.IntFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) IntConfigHidden(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.IntConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) IntEnvHidden(name, env string, value int, usage string) *int {
	p := new(int)
	f.IntEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) IntHidden(name string, value int, usage string) *int {
	p := new(int)
	f.IntHiddenVar(p, name, value, usage)
	return p
}

// Int64类型定义
// 定义一个Int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) Int64FullHiddenVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) Int64ConfigHiddenVar(p *int64, name string, title, key string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个Int64类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) Int64EnvHiddenVar(p *int64, name string, env string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个Int64类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) Int64HiddenVar(p *int64, name string, value int64, usage string) {
	f.HiddenVar(newInt64Value(value, p), name, usage)
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Int64FullHidden(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Int64ConfigHidden(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Int64EnvHidden(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Int64Hidden(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64HiddenVar(p, name, value, usage)
	return p
}

// Uint类型定义
// 定义一个Uint类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) UintFullHiddenVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个Uint类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) UintConfigHiddenVar(p *uint, name string, title, key string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个Uint类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) UintEnvHiddenVar(p *uint, name string, env string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个Uint类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) UintHiddenVar(p *uint, name string, value uint, usage string) {
	f.HiddenVar(newUintValue(value, p), name, usage)
}

// Uint类型定义
func (f *ArgsFlag) UintFullHidden(name, title, key, env string, value uint, usage string) *uint {
	// 创建一个新的uint类型的指针
	p := new(uint)
	// 调用UintFullHiddenVar方法，将指针、名称、标题、键、环境、值和用法作为参数传入
	f.UintFullHiddenVar(p, name, title, key, env, value, usage)
	// 返回指针
	return p
}

func (f *ArgsFlag) UintConfigHidden(name, title, key string, value uint, usage string) *uint {
	// 创建一个新的uint类型的指针
	p := new(uint)
	// 调用UintConfigHiddenVar方法，将指针、名称、标题、键、值和用法作为参数传入
	f.UintConfigHiddenVar(p, name, title, key, value, usage)
	// 返回指针
	return p
}

func (f *ArgsFlag) UintEnvHidden(name, env string, value uint, usage string) *uint {
	// 创建一个新的uint类型的指针
	p := new(uint)
	// 调用UintEnvHiddenVar方法，将指针、名称、环境、值和用法作为参数传入
	f.UintEnvHiddenVar(p, name, env, value, usage)
	// 返回指针
	return p
}

func (f *ArgsFlag) UintHidden(name string, value uint, usage string) *uint {
	// 创建一个新的uint类型的指针
	p := new(uint)
	// 调用UintHiddenVar方法，将指针、名称和用法作为参数传入
	f.UintHiddenVar(p, name, value, usage)
	// 返回指针
	return p
}

// Uint64类型定义
func (f *ArgsFlag) Uint64FullHiddenVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	// 调用FullHiddenVar方法，将Uint64类型的值和指针作为参数传入
	f.FullHiddenVar(newUint64Value(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) Uint64ConfigHiddenVar(p *uint64, name string, title, key string, value uint64, usage string) {
	// 调用FullHiddenVar方法，将Uint64类型的值和指针作为参数传入
	f.FullHiddenVar(newUint64Value(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) Uint64EnvHiddenVar(p *uint64, name string, env string, value uint64, usage string) {
	// 调用FullHiddenVar方法，将Uint64类型的值和指针作为参数传入
	f.FullHiddenVar(newUint64Value(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) Uint64HiddenVar(p *uint64, name string, value uint64, usage string) {
	// 调用HiddenVar方法，将Uint64类型的值和指针作为参数传入
	f.HiddenVar(newUint64Value(value, p), name, usage)
}

func (f *ArgsFlag) Uint64FullHidden(name, title, key, env string, value uint64, usage string) *uint64 {
	// 创建一个新的uint64类型的指针
	p := new(uint64)
	// 调用Uint64FullHiddenVar方法，将指针、名称、标题、键、环境、值和用法作为参数传入
	f.Uint64FullHiddenVar(p, name, title, key, env, value, usage)
	// 返回指针
	return p
}

func (f *ArgsFlag) Uint64ConfigHidden(name, title, key string, value uint64, usage string) *uint64 {
	// 创建一个新的uint64类型的指针
	p := new(uint64)
	// 调用Uint64ConfigHiddenVar方法，将指针、名称、标题、键、值和用法作为参数传入
	f.Uint64ConfigHiddenVar(p, name, title, key, value, usage)
	// 返回指针
	return p
}

// Uint64EnvHidden函数用于设置一个uint64类型的参数，该参数可以通过环境变量设置，并且在帮助信息中隐藏
func (f *ArgsFlag) Uint64EnvHidden(name, env string, value uint64, usage string) *uint64 {
	// 创建一个新的uint64类型的指针
	p := new(uint64)
	// 调用Uint64EnvHiddenVar方法，将指针、名称、环境、值和用法作为参数传入
	f.Uint64EnvHiddenVar(p, name, env, value, usage)
	// 返回指针
	return p
}

// Uint64Hidden函数用于设置一个uint64类型的隐藏参数
func (f *ArgsFlag) Uint64Hidden(name string, value uint64, usage string) *uint64 {
	// 创建一个新的uint64类型的指针
	p := new(uint64)
	// 调用Uint64HiddenVar方法，将指针、名称和用法作为参数传入
	f.Uint64HiddenVar(p, name, value, usage)
	// 返回指针
	return p
}
