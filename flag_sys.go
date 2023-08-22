package flag

import (
	"fmt"
	"io"
	"os"
	"time"
)

var sysflag = ArgsFlag{NewFlagSet(os.Args[0], ExitOnError)}

func Parse() {
	sysflag.f.Parse(os.Args[1:])
}

// Bool类型定义
func BoolFullVar(p *bool, name string, title, key, env string, value bool, usage string) {
	sysflag.f.FullVar(newBoolValue(value, p), name, title, key, env, usage)
}

func BoolConfigVar(p *bool, name string, title, key string, value bool, usage string) {
	sysflag.f.FullVar(newBoolValue(value, p), name, title, key, "", usage)
}

func BoolEnvVar(p *bool, name string, env string, value bool, usage string) {
	sysflag.f.FullVar(newBoolValue(value, p), name, "", "", env, usage)
}

func BoolVar(p *bool, name string, value bool, usage string) {
	sysflag.f.Var(newBoolValue(value, p), name, usage)
}

func BoolFull(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolFullVar(p, name, title, key, env, value, usage)
	return p
}

func BoolConfig(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolConfigVar(p, name, title, key, value, usage)
	return p
}

func BoolEnv(name, env string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolEnvVar(p, name, env, value, usage)
	return p
}

func Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolVar(p, name, value, usage)
	return p
}

// String类型定义
func StringFullVar(p *string, name string, title, key, env string, value string, usage string) {
	sysflag.f.FullVar(newStringValue(value, p), name, title, key, env, usage)
}

func StringConfigVar(p *string, name string, title, key string, value string, usage string) {
	sysflag.f.FullVar(newStringValue(value, p), name, title, key, "", usage)
}

func StringEnvVar(p *string, name string, env string, value string, usage string) {
	sysflag.f.FullVar(newStringValue(value, p), name, "", "", env, usage)
}

func StringVar(p *string, name string, value string, usage string) {
	sysflag.f.Var(newStringValue(value, p), name, usage)
}

func StringFull(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	sysflag.StringFullVar(p, name, title, key, env, value, usage)
	return p
}

func StringConfig(name, title, key string, value string, usage string) *string {
	p := new(string)
	sysflag.StringConfigVar(p, name, title, key, value, usage)
	return p
}

func StringEnv(name, env string, value string, usage string) *string {
	p := new(string)
	sysflag.StringEnvVar(p, name, env, value, usage)
	return p
}

func String(name string, value string, usage string) *string {
	p := new(string)
	sysflag.StringVar(p, name, value, usage)
	return p
}

// 浮点类型定义
func Float64FullVar(p *float64, name string, title, key, env string, value float64, usage string) {
	sysflag.f.FullVar(newFloat64Value(value, p), name, title, key, env, usage)
}

func Float64ConfigVar(p *float64, name string, title, key string, value float64, usage string) {
	sysflag.f.FullVar(newFloat64Value(value, p), name, title, key, "", usage)
}

func Float64EnvVar(p *float64, name string, env string, value float64, usage string) {
	sysflag.f.FullVar(newFloat64Value(value, p), name, "", "", env, usage)
}

func Float64Var(p *float64, name string, value float64, usage string) {
	sysflag.f.Var(newFloat64Value(value, p), name, usage)
}

func Float64Full(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64FullVar(p, name, title, key, env, value, usage)
	return p
}

func Float64Config(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64ConfigVar(p, name, title, key, value, usage)
	return p
}

func Float64Env(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64EnvVar(p, name, env, value, usage)
	return p
}

func Float64(name string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64Var(p, name, value, usage)
	return p
}

// 延迟定义
func DurationFullVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	sysflag.f.FullVar(newDurationValue(value, p), name, title, key, env, usage)
}

func DurationConfigVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	sysflag.f.FullVar(newDurationValue(value, p), name, title, key, "", usage)
}

func DurationEnvVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	sysflag.f.FullVar(newDurationValue(value, p), name, "", "", env, usage)
}

func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	sysflag.f.Var(newDurationValue(value, p), name, usage)
}

func DurationFull(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationFullVar(p, name, title, key, env, value, usage)
	return p
}

func DurationConfig(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationConfigVar(p, name, title, key, value, usage)
	return p
}

func DurationEnv(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationEnvVar(p, name, env, value, usage)
	return p
}

func Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationVar(p, name, value, usage)
	return p
}

// Int类型定义
func IntFullVar(p *int, name string, title, key, env string, value int, usage string) {
	sysflag.f.FullVar(newIntValue(value, p), name, title, key, env, usage)
}

