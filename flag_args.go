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

func NewArgsFlag() *ArgsFlag {
	return &ArgsFlag{NewFlagSet(os.Args[0], ExitOnError)}
}

func (f *ArgsFlag) Parse() {
	f.f.Parse(os.Args[1:])
}

// Bool类型定义
func (f *ArgsFlag) BoolFullVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.f.FullVar(newBoolValue(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) BoolConfigVar(p *bool, name string, title, key string, value bool, usage string) {
	f.f.FullVar(newBoolValue(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) BoolEnvVar(p *bool, name string, env string, value bool, usage string) {
	f.f.FullVar(newBoolValue(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) BoolVar(p *bool, name string, value bool, usage string) {
	f.f.Var(newBoolValue(value, p), name, usage)
}

func (f *ArgsFlag) BoolFull(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolFullVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) BoolConfig(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolConfigVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) BoolEnv(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolEnvVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVar(p, name, value, usage)
	return p
}

// String类型定义
func (f *ArgsFlag) StringFullVar(p *string, name string, title, key, env string, value string, usage string) {
	f.f.FullVar(newStringValue(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) StringConfigVar(p *string, name string, title, key string, value string, usage string) {
	f.f.FullVar(newStringValue(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) StringEnvVar(p *string, name string, env string, value string, usage string) {
	f.f.FullVar(newStringValue(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) StringVar(p *string, name string, value string, usage string) {
	f.f.Var(newStringValue(value, p), name, usage)
}

func (f *ArgsFlag) StringFull(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.StringFullVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) StringConfig(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.StringConfigVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) StringEnv(name, env string, value string, usage string) *string {
	p := new(string)
	f.StringEnvVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) String(name string, value string, usage string) *string {
	p := new(string)
	f.StringVar(p, name, value, usage)
	return p
}

// 浮点类型定义
func (f *ArgsFlag) Float64FullVar(p *float64, name string, title, key, env string, value float64, usage string) {
	f.f.FullVar(newFloat64Value(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) Float64ConfigVar(p *float64, name string, title, key string, value float64, usage string) {
	f.f.FullVar(newFloat64Value(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) Float64EnvVar(p *float64, name string, env string, value float64, usage string) {
	f.f.FullVar(newFloat64Value(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) Float64Var(p *float64, name string, value float64, usage string) {
	f.f.Var(newFloat64Value(value, p), name, usage)
}

func (f *ArgsFlag) Float64Full(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64FullVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) Float64Config(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64ConfigVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) Float64Env(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64EnvVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) Float64(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64Var(p, name, value, usage)
	return p
}

// 延迟定义
func (f *ArgsFlag) DurationFullVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	f.f.FullVar(newDurationValue(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) DurationConfigVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	f.f.FullVar(newDurationValue(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) DurationEnvVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	f.f.FullVar(newDurationValue(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.f.Var(newDurationValue(value, p), name, usage)
}

func (f *ArgsFlag) DurationFull(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationFullVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) DurationConfig(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationConfigVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) DurationEnv(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationEnvVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, value, usage)
	return p
}

// Int类型定义
func (f *ArgsFlag) IntFullVar(p *int, name string, title, key, env string, value int, usage string) {
	f.f.FullVar(newIntValue(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) IntConfigVar(p *int, name string, title, key string, value int, usage string) {
	f.f.FullVar(newIntValue(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) IntEnvVar(p *int, name string, env string, value int, usage string) {
	f.f.FullVar(newIntValue(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) IntVar(p *int, name string, value int, usage string) {
	f.f.Var(newIntValue(value, p), name, usage)
}

func (f *ArgsFlag) IntFull(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.IntFullVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) IntConfig(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.IntConfigVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) IntEnv(name, env string, value int, usage string) *int {
	p := new(int)
	f.IntEnvVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) Int(name string, value int, usage string) *int {
	p := new(int)
	f.IntVar(p, name, value, usage)
	return p
}

// Int64类型定义
func (f *ArgsFlag) Int64FullVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.f.FullVar(newInt64Value(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) Int64ConfigVar(p *int64, name string, title, key string, value int64, usage string) {
	f.f.FullVar(newInt64Value(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) Int64EnvVar(p *int64, name string, env string, value int64, usage string) {
	f.f.FullVar(newInt64Value(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) Int64Var(p *int64, name string, value int64, usage string) {
	f.f.Var(newInt64Value(value, p), name, usage)
}

func (f *ArgsFlag) Int64Full(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) Int64Config(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) Int64Env(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64Var(p, name, value, usage)
	return p
}

// Uint类型定义
func (f *ArgsFlag) UintFullVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.f.FullVar(newUintValue(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) UintConfigVar(p *uint, name string, title, key string, value uint, usage string) {
	f.f.FullVar(newUintValue(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) UintEnvVar(p *uint, name string, env string, value uint, usage string) {
	f.f.FullVar(newUintValue(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) UintVar(p *uint, name string, value uint, usage string) {
	f.f.Var(newUintValue(value, p), name, usage)
}

func (f *ArgsFlag) UintFull(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintFullVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) UintConfig(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	f.UintConfigVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) UintEnv(name, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintEnvVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) Uint(name string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVar(p, name, value, usage)
	return p
}

// Uint64类型定义
func (f *ArgsFlag) Uint64FullVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	f.f.FullVar(newUint64Value(value, p), name, title, key, env, usage)
}

func (f *ArgsFlag) Uint64ConfigVar(p *uint64, name string, title, key string, value uint64, usage string) {
	f.f.FullVar(newUint64Value(value, p), name, title, key, "", usage)
}

func (f *ArgsFlag) Uint64EnvVar(p *uint64, name string, env string, value uint64, usage string) {
	f.f.FullVar(newUint64Value(value, p), name, "", "", env, usage)
}

func (f *ArgsFlag) Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.f.Var(newUint64Value(value, p), name, usage)
}

func (f *ArgsFlag) Uint64Full(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64FullVar(p, name, title, key, env, value, usage)
	return p
}

func (f *ArgsFlag) Uint64Config(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64ConfigVar(p, name, title, key, value, usage)
	return p
}

func (f *ArgsFlag) Uint64Env(name, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64EnvVar(p, name, env, value, usage)
	return p
}

func (f *ArgsFlag) Uint64(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64Var(p, name, value, usage)
	return p
}

func (f *ArgsFlag) Args() []string { return f.f.args }
func (f *ArgsFlag) ArgsLen() int   { return len(f.f.args) }

func (f *ArgsFlag) Usage() {
	Greenf(fmt.Sprintf("%s的使用方法", os.Args[0]), "\n\n")
	f.f.PrintDefaults()
}

func (f *ArgsFlag) PrintAll() { f.f.PrintAll() }

func (f *ArgsFlag) SetOutput(output io.Writer) { f.f.SetOutput(output) }
