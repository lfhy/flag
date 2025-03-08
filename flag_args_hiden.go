package flag

import "time"

// Bool类型定义
// 定义一个Bool类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) BoolFullHidenVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.FullHidenVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) BoolConfigHidenVar(p *bool, name string, title, key string, value bool, usage string) {
	f.FullHidenVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个Bool类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) BoolEnvHidenVar(p *bool, name string, env string, value bool, usage string) {
	f.FullHidenVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个Bool类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) BoolHidenVar(p *bool, name string, value bool, usage string) {
	f.HidenVar(newBoolValue(value, p), name, usage)
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) BoolFullHiden(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) BoolConfigHiden(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolConfigHidenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) BoolEnvHiden(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolEnvHidenVar(p, name, env, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) BoolHiden(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolHidenVar(p, name, value, usage)
	return p
}

// String类型定义
// 定义一个String类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) StringFullHidenVar(p *string, name string, title, key, env string, value string, usage string) {
	f.FullHidenVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个String类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) StringConfigHidenVar(p *string, name string, title, key string, value string, usage string) {
	f.FullHidenVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个String类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) StringEnvHidenVar(p *string, name string, env string, value string, usage string) {
	f.FullHidenVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个String类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) StringHidenVar(p *string, name string, value string, usage string) {
	f.HidenVar(newStringValue(value, p), name, usage)
}

// 定义一个String类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) StringFullHiden(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.StringFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) StringConfigHiden(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.StringConfigHidenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) StringEnvHiden(name, env string, value string, usage string) *string {
	p := new(string)
	f.StringEnvHidenVar(p, name, env, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) StringHiden(name string, value string, usage string) *string {
	p := new(string)
	f.StringHidenVar(p, name, value, usage)
	return p
}

