package flag

import (
	"fmt"
	"sort"
	"strings"
)

type Cmd interface {
	Name() string
	Run(args ...string) error
}

type CmdInitWithError interface {
	Init(args ...string) error
}

type CmdInit interface {
	Init(args ...string)
}

type CmdUsage interface {
	Usage() string
}

type CmdHide interface {
	Hide() bool
}

type CmdHelp interface {
	Help() string
}

type CmdHelpFn interface {
	Help()
}

type CmdUsageFn interface {
	Usage()
}

type CmdPrintDefaults interface {
	PrintDefaults()
}

func (f *FlagSet) RegisterCommand(cmd Cmd) {
	if f.LookupCmd(cmd.Name()) {
		msg := fmt.Sprintf("命令重复定义: %s", cmd.Name())
		fmt.Fprintln(f.out(), msg)
		panic(msg)
	}
	f.cmds = append(f.cmds, cmd)
}

func (f *FlagSet) Run(args ...string) error {
	if !f.parsed {
		err := f.Parse(args)
		if err != nil {
			return err
		}

	}
	if f.ArgsLen() > 0 {
		return f.RunCmd(f.Args()[0], f.Args()[1:]...)
	}
	if f.LookupCmd("help") {
		return f.RunCmd("help")
	}
	return fmt.Errorf("使用 --help 查看帮助信息")
}

func (f *FlagSet) LookupCmd(cmd string) bool {
	_, _, ok := f.lookupCommand(cmd)
	return ok
}

func (f *FlagSet) RunCmd(cmd string, args ...string) error {
	isHelp := cmd == "help" && !f.LookupCmd("help")
	if isHelp {
		if len(args) == 0 {
			f.PrintDefaults()
			return nil
		}
		cmd = args[0]
	}
	realCmd, canonicalName, ok := f.lookupCommand(cmd)
	if !ok {
		return fmt.Errorf("未知命令 %s 使用 --help 查看帮助信息", cmd)
	}
	cmd = canonicalName

	for _, c := range f.cmds {
		if c != realCmd {
			continue
		}
		if isHelp {
			switch c := c.(type) {
			case CmdHelpFn:
				c.Help()
			case CmdUsageFn:
				c.Usage()
			case CmdHelp:
				fmt.Println(safeGetHelp(cmd, c.Help))
			case CmdUsage:
				fmt.Println(safeGetHelp(cmd, c.Usage))
			default:
				fmt.Println("运行", c.Name())
			}
			return nil
		}
		if cmdInit, ok := c.(CmdInitWithError); ok {
			err := cmdInit.Init(args...)
			if err != nil {
				// 子命令在解析 -h/--help 时已经输出过帮助信息了。
				if err == ErrHelp {
					return nil
				}
				switch c := c.(type) {
				case CmdHelp:
					fmt.Println(safeGetHelp(cmd, c.Help))
				case CmdUsage:
					fmt.Println(safeGetHelp(cmd, c.Usage))
				}
				pfn, ok := c.(CmdPrintDefaults)
				if ok {
					fmt.Println()
					pfn.PrintDefaults()
				}
				return err
			}
		} else if cmdInit, ok := c.(CmdInit); ok {
			cmdInit.Init(args...)
		}
		return c.Run(args...)
	}
	return fmt.Errorf("未知命令 %s 使用 --help 查看帮助信息", cmd)
}

func (f *FlagSet) AliasCmd(name, alias string) {
	if name == "" || alias == "" {
		msg := "命令别名不能为空"
		fmt.Fprintln(f.out(), msg)
		panic(msg)
	}
	if name == alias {
		return
	}
	if _, _, ok := f.lookupCommand(name); !ok {
		msg := fmt.Sprintf("命令不存在，无法设置别名: %s", name)
		fmt.Fprintln(f.out(), msg)
		panic(msg)
	}
	if _, _, ok := f.lookupCommand(alias); ok {
		msg := fmt.Sprintf("命令别名重复定义: %s", alias)
		fmt.Fprintln(f.out(), msg)
		panic(msg)
	}
	if f.cmdAliases == nil {
		f.cmdAliases = make(map[string]string)
	}
	f.cmdAliases[alias] = name
}

func (f *FlagSet) lookupCommand(name string) (Cmd, string, bool) {
	canonicalName := name
	if realName, ok := f.cmdAliases[name]; ok {
		canonicalName = realName
	}
	for _, c := range f.cmds {
		if c.Name() == canonicalName {
			return c, canonicalName, true
		}
	}
	return nil, "", false
}

func (f *FlagSet) commandAliases(name string) []string {
	var aliases []string
	for alias, canonicalName := range f.cmdAliases {
		if canonicalName == name {
			aliases = append(aliases, alias)
		}
	}
	sort.Strings(aliases)
	return aliases
}

func (f *FlagSet) commandDisplayName(name string) string {
	names := []string{name}
	names = append(names, f.commandAliases(name)...)
	return strings.Join(names, ", ")
}
