package flag

import "time"

// Bool类型定义
func (f *ArgsFlag) BoolFullHidenVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.f.FullHidenVar(newBoolValue(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) BoolConfigHidenVar(p *bool, name string, title, key string, value bool, usage string) {
	f.f.FullHidenVar(newBoolValue(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) BoolEnvHidenVar(p *bool, name string, env string, value bool, usage string) {
	f.f.FullHidenVar(newBoolValue(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) BoolHidenVar(p *bool, name string, value bool, usage string) {
	f.f.HidenVar(newBoolValue(value, p), name, usage)
}

func (f *ArgsFlag) BoolFullHiden(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.f.BoolFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) BoolConfigHiden(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.f.BoolConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) BoolEnvHiden(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.f.BoolEnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) BoolHiden(name string, value bool, usage string) *bool {
	p := new(bool)
	f.f.BoolHidenVar(p, name, value, usage)
	return p
}

// String类型定义
func (f *ArgsFlag) StringFullHidenVar(p *string, name string, title, key, env string, value string, usage string) {
	f.f.FullHidenVar(newStringValue(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) StringConfigHidenVar(p *string, name string, title, key string, value string, usage string) {
	f.f.FullHidenVar(newStringValue(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) StringEnvHidenVar(p *string, name string, env string, value string, usage string) {
	f.f.FullHidenVar(newStringValue(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) StringHidenVar(p *string, name string, value string, usage string) {
	f.f.HidenVar(newStringValue(value, p), name, usage)
}

func (f *ArgsFlag) StringFullHiden(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.f.StringFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) StringConfigHiden(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.f.StringConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) StringEnvHiden(name, env string, value string, usage string) *string {
	p := new(string)
	f.f.StringEnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) StringHiden(name string, value string, usage string) *string {
	p := new(string)
	f.f.StringHidenVar(p, name, value, usage)
	return p
}

// 浮点类型定义
func (f *ArgsFlag) Float64FullHidenVar(p *float64, name string, title, key, env string, value float64, usage string) {
	f.f.FullHidenVar(newFloat64Value(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) Float64ConfigHidenVar(p *float64, name string, title, key string, value float64, usage string) {
	f.f.FullHidenVar(newFloat64Value(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) Float64EnvHidenVar(p *float64, name string, env string, value float64, usage string) {
	f.f.FullHidenVar(newFloat64Value(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) Float64HidenVar(p *float64, name string, value float64, usage string) {
	f.f.HidenVar(newFloat64Value(value, p), name, usage)
}

func (f *ArgsFlag) Float64FullHiden(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.f.Float64FullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) Float64ConfigHiden(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	f.f.Float64ConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) Float64EnvHiden(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.f.Float64EnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) Float64Hiden(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.f.Float64HidenVar(p, name, value, usage)
	return p
}

// 延迟定义
func (f *ArgsFlag) DurationFullHidenVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	f.f.FullHidenVar(newDurationValue(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) DurationConfigHidenVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	f.f.FullHidenVar(newDurationValue(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) DurationEnvHidenVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	f.f.FullHidenVar(newDurationValue(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) DurationHidenVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.f.HidenVar(newDurationValue(value, p), name, usage)
}

func (f *ArgsFlag) DurationFullHiden(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.f.DurationFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) DurationConfigHiden(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.f.DurationConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) DurationEnvHiden(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.f.DurationEnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) DurationHiden(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.f.DurationHidenVar(p, name, value, usage)
	return p
}

// Int类型定义
func (f *ArgsFlag) IntFullHidenVar(p *int, name string, title, key, env string, value int, usage string) {
	f.f.FullHidenVar(newIntValue(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) IntConfigHidenVar(p *int, name string, title, key string, value int, usage string) {
	f.f.FullHidenVar(newIntValue(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) IntEnvHidenVar(p *int, name string, env string, value int, usage string) {
	f.f.FullHidenVar(newIntValue(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) IntHidenVar(p *int, name string, value int, usage string) {
	f.f.HidenVar(newIntValue(value, p), name, usage)
}

func (f *ArgsFlag) IntFullHiden(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.f.IntFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) IntConfigHiden(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.f.IntConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) IntEnvHiden(name, env string, value int, usage string) *int {
	p := new(int)
	f.f.IntEnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) IntHiden(name string, value int, usage string) *int {
	p := new(int)
	f.f.IntHidenVar(p, name, value, usage)
	return p
}

// Int64类型定义
func (f *ArgsFlag) Int64FullHidenVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.f.FullHidenVar(newInt64Value(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) Int64ConfigHidenVar(p *int64, name string, title, key string, value int64, usage string) {
	f.f.FullHidenVar(newInt64Value(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) Int64EnvHidenVar(p *int64, name string, env string, value int64, usage string) {
	f.f.FullHidenVar(newInt64Value(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) Int64HidenVar(p *int64, name string, value int64, usage string) {
	f.f.HidenVar(newInt64Value(value, p), name, usage)
}

func (f *ArgsFlag) Int64FullHiden(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.f.Int64FullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) Int64ConfigHiden(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.f.Int64ConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) Int64EnvHiden(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.f.Int64EnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) Int64Hiden(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.f.Int64HidenVar(p, name, value, usage)
	return p
}

// Uint类型定义
func (f *ArgsFlag) UintFullHidenVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.f.FullHidenVar(newUintValue(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) UintConfigHidenVar(p *uint, name string, title, key string, value uint, usage string) {
	f.f.FullHidenVar(newUintValue(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) UintEnvHidenVar(p *uint, name string, env string, value uint, usage string) {
	f.f.FullHidenVar(newUintValue(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) UintHidenVar(p *uint, name string, value uint, usage string) {
	f.f.HidenVar(newUintValue(value, p), name, usage)
}

func (f *ArgsFlag) UintFullHiden(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	f.f.UintFullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) UintConfigHiden(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	f.f.UintConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) UintEnvHiden(name, env string, value uint, usage string) *uint {
	p := new(uint)
	f.f.UintEnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) UintHiden(name string, value uint, usage string) *uint {
	p := new(uint)
	f.f.UintHidenVar(p, name, value, usage)
	return p
}

// Uint64类型定义
func (f *ArgsFlag) Uint64FullHidenVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	f.f.FullHidenVar(newUint64Value(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) Uint64ConfigHidenVar(p *uint64, name string, title, key string, value uint64, usage string) {
	f.f.FullHidenVar(newUint64Value(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) Uint64EnvHidenVar(p *uint64, name string, env string, value uint64, usage string) {
	f.f.FullHidenVar(newUint64Value(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) Uint64HidenVar(p *uint64, name string, value uint64, usage string) {
	f.f.HidenVar(newUint64Value(value, p), name, usage)
}

func (f *ArgsFlag) Uint64FullHiden(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.f.Uint64FullHidenVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) Uint64ConfigHiden(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.f.Uint64ConfigHidenVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) Uint64EnvHiden(name, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.f.Uint64EnvHidenVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) Uint64Hiden(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.f.Uint64HidenVar(p, name, value, usage)
	return p
}
