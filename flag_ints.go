package flag

import "strconv"

// parseIntsDefault 将默认值字符串按逗号切分为 []int
func parseIntsDefault(value string) []int {
	if value == "" {
		return []int{}
	}
	parts := splitStrings(value)
	out := make([]int, 0, len(parts))
	for _, p := range parts {
		n, _ := strconv.ParseInt(p, 0, 64)
		out = append(out, int(n))
	}
	return out
}

// parseInt64sDefault 将默认值字符串按逗号切分为 []int64
func parseInt64sDefault(value string) []int64 {
	if value == "" {
		return []int64{}
	}
	parts := splitStrings(value)
	out := make([]int64, 0, len(parts))
	for _, p := range parts {
		n, _ := strconv.ParseInt(p, 0, 64)
		out = append(out, n)
	}
	return out
}

// ============================== FlagSet (Ints) ==============================

// Ints类型定义
// 定义一个[]int类型的Flag，并设置其全名、标题、键、环境变量、默认值（逗号分隔字符串）和用法
func (f *FlagSet) IntsFullVar(p *[]int, name string, title, key, env string, value string, usage string) {
	f.FullVar(newIntsValue(parseIntsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) IntsConfigVar(p *[]int, name string, title, key string, value string, usage string) {
	f.FullVar(newIntsValue(parseIntsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) IntsEnvVar(p *[]int, name string, env string, value string, usage string) {
	f.FullVar(newIntsValue(parseIntsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) IntsVar(p *[]int, name string, value string, usage string) {
	f.FullVar(newIntsValue(parseIntsDefault(value), p), name, "", "", "", usage)
}

// 定义一个IntSlice类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) IntSliceVar(p *IntSlice, name string, value []int, usage string) {
	*p = IntSlice(value)
	f.FullVar(p, name, "", "", "", usage)
}

// 定义一个[]int类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) IntsFull(name, title, key, env string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) IntsConfig(name, title, key string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) IntsEnv(name, env string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Ints(name string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]int类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) IntsFullHiddenVar(p *[]int, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newIntsValue(parseIntsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) IntsConfigHiddenVar(p *[]int, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newIntsValue(parseIntsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) IntsEnvHiddenVar(p *[]int, name string, env string, value string, usage string) {
	f.FullHiddenVar(newIntsValue(parseIntsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) IntsHiddenVar(p *[]int, name string, value string, usage string) {
	f.HiddenVar(newIntsValue(parseIntsDefault(value), p), name, usage)
}

// 定义一个[]int类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) IntsFullHidden(name, title, key, env string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) IntsConfigHidden(name, title, key string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) IntsEnvHidden(name, env string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) IntsHidden(name string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsHiddenVar(p, name, value, usage)
	return p
}

// ============================== FlagSet (Int64s) ==============================

// 定义一个[]int64类型的Flag，并设置其全名、标题、键、环境变量、默认值（逗号分隔字符串）和用法
func (f *FlagSet) Int64sFullVar(p *[]int64, name string, title, key, env string, value string, usage string) {
	f.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int64类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) Int64sConfigVar(p *[]int64, name string, title, key string, value string, usage string) {
	f.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int64类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) Int64sEnvVar(p *[]int64, name string, env string, value string, usage string) {
	f.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int64类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) Int64sVar(p *[]int64, name string, value string, usage string) {
	f.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]int64类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) Int64sFull(name, title, key, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int64类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Int64sConfig(name, title, key string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int64类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) Int64sEnv(name, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int64类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Int64s(name string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]int64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) Int64sFullHiddenVar(p *[]int64, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int64类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) Int64sConfigHiddenVar(p *[]int64, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int64类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) Int64sEnvHiddenVar(p *[]int64, name string, env string, value string, usage string) {
	f.FullHiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int64类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) Int64sHiddenVar(p *[]int64, name string, value string, usage string) {
	f.HiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, usage)
}

// 定义一个[]int64类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Int64sFullHidden(name, title, key, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int64类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Int64sConfigHidden(name, title, key string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int64类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Int64sEnvHidden(name, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int64类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) Int64sHidden(name string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sHiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag (Ints) ==============================

// 定义一个[]int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) IntsFullVar(p *[]int, name string, title, key, env string, value string, usage string) {
	f.FullVar(newIntsValue(parseIntsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) IntsConfigVar(p *[]int, name string, title, key string, value string, usage string) {
	f.FullVar(newIntsValue(parseIntsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) IntsEnvVar(p *[]int, name string, env string, value string, usage string) {
	f.FullVar(newIntsValue(parseIntsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) IntsVar(p *[]int, name string, value string, usage string) {
	f.FullVar(newIntsValue(parseIntsDefault(value), p), name, "", "", "", usage)
}

// 定义一个IntSlice类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) IntSliceVar(p *IntSlice, name string, value []int, usage string) {
	*p = IntSlice(value)
	f.FullVar(p, name, "", "", "", usage)
}

// 定义一个[]int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) IntsFull(name, title, key, env string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) IntsConfig(name, title, key string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) IntsEnv(name, env string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Ints(name string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) IntsFullHiddenVar(p *[]int, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newIntsValue(parseIntsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) IntsConfigHiddenVar(p *[]int, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newIntsValue(parseIntsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) IntsEnvHiddenVar(p *[]int, name string, env string, value string, usage string) {
	f.FullHiddenVar(newIntsValue(parseIntsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) IntsHiddenVar(p *[]int, name string, value string, usage string) {
	f.HiddenVar(newIntsValue(parseIntsDefault(value), p), name, usage)
}

// 定义一个[]int类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) IntsFullHidden(name, title, key, env string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) IntsConfigHidden(name, title, key string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) IntsEnvHidden(name, env string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) IntsHidden(name string, value string, usage string) *[]int {
	p := new([]int)
	f.IntsHiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag (Int64s) ==============================

// 定义一个[]int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Int64sFullVar(p *[]int64, name string, title, key, env string, value string, usage string) {
	f.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int64类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Int64sConfigVar(p *[]int64, name string, title, key string, value string, usage string) {
	f.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int64类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Int64sEnvVar(p *[]int64, name string, env string, value string, usage string) {
	f.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int64类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) Int64sVar(p *[]int64, name string, value string, usage string) {
	f.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Int64sFull(name, title, key, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Int64sConfig(name, title, key string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Int64sEnv(name, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Int64s(name string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64sFullHiddenVar(p *[]int64, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int64类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64sConfigHiddenVar(p *[]int64, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int64类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64sEnvHiddenVar(p *[]int64, name string, env string, value string, usage string) {
	f.FullHiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int64类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64sHiddenVar(p *[]int64, name string, value string, usage string) {
	f.HiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, usage)
}

// 定义一个[]int64类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64sFullHidden(name, title, key, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int64类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64sConfigHidden(name, title, key string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int64类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64sEnvHidden(name, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int64类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) Int64sHidden(name string, value string, usage string) *[]int64 {
	p := new([]int64)
	f.Int64sHiddenVar(p, name, value, usage)
	return p
}

// ============================== 包级函数（sysflag） ==============================

// Ints类型定义
// 定义一个[]int类型的标志，并设置其名称、标题、键、环境变量、默认值（逗号分隔）和用法
func IntsFullVar(p *[]int, name string, title, key, env string, value string, usage string) {
	sysflag.FullVar(newIntsValue(parseIntsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int类型的标志，并设置其名称、标题、键、默认值和用法
func IntsConfigVar(p *[]int, name string, title, key string, value string, usage string) {
	sysflag.FullVar(newIntsValue(parseIntsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int类型的标志，并设置其名称、环境变量、默认值和用法
func IntsEnvVar(p *[]int, name string, env string, value string, usage string) {
	sysflag.FullVar(newIntsValue(parseIntsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int类型的标志，并设置其名称、默认值和用法
func IntsVar(p *[]int, name string, value string, usage string) {
	sysflag.FullVar(newIntsValue(parseIntsDefault(value), p), name, "", "", "", usage)
}

// 定义一个IntSlice类型的标志，并设置其名称、默认值和用法
func IntSliceVar(p *IntSlice, name string, value []int, usage string) {
	sysflag.IntSliceVar(p, name, value, usage)
}

// 定义一个[]int类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func IntsFull(name, title, key, env string, value string, usage string) *[]int {
	p := new([]int)
	sysflag.IntsFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func IntsConfig(name, title, key string, value string, usage string) *[]int {
	p := new([]int)
	sysflag.IntsConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func IntsEnv(name, env string, value string, usage string) *[]int {
	p := new([]int)
	sysflag.IntsEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int类型的标志，并设置其名称、默认值和用法，并返回指针
func Ints(name string, value string, usage string) *[]int {
	p := new([]int)
	sysflag.IntsVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]int类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func IntsFullHiddenVar(p *[]int, name string, title, key, env string, value string, usage string) {
	sysflag.FullHiddenVar(newIntsValue(parseIntsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int类型的隐藏标志，并设置其名称、标题、键、默认值和用法
func IntsConfigHiddenVar(p *[]int, name string, title, key string, value string, usage string) {
	sysflag.FullHiddenVar(newIntsValue(parseIntsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int类型的隐藏标志，并设置其名称、环境变量、默认值和用法
func IntsEnvHiddenVar(p *[]int, name string, env string, value string, usage string) {
	sysflag.FullHiddenVar(newIntsValue(parseIntsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int类型的隐藏标志，并设置其名称、默认值和用法
func IntsHiddenVar(p *[]int, name string, value string, usage string) {
	sysflag.HiddenVar(newIntsValue(parseIntsDefault(value), p), name, usage)
}

// 定义一个[]int类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func IntsFullHidden(name, title, key, env string, value string, usage string) *[]int {
	p := new([]int)
	sysflag.IntsFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int类型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func IntsConfigHidden(name, title, key string, value string, usage string) *[]int {
	p := new([]int)
	sysflag.IntsConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int类型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func IntsEnvHidden(name, env string, value string, usage string) *[]int {
	p := new([]int)
	sysflag.IntsEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int类型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func IntsHidden(name string, value string, usage string) *[]int {
	p := new([]int)
	sysflag.IntsHiddenVar(p, name, value, usage)
	return p
}

// Int64s类型定义
// 定义一个[]int64类型的标志，并设置其名称、标题、键、环境变量、默认值（逗号分隔）和用法
func Int64sFullVar(p *[]int64, name string, title, key, env string, value string, usage string) {
	sysflag.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int64类型的标志，并设置其名称、标题、键、默认值和用法
func Int64sConfigVar(p *[]int64, name string, title, key string, value string, usage string) {
	sysflag.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int64类型的标志，并设置其名称、环境变量、默认值和用法
func Int64sEnvVar(p *[]int64, name string, env string, value string, usage string) {
	sysflag.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int64类型的标志，并设置其名称、默认值和用法
func Int64sVar(p *[]int64, name string, value string, usage string) {
	sysflag.FullVar(newInt64sValue(parseInt64sDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]int64类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func Int64sFull(name, title, key, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	sysflag.Int64sFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int64类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func Int64sConfig(name, title, key string, value string, usage string) *[]int64 {
	p := new([]int64)
	sysflag.Int64sConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int64类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func Int64sEnv(name, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	sysflag.Int64sEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int64类型的标志，并设置其名称、默认值和用法，并返回指针
func Int64s(name string, value string, usage string) *[]int64 {
	p := new([]int64)
	sysflag.Int64sVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]int64类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func Int64sFullHiddenVar(p *[]int64, name string, title, key, env string, value string, usage string) {
	sysflag.FullHiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]int64类型的隐藏标志，并设置其名称、标题、键、默认值和用法
func Int64sConfigHiddenVar(p *[]int64, name string, title, key string, value string, usage string) {
	sysflag.FullHiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]int64类型的隐藏标志，并设置其名称、环境变量、默认值和用法
func Int64sEnvHiddenVar(p *[]int64, name string, env string, value string, usage string) {
	sysflag.FullHiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]int64类型的隐藏标志，并设置其名称、默认值和用法
func Int64sHiddenVar(p *[]int64, name string, value string, usage string) {
	sysflag.HiddenVar(newInt64sValue(parseInt64sDefault(value), p), name, usage)
}

// 定义一个[]int64类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func Int64sFullHidden(name, title, key, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	sysflag.Int64sFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]int64类型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func Int64sConfigHidden(name, title, key string, value string, usage string) *[]int64 {
	p := new([]int64)
	sysflag.Int64sConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]int64类型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func Int64sEnvHidden(name, env string, value string, usage string) *[]int64 {
	p := new([]int64)
	sysflag.Int64sEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]int64类型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func Int64sHidden(name string, value string, usage string) *[]int64 {
	p := new([]int64)
	sysflag.Int64sHiddenVar(p, name, value, usage)
	return p
}
