package flag

// ============================== FlagSet (Int) ==============================

// Int类型定义
// 定义一个Int类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) IntFullVar(p *int, name string, title, key, env string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个Int类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) IntConfigVar(p *int, name string, title, key string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个Int类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) IntEnvVar(p *int, name string, env string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个Int类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) IntVar(p *int, name string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, "", "", "", usage)
}

// 定义一个Int类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) IntFull(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.IntFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) IntConfig(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.IntConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) IntEnv(name, env string, value int, usage string) *int {
	p := new(int)
	f.IntEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Int(name string, value int, usage string) *int {
	p := new(int)
	f.IntVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个Int类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) IntFullHiddenVar(p *int, name string, title, key, env string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个Int类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) IntConfigHiddenVar(p *int, name string, title, key string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个Int类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) IntEnvHiddenVar(p *int, name string, env string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个Int类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) IntHiddenVar(p *int, name string, value int, usage string) {
	f.HiddenVar(newIntValue(value, p), name, usage)
}

// 定义一个Int类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) IntFullHidden(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.IntFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) IntConfigHidden(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.IntConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) IntEnvHidden(name, env string, value int, usage string) *int {
	p := new(int)
	f.IntEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Int类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) IntHidden(name string, value int, usage string) *int {
	p := new(int)
	f.IntHiddenVar(p, name, value, usage)
	return p
}

// ============================== FlagSet (Int64) ==============================

// Int64类型定义
// 定义一个Int64类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) Int64FullVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个Int64类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) Int64ConfigVar(p *int64, name string, title, key string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个Int64类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) Int64EnvVar(p *int64, name string, env string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个Int64类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, "", "", "", usage)
}

// 定义一个Int64类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) Int64Full(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Int64Config(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) Int64Env(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64Var(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个Int64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) Int64FullHiddenVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个Int64类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) Int64ConfigHiddenVar(p *int64, name string, title, key string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个Int64类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) Int64EnvHiddenVar(p *int64, name string, env string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个Int64类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) Int64HiddenVar(p *int64, name string, value int64, usage string) {
	f.HiddenVar(newInt64Value(value, p), name, usage)
}

// 定义一个Int64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Int64FullHidden(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Int64ConfigHidden(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Int64EnvHidden(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Int64类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Int64Hidden(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64HiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag (Int) ==============================

// 定义一个Int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) IntFullVar(p *int, name string, title, key, env string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个Int类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) IntConfigVar(p *int, name string, title, key string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个Int类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) IntEnvVar(p *int, name string, env string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个Int类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) IntVar(p *int, name string, value int, usage string) {
	f.FullVar(newIntValue(value, p), name, "", "", "", usage)
}

// 定义一个Int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) IntFull(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.IntFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) IntConfig(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.IntConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) IntEnv(name, env string, value int, usage string) *int {
	p := new(int)
	f.IntEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Int(name string, value int, usage string) *int {
	p := new(int)
	f.IntVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个Int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) IntFullHiddenVar(p *int, name string, title, key, env string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个Int类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) IntConfigHiddenVar(p *int, name string, title, key string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个Int类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) IntEnvHiddenVar(p *int, name string, env string, value int, usage string) {
	f.FullHiddenVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个Int类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) IntHiddenVar(p *int, name string, value int, usage string) {
	f.HiddenVar(newIntValue(value, p), name, usage)
}

