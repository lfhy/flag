package flag

// ============================== FlagSet ==============================

// Bool类型定义
// 定义一个Bool类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法
func (f *FlagSet) BoolFullVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个Bool类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) BoolConfigVar(p *bool, name string, title, key string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个Bool类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) BoolEnvVar(p *bool, name string, env string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个Bool类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, "", "", "", usage)
}

// 定义一个Bool类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) BoolFull(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) BoolConfig(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) BoolEnv(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个Bool类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) BoolFullHiddenVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个Bool类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) BoolConfigHiddenVar(p *bool, name string, title, key string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个Bool类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) BoolEnvHiddenVar(p *bool, name string, env string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个Bool类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) BoolHiddenVar(p *bool, name string, value bool, usage string) {
	f.HiddenVar(newBoolValue(value, p), name, usage)
}

// 定义一个Bool类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) BoolFullHidden(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) BoolConfigHidden(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) BoolEnvHidden(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Bool类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) BoolHidden(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolHiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag ==============================

// 定义一个Bool类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) BoolFullVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) BoolConfigVar(p *bool, name string, title, key string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个Bool类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) BoolEnvVar(p *bool, name string, env string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个Bool类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) BoolVar(p *bool, name string, value bool, usage string) {
	f.FullVar(newBoolValue(value, p), name, "", "", "", usage)
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) BoolFull(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) BoolConfig(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) BoolEnv(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个Bool类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) BoolFullHiddenVar(p *bool, name string, title, key, env string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) BoolConfigHiddenVar(p *bool, name string, title, key string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个Bool类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) BoolEnvHiddenVar(p *bool, name string, env string, value bool, usage string) {
	f.FullHiddenVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个Bool类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) BoolHiddenVar(p *bool, name string, value bool, usage string) {
	f.HiddenVar(newBoolValue(value, p), name, usage)
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) BoolFullHidden(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) BoolConfigHidden(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) BoolEnvHidden(name, env string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个Bool类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) BoolHidden(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolHiddenVar(p, name, value, usage)
	return p
}

// ============================== 包级函数（sysflag） ==============================

// 定义一个布尔类型的标志，并设置其名称、标题、键、环境变量、默认值和用法
func BoolFullVar(p *bool, name string, title, key, env string, value bool, usage string) {
	sysflag.FullVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个布尔类型的标志，并设置其名称、标题、键、默认值和用法
func BoolConfigVar(p *bool, name string, title, key string, value bool, usage string) {
	sysflag.FullVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个布尔类型的标志，并设置其名称、环境变量、默认值和用法
func BoolEnvVar(p *bool, name string, env string, value bool, usage string) {
	sysflag.FullVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个布尔类型的标志，并设置其名称、默认值和用法
func BoolVar(p *bool, name string, value bool, usage string) {
	sysflag.FullVar(newBoolValue(value, p), name, "", "", "", usage)
}

// 定义一个布尔类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func BoolFull(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个布尔类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func BoolConfig(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个布尔类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func BoolEnv(name, env string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个布尔类型的标志，并设置其名称、默认值和用法，并返回指针
func Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个布尔类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func BoolFullHiddenVar(p *bool, name string, title, key, env string, value bool, usage string) {
	sysflag.FullHiddenVar(newBoolValue(value, p), name, title, key, env, usage)
}

// 定义一个布尔类型的隐藏标志，并设置其名称、标题、键、默认值和用法
func BoolConfigHiddenVar(p *bool, name string, title, key string, value bool, usage string) {
	sysflag.FullHiddenVar(newBoolValue(value, p), name, title, key, "", usage)
}

// 定义一个布尔类型的隐藏标志，并设置其名称、环境变量、默认值和用法
func BoolEnvHiddenVar(p *bool, name string, env string, value bool, usage string) {
	sysflag.FullHiddenVar(newBoolValue(value, p), name, "", "", env, usage)
}

// 定义一个布尔类型的隐藏标志，并设置其名称、默认值和用法
func BoolHiddenVar(p *bool, name string, value bool, usage string) {
	sysflag.HiddenVar(newBoolValue(value, p), name, usage)
}

// 定义一个布尔类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func BoolFullHidden(name, title, key, env string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个布尔类型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func BoolConfigHidden(name, title, key string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个布尔类型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func BoolEnvHidden(name, env string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个布尔类型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func BoolHidden(name string, value bool, usage string) *bool {
	p := new(bool)
	sysflag.BoolHiddenVar(p, name, value, usage)
	return p
}
