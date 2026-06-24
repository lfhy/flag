package flag

// ============================== FlagSet ==============================

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
	f.FullVar(newFloat64Value(value, p), name, "", "", "", usage)
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

// ----------------------------- Hidden -----------------------------

// 定义一个Float64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) Float64FullHiddenVar(p *float64, name string, title, key, env string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// 定义一个Float64类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) Float64ConfigHiddenVar(p *float64, name string, title, key string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// 定义一个Float64类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) Float64EnvHiddenVar(p *float64, name string, env string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// 定义一个Float64类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) Float64HiddenVar(p *float64, name string, value float64, usage string) {
	f.HiddenVar(newFloat64Value(value, p), name, usage)
}

// 定义一个Float64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Float64FullHidden(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Float64类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Float64ConfigHidden(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Float64类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Float64EnvHidden(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Float64类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Float64Hidden(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64HiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag ==============================

// 定义一个Float64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Float64FullVar(p *float64, name string, title, key, env string, value float64, usage string) {
	f.FullVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Float64ConfigVar(p *float64, name string, title, key string, value float64, usage string) {
	f.FullVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// 定义一个Float64类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Float64EnvVar(p *float64, name string, env string, value float64, usage string) {
	f.FullVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// 定义一个Float64类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Float64Var(p *float64, name string, value float64, usage string) {
	f.FullVar(newFloat64Value(value, p), name, "", "", "", usage)
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Float64Full(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Float64Config(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64ConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Float64Env(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64EnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Float64(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64Var(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个Float64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Float64FullHiddenVar(p *float64, name string, title, key, env string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Float64ConfigHiddenVar(p *float64, name string, title, key string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// 定义一个Float64类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Float64EnvHiddenVar(p *float64, name string, env string, value float64, usage string) {
	f.FullHiddenVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// 定义一个Float64类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Float64HiddenVar(p *float64, name string, value float64, usage string) {
	f.HiddenVar(newFloat64Value(value, p), name, usage)
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Float64FullHidden(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Float64ConfigHidden(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Float64EnvHidden(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Float64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Float64Hidden(name string, value float64, usage string) *float64 {
	p := new(float64)
	f.Float64HiddenVar(p, name, value, usage)
	return p
}

// ============================== 包级函数（sysflag） ==============================

// 定义一个浮点类型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func Float64FullVar(p *float64, name string, title, key, env string, value float64, usage string) {
	sysflag.FullVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// 定义一个浮点类型的标志，并设置其名称、标题、键、默认值和用法
func Float64ConfigVar(p *float64, name string, title, key string, value float64, usage string) {
	sysflag.FullVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// 定义一个浮点类型的标志，并设置其名称、环境变量、默认值和用法
func Float64EnvVar(p *float64, name string, env string, value float64, usage string) {
	sysflag.FullVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// 定义一个浮点类型的标志，并设置其名称、默认值和用法
func Float64Var(p *float64, name string, value float64, usage string) {
	sysflag.FullVar(newFloat64Value(value, p), name, "", "", "", usage)
}

// 定义一个浮点类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func Float64Full(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个浮点类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func Float64Config(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64ConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个浮点类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func Float64Env(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64EnvVar(p, name, env, value, usage)
	return p
}

// 定义一个浮点类型的标志，并设置其名称、默认值和用法，并返回指针
func Float64(name string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64Var(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个浮点类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func Float64FullHiddenVar(p *float64, name string, title, key, env string, value float64, usage string) {
	sysflag.FullHiddenVar(newFloat64Value(value, p), name, title, key, env, usage)
}

// 定义一个浮点类型的隐藏标志，并设置其名称、标题、键、默认值和用法
func Float64ConfigHiddenVar(p *float64, name string, title, key string, value float64, usage string) {
	sysflag.FullHiddenVar(newFloat64Value(value, p), name, title, key, "", usage)
}

// 定义一个浮点类型的隐藏标志，并设置其名称、环境变量、默认值和用法
func Float64EnvHiddenVar(p *float64, name string, env string, value float64, usage string) {
	sysflag.FullHiddenVar(newFloat64Value(value, p), name, "", "", env, usage)
}

// 定义一个浮点类型的隐藏标志，并设置其名称、默认值和用法
func Float64HiddenVar(p *float64, name string, value float64, usage string) {
	sysflag.HiddenVar(newFloat64Value(value, p), name, usage)
}

// 定义一个浮点类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func Float64FullHidden(name, title, key, env string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个浮点类型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func Float64ConfigHidden(name, title, key string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个浮点类型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func Float64EnvHidden(name, env string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个浮点类型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func Float64Hidden(name string, value float64, usage string) *float64 {
	p := new(float64)
	sysflag.Float64HiddenVar(p, name, value, usage)
	return p
}
