package flag

import (
	"fmt"
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
	for _, c := range f.cmds {
		if c.Name() == cmd {
			return true
		}
	}
	return false
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

	for _, c := range f.cmds {
		if c.Name() == cmd {
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
	}
	return fmt.Errorf("未知命令 %s 使用 --help 查看帮助信息", cmd)
}