// 浮点类型定义
// 定义一个Float64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) Float64FullHidenVar(p *float64, name string, title, key, env string, value float64, usage string) {
	f.FullHidenVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) Float64ConfigHidenVar(p *float64, name string, title, key string, value float64, usage string) {
	f.FullHidenVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// 定义一个Float64类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) Float64EnvHidenVar(p *float64, name string, env string, value float64, usage string) {
	f.FullHidenVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// 定义一个Float64类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) Float64HidenVar(p *float64, name string, value float64, usage string) {
	f.HidenVar(newFloat64Value(value, p), name, usage)
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Float64FullHiden(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64FullHidenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Float64ConfigHiden(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64ConfigHidenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Float64EnvHiden(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64EnvHidenVar(p, name, env, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Float64Hiden(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64HidenVar(p, name, value, usage)
	return p
}

// 延迟定义
// 定义一个time.Duration类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) DurationFullHidenVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	f.FullHidenVar(newDurationValue(value, p), name, title, key, env, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) DurationConfigHidenVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	f.FullHidenVar(newDurationValue(value, p), name, title, key, "", usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) DurationEnvHidenVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	f.FullHidenVar(newDurationValue(value, p), name, "", "", env, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) DurationHidenVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.HidenVar(newDurationValue(value, p), name, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) DurationFullHiden(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) DurationConfigHiden(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationConfigHidenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) DurationEnvHiden(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationEnvHidenVar(p, name, env, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) DurationHiden(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationHidenVar(p, name, value, usage)
	return p
}

// Int类型定义
// 定义一个Int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) IntFullHidenVar(p *int, name string, title, key, env string, value int, usage string) {
	f.FullHidenVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个Int类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) IntConfigHidenVar(p *int, name string, title, key string, value int, usage string) {
	f.FullHidenVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个Int类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) IntEnvHidenVar(p *int, name string, env string, value int, usage string) {
	f.FullHidenVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个Int类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) IntHidenVar(p *int, name string, value int, usage string) {
	f.HidenVar(newIntValue(value, p), name, usage)
}

// 定义一个Int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) IntFullHiden(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.IntFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) IntConfigHiden(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.IntConfigHidenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) IntEnvHiden(name, env string, value int, usage string) *int {
	p := new(int)
	f.IntEnvHidenVar(p, name, env, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) IntHiden(name string, value int, usage string) *int {
	p := new(int)
	f.IntHidenVar(p, name, value, usage)
	return p
}

// Int64类型定义
// 定义一个Int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) Int64FullHidenVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.FullHidenVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) Int64ConfigHidenVar(p *int64, name string, title, key string, value int64, usage string) {
	f.FullHidenVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个Int64类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) Int64EnvHidenVar(p *int64, name string, env string, value int64, usage string) {
	f.FullHidenVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个Int64类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) Int64HidenVar(p *int64, name string, value int64, usage string) {
	f.HidenVar(newInt64Value(value, p), name, usage)
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Int64FullHiden(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullHidenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Int64ConfigHiden(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigHidenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Int64EnvHiden(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvHidenVar(p, name, env, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针
func (f *ArgsFlag) Int64Hiden(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64HidenVar(p, name, value, usage)
	return p
}

// Uint类型定义
// 定义一个Uint类型的变量，并设置其名称、标题、键、环境变量、默认值和用法
func (f *ArgsFlag) UintFullHidenVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.FullHidenVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个Uint类型的变量，并设置其名称、标题、键、默认值和用法
func (f *ArgsFlag) UintConfigHidenVar(p *uint, name string, title, key string, value uint, usage string) {
	f.FullHidenVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个Uint类型的变量，并设置其名称、环境变量、默认值和用法
func (f *ArgsFlag) UintEnvHidenVar(p *uint, name string, env string, value uint, usage string) {
	f.FullHidenVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个Uint类型的变量，并设置其名称、默认值和用法
func (f *ArgsFlag) UintHidenVar(p *uint, name string, value uint, usage string) {
	f.HidenVar(newUintValue(value, p), name, usage)
}

// Uint类型定义
func (f *ArgsFlag) UintFullHiden(name, title, key, env string, value uint, usage string) *uint {
	// 创建一个新的uint类型的指针
	p := new(uint)
	// 调用UintFullHidenVar方法，将指针、名称、标题、键、环境、值和用法作为参数传入
	f.UintFullHidenVar(p, name, title, key, env, value, usage)
	// 返回指针
	return p
}

func (f *ArgsFlag) UintConfigHiden(name, title, key string, value uint, usage string) *uint {
	// 创建一个新的uint类型的指针
	p := new(uint)
	// 调用UintConfigHidenVar方法，将指针、名称、标题、键、值和用法作为参数传入
	f.UintConfigHidenVar(p, name, title, key, value, usage)
	// 返回指针
	return p
}

func (f *ArgsFlag) UintEnvHiden(name, env string, value uint, usage string) *uint {
	// 创建一个新的uint类型的指针
	p := new(uint)
	// 调用UintEnvHidenVar方法，将指针、名称、环境、值和用法作为参数传入
	f.UintEnvHidenVar(p, name, env, value, usage)
	// 返回指针
	return p
}

func (f *ArgsFlag) UintHiden(name string, value uint, usage string) *uint {
	// 创建一个新的uint类型的指针
	p := new(uint)
	// 调用UintHidenVar方法，将指针、名称和用法作为参数传入
	f.UintHidenVar(p, name, value, usage)
	// 返回指针
	return p
}

// Uint64类型定义
func (f *ArgsFlag) Uint64FullHidenVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	// 调用FullHidenVar方法，将Uint64类型的值和指针作为参数传入
	f.FullHidenVar(newUint64Value(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) Uint64ConfigHidenVar(p *uint64, name string, title, key string, value uint64, usage string) {
	// 调用FullHidenVar方法，将Uint64类型的值和指针作为参数传入
	f.FullHidenVar(newUint64Value(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) Uint64EnvHidenVar(p *uint64, name string, env string, value uint64, usage string) {
	// 调用FullHidenVar方法，将Uint64类型的值和指针作为参数传入
	f.FullHidenVar(newUint64Value(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) Uint64HidenVar(p *uint64, name string, value uint64, usage string) {
	// 调用HidenVar方法，将Uint64类型的值和指针作为参数传入
	f.HidenVar(newUint64Value(value, p), name, usage)
}

func (f *ArgsFlag) Uint64FullHiden(name, title, key, env string, value uint64, usage string) *uint64 {
	// 创建一个新的uint64类型的指针
	p := new(uint64)
	// 调用Uint64FullHidenVar方法，将指针、名称、标题、键、环境、值和用法作为参数传入
	f.Uint64FullHidenVar(p, name, title, key, env, value, usage)
	// 返回指针
	return p
}

func (f *ArgsFlag) Uint64ConfigHiden(name, title, key string, value uint64, usage string) *uint64 {
	// 创建一个新的uint64类型的指针
	p := new(uint64)
	// 调用Uint64ConfigHidenVar方法，将指针、名称、标题、键、值和用法作为参数传入
	f.Uint64ConfigHidenVar(p, name, title, key, value, usage)
	// 返回指针
	return p
}

// Uint64EnvHiden函数用于设置一个uint64类型的参数，该参数可以通过环境变量设置，并且在帮助信息中隐藏
func (f *ArgsFlag) Uint64EnvHiden(name, env string, value uint64, usage string) *uint64 {
	// 创建一个新的uint64类型的指针
	p := new(uint64)
	// 调用Uint64EnvHidenVar方法，将指针、名称、环境、值和用法作为参数传入
	f.Uint64EnvHidenVar(p, name, env, value, usage)
	// 返回指针
	return p
}

// Uint64Hiden函数用于设置一个uint64类型的隐藏参数
func (f *ArgsFlag) Uint64Hiden(name string, value uint64, usage string) *uint64 {
	// 创建一个新的uint64类型的指针
	p := new(uint64)
	// 调用Uint64HidenVar方法，将指针、名称和用法作为参数传入
	f.Uint64HidenVar(p, name, value, usage)
	// 返回指针
	return p
}
