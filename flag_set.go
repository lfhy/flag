package flag

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type FlagSet struct {
	Usage func() //打印使用方法

	name          string // 参数名称
	parsed        bool   // 是否解析
	actual        map[string]*Flag
	formal        map[string]*Flag
	args          []string // 参数
	errorHandling ErrorHandling
	output        io.Writer //如果没有设置则输出至Stdout
	config        *Config
	cmds          []Cmd
}

func (f *FlagSet) GetConfig() *Config {
	// 避免无初始化
	if f.config == nil {
		config := newConfig("config.toml")
		f.config = &config
	}
	return f.config
}

func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
	f := &FlagSet{
		name:          name,
		errorHandling: errorHandling,
	}
	return f
}

// 设置解析错误时的处理方式
func (f *FlagSet) SetErrorHandling(h ErrorHandling) {
	f.errorHandling = h
}

// 打印默认参数变量
// PrintDefaults将所有的默认值打印到标准错误
func (f *FlagSet) PrintDefaults() {
	//遍历全部的参数变量
	f.VisitAll(func(flag *Flag) {
		if flag.Hidden {
			return
		}
		//-之前有两个空格进行区分参数
		s := fmt.Sprintf("  -%s", flag.Name) //打印参数名称
		//获取参数名称和用法
		name, usage := UnquoteUsage(flag)
		//如果名称大于0则添加
		if len(name) > 0 {
			s += " 参数值类型:" + name
		} else {
			s += " 参数值类型:bool"
		}

		//添加换行符
		// 一个ASCII字母的布尔标志非常常见，我们特别对待它们，将它们的用法放在同一行。
		if len(s) <= 4 { // space, space, '-', 'x'.
			s += "\t"
		} else {
			// 制表符之前的四个空格触发4和8空格制表符停止的良好对齐。
			s += "\n    \t"
		}

		//添加使用方法
		s += usage
		//如果不是空值
		if !isZeroValue(flag, flag.DefValue) {
			if _, ok := flag.Value.(*stringValue); ok {
				// 在值上加引号
				s += fmt.Sprintf(" (默认值: %q)", flag.DefValue)
			} else {
				s += fmt.Sprintf(" (默认值: %v)", flag.DefValue)
			}
		}

		//末尾添加换行 换两行
		fmt.Fprint(f.out(), s, "\n\n")
	})
	// cmd
	if len(f.cmds) > 0 {
		fmt.Fprint(f.out(), "可用命令:\n")
		cmdTable := newCmdTable("命令", "说明")
		for _, cmd := range f.cmds {
			cmdHide, ok := cmd.(CmdHide)
			if ok && cmdHide.Hide() {
				continue
			}

			var description string
			// 获取帮助信息
			switch cmd := cmd.(type) {
			case CmdUsage:
				description = cmd.Usage()
			case CmdUsage2:
				description = cmd.usage()
			case CmdHelp:
				description = cmd.Help()
			default:
				description = "运行 " + cmd.Name()
			}

			cmdTable.Add(cmd.Name(), description)
		}
		cmdTable.Print()
		fmt.Fprint(f.out(), "\n") // 添加额外换行以保持一致性
	}

}

// 设置参数
func (f *FlagSet) Set(name, value string) error {
	flag, ok := f.formal[name]
	if !ok {
		return fmt.Errorf("flag不存在 -%v", name)
	}
	err := flag.Value.Set(value)
	if err != nil {
		return err
	}
	if f.actual == nil {
		f.actual = make(map[string]*Flag)
	}
	f.actual[name] = flag
	return nil
}

// 设置输出
func (f *FlagSet) out() io.Writer {
	if f.output == nil {
		return os.Stdout
	}
	return f.output
}

func (f *FlagSet) SetOutput(output io.Writer) { f.output = output }

// 参数排序
func sortFlags(flags map[string]*Flag) []*Flag {
	list := make(sort.StringSlice, len(flags))
	i := 0
	for _, f := range flags {
		list[i] = f.Name
		i++
	}
	list.Sort()
	result := make([]*Flag, len(list))
	for i, name := range list {
		result[i] = flags[name]
	}
	return result
}

// 按顺序访问全部flag
func (f *FlagSet) VisitAll(fn func(*Flag)) {
	for _, flag := range sortFlags(f.formal) {
		fn(flag)
	}
}

// 只访问有设置值的flag
func (f *FlagSet) Visit(fn func(*Flag)) {
	for _, flag := range sortFlags(f.actual) {
		fn(flag)
	}
}

// 查找参数是否存在
func (f *FlagSet) Lookup(name string) *Flag { return f.formal[name] }

// NFlag返回已设置的标志数
func (f *FlagSet) NFlag() int { return len(f.actual) }