func IntConfigVar(p *int, name string, title, key string, value int, usage string) {
	sysflag.f.FullVar(newIntValue(value, p), name, title, key, "", usage)
}

func IntEnvVar(p *int, name string, env string, value int, usage string) {
	sysflag.f.FullVar(newIntValue(value, p), name, "", "", env, usage)
}

func IntVar(p *int, name string, value int, usage string) {
	sysflag.f.Var(newIntValue(value, p), name, usage)
}

func IntFull(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	sysflag.IntFullVar(p, name, title, key, env, value, usage)
	return p
}

func IntConfig(name, title, key string, value int, usage string) *int {
	p := new(int)
	sysflag.IntConfigVar(p, name, title, key, value, usage)
	return p
}

func IntEnv(name, env string, value int, usage string) *int {
	p := new(int)
	sysflag.IntEnvVar(p, name, env, value, usage)
	return p
}

func Int(name string, value int, usage string) *int {
	p := new(int)
	sysflag.IntVar(p, name, value, usage)
	return p
}

// Int64类型定义
func Int64FullVar(p *int64, name string, title, key, env string, value int64, usage string) {
	sysflag.f.FullVar(newInt64Value(value, p), name, title, key, env, usage)
}

func Int64ConfigVar(p *int64, name string, title, key string, value int64, usage string) {
	sysflag.f.FullVar(newInt64Value(value, p), name, title, key, "", usage)
}

func Int64EnvVar(p *int64, name string, env string, value int64, usage string) {
	sysflag.f.FullVar(newInt64Value(value, p), name, "", "", env, usage)
}

func Int64Var(p *int64, name string, value int64, usage string) {
	sysflag.f.Var(newInt64Value(value, p), name, usage)
}

func Int64Full(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64FullVar(p, name, title, key, env, value, usage)
	return p
}

func Int64Config(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64ConfigVar(p, name, title, key, value, usage)
	return p
}

func Int64Env(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64EnvVar(p, name, env, value, usage)
	return p
}

func Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64Var(p, name, value, usage)
	return p
}

// Uint类型定义
func UintFullVar(p *uint, name string, title, key, env string, value uint, usage string) {
	sysflag.f.FullVar(newUintValue(value, p), name, title, key, env, usage)
}

func UintConfigVar(p *uint, name string, title, key string, value uint, usage string) {
	sysflag.f.FullVar(newUintValue(value, p), name, title, key, "", usage)
}

func UintEnvVar(p *uint, name string, env string, value uint, usage string) {
	sysflag.f.FullVar(newUintValue(value, p), name, "", "", env, usage)
}

func UintVar(p *uint, name string, value uint, usage string) {
	sysflag.f.Var(newUintValue(value, p), name, usage)
}

func UintFull(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintFullVar(p, name, title, key, env, value, usage)
	return p
}

func UintConfig(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintConfigVar(p, name, title, key, value, usage)
	return p
}

func UintEnv(name, env string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintEnvVar(p, name, env, value, usage)
	return p
}

func Uint(name string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintVar(p, name, value, usage)
	return p
}

// Uint64类型定义
func Uint64FullVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	sysflag.f.FullVar(newUint64Value(value, p), name, title, key, env, usage)
}

func Uint64ConfigVar(p *uint64, name string, title, key string, value uint64, usage string) {
	sysflag.f.FullVar(newUint64Value(value, p), name, title, key, "", usage)
}

func Uint64EnvVar(p *uint64, name string, env string, value uint64, usage string) {
	sysflag.f.FullVar(newUint64Value(value, p), name, "", "", env, usage)
}

func Uint64Var(p *uint64, name string, value uint64, usage string) {
	sysflag.f.Var(newUint64Value(value, p), name, usage)
}

func Uint64Full(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64FullVar(p, name, title, key, env, value, usage)
	return p
}

func Uint64Config(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64ConfigVar(p, name, title, key, value, usage)
	return p
}

func Uint64Env(name, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64EnvVar(p, name, env, value, usage)
	return p
}

func Uint64(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64Var(p, name, value, usage)
	return p
}

func Args() []string { return sysflag.f.args }
func ArgsLen() int   { return len(sysflag.f.args) }

func Usage() {
	Greenf(fmt.Sprintf("%s的使用方法", os.Args[0]), "\n\n")
	sysflag.f.PrintDefaults()
}

func PrintAll() { sysflag.f.PrintAll() }

func SetOutput(output io.Writer) { sysflag.f.SetOutput(output) }
