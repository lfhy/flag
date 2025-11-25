package flag

import "time"

// Bool类型定义
// 定义一个Bool类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法
func (f *FlagSet) BoolFullHiddenVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个Bool类型的Flag，并设置其名称、标题、键、默认值和用法
func (f *FlagSet) BoolConfigHiddenVar(p *bool, name string, title, key string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个Bool类型的Flag，并设置其名称、环境变量、默认值和用法
func (f *FlagSet) BoolEnvHiddenVar(p *bool, name string, env string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个Bool类型的Flag，并设置其名称、默认值和用法
func (f *FlagSet) BoolHiddenVar(p *bool, name string, value bool, usage string) {
	f.HiddenVar(newBoolValue(value, p), name, usage)
}

// 定义一个Bool类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func (f *FlagSet) BoolFullHidden(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针
func (f *FlagSet) BoolConfigHidden(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针
func (f *FlagSet) BoolEnvHidden(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其名称、默认值和用法，并返回指针
func (f *FlagSet) BoolHidden(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolHiddenVar(p, name, value, usage)
	return p
}

// String类型定义
// 定义一个String类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法
func (f *FlagSet) StringFullHiddenVar(p *string, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个String类型的Flag，并设置其名称、标题、键、默认值和用法
func (f *FlagSet) StringConfigHiddenVar(p *string, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个String类型的Flag，并设置其名称、环境变量、默认值和用法
func (f *FlagSet) StringEnvHiddenVar(p *string, name string, env string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个String类型的Flag，并设置其名称、默认值和用法
func (f *FlagSet) StringHiddenVar(p *string, name string, value string, usage string) {
	f.HiddenVar(newStringValue(value, p), name, usage)
}

// 定义一个String类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func (f *FlagSet) StringFullHidden(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.StringFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针
func (f *FlagSet) StringConfigHidden(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.StringConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针
func (f *FlagSet) StringEnvHidden(name, env string, value string, usage string) *string {
	p := new(string)
	f.StringEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其名称、默认值和用法，并返回指针
func (f *FlagSet) StringHidden(name string, value string, usage string) *string {
	p := new(string)
	f.StringHiddenVar(p, name, value, usage)
	return p
}

// 浮点类型定义
// 定义一个Float64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法
func (f *FlagSet) Float64FullHiddenVar(p *float64, name string, title, key, env string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// 定义一个Float64类型的Flag，并设置其名称、标题、键、默认值和用法
func (f *FlagSet) Float64ConfigHiddenVar(p *float64, name string, title, key string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// 定义一个Float64类型的Flag，并设置其名称、环境变量、默认值和用法
func (f *FlagSet) Float64EnvHiddenVar(p *float64, name string, env string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// 定义一个Float64类型的Flag，并设置其名称、默认值和用法
func (f *FlagSet) Float64HiddenVar(p *float64, name string, value float64, usage string) {
	f.HiddenVar(newFloat64Value(value, p), name, usage)
}

// 定义一个Float64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func (f *FlagSet) Float64FullHidden(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Float64类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针
func (f *FlagSet) Float64ConfigHidden(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Float64类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针
func (f *FlagSet) Float64EnvHidden(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Float64类型的Flag，并设置其名称、默认值和用法，并返回指针
func (f *FlagSet) Float64Hidden(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64HiddenVar(p, name, value, usage)
	return p
}

// 延迟定义
// 定义一个Duration类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法
func (f *FlagSet) DurationFullHiddenVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, title, key, env, usage)
}

// 定义一个Duration类型的Flag，并设置其名称、标题、键、默认值和用法
func (f *FlagSet) DurationConfigHiddenVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, title, key, "", usage)
}

// 定义一个Duration类型的Flag，并设置其名称、环境变量、默认值和用法
func (f *FlagSet) DurationEnvHiddenVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, "", "", env, usage)
}

// 定义一个Duration类型的Flag，并设置其名称、默认值和用法
func (f *FlagSet) DurationHiddenVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.HiddenVar(newDurationValue(value, p), name, usage)
}

// 定义一个Duration类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func (f *FlagSet) DurationFullHidden(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Duration类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针
func (f *FlagSet) DurationConfigHidden(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Duration类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针
func (f *FlagSet) DurationEnvHidden(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Duration类型的Flag，并设置其名称、默认值和用法，并返回指针
func (f *FlagSet) DurationHidden(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationHiddenVar(p, name, value, usage)
	return p
}

// Int类型定义
// 定义一个Int类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法
func (f *FlagSet) IntFullHiddenVar(p *int, name string, title, key, env string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个Int类型的Flag，并设置其名称、标题、键、默认值和用法
func (f *FlagSet) IntConfigHiddenVar(p *int, name string, title, key string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个Int类型的Flag，并设置其名称、环境变量、默认值和用法
func (f *FlagSet) IntEnvHiddenVar(p *int, name string, env string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个Int类型的Flag，并设置其名称、默认值和用法
func (f *FlagSet) IntHiddenVar(p *int, name string, value int, usage string) {
	f.HiddenVar(newIntValue(value, p), name, usage)
}

// 定义一个Int类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func (f *FlagSet) IntFullHidden(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.IntFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针
func (f *FlagSet) IntConfigHidden(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.IntConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针
func (f *FlagSet) IntEnvHidden(name, env string, value int, usage string) *int {
	p := new(int)
	f.IntEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其名称、默认值和用法，并返回指针
func (f *FlagSet) IntHidden(name string, value int, usage string) *int {
	p := new(int)
	f.IntHiddenVar(p, name, value, usage)
	return p
}

// Int64类型定义
// 定义一个Int64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法
func (f *FlagSet) Int64FullHiddenVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个Int64类型的Flag，并设置其名称、标题、键、默认值和用法
func (f *FlagSet) Int64ConfigHiddenVar(p *int64, name string, title, key string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个Int64类型的Flag，并设置其名称、环境变量、默认值和用法
func (f *FlagSet) Int64EnvHiddenVar(p *int64, name string, env string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个Int64类型的Flag，并设置其名称、默认值和用法
func (f *FlagSet) Int64HiddenVar(p *int64, name string, value int64, usage string) {
	f.HiddenVar(newInt64Value(value, p), name, usage)
}

// 定义一个Int64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func (f *FlagSet) Int64FullHidden(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针
func (f *FlagSet) Int64ConfigHidden(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针
func (f *FlagSet) Int64EnvHidden(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其名称、默认值和用法，并返回指针
func (f *FlagSet) Int64Hidden(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64HiddenVar(p, name, value, usage)
	return p
}

// Uint类型定义
// 定义一个Uint类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法
func (f *FlagSet) UintFullHiddenVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个Uint类型的Flag，并设置其名称、标题、键、默认值和用法
func (f *FlagSet) UintConfigHiddenVar(p *uint, name string, title, key string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个Uint类型的Flag，并设置其名称、环境变量、默认值和用法
func (f *FlagSet) UintEnvHiddenVar(p *uint, name string, env string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个Uint类型的Flag，并设置其名称、默认值和用法
func (f *FlagSet) UintHiddenVar(p *uint, name string, value uint, usage string) {
	f.HiddenVar(newUintValue(value, p), name, usage)
}

// 定义一个Uint类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func (f *FlagSet) UintFullHidden(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Uint类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针
func (f *FlagSet) UintConfigHidden(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	f.UintConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// Uint类型定义
// 定义Uint类型的FlagSet，并返回指针
func (f *FlagSet) UintEnvHidden(name, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义Uint类型的FlagSet，并返回指针
func (f *FlagSet) UintHidden(name string, value uint, usage string) *uint {
	p := new(uint)
	f.UintHiddenVar(p, name, value, usage)
	return p
}

// Uint64类型定义
// 定义Uint64类型的FlagSet，并返回指针
func (f *FlagSet) Uint64FullHiddenVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	f.FullHiddenVar(newUint64Value(value, p), name, title, key, env, usage)
}

// 定义Uint64类型的FlagSet，并返回指针
func (f *FlagSet) Uint64ConfigHiddenVar(p *uint64, name string, title, key string, value uint64, usage string) {
	f.FullHiddenVar(newUint64Value(value, p), name, title, key, "", usage)
}

// 定义Uint64类型的FlagSet，并返回指针
func (f *FlagSet) Uint64EnvHiddenVar(p *uint64, name string, env string, value uint64, usage string) {
	f.FullHiddenVar(newUint64Value(value, p), name, "", "", env, usage)
}

// 定义Uint64类型的FlagSet，并返回指针
func (f *FlagSet) Uint64HiddenVar(p *uint64, name string, value uint64, usage string) {
	f.HiddenVar(newUint64Value(value, p), name, usage)
}

// 定义Uint64类型的FlagSet，并返回指针
func (f *FlagSet) Uint64FullHidden(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义Uint64类型的FlagSet，并返回指针
func (f *FlagSet) Uint64ConfigHidden(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义Uint64类型的FlagSet，并返回指针
func (f *FlagSet) Uint64EnvHidden(name, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义Uint64类型的FlagSet，并返回指针
func (f *FlagSet) Uint64Hidden(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64HiddenVar(p, name, value, usage)
	return p
}
