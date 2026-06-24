package flag

// stringsDefault 将默认值字符串按逗号切分为切片
func stringsDefault(value string) []string {
	if value == "" {
		return []string{}
	}
	return splitStrings(value)
}

// splitStrings 按逗号切分字符串
func splitStrings(value string) []string {
	if value == "" {
		return []string{}
	}
	parts := []string{}
	start := 0
	for i := 0; i < len(value); i++ {
		if value[i] == ',' {
			parts = append(parts, value[start:i])
			start = i + 1
		}
	}
	parts = append(parts, value[start:])
	return parts
}

// ============================== FlagSet ==============================

// Strings类型定义
// 定义一个[]string类型的Flag，并设置其全名、标题、键、环境变量、默认值（逗号分隔字符串）和用法
func (f *FlagSet) StringsFullVar(p *[]string, name string, title, key, env string, value string, usage string) {
	f.FullVar(newStringsValue(stringsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]string类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量
func (f *FlagSet) StringsConfigVar(p *[]string, name string, title, key string, value string, usage string) {
	f.FullVar(newStringsValue(stringsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]string类型的Flag，并设置其全名和环境变量，不设置标题和键
func (f *FlagSet) StringsEnvVar(p *[]string, name string, env string, value string, usage string) {
	f.FullVar(newStringsValue(stringsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]string类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量
func (f *FlagSet) StringsVar(p *[]string, name string, value string, usage string) {
	f.FullVar(newStringsValue(stringsDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]string类型的Flag，并设置其全名、标题、键、环境变量、默认值和用法，返回一个指向该Flag的指针
func (f *FlagSet) StringsFull(name, title, key, env string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]string类型的Flag，并设置其全名、标题、键和默认值，不设置环境变量，返回一个指向该Flag的指针
func (f *FlagSet) StringsConfig(name, title, key string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]string类型的Flag，并设置其全名和环境变量，不设置标题和键，返回一个指向该Flag的指针
func (f *FlagSet) StringsEnv(name, env string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]string类型的Flag，并设置其全名和默认值，不设置标题、键和环境变量，返回一个指向该Flag的指针
func (f *FlagSet) Strings(name string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]string类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法（隐藏）
func (f *FlagSet) StringsFullHiddenVar(p *[]string, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newStringsValue(stringsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]string类型的Flag，并设置其名称、标题、键、默认值和用法（隐藏）
func (f *FlagSet) StringsConfigHiddenVar(p *[]string, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newStringsValue(stringsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]string类型的Flag，并设置其名称、环境变量、默认值和用法（隐藏）
func (f *FlagSet) StringsEnvHiddenVar(p *[]string, name string, env string, value string, usage string) {
	f.FullHiddenVar(newStringsValue(stringsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]string类型的Flag，并设置其名称、默认值和用法（隐藏）
func (f *FlagSet) StringsHiddenVar(p *[]string, name string, value string, usage string) {
	f.HiddenVar(newStringsValue(stringsDefault(value), p), name, usage)
}

// 定义一个[]string类型的Flag，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) StringsFullHidden(name, title, key, env string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]string类型的Flag，并设置其名称、标题、键、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) StringsConfigHidden(name, title, key string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]string类型的Flag，并设置其名称、环境变量、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) StringsEnvHidden(name, env string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]string类型的Flag，并设置其名称、默认值和用法，并返回指针（隐藏）
func (f *FlagSet) StringsHidden(name string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsHiddenVar(p, name, value, usage)
	return p
}

// ============================== ArgsFlag ==============================

// 定义一个[]string类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) StringsFullVar(p *[]string, name string, title, key, env string, value string, usage string) {
	f.FullVar(newStringsValue(stringsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]string类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag）
func (f *ArgsFlag) StringsConfigVar(p *[]string, name string, title, key string, value string, usage string) {
	f.FullVar(newStringsValue(stringsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]string类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag）
func (f *ArgsFlag) StringsEnvVar(p *[]string, name string, env string, value string, usage string) {
	f.FullVar(newStringsValue(stringsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]string类型的变量，并设置其名称、默认值和用法（ArgsFlag）
func (f *ArgsFlag) StringsVar(p *[]string, name string, value string, usage string) {
	f.FullVar(newStringsValue(stringsDefault(value), p), name, "", "", "", usage)
}

// 定义一个[]string类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) StringsFull(name, title, key, env string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]string类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) StringsConfig(name, title, key string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]string类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) StringsEnv(name, env string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个[]string类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag）
func (f *ArgsFlag) Strings(name string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个[]string类型的变量，并设置其名称、标题、键、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) StringsFullHiddenVar(p *[]string, name string, title, key, env string, value string, usage string) {
	f.FullHiddenVar(newStringsValue(stringsDefault(value), p), name, title, key, env, usage)
}

// 定义一个[]string类型的变量，并设置其名称、标题、键、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) StringsConfigHiddenVar(p *[]string, name string, title, key string, value string, usage string) {
	f.FullHiddenVar(newStringsValue(stringsDefault(value), p), name, title, key, "", usage)
}

// 定义一个[]string类型的变量，并设置其名称、环境变量、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) StringsEnvHiddenVar(p *[]string, name string, env string, value string, usage string) {
	f.FullHiddenVar(newStringsValue(stringsDefault(value), p), name, "", "", env, usage)
}

// 定义一个[]string类型的变量，并设置其名称、默认值和用法（ArgsFlag，隐藏）
func (f *ArgsFlag) StringsHiddenVar(p *[]string, name string, value string, usage string) {
	f.HiddenVar(newStringsValue(stringsDefault(value), p), name, usage)
}

// 定义一个[]string类型的变量，并设置其名称、标题、键、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) StringsFullHidden(name, title, key, env string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个[]string类型的变量，并设置其名称、标题、键、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) StringsConfigHidden(name, title, key string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个[]string类型的变量，并设置其名称、环境变量、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) StringsEnvHidden(name, env string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个[]string类型的变量，并设置其名称、默认值和用法，并返回该变量的指针（ArgsFlag，隐藏）
func (f *ArgsFlag) StringsHidden(name string, value string, usage string) *[]string {
	p := new([]string)
	f.StringsHiddenVar(p, name, value, usage)
	return p
}

// ============================== 包级函数（sysflag） ==============================

// 定义一个字符串切片类型的标志，并设置其名称、标题、键、环境变量、默认值（逗号分隔）和用法
func StringsFullVar(p *[]string, name string, title, key, env string, value string, usage string) {
	sysflag.FullVar(newStringsValue(stringsDefault(value), p), name, title, key, env, usage)
}

// 定义一个字符串切片类型的标志，并设置其名称、标题、键、默认值和用法
func StringsConfigVar(p *[]string, name string, title, key string, value string, usage string) {
	sysflag.FullVar(newStringsValue(stringsDefault(value), p), name, title, key, "", usage)
}

// 定义一个字符串切片类型的标志，并设置其名称、环境变量、默认值和用法
func StringsEnvVar(p *[]string, name string, env string, value string, usage string) {
	sysflag.FullVar(newStringsValue(stringsDefault(value), p), name, "", "", env, usage)
}

// 定义一个字符串切片类型的标志，并设置其名称、默认值和用法
func StringsVar(p *[]string, name string, value string, usage string) {
	sysflag.FullVar(newStringsValue(stringsDefault(value), p), name, "", "", "", usage)
}

// 定义一个字符串切片类型的标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func StringsFull(name, title, key, env string, value string, usage string) *[]string {
	p := new([]string)
	sysflag.StringsFullVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个字符串切片类型的标志，并设置其名称、标题、键、默认值和用法，并返回指针
func StringsConfig(name, title, key string, value string, usage string) *[]string {
	p := new([]string)
	sysflag.StringsConfigVar(p, name, title, key, value, usage)
	return p
}

// 定义一个字符串切片类型的标志，并设置其名称、环境变量、默认值和用法，并返回指针
func StringsEnv(name, env string, value string, usage string) *[]string {
	p := new([]string)
	sysflag.StringsEnvVar(p, name, env, value, usage)
	return p
}

// 定义一个字符串切片类型的标志，并设置其名称、默认值和用法，并返回指针
func Strings(name string, value string, usage string) *[]string {
	p := new([]string)
	sysflag.StringsVar(p, name, value, usage)
	return p
}

// ----------------------------- Hidden -----------------------------

// 定义一个字符串切片类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法
func StringsFullHiddenVar(p *[]string, name string, title, key, env string, value string, usage string) {
	sysflag.FullHiddenVar(newStringsValue(stringsDefault(value), p), name, title, key, env, usage)
}

// 定义一个字符串切片类型的隐藏标志，并设置其名称、标题、键、默认值和用法
func StringsConfigHiddenVar(p *[]string, name string, title, key string, value string, usage string) {
	sysflag.FullHiddenVar(newStringsValue(stringsDefault(value), p), name, title, key, "", usage)
}

// 定义一个字符串切片类型的隐藏标志，并设置其名称、环境变量、默认值和用法
func StringsEnvHiddenVar(p *[]string, name string, env string, value string, usage string) {
	sysflag.FullHiddenVar(newStringsValue(stringsDefault(value), p), name, "", "", env, usage)
}

// 定义一个字符串切片类型的隐藏标志，并设置其名称、默认值和用法
func StringsHiddenVar(p *[]string, name string, value string, usage string) {
	sysflag.HiddenVar(newStringsValue(stringsDefault(value), p), name, usage)
}

// 定义一个字符串切片类型的隐藏标志，并设置其名称、标题、键、环境变量、默认值和用法，并返回指针
func StringsFullHidden(name, title, key, env string, value string, usage string) *[]string {
	p := new([]string)
	sysflag.StringsFullHiddenVar(p, name, title, key, env, value, usage)
	return p
}

// 定义一个字符串切片类型的隐藏标志，并设置其名称、标题、键、默认值和用法，并返回指针
func StringsConfigHidden(name, title, key string, value string, usage string) *[]string {
	p := new([]string)
	sysflag.StringsConfigHiddenVar(p, name, title, key, value, usage)
	return p
}

// 定义一个字符串切片类型的隐藏标志，并设置其名称、环境变量、默认值和用法，并返回指针
func StringsEnvHidden(name, env string, value string, usage string) *[]string {
	p := new([]string)
	sysflag.StringsEnvHiddenVar(p, name, env, value, usage)
	return p
}

// 定义一个字符串切片类型的隐藏标志，并设置其名称、默认值和用法，并返回指针
func StringsHidden(name string, value string, usage string) *[]string {
	p := new([]string)
	sysflag.StringsHiddenVar(p, name, value, usage)
	return p
}