// 定义一个Int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) IntFullHidden(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	f.IntFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) IntConfigHidden(name, title, key string, value int, usage string) *int {
	p := new(int)
	f.IntConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) IntEnvHidden(name, env string, value int, usage string) *int {
	p := new(int)
	f.IntEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Int类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) IntHidden(name string, value int, usage string) *int {
	p := new(int)
	f.IntHiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag (Int64) ==============================

// 定义一个Int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Int64FullVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Int64ConfigVar(p *int64, name string, title, key string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个Int64类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Int64EnvVar(p *int64, name string, env string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个Int64类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Int64Var(p *int64, name string, value int64, usage string) {
	f.FullVar(newInt64Value(value, p), name, "", "", "", usage)
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Int64Full(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Int64Config(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Int64Env(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64Var(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个Int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64FullHiddenVar(p *int64, name string, title, key, env string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64ConfigHiddenVar(p *int64, name string, title, key string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个Int64类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64EnvHiddenVar(p *int64, name string, env string, value int64, usage string) {
	f.FullHiddenVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个Int64类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64HiddenVar(p *int64, name string, value int64, usage string) {
	f.HiddenVar(newInt64Value(value, p), name, usage)
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64FullHidden(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64ConfigHidden(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64EnvHidden(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Int64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64Hidden(name string, value int64, usage string) *int64 {
	p := new(int64)
	f.Int64HiddenVar(p, name, value, usage)
	return p
}

// ============================== 包级函数（sysflag） ==============================

// Int类型定义
// 定义一个整型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func IntFullVar(p *int, name string, title, key, env string, value int, usage string) {
	sysflag.FullVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个整型的标志，并设置其名称、标题、键、默认值和用法
func IntConfigVar(p *int, name string, title, key string, value int, usage string) {
	sysflag.FullVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个整型的标志，并设置其名称、环境变量、默认值和用法
func IntEnvVar(p *int, name string, env string, value int, usage string) {
	sysflag.FullVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个整型的标志，并设置其名称、默认值和用法
func IntVar(p *int, name string, value int, usage string) {
	sysflag.FullVar(newIntValue(value, p), name, "", "", "", usage)
}

// 定义一个整型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func IntFull(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	sysflag.IntFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个整型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func IntConfig(name, title, key string, value int, usage string) *int {
	p := new(int)
	sysflag.IntConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个整型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func IntEnv(name, env string, value int, usage string) *int {
	p := new(int)
	sysflag.IntEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个整型的标志，并设置其名称、默认值和用法，并返回指针
func Int(name string, value int, usage string) *int {
	p := new(int)
	sysflag.IntVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个整型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func IntFullHiddenVar(p *int, name string, title, key, env string, value int, usage string) {
	sysflag.FullHiddenVar(newIntValue(value, p), name, title, key, env, usage)
}

// 定义一个整型的隐藏标志，并设置其名称、标题、键、默认值和用法
func IntConfigHiddenVar(p *int, name string, title, key string, value int, usage string) {
	sysflag.FullHiddenVar(newIntValue(value, p), name, title, key, "", usage)
}

// 定义一个整型的隐藏标志，并设置其名称、环境变量、默认值和用法
func IntEnvHiddenVar(p *int, name string, env string, value int, usage string) {
	sysflag.FullHiddenVar(newIntValue(value, p), name, "", "", env, usage)
}

// 定义一个整型的隐藏标志，并设置其名称、默认值和用法
func IntHiddenVar(p *int, name string, value int, usage string) {
	sysflag.HiddenVar(newIntValue(value, p), name, usage)
}

// 定义一个整型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func IntFullHidden(name, title, key, env string, value int, usage string) *int {
	p := new(int)
	sysflag.IntFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个整型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func IntConfigHidden(name, title, key string, value int, usage string) *int {
	p := new(int)
	sysflag.IntConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个整型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func IntEnvHidden(name, env string, value int, usage string) *int {
	p := new(int)
	sysflag.IntEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个整型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func IntHidden(name string, value int, usage string) *int {
	p := new(int)
	sysflag.IntHiddenVar(p, name, value, usage)
	return p
}

// Int64类型定义
// 定义一个64位整型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func Int64FullVar(p *int64, name string, title, key, env string, value int64, usage string) {
	sysflag.FullVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个64位整型的标志，并设置其名称、标题、键、默认值和用法
func Int64ConfigVar(p *int64, name string, title, key string, value int64, usage string) {
	sysflag.FullVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个64位整型的标志，并设置其名称、环境变量、默认值和用法
func Int64EnvVar(p *int64, name string, env string, value int64, usage string) {
	sysflag.FullVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个64位整型的标志，并设置其名称、默认值和用法
func Int64Var(p *int64, name string, value int64, usage string) {
	sysflag.FullVar(newInt64Value(value, p), name, "", "", "", usage)
}

// 定义一个64位整型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func Int64Full(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64FullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个64位整型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func Int64Config(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64ConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个64位整型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func Int64Env(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64EnvVar(p, name, env, value, usage)
	return p
}

// 定义一个64位整型的标志，并设置其名称、默认值和用法，并返回指针
func Int64(name string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64Var(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个64位整型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func Int64FullHiddenVar(p *int64, name string, title, key, env string, value int64, usage string) {
	sysflag.FullHiddenVar(newInt64Value(value, p), name, title, key, env, usage)
}

// 定义一个64位整型的隐藏标志，并设置其名称、标题、键、默认值和用法
func Int64ConfigHiddenVar(p *int64, name string, title, key string, value int64, usage string) {
	sysflag.FullHiddenVar(newInt64Value(value, p), name, title, key, "", usage)
}

// 定义一个64位整型的隐藏标志，并设置其名称、环境变量、默认值和用法
func Int64EnvHiddenVar(p *int64, name string, env string, value int64, usage string) {
	sysflag.FullHiddenVar(newInt64Value(value, p), name, "", "", env, usage)
}

// 定义一个64位整型的隐藏标志，并设置其名称、默认值和用法
func Int64HiddenVar(p *int64, name string, value int64, usage string) {
	sysflag.HiddenVar(newInt64Value(value, p), name, usage)
}

// 定义一个64位整型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func Int64FullHidden(name, title, key, env string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64FullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个64位整型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func Int64ConfigHidden(name, title, key string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64ConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个64位整型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func Int64EnvHidden(name, env string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64EnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个64位整型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func Int64Hidden(name string, value int64, usage string) *int64 {
	p := new(int64)
	sysflag.Int64HiddenVar(p, name, value, usage)
	return p
}
