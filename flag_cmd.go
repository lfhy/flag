package flag

import (
	"fmt"
)

type Cmd interface {
	Name() string
	Init(args ...string) error
	Run(args ...string) error
}

type CmdUsage interface {
	Usage() string
}

type CmdHide interface {
	Hide() bool
}

type CmdPrintDefaults interface {
	PrintDefaults()
}

// 内部使用
type CmdUsage2 interface {
	usage() string
}

type CmdHelp interface {
	Help() string
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
				case CmdUsage:
					fmt.Println(c.Usage())
				case CmdPrintDefaults:
					c.PrintDefaults()
				case CmdUsage2:
					fmt.Println(c.usage())
				case CmdHelp:
					fmt.Println(c.Help())
				default:
					fmt.Println("运行", c.Name())
				}
				return nil
			}
			err := c.Init(args...)
			if err != nil {
				switch c := c.(type) {
				case CmdUsage:
					fmt.Println(c.Usage())
				case CmdUsage2:
					fmt.Println(c.usage())
				case CmdHelp:
					fmt.Println(c.Help())
				}
				return err
			}
			return c.Run(args...)
		}
	}
	return fmt.Errorf("未知命令 %s 使用 --help 查看帮助信息", cmd)
}
