package flag

// ============================== FlagSet (Uint) ==============================

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
	f.FullVar(newUintValue(value, p), name, "", "", "", usage)
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

// ----------------------------- Hidden -----------------------------

// 定义一个Uint类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) UintFullHiddenVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个Uint类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) UintConfigHiddenVar(p *uint, name string, title, key string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个Uint类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) UintEnvHiddenVar(p *uint, name string, env string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个Uint类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) UintHiddenVar(p *uint, name string, value uint, usage string) {
	f.HiddenVar(newUintValue(value, p), name, usage)
}

// 定义一个Uint类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) UintFullHidden(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Uint类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) UintConfigHidden(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	f.UintConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// Uint类型定义（隐藏）
func (f *FlagSet) UintEnvHidden(name, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintEnvHiddenVar(p, name, env, value, usage)
	return p
}

// Uint类型定义（隐藏）
func (f *FlagSet) UintHidden(name string, value uint, usage string) *uint {
	p := new(uint)
	f.UintHiddenVar(p, name, value, usage)
	return p
}

// ============================== FlagSet (Uint64) ==============================

// Uint64类型定义
func (f *FlagSet) Uint64FullVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	f.FullVar(newUint64Value(value, p), name, title, key, env, usage)
}

// Uint64ConfigVar 函数用于将一个 uint64 类型的变量与一个配置文件中的变量进行绑定
func (f *FlagSet) Uint64ConfigVar(p *uint64, name string, title, key string, value uint64, usage string) {
	f.FullVar(newUint64Value(value, p), name, title, key, "", usage)
}

// Uint64EnvVar 函数用于将一个 uint64 类型的变量与一个环境变量进行绑定
func (f *FlagSet) Uint64EnvVar(p *uint64, name string, env string, value uint64, usage string) {
	f.FullVar(newUint64Value(value, p), name, "", "", env, usage)
}

// Uint64Var 函数用于将一个 uint64 类型的变量与一个命令行参数进行绑定
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.FullVar(newUint64Value(value, p), name, "", "", "", usage)
}

// Uint64Full 函数用于将一个 uint64 类型的变量与一个配置文件、环境变量和命令行参数进行绑定
func (f *FlagSet) Uint64Full(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64FullVar(p, name, title, key, env, value, usage)
	return p
}

// Uint64Config 函数用于将一个 uint64 类型的变量与一个配置文件和命令行参数进行绑定
func (f *FlagSet) Uint64Config(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64ConfigVar(p, name, title, key, value, usage)
	return p
}

// Uint64Env 函数用于将一个 uint64 类型的变量与一个环境变量和命令行参数进行绑定
func (f *FlagSet) Uint64Env(name, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64EnvVar(p, name, env, value, usage)
	return p
}

// Uint64函数用于设置一个uint64类型的flag，并返回一个指向该flag的指针
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64Var(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// Uint64类型定义（隐藏）
func (f *FlagSet) Uint64FullHiddenVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	f.FullHiddenVar(newUint64Value(value, p), name, title, key, env, usage)
}

// Uint64类型定义（隐藏）
func (f *FlagSet) Uint64ConfigHiddenVar(p *uint64, name string, title, key string, value uint64, usage string) {
	f.FullHiddenVar(newUint64Value(value, p), name, title, key, "", usage)
}

// Uint64类型定义（隐藏）
func (f *FlagSet) Uint64EnvHiddenVar(p *uint64, name string, env string, value uint64, usage string) {
	f.FullHiddenVar(newUint64Value(value, p), name, "", "", env, usage)
}

// Uint64类型定义（隐藏）
func (f *FlagSet) Uint64HiddenVar(p *uint64, name string, value uint64, usage string) {
	f.HiddenVar(newUint64Value(value, p), name, usage)
}

