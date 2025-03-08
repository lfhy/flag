package cmd

import (
	"fmt"
	"strings"
)

type Cmd interface {
	// 定位命令
	Name() []string
	// 初始化命令
	Init(arg ...string) error
	// 短的帮助信息 用于命令行送显
	UsageShort() string
	// 自动生成的帮助信息
	Usage()
	// 执行
	Run() error
}

var Cmds []Cmd

func RunCmd(cmd string, arg ...string) error {
	isHelp := false
	if cmd == "help" {
		isHelp = true
	} else {
		for _, v := range arg {
			if v == "--help" || v == "-help" || v == "-h" {
				isHelp = true
				arg[0] = cmd
				break
			}
		}
	}

	if isHelp {
		// 无参数的情况
		if len(arg) == 0 {
			ListCommand()
			return nil
		}
		// 有参数查找指定参数
		hcmd := arg[0]
		for _, c := range Cmds {
			for _, alias := range c.Name() {
				if alias == hcmd {
					c.Init()
					c.Usage()
					return nil
				}
			}
		}
		return fmt.Errorf("未知命令 %s 使用 --help 查看帮助信息", hcmd)
	}

	for _, c := range Cmds {
		for _, alias := range c.Name() {
			if alias == cmd {
				err := c.Init(SortArgs(arg...)...)
				if err != nil {
					// c.Usage()
					return err
				}
				return c.Run()
			}
		}
	}
	return fmt.Errorf("未知命令 %s 使用 --help 查看帮助信息", cmd)
}

// 注册命令
func RegisterCommand(cmd Cmd) {
	Cmds = append(Cmds, cmd)
}

type HideModel interface {
	Hide() bool
}

const (
	listFormat = "  %-8s %-15s   %s\n"
)

// 命令列表
func ListCommand() {
	fmt.Println("子命令列表:")
	fmt.Printf(listFormat, "命令", "别名", "作用")
	for _, c := range Cmds {
		canHide, ok := c.(HideModel)
		if ok && canHide.Hide() {
			continue
		}
		if len(c.Name()) > 1 {
			fmt.Printf(listFormat, c.Name()[0], strings.Join(c.Name()[1:], ","), c.UsageShort())
		} else {
			fmt.Printf(listFormat, c.Name()[0], "", c.UsageShort())
		}

	}
	fmt.Printf("\n使用 help [子命令] 查看子命令详细信息\n")
}

func SortArgs(args ...string) []string {
	var res []string
	kv, other := ParseToKV(args...)
	for k, v := range kv {
		res = append(res, fmt.Sprintf("%v=%v", k, v))
	}
	res = append(res, other...)
	return res
}

func ParseToKV(args ...string) (map[string]string, []string) {
	kvargs := make(map[string]string)
	var otherArgs []string
	for index, arg := range args {
		if strings.HasPrefix(arg, "-") {
			// 判断是否有等号
			argInfo := strings.Split(arg, "=")
			if len(argInfo) >= 2 {
				kvargs[argInfo[0]] = argInfo[1]
				// 否则取下一个值
			} else {
				next := index + 1
				if len(args) > next {
					info := args[next]
					if strings.HasPrefix(info, "-") {
						kvargs[argInfo[0]] = "true"
					} else {
						kvargs[argInfo[0]] = info
					}
				} else {
					kvargs[argInfo[0]] = "true"
				}
			}
		} else {
			// 获取之前那个
			prev := index - 1
			if prev >= 0 {
				info := args[prev]
				if strings.HasPrefix(info, "-") && len(strings.Split(info, "=")) == 1 {
					continue
				}
			}
			otherArgs = append(otherArgs, arg)
		}
	}
	return kvargs, otherArgs
}
