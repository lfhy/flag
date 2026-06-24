package flag

import "strconv"

// parseUintsDefault 将默认值字符串按逗号切分为 []uint
func parseUintsDefault(value string) []uint {
	if value == "" {
		return []uint{}
	}
	parts := splitStrings(value)
	out := make([]uint, 0, len(parts))
	for _, p := range parts {
		n, _ := strconv.ParseUint(p, 0, 64)
		out = append(out, uint(n))
	}
	return out
}

// parseUint64sDefault 将默认值字符串按逗号切分为 []uint64
func parseUint64sDefault(value string) []uint64 {
	if value == "" {
		return []uint64{}
	}
	parts := splitStrings(value)
	out := make([]uint64, 0, len(parts))
	for _, p := range parts {
		n, _ := strconv.ParseUint(p, 0, 64)
		out = append(out, n)
	}
	return out
}

// ============================== FlagSet (Uints) ==============================

// Uints类型定义
// 定义一个[]uint类型的Flag，并设置其全名、标题、键、环境变量、默认值（逗号分隔字符串）和用法
func (f *FlagSet) UintsFullVar(p *[]uint, name string, title, key, env string, value string, usage string) {
	f.FullVar(newUintsValue(parseUintsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) UintsConfigVar(p *[]uint, name string, title, key string, value string, usage string) {
	f.FullVar(newUintsValue(parseUintsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) UintsEnvVar(p *[]uint, name string, env string, value string, usage string) {
	f.FullVar(newUintsValue(parseUintsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) UintsVar(p *[]uint, name string, value string, usage string) {
	f.FullVar(newUintsValue(parseUintsDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]uint类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) UintsFull(name, title, key, env string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) UintsConfig(name, title, key string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) UintsEnv(name, env string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Uints(name string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]uint类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) UintsFullHiddenVar(p *[]uint, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newUintsValue(parseUintsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) UintsConfigHiddenVar(p *[]uint, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newUintsValue(parseUintsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) UintsEnvHiddenVar(p *[]uint, name string, env string, value string, usage string) {
	f.FullHiddenVar(newUintsValue(parseUintsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) UintsHiddenVar(p *[]uint, name string, value string, usage string) {
	f.HiddenVar(newUintsValue(parseUintsDefault(value), p), name, usage)
}

// 定义一个[]uint类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) UintsFullHidden(name, title, key, env string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) UintsConfigHidden(name, title, key string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) UintsEnvHidden(name, env string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) UintsHidden(name string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsHiddenVar(p, name, value, usage)
	return p
}

// ============================== FlagSet (Uint64s) ==============================

// 定义一个[]uint64类型的Flag，并设置其全名、标题、键、环境变量、默认值（逗号分隔字符串）和用法
func (f *FlagSet) Uint64sFullVar(p *[]uint64, name string, title, key, env string, value string, usage string) {
	f.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint64类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) Uint64sConfigVar(p *[]uint64, name string, title, key string, value string, usage string) {
	f.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint64类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) Uint64sEnvVar(p *[]uint64, name string, env string, value string, usage string) {
	f.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint64类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) Uint64sVar(p *[]uint64, name string, value string, usage string) {
	f.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]uint64类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) Uint64sFull(name, title, key, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint64类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Uint64sConfig(name, title, key string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint64类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) Uint64sEnv(name, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint64类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Uint64s(name string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]uint64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) Uint64sFullHiddenVar(p *[]uint64, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint64类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) Uint64sConfigHiddenVar(p *[]uint64, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint64类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) Uint64sEnvHiddenVar(p *[]uint64, name string, env string, value string, usage string) {
	f.FullHiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint64类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) Uint64sHiddenVar(p *[]uint64, name string, value string, usage string) {
	f.HiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, usage)
}

// 定义一个[]uint64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Uint64sFullHidden(name, title, key, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint64类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Uint64sConfigHidden(name, title, key string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint64类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Uint64sEnvHidden(name, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint64类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Uint64sHidden(name string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sHiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag (Uints) ==============================

// 定义一个[]uint类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) UintsFullVar(p *[]uint, name string, title, key, env string, value string, usage string) {
	f.FullVar(newUintsValue(parseUintsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) UintsConfigVar(p *[]uint, name string, title, key string, value string, usage string) {
	f.FullVar(newUintsValue(parseUintsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) UintsEnvVar(p *[]uint, name string, env string, value string, usage string) {
	f.FullVar(newUintsValue(parseUintsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) UintsVar(p *[]uint, name string, value string, usage string) {
	f.FullVar(newUintsValue(parseUintsDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]uint类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) UintsFull(name, title, key, env string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) UintsConfig(name, title, key string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) UintsEnv(name, env string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Uints(name string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]uint类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) UintsFullHiddenVar(p *[]uint, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newUintsValue(parseUintsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) UintsConfigHiddenVar(p *[]uint, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newUintsValue(parseUintsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) UintsEnvHiddenVar(p *[]uint, name string, env string, value string, usage string) {
	f.FullHiddenVar(newUintsValue(parseUintsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) UintsHiddenVar(p *[]uint, name string, value string, usage string) {
	f.HiddenVar(newUintsValue(parseUintsDefault(value), p), name, usage)
}

// 定义一个[]uint类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) UintsFullHidden(name, title, key, env string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) UintsConfigHidden(name, title, key string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) UintsEnvHidden(name, env string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) UintsHidden(name string, value string, usage string) *[]uint {
	p := new([]uint)
	f.UintsHiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag (Uint64s) ==============================

// 定义一个[]uint64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Uint64sFullVar(p *[]uint64, name string, title, key, env string, value string, usage string) {
	f.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint64类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Uint64sConfigVar(p *[]uint64, name string, title, key string, value string, usage string) {
	f.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint64类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Uint64sEnvVar(p *[]uint64, name string, env string, value string, usage string) {
	f.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint64类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Uint64sVar(p *[]uint64, name string, value string, usage string) {
	f.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]uint64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Uint64sFull(name, title, key, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Uint64sConfig(name, title, key string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Uint64sEnv(name, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Uint64s(name string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]uint64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Uint64sFullHiddenVar(p *[]uint64, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint64类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Uint64sConfigHiddenVar(p *[]uint64, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint64类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Uint64sEnvHiddenVar(p *[]uint64, name string, env string, value string, usage string) {
	f.FullHiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint64类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Uint64sHiddenVar(p *[]uint64, name string, value string, usage string) {
	f.HiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, usage)
}

// 定义一个[]uint64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Uint64sFullHidden(name, title, key, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Uint64sConfigHidden(name, title, key string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Uint64sEnvHidden(name, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Uint64sHidden(name string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	f.Uint64sHiddenVar(p, name, value, usage)
	return p
}

// ============================== 包级函数（sysflag） ==============================

// Uints类型定义
// 定义一个[]uint类型的标志，并设置其名称、标题、键、环境变量、默认值（逗号分隔）和用法
func UintsFullVar(p *[]uint, name string, title, key, env string, value string, usage string) {
	sysflag.FullVar(newUintsValue(parseUintsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint类型的标志，并设置其名称、标题、键、默认值和用法
func UintsConfigVar(p *[]uint, name string, title, key string, value string, usage string) {
	sysflag.FullVar(newUintsValue(parseUintsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint类型的标志，并设置其名称、环境变量、默认值和用法
func UintsEnvVar(p *[]uint, name string, env string, value string, usage string) {
	sysflag.FullVar(newUintsValue(parseUintsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint类型的标志，并设置其名称、默认值和用法
func UintsVar(p *[]uint, name string, value string, usage string) {
	sysflag.FullVar(newUintsValue(parseUintsDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]uint类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func UintsFull(name, title, key, env string, value string, usage string) *[]uint {
	p := new([]uint)
	sysflag.UintsFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func UintsConfig(name, title, key string, value string, usage string) *[]uint {
	p := new([]uint)
	sysflag.UintsConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func UintsEnv(name, env string, value string, usage string) *[]uint {
	p := new([]uint)
	sysflag.UintsEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint类型的标志，并设置其名称、默认值和用法，并返回指针
func Uints(name string, value string, usage string) *[]uint {
	p := new([]uint)
	sysflag.UintsVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]uint类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func UintsFullHiddenVar(p *[]uint, name string, title, key, env string, value string, usage string) {
	sysflag.FullHiddenVar(newUintsValue(parseUintsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint类型的隐藏标志，并设置其名称、标题、键、默认值和用法
func UintsConfigHiddenVar(p *[]uint, name string, title, key string, value string, usage string) {
	sysflag.FullHiddenVar(newUintsValue(parseUintsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint类型的隐藏标志，并设置其名称、环境变量、默认值和用法
func UintsEnvHiddenVar(p *[]uint, name string, env string, value string, usage string) {
	sysflag.FullHiddenVar(newUintsValue(parseUintsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint类型的隐藏标志，并设置其名称、默认值和用法
func UintsHiddenVar(p *[]uint, name string, value string, usage string) {
	sysflag.HiddenVar(newUintsValue(parseUintsDefault(value), p), name, usage)
}

// 定义一个[]uint类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func UintsFullHidden(name, title, key, env string, value string, usage string) *[]uint {
	p := new([]uint)
	sysflag.UintsFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint类型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func UintsConfigHidden(name, title, key string, value string, usage string) *[]uint {
	p := new([]uint)
	sysflag.UintsConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint类型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func UintsEnvHidden(name, env string, value string, usage string) *[]uint {
	p := new([]uint)
	sysflag.UintsEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint类型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func UintsHidden(name string, value string, usage string) *[]uint {
	p := new([]uint)
	sysflag.UintsHiddenVar(p, name, value, usage)
	return p
}

// Uint64s类型定义
// 定义一个[]uint64类型的标志，并设置其名称、标题、键、环境变量、默认值（逗号分隔）和用法
func Uint64sFullVar(p *[]uint64, name string, title, key, env string, value string, usage string) {
	sysflag.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint64类型的标志，并设置其名称、标题、键、默认值和用法
func Uint64sConfigVar(p *[]uint64, name string, title, key string, value string, usage string) {
	sysflag.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint64类型的标志，并设置其名称、环境变量、默认值和用法
func Uint64sEnvVar(p *[]uint64, name string, env string, value string, usage string) {
	sysflag.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint64类型的标志，并设置其名称、默认值和用法
func Uint64sVar(p *[]uint64, name string, value string, usage string) {
	sysflag.FullVar(newUint64sValue(parseUint64sDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]uint64类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func Uint64sFull(name, title, key, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	sysflag.Uint64sFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint64类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func Uint64sConfig(name, title, key string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	sysflag.Uint64sConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint64类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func Uint64sEnv(name, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	sysflag.Uint64sEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint64类型的标志，并设置其名称、默认值和用法，并返回指针
func Uint64s(name string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	sysflag.Uint64sVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]uint64类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func Uint64sFullHiddenVar(p *[]uint64, name string, title, key, env string, value string, usage string) {
	sysflag.FullHiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]uint64类型的隐藏标志，并设置其名称、标题、键、默认值和用法
func Uint64sConfigHiddenVar(p *[]uint64, name string, title, key string, value string, usage string) {
	sysflag.FullHiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]uint64类型的隐藏标志，并设置其名称、环境变量、默认值和用法
func Uint64sEnvHiddenVar(p *[]uint64, name string, env string, value string, usage string) {
	sysflag.FullHiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]uint64类型的隐藏标志，并设置其名称、默认值和用法
func Uint64sHiddenVar(p *[]uint64, name string, value string, usage string) {
	sysflag.HiddenVar(newUint64sValue(parseUint64sDefault(value), p), name, usage)
}

// 定义一个[]uint64类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func Uint64sFullHidden(name, title, key, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	sysflag.Uint64sFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]uint64类型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func Uint64sConfigHidden(name, title, key string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	sysflag.Uint64sConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]uint64类型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func Uint64sEnvHidden(name, env string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	sysflag.Uint64sEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]uint64类型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func Uint64sHidden(name string, value string, usage string) *[]uint64 {
	p := new([]uint64)
	sysflag.Uint64sHiddenVar(p, name, value, usage)
	return p
}