// Arg返回第i个参数 第0个参数是自己
func (f *FlagSet) Arg(i int) string {
	if i < 0 || i >= len(f.args) {
		return ""
	}
	return f.args[i]
}

// ArgsLen函数用于获取参数的长度
func (f *FlagSet) ArgsLen() int { return len(f.args) }

// 打印帮助信息
func (f *FlagSet) usage() {
	if f.Usage == nil {
		defaultUsage(f)
	} else {
		f.Usage()
	}
}

// 打印帮助并结束
func (f *FlagSet) failf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	fmt.Fprintln(f.out(), err)
	f.usage()
	return err
}

// 定义默认的配置文件标志名
var DefaultConfigFlagName = "c"

// 解析参数
func (f *FlagSet) Parse(arguments []string) error {
	// 如果没有定义默认的配置文件标志名，则添加一个隐藏的标志名
	if _, ok := f.formal[DefaultConfigFlagName]; !ok {
		f.StringHidden(DefaultConfigFlagName, "", "默认读取的配置文件 如果参数没有值则会读取配置文件中的值")
	}

	// 标记已经解析过
	f.parsed = true
	// 保存参数
	f.args = arguments
	// 循环解析参数
	for {
		seen, err := f.parseOne()
		if seen {
			continue
		}
		if err == nil {
			break
		}
		// 根据错误处理方式处理错误
		switch f.errorHandling {
		case ContinueOnError:
			return err
		case ExitOnError:
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}

	// 读取配置文件
	var cFile string
	// 如果定义了默认的配置文件标志名，则读取其值
	if cf := f.formal[DefaultConfigFlagName]; cf != nil {
		cFile = cf.Value.String()
	}
	// 如果实际解析时定义了默认的配置文件标志名，则读取其值
	if cf := f.actual[DefaultConfigFlagName]; cf != nil {
		cFile = cf.Value.String()
	}

	// 如果没有定义默认的配置文件标志名，则从未解析的参数中查找
	if cFile == "" {
		cFile = f.findConfigArgInUnresolved()
	}

	// 如果找到了配置文件，则进行解析
	if cFile != "" {
		_, err := os.Stat(cFile)
		// 如果文件存在再进行解析
		if err == nil {
			if err := f.ParseFile(cFile, false); err != nil {
				// 根据错误处理方式处理错误
				switch f.errorHandling {
				case ContinueOnError:
					return err
				case ExitOnError:
					os.Exit(2)
				case PanicOnError:
					panic(err)
				}
				return err
			}
		}
	}

	// 应用到环境变量
	f.parseEnv()
	return nil
}

// 从配置文件读取
// 参数1:配置文件路径
// 参数2:覆盖传参
func (f *FlagSet) ParseFile(config string, rewritevalue bool) error {
	fconfig := newConfig(config)
	f.config = &fconfig
	for name, flag := range f.formal {
		//忽略已设置的参数；参数优先于文件
		if f.actual[name] != nil && !rewritevalue {
			continue
		}
		if name == "help" || name == "h" {
			f.usage()
			return ErrHelp
		}

		// 为空不做解析
		title := flag.ConfigTitle
		key := flag.ConfigKey
		if title == "" && key == "" {
			continue
		}
		value := fconfig.ReadConfigToString(title, key)
		if value == "" {
			continue
		}

		if err := flag.Value.Set(value); err != nil {
			return f.failf("无效的值 %q 配置文件参数 %s 错误: %v", value, name, err)
		}

		// 更新有效参数
		if f.actual == nil {
			f.actual = make(map[string]*Flag)
		}
		f.actual[name] = flag
	}
	return nil
}

// 定义一个错误变量，表示帮助请求
var ErrHelp = errors.New("flag: help requested")

// 解析
// 返回1:是否定义
// 返回2:是否出错
func (f *FlagSet) parseOne() (bool, error) {
	if len(f.args) == 0 {
		return false, nil
	}
	s := f.args[0]
	if len(s) < 2 || s[0] != '-' {
		return false, nil
	}
	numMinuses := 1
	if s[1] == '-' {
		numMinuses++
		if len(s) == 2 { // "--" terminates the flags
			f.args = f.args[1:]
			return false, nil
		}
	}
	name := s[numMinuses:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return false, f.failf("错误的输入: %s", s)
	}

	// 忽略go test解析
	if strings.HasPrefix(name, "test.") {
		return false, nil
	}

	// it's a flag. does it have an argument?
	f.args = f.args[1:]
	hasValue := false
	value := ""
	for i := 1; i < len(name); i++ { // equals cannot be first
		if name[i] == '=' {
			value = name[i+1:]
			hasValue = true
			name = name[0:i]
			break
		}
	}
	m := f.formal
	flag, alreadythere := m[name] // BUG
	if !alreadythere {
		if name == "help" || name == "h" { // special case for nice help message.
			f.usage()
			return false, ErrHelp
		}
		return false, f.failf("不存在该参数: -%s\n", name)
	}

	// 判断bool的情况
	if fv, ok := flag.Value.(boolFlag); ok && fv.IsBoolFlag() {
		if hasValue {
			if err := fv.Set(value); err != nil {
				return false, f.failf("无效的 boolean 值 %q for -%s: %v", value, name, err)
			}
		} else {
			if err := fv.Set("true"); err != nil {
				return false, f.failf("无效的 boolean 参数 %s: %v", name, err)
			}
		}
	} else {
		// It must have a value, which might be the next argument.
		if !hasValue && len(f.args) > 0 {
			// value is the next arg
			hasValue = true
			value, f.args = f.args[0], f.args[1:]
		}
		if !hasValue {
			return false, f.failf("flag 需要一个参数: -%s", name)
		}
		if err := flag.Value.Set(value); err != nil {
			return false, f.failf("无效的值 %q 参数: -%s: %v", value, name, err)
		}
	}

	if f.actual == nil {
		f.actual = make(map[string]*Flag)
	}
	f.actual[name] = flag
	return true, nil
}

// 定义类型
// 参数1:值
// 参数2:解析的参数flag
// 参数3:解析的配置文件父类
// 参数4:解析的配置文件key
// 参数5:环境变量名称
// 参数6:使用方法
func (f *FlagSet) FullVar(value Value, name string, title, key string, env string, usage string) {
	f.VarFlag(value, name, title, key, env, false, usage)
}

func (f *FlagSet) FullHiddenVar(value Value, name string, title, key string, env string, usage string) {
	f.VarFlag(value, name, title, key, env, true, usage)
}

// 定义Flag
// 参数1:值
// 参数2:解析的参数flag
// 参数3:解析的配置文件父类
// 参数4:解析的配置文件key
// 参数5:环境变量名称
// 参数6:是否隐藏
// 参数7:使用方法
func (f *FlagSet) VarFlag(value Value, name string, title, key string, env string, hidden bool, usage string) {
	flag := &Flag{Name: name, Usage: usage, Value: value, DefValue: value.String(), ConfigTitle: title, ConfigKey: key, EnvName: env, Hidden: hidden}
	_, alreadythere := f.formal[name]
	if alreadythere {
		var msg string
		if f.name == "" {
			msg = fmt.Sprintf("flag重复定义: %s", name)
		} else {
			msg = fmt.Sprintf("%s flag已经被定义: %s", f.name, name)
		}
		fmt.Fprintln(f.out(), msg)
		panic(msg)
	}
	if f.formal == nil {
		f.formal = make(map[string]*Flag)
	}
	f.formal[name] = flag
}

// 定义隐藏类型类型
// 参数1:值
// 参数2:解析的参数flag
// 参数3:使用方法
func (f *FlagSet) HiddenVar(value Value, name string, usage string) {
	f.FullHiddenVar(value, name, "", "", "", usage)
}

// 查找是否配置文件没有被解析
func (f *FlagSet) findConfigArgInUnresolved() string {
	configArg := "-" + DefaultConfigFlagName
	for i := 0; i < len(f.args); i++ {
		if strings.HasPrefix(f.args[i], configArg) {
			if f.args[i] == configArg && i+1 < len(f.args) {
				return f.args[i+1]
			}

			if strings.HasPrefix(f.args[i], configArg+"=") {
				return f.args[i][len(configArg)+1:]
			}
		}
	}
	return ""
}

// 应用到环境变量
func (f *FlagSet) parseEnv() {
	for _, f2 := range f.formal {
		if f2 != nil && f2.EnvName != "" {
			os.Setenv(f2.EnvName, f2.Value.String())
		}
	}
}

// 返回参数
func (f *FlagSet) Args() []string { return f.args }

// 打印全部参数值
func (f *FlagSet) PrintAll() {
	tabs := newCmdTable("参数", "当前值", "配置文件参数", "环境变量名", "隐藏参数", "默认值")
	for _, f2 := range sortFlags(f.formal) {
		if f2 != nil {
			configName := f2.ConfigTitle + "." + f2.ConfigKey
			if configName == "." {
				configName = ""
			}
			value := f2.Value.String()
			if len(strings.Split(value, ",")) > 1 {
				value = strings.Join(strings.Split(value, ","), "\n")
			}
			tabs.Add(f2.Name, value, configName, f2.EnvName, fmt.Sprint(f2.Hidden), f2.DefValue)
		}
	}
	tabs.Print()
}
