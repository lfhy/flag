package flag

import "time"

// ============================== FlagSet ==============================

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
	f.FullVar(newDurationValue(value, p), name, "", "", "", usage)
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

// ----------------------------- Hidden -----------------------------

// 定义一个Duration类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) DurationFullHiddenVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, title, key, env, usage)
}

// 定义一个Duration类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) DurationConfigHiddenVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, title, key, "", usage)
}

// 定义一个Duration类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) DurationEnvHiddenVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, "", "", env, usage)
}

// 定义一个Duration类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) DurationHiddenVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.HiddenVar(newDurationValue(value, p), name, usage)
}

// 定义一个Duration类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) DurationFullHidden(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Duration类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) DurationConfigHidden(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Duration类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) DurationEnvHidden(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Duration类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) DurationHidden(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationHiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag ==============================

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) DurationFullVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	f.FullVar(newDurationValue(value, p), name, title, key, env, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) DurationConfigVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	f.FullVar(newDurationValue(value, p), name, title, key, "", usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) DurationEnvVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	f.FullVar(newDurationValue(value, p), name, "", "", env, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.FullVar(newDurationValue(value, p), name, "", "", "", usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) DurationFull(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) DurationConfig(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) DurationEnv(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Duration(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) DurationFullHiddenVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, title, key, env, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) DurationConfigHiddenVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, title, key, "", usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) DurationEnvHiddenVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	f.FullHiddenVar(newDurationValue(value, p), name, "", "", env, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) DurationHiddenVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.HiddenVar(newDurationValue(value, p), name, usage)
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) DurationFullHidden(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) DurationConfigHidden(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) DurationEnvHidden(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个time.Duration类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) DurationHidden(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	f.DurationHiddenVar(p, name, value, usage)
	return p
}

// ============================== 包级函数（sysflag） ==============================

// 定义一个延迟类型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func DurationFullVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	sysflag.FullVar(newDurationValue(value, p), name, title, key, env, usage)
}

// 定义一个延迟类型的标志，并设置其名称、标题、键、默认值和用法
func DurationConfigVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	sysflag.FullVar(newDurationValue(value, p), name, title, key, "", usage)
}

// 定义一个延迟类型的标志，并设置其名称、环境变量、默认值和用法
func DurationEnvVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	sysflag.FullVar(newDurationValue(value, p), name, "", "", env, usage)
}

// 定义一个延迟类型的标志，并设置其名称、默认值和用法
func DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	sysflag.FullVar(newDurationValue(value, p), name, "", "", "", usage)
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

// ----------------------------- Hidden -----------------------------

// 定义一个延迟类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func DurationFullHiddenVar(p *time.Duration, name string, title, key, env string, value time.Duration, usage string) {
	sysflag.FullHiddenVar(newDurationValue(value, p), name, title, key, env, usage)
}

// 定义一个延迟类型的隐藏标志，并设置其名称、标题、键、默认值和用法
func DurationConfigHiddenVar(p *time.Duration, name string, title, key string, value time.Duration, usage string) {
	sysflag.FullHiddenVar(newDurationValue(value, p), name, title, key, "", usage)
}

// 定义一个延迟类型的隐藏标志，并设置其名称、环境变量、默认值和用法
func DurationEnvHiddenVar(p *time.Duration, name string, env string, value time.Duration, usage string) {
	sysflag.FullHiddenVar(newDurationValue(value, p), name, "", "", env, usage)
}

// 定义一个延迟类型的隐藏标志，并设置其名称、默认值和用法
func DurationHiddenVar(p *time.Duration, name string, value time.Duration, usage string) {
	sysflag.HiddenVar(newDurationValue(value, p), name, usage)
}

// 定义一个延迟类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func DurationFullHidden(name, title, key, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个延迟类型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func DurationConfigHidden(name, title, key string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个延迟类型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func DurationEnvHidden(name, env string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个延迟类型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func DurationHidden(name string, value time.Duration, usage string) *time.Duration {
	p := new(time.Duration)
	sysflag.DurationHiddenVar(p, name, value, usage)
	return p
}
