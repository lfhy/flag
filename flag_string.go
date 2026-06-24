package flag

// ============================== FlagSet ==============================

// String类型定义
// 定义一个String类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) StringFullVar(p *string, name string, title, key, env string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个String类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) StringConfigVar(p *string, name string, title, key string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个String类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) StringEnvVar(p *string, name string, env string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个String类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, "", "", "", usage)
}

// 定义一个String类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) StringFull(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.StringFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) StringConfig(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.StringConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) StringEnv(name, env string, value string, usage string) *string {
	p := new(string)
	f.StringEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) String(name string, value string, usage string) *string {
	p := new(string)
	f.StringVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个String类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) StringFullHiddenVar(p *string, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个String类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) StringConfigHiddenVar(p *string, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个String类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) StringEnvHiddenVar(p *string, name string, env string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个String类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) StringHiddenVar(p *string, name string, value string, usage string) {
	f.HiddenVar(newStringValue(value, p), name, usage)
}

// 定义一个String类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) StringFullHidden(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.StringFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) StringConfigHidden(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.StringConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) StringEnvHidden(name, env string, value string, usage string) *string {
	p := new(string)
	f.StringEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个String类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) StringHidden(name string, value string, usage string) *string {
	p := new(string)
	f.StringHiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag ==============================

// 定义一个String类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) StringFullVar(p *string, name string, title, key, env string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个String类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) StringConfigVar(p *string, name string, title, key string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个String类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) StringEnvVar(p *string, name string, env string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个String类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) StringVar(p *string, name string, value string, usage string) {
	f.FullVar(newStringValue(value, p), name, "", "", "", usage)
}

// 定义一个String类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) StringFull(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.StringFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) StringConfig(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.StringConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) StringEnv(name, env string, value string, usage string) *string {
	p := new(string)
	f.StringEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) String(name string, value string, usage string) *string {
	p := new(string)
	f.StringVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个String类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) StringFullHiddenVar(p *string, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个String类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) StringConfigHiddenVar(p *string, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个String类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) StringEnvHiddenVar(p *string, name string, env string, value string, usage string) {
	f.FullHiddenVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个String类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) StringHiddenVar(p *string, name string, value string, usage string) {
	f.HiddenVar(newStringValue(value, p), name, usage)
}

// 定义一个String类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) StringFullHidden(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	f.StringFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) StringConfigHidden(name, title, key string, value string, usage string) *string {
	p := new(string)
	f.StringConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) StringEnvHidden(name, env string, value string, usage string) *string {
	p := new(string)
	f.StringEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个String类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) StringHidden(name string, value string, usage string) *string {
	p := new(string)
	f.StringHiddenVar(p, name, value, usage)
	return p
}

// ============================== 包级函数（sysflag） ==============================

// 定义一个字符串类型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func StringFullVar(p *string, name string, title, key, env string, value string, usage string) {
	sysflag.FullVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个字符串类型的标志，并设置其名称、标题、键、默认值和用法
func StringConfigVar(p *string, name string, title, key string, value string, usage string) {
	sysflag.FullVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个字符串类型的标志，并设置其名称、环境变量、默认值和用法
func StringEnvVar(p *string, name string, env string, value string, usage string) {
	sysflag.FullVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个字符串类型的标志，并设置其名称、默认值和用法
func StringVar(p *string, name string, value string, usage string) {
	sysflag.FullVar(newStringValue(value, p), name, "", "", "", usage)
}

// 定义一个字符串类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func StringFull(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	sysflag.StringFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个字符串类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func StringConfig(name, title, key string, value string, usage string) *string {
	p := new(string)
	sysflag.StringConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个字符串类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func StringEnv(name, env string, value string, usage string) *string {
	p := new(string)
	sysflag.StringEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个字符串类型的标志，并设置其名称、默认值和用法，并返回指针
func String(name string, value string, usage string) *string {
	p := new(string)
	sysflag.StringVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个字符串类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func StringFullHiddenVar(p *string, name string, title, key, env string, value string, usage string) {
	sysflag.FullHiddenVar(newStringValue(value, p), name, title, key, env, usage)
}

// 定义一个字符串类型的隐藏标志，并设置其名称、标题、键、默认值和用法
func StringConfigHiddenVar(p *string, name string, title, key string, value string, usage string) {
	sysflag.FullHiddenVar(newStringValue(value, p), name, title, key, "", usage)
}

// 定义一个字符串类型的隐藏标志，并设置其名称、环境变量、默认值和用法
func StringEnvHiddenVar(p *string, name string, env string, value string, usage string) {
	sysflag.FullHiddenVar(newStringValue(value, p), name, "", "", env, usage)
}

// 定义一个字符串类型的隐藏标志，并设置其名称、默认值和用法
func StringHiddenVar(p *string, name string, value string, usage string) {
	sysflag.HiddenVar(newStringValue(value, p), name, usage)
}

// 定义一个字符串类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func StringFullHidden(name, title, key, env string, value string, usage string) *string {
	p := new(string)
	sysflag.StringFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个字符串类型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func StringConfigHidden(name, title, key string, value string, usage string) *string {
	p := new(string)
	sysflag.StringConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个字符串类型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func StringEnvHidden(name, env string, value string, usage string) *string {
	p := new(string)
	sysflag.StringEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个字符串类型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func StringHidden(name string, value string, usage string) *string {
	p := new(string)
	sysflag.StringHiddenVar(p, name, value, usage)
	return p
}
