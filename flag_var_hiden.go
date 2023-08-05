package flag

import "time"

// Bool类型定义
func (f *FlagSet) BoolFullHidenVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.FullHidenVar(newBoolValue(value, p), name, title, key, env, usage)
}

func (f *FlagSet) BoolConfigHidenVar(p *bool, name string, title, key string, value bool, usage string) {
	f.FullHidenVar(newBoolValue(value, p), name, title, key, "", usage)
}

func (f *FlagSet) BoolEnvHidenVar(p *bool, name string, env string, value bool, usage string) {
	f.FullHidenVar(newBoolValue(value, p), name, "", "", env, usage)
}

func (f *FlagSet) BoolHidenVar(p *bool, name string, value bool, usage string) {
	f.HidenVar(newBoolValue(value, p), name, usage)
}

func (f *FlagSet) BoolFullHiden(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *FlagSet) BoolConfigHiden(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *FlagSet) BoolEnvHiden(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolEnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *FlagSet) BoolHiden(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolHidenVar(p, name, value, usage)
	return p
}

// String类型定义
func (f *FlagSet) StringFullHidenVar(p *string, name string, title, key, env string, value string, usage string) {
	f.FullHidenVar(newStringValue(value, p), name, title, key, env, usage)
}

func (f *FlagSet) StringConfigHidenVar(p *string, name string, title, key string, value string, usage string) {
	f.FullHidenVar(newStringValue(value, p), name, title, key, "", usage)
}

func (f *FlagSet) StringEnvHidenVar(p *string, name string, env string, value string, usage string) {
	f.FullHidenVar(newStringValue(value, p), name, "", "", env, usage)
}

func (f *FlagSet) StringHidenVar(p *string, name string, value string, usage string) {
	f.HidenVar(newStringValue(value, p), name, usage)
}

func (f *FlagSet) StringFullHiden(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.StringFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *FlagSet) StringConfigHiden(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.StringConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *FlagSet) StringEnvHiden(name, env string, value string, usage string) *string {
	p := new(string)
	f.StringEnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *FlagSet) StringHiden(name string, value string, usage string) *string {
	p := new(string)
	f.StringHidenVar(p, name, value, usage)
	return p
}

// 浮点类型定义
func (f *FlagSet) Float64FullHidenVar(p *float64, name string, title, key, env string, value float64, usage string) {
	f.FullHidenVar(newFloat64Value(value, p), name, title, key, env, usage)
}

func (f *FlagSet) Float64ConfigHidenVar(p *float64, name string, title, key string, value float64, usage string) {
	f.FullHidenVar(newFloat64Value(value, p), name, title, key, "", usage)
}

func (f *FlagSet) Float64EnvHidenVar(p *float64, name string, env string, value float64, usage string) {
	f.FullHidenVar(newFloat64Value(value, p), name, "", "", env, usage)
}

func (f *FlagSet) Float64HidenVar(p *float64, name string, value float64, usage string) {
	f.HidenVar(newFloat64Value(value, p), name, usage)
}

func (f *FlagSet) Float64FullHiden(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64FullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *FlagSet) Float64ConfigHiden(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64ConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *FlagSet) Float64EnvHiden(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64EnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *FlagSet) Float64Hiden(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64HidenVar(p, name, value, usage)
	return p
}

// 延迟定义
func (f *FlagSet) DurationFullHidenVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	f.FullHidenVar(newDurationValue(value, p), name, title, key, env, usage)
}

func (f *FlagSet) DurationConfigHidenVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	f.FullHidenVar(newDurationValue(value, p), name, title, key, "", usage)
}

func (f *FlagSet) DurationEnvHidenVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	f.FullHidenVar(newDurationValue(value, p), name, "", "", env, usage)
}

func (f *FlagSet) DurationHidenVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.HidenVar(newDurationValue(value, p), name, usage)
}

func (f *FlagSet) DurationFullHiden(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *FlagSet) DurationConfigHiden(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *FlagSet) DurationEnvHiden(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationEnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *FlagSet) DurationHiden(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationHidenVar(p, name, value, usage)
	return p
}