// Uint64类型定义（隐藏）
func (f *FlagSet) Uint64FullHidden(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// Uint64类型定义（隐藏）
func (f *FlagSet) Uint64ConfigHidden(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// Uint64类型定义（隐藏）
func (f *FlagSet) Uint64EnvHidden(name, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// Uint64类型定义（隐藏）
func (f *FlagSet) Uint64Hidden(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64HiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag (Uint) ==============================

// 定义一个Uint类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) UintFullVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.FullVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个Uint类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) UintConfigVar(p *uint, name string, title, key string, value uint, usage string) {
	f.FullVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个Uint类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) UintEnvVar(p *uint, name string, env string, value uint, usage string) {
	f.FullVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个Uint类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) UintVar(p *uint, name string, value uint, usage string) {
	f.FullVar(newUintValue(value, p), name, "", "", "", usage)
}

// 定义一个Uint类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) UintFull(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Uint类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) UintConfig(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	f.UintConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Uint类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) UintEnv(name, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Uint类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Uint(name string, value uint, usage string) *uint {
	p := new(uint)
	f.UintVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个Uint类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) UintFullHiddenVar(p *uint, name string, title, key, env string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个Uint类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) UintConfigHiddenVar(p *uint, name string, title, key string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个Uint类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) UintEnvHiddenVar(p *uint, name string, env string, value uint, usage string) {
	f.FullHiddenVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个Uint类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) UintHiddenVar(p *uint, name string, value uint, usage string) {
	f.HiddenVar(newUintValue(value, p), name, usage)
}

// Uint类型定义（ArgsFlag，隐藏）
func (f *ArgsFlag) UintFullHidden(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// Uint类型定义（ArgsFlag，隐藏）
func (f *ArgsFlag) UintConfigHidden(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	f.UintConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// Uint类型定义（ArgsFlag，隐藏）
func (f *ArgsFlag) UintEnvHidden(name, env string, value uint, usage string) *uint {
	p := new(uint)
	f.UintEnvHiddenVar(p, name, env, value, usage)
	return p
}

// Uint类型定义（ArgsFlag，隐藏）
func (f *ArgsFlag) UintHidden(name string, value uint, usage string) *uint {
	p := new(uint)
	f.UintHiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag (Uint64) ==============================

// Uint64类型定义（ArgsFlag）
func (f *ArgsFlag) Uint64FullHiddenVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	f.FullHiddenVar(newUint64Value(value, p), name, title, key, env, usage)
}

// Uint64类型定义（ArgsFlag）
func (f *ArgsFlag) Uint64ConfigHiddenVar(p *uint64, name string, title, key string, value uint64, usage string) {
	f.FullHiddenVar(newUint64Value(value, p), name, title, key, "", usage)
}

// Uint64类型定义（ArgsFlag）
func (f *ArgsFlag) Uint64EnvHiddenVar(p *uint64, name string, env string, value uint64, usage string) {
	f.FullHiddenVar(newUint64Value(value, p), name, "", "", env, usage)
}

// Uint64类型定义（ArgsFlag）
func (f *ArgsFlag) Uint64HiddenVar(p *uint64, name string, value uint64, usage string) {
	f.HiddenVar(newUint64Value(value, p), name, usage)
}

// Uint64类型定义（ArgsFlag）
func (f *ArgsFlag) Uint64FullHidden(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// Uint64类型定义（ArgsFlag）
func (f *ArgsFlag) Uint64ConfigHidden(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// Uint64EnvHidden函数用于设置一个uint64类型的参数，该参数可以通过环境变量设置，并且在帮助信息中隐藏（ArgsFlag）
func (f *ArgsFlag) Uint64EnvHidden(name, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// Uint64Hidden函数用于设置一个uint64类型的隐藏参数（ArgsFlag）
func (f *ArgsFlag) Uint64Hidden(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	f.Uint64HiddenVar(p, name, value, usage)
	return p
}

// 注意：ArgsFlag 的 Uint64 非隐藏方法使用 *FlagSet 的 Uint64* 方法（通过内嵌自动继承）

// ============================== 包级函数（sysflag） ==============================

// Uint类型定义
// 定义一个无符号整型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func UintFullVar(p *uint, name string, title, key, env string, value uint, usage string) {
	sysflag.FullVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个无符号整型的标志，并设置其名称、标题、键、默认值和用法
func UintConfigVar(p *uint, name string, title, key string, value uint, usage string) {
	sysflag.FullVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个无符号整型的标志，并设置其名称、环境变量、默认值和用法
func UintEnvVar(p *uint, name string, env string, value uint, usage string) {
	sysflag.FullVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个无符号整型的标志，并设置其名称、默认值和用法
func UintVar(p *uint, name string, value uint, usage string) {
	sysflag.FullVar(newUintValue(value, p), name, "", "", "", usage)
}

// 定义一个无符号整型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func UintFull(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个无符号整型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func UintConfig(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个无符号整型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func UintEnv(name, env string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个无符号整型的标志，并设置其名称、默认值和用法，并返回指针
func Uint(name string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个无符号整型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func UintFullHiddenVar(p *uint, name string, title, key, env string, value uint, usage string) {
	sysflag.FullHiddenVar(newUintValue(value, p), name, title, key, env, usage)
}

// 定义一个无符号整型的隐藏标志，并设置其名称、标题、键、默认值和用法
func UintConfigHiddenVar(p *uint, name string, title, key string, value uint, usage string) {
	sysflag.FullHiddenVar(newUintValue(value, p), name, title, key, "", usage)
}

// 定义一个无符号整型的隐藏标志，并设置其名称、环境变量、默认值和用法
func UintEnvHiddenVar(p *uint, name string, env string, value uint, usage string) {
	sysflag.FullHiddenVar(newUintValue(value, p), name, "", "", env, usage)
}

// 定义一个无符号整型的隐藏标志，并设置其名称、默认值和用法
func UintHiddenVar(p *uint, name string, value uint, usage string) {
	sysflag.HiddenVar(newUintValue(value, p), name, usage)
}

// 定义一个无符号整型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func UintFullHidden(name, title, key, env string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个无符号整型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func UintConfigHidden(name, title, key string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个无符号整型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func UintEnvHidden(name, env string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个无符号整型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func UintHidden(name string, value uint, usage string) *uint {
	p := new(uint)
	sysflag.UintHiddenVar(p, name, value, usage)
	return p
}

// Uint64类型定义
// 定义一个64位无符号整型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func Uint64FullVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	sysflag.FullVar(newUint64Value(value, p), name, title, key, env, usage)
}

// 定义64位无符号整型的标志，并设置其名称、标题、键、默认值和用法
func Uint64ConfigVar(p *uint64, name string, title, key string, value uint64, usage string) {
	sysflag.FullVar(newUint64Value(value, p), name, title, key, "", usage)
}

// 定义一个64位无符号整型的标志，并设置其名称、环境变量、默认值和用法
func Uint64EnvVar(p *uint64, name string, env string, value uint64, usage string) {
	sysflag.FullVar(newUint64Value(value, p), name, "", "", env, usage)
}

// Uint64Var函数用于将一个uint64类型的变量注册到命令行参数中
func Uint64Var(p *uint64, name string, value uint64, usage string) {
	sysflag.FullVar(newUint64Value(value, p), name, "", "", "", usage)
}

// Uint64Full函数用于将一个uint64类型的变量注册到命令行参数中，并设置参数的名称、标题、键和值
func Uint64Full(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64FullVar(p, name, title, key, env, value, usage)
	return p
}

// Uint64Config函数用于将一个uint64类型的变量注册到命令行参数中，并设置参数的名称、标题、键和值
func Uint64Config(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64ConfigVar(p, name, title, key, value, usage)
	return p
}

// Uint64Env函数用于将一个uint64类型的变量注册到命令行参数中，并设置参数的名称、环境变量和值
func Uint64Env(name, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64EnvVar(p, name, env, value, usage)
	return p
}

// Uint64函数用于将一个uint64类型的变量注册到命令行参数中，并设置参数的名称和值
func Uint64(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64Var(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// Uint64类型定义（隐藏）
func Uint64FullHiddenVar(p *uint64, name string, title, key, env string, value uint64, usage string) {
	sysflag.FullHiddenVar(newUint64Value(value, p), name, title, key, env, usage)
}

// Uint64类型定义（隐藏）
func Uint64ConfigHiddenVar(p *uint64, name string, title, key string, value uint64, usage string) {
	sysflag.FullHiddenVar(newUint64Value(value, p), name, title, key, "", usage)
}

// Uint64类型定义（隐藏）
func Uint64EnvHiddenVar(p *uint64, name string, env string, value uint64, usage string) {
	sysflag.FullHiddenVar(newUint64Value(value, p), name, "", "", env, usage)
}

// Uint64类型定义（隐藏）
func Uint64HiddenVar(p *uint64, name string, value uint64, usage string) {
	sysflag.HiddenVar(newUint64Value(value, p), name, usage)
}

// Uint64类型定义（隐藏）
func Uint64FullHidden(name, title, key, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// Uint64类型定义（隐藏）
func Uint64ConfigHidden(name, title, key string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// Uint64类型定义（隐藏）
func Uint64EnvHidden(name, env string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// Uint64类型定义（隐藏）
func Uint64Hidden(name string, value uint64, usage string) *uint64 {
	p := new(uint64)
	sysflag.Uint64HiddenVar(p, name, value, usage)
	return p
}