// Int类型定义
func (f *FlagSet) IntFullHidenVar(p *int, name string, title, key, env string, value int, usage string) {
	f.FullHidenVar(newIntValue(value, p), name, title, key, env, usage)
}

func (f *FlagSet) IntConfigHidenVar(p *int, name string, title, key string, value int, usage string) {
	f.FullHidenVar(newIntValue(value, p), name, title, key, "", usage)
}

func (f *FlagSet) IntEnvHidenVar(p *int, name string, env string, value int, usage string) {
	f.FullHidenVar(newIntValue(value, p), name, "", "", env, usage)
}

func (f *FlagSet) IntHidenVar(p *int, name string, value int, usage string) {
	f.HidenVar(newIntValue(value, p), name, usage)
}

func (f *FlagSet) IntFullHiden(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.IntFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *FlagSet) IntConfigHiden(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.IntConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *FlagSet) IntEnvHiden(name, env string, value int, usage string) *int {
	p := new(int)
	f.IntEnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *FlagSet) IntHiden(name string, value int, usage string) *int {
	p := new(int)
	f.IntHidenVar(p, name, value, usage)
	return p
}

// Int64类型定义
func (f *FlagSet) Int64FullHidenVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.FullHidenVar(newInt64Value(value, p), name, title, key, env, usage)
}

func (f *FlagSet) Int64ConfigHidenVar(p *int64, name string, title, key string, value int64, usage string) {
	f.FullHidenVar(newInt64Value(value, p), name, title, key, "", usage)
}

func (f *FlagSet) Int64EnvHidenVar(p *int64, name string, env string, value int64, usage string) {
	f.FullHidenVar(newInt64Value(value, p), name, "", "", env, usage)
}

func (f *FlagSet) Int64HidenVar(p *int64, name string, value int64, usage string) {
	f.HidenVar(newInt64Value(value, p), name, usage)
}

func (f *FlagSet) Int64FullHiden(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *FlagSet) Int64ConfigHiden(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *FlagSet) Int64EnvHiden(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *FlagSet) Int64Hiden(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64HidenVar(p, name, value, usage)
	return p
}

// Uint类型定义
func (f *FlagSet) UintFullHidenVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.FullHidenVar(newUintValue(value, p), name, title, key, env, usage)
}

func (f *FlagSet) UintConfigHidenVar(p *uint, name string, title, key string, value uint, usage string) {
	f.FullHidenVar(newUintValue(value, p), name, title, key, "", usage)
}

func (f *FlagSet) UintEnvHidenVar(p *uint, name string, env string, value uint, usage string) {
	f.FullHidenVar(newUintValue(value, p), name, "", "", env, usage)
}

func (f *FlagSet) UintHidenVar(p *uint, name string, value uint, usage string) {
	f.HidenVar(newUintValue(value, p), name, usage)
}

func (f *FlagSet) UintFullHiden(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *FlagSet) UintConfigHiden(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	f.UintConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *FlagSet) UintEnvHiden(name, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintEnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *FlagSet) UintHiden(name string, value uint, usage string) *uint {
	p := new(uint)
	f.UintHidenVar(p, name, value, usage)
	return p
}

// Uint64类型定义
func (f *FlagSet) Uint64FullHidenVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	f.FullHidenVar(newUint64Value(value, p), name, title, key, env, usage)
}

func (f *FlagSet) Uint64ConfigHidenVar(p *uint64, name string, title, key string, value uint64, usage string) {
	f.FullHidenVar(newUint64Value(value, p), name, title, key, "", usage)
}

func (f *FlagSet) Uint64EnvHidenVar(p *uint64, name string, env string, value uint64, usage string) {
	f.FullHidenVar(newUint64Value(value, p), name, "", "", env, usage)
}

func (f *FlagSet) Uint64HidenVar(p *uint64, name string, value uint64, usage string) {
	f.HidenVar(newUint64Value(value, p), name, usage)
}

func (f *FlagSet) Uint64FullHiden(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64FullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *FlagSet) Uint64ConfigHiden(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64ConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *FlagSet) Uint64EnvHiden(name, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64EnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *FlagSet) Uint64Hiden(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64HidenVar(p, name, value, usage)
	return p
}
