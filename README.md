# flag

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

一个增强版的 Go 命令行参数解析库，兼容标准库 `flag` 的用法，并在此基础上扩展了配置文件、环境变量、隐藏参数、子命令、别名、字符串/数值切片等能力。

> 📖 [English Documentation](./README.en.md)

## ✨ 特性

- **多来源优先级**：命令行参数 > 环境变量 > 配置文件 > 默认值
- **🌟 统一配置入口 `-c`**：`Parse()` 自动注册隐藏参数 `-c`，传入配置文件后命令行/环境变量/配置文件按优先级一次性合并，无需手动调用 viper
- **配置文件集成**：基于 [viper](https://github.com/spf13/viper)，支持 JSON / TOML / YAML 等格式，参数变更可回写配置
- **隐藏参数**：支持注册不在 `--help` 中显示的参数，但仍可被解析
- **子命令 & 别名**：支持注册子命令（如 `git` 风格），并为参数/子命令设置别名（如 `-v` ↔ `--verbose`）
- **位置参数与选项可交错**：`FlagSet.Parse` 支持选项写在位置参数前后；布尔选项不会吞掉路径，`--` 后内容始终按字面位置参数保留。顶层子命令分发可使用 `ParseStandard` 在首个位置参数处停止
- **切片类型**：内置 `Strings / Ints / Int64s / Uints / Uint64s` 等切片类型，逗号分隔解析、多次传参自动累加
- **反射通用入口 `Var()`**：一个 API 覆盖所有基础类型和切片类型
- **彩色帮助输出**：内置 ANSI 彩色打印，`PrintAll` 以表格形式列出全部参数

## 📦 安装

```bash
go get github.com/lfhy/flag
```

## 🚀 快速开始

```go
package main

import (
    "fmt"
    "github.com/lfhy/flag"
)

var (
    Port    int
    Host    string
    Tags    []string
    Debug   bool
)

func init() {
    flag.IntVar(&Port, "port", 8080, "服务端口")
    flag.StringVar(&Host, "host", "127.0.0.1", "监听地址")
    flag.StringsVar(&Tags, "tags", "go,linux", "标签（逗号分隔）")
    flag.BoolVar(&Debug, "debug", false, "调试模式")
    flag.Parse()
}

func main() {
    fmt.Printf("Port=%d Host=%s Tags=%v Debug=%v\n", Port, Host, Tags, Debug)
}
```

运行：

```bash
$ go run main.go -port 9000 -tags a,b -tags c
Port=9000 Host=127.0.0.1 Tags=[a b c] Debug=false
```

## 📚 使用指南

### 1. 参数定义方式

包级别提供了与标准库一致风格的 `XxxVar` / `Xxx` 函数；也可通过 `FlagSet` 在子命令中独立使用。

| 方式 | 说明 |
|---|---|
| `flag.IntVar(p, "n", 1, "...")` | 基础类型，仅命令行 + 默认值 |
| `flag.IntEnvVar(p, "n", "APP_N", 1, "...")` | 同时绑定环境变量 |
| `flag.IntConfigVar(p, "n", "server", "port", 1, "...")` | 同时绑定配置文件 |
| `flag.IntFullVar(p, "n", "server", "port", "APP_N", 1, "...")` | 全部来源 |
| `flag.Int("n", 1, "...")` | 返回 `*int`，无需预先定义变量 |
| 函数名加 `Hidden`（如 `IntHiddenVar`） | 注册为隐藏参数 |

> 所有类型（Bool / String / Int / Int64 / Uint / Uint64 / Float64 / Duration 以及对应的切片 Strings/Ints/... ）均遵循同一套命名规则。

### 2. 切片类型

切片类型支持逗号分隔解析，**多次传参会累加**（与 `pflag` 的 `StringSlice` 行为一致）：

```go
var ids []int
flag.IntsVar(&ids, "ids", "1,2,3", "ID 列表")
// -ids=4,5 -ids=6  →  [1 2 3 4 5 6]
```

支持：`StringsVar` / `IntsVar` / `Int64sVar` / `UintsVar` / `Uint64sVar`。

### 3. 通用反射入口 Var()

`Var()` 接收一个 `FlagVar` 结构体，通过反射自动识别基础类型和切片类型，无需为每种类型单独调用 API：

```go
var (
    name string
    age  int
    tags []string
    ids  []int64
)

flag.Var(&flag.FlagVar{Value: &name, Name: "name", Env: "APP_NAME", DefaultValue: "default"})
flag.Var(&flag.FlagVar{Value: &age, Name: "age", DefaultValue: 18})
flag.Var(&flag.FlagVar{Value: &tags, Name: "tags", DefaultValue: "a,b,c"})
flag.Var(&flag.FlagVar{Value: &ids, Name: "ids", DefaultValue: []int64{1, 2}})
```

`Var()` 支持的类型：`bool / string / int* / uint* / float64 / time.Duration` 以及对应的切片。

### 4. 环境变量与配置文件（手动绑定）

```go
// 绑定环境变量：未在命令行指定时从环境变量读取
flag.StringEnvVar(&host, "host", "APP_HOST", "0.0.0.0", "监听地址")

// 绑定配置文件：从 title.key 读取
flag.IntConfigVar(&port, "port", "server", "port", 8080, "服务端口")

// 读取/回写配置
cfg := flag.GetConfig()
cfg.WriteToConfig("server", "port", 9000)
```

优先级：**命令行 > 环境变量 > 配置文件 > 默认值**。

### 5. 🌟 统一配置入口：`-c` 指定配置文件（特有能力）

这是本库区别于标准库 `flag` 的核心特性。`Parse()` 时会自动注册一个隐藏参数 `-c`，传入配置文件路径后，**命令行参数、配置文件、环境变量会在一次解析中按优先级统一合并**，无需手动调用 viper。

```go
// app.go
var (
    Host string
    Port int
)

func init() {
    flag.StringVar(&Host, "host", "127.0.0.1", "监听地址")
    flag.IntConfigVar(&Port, "port", "server", "port", 8080, "服务端口")
    flag.Parse()
}
```

假设有配置文件 `config.toml`：

```toml
[server]
port = 9000
```

三种来源合并：

```bash
# 1. 完全使用配置文件
$ go run app.go -c config.toml
# Host=127.0.0.1（默认值）  Port=9000（来自配置文件）

# 2. 命令行覆盖配置文件
$ go run app.go -c config.toml -port 8888
# Port=8888（命令行 > 配置文件）

# 3. 环境变量也参与（未在命令行/配置中出现时生效）
$ APP_HOST=0.0.0.0 go run app.go -c config.toml
# Host=0.0.0.0（环境变量 > 默认值）
```

要点：

- `-c` 是自动注册的隐藏参数，无需手动定义；未指定时跳过文件解析
- 可通过 `fs.SetConfigFlagName("conf")` 自定义参数名（FlagSet 级别，互不影响，默认 `"c"`）
- 配置文件按 `ConfigTitle.ConfigKey`（即 `IntConfigVar` 第二、三个参数）的节/键读取
- 可通过 `flag.DefaultConfigFlagName` 修改默认标志名（默认 `"c"`）
- 解析完成后用 `flag.GetConfig()` 获取 `*Config`，支持读写回文件

### 6. 别名

```go
flag.String("verbose", "false", "详细输出")
flag.Alias("verbose", "v")     // -v 等价于 --verbose
```

### 7. 隐藏参数

```go
flag.StringHiddenVar(&token, "token", "", "内部鉴权 Token")
// 不会出现在 --help 输出中，但 -token=xxx 仍然生效
```

### 8. 子命令

实现 `Cmd` 接口（`Name / Init / Run / Help`）后注册即可：

```go
type ServeCmd struct { *flag.FlagSet }

func (*ServeCmd) Name() string { return "serve" }
func (s *ServeCmd) Init(args ...string) error {
    s.FlagSet = flag.NewFlagSet("serve", flag.ContinueOnError)
    // ... 定义子命令参数
    return s.Parse(args)
}
func (*ServeCmd) Run(args ...string) error { /* ... */ return nil }
func (*ServeCmd) Help() string             { return "启动服务" }

flag.RegisterCommand(&ServeCmd{})
flag.Parse()
```

## 📖 API 速览

每个类型在每个层级（`FlagSet` 方法 / `ArgsFlag` 方法 / 包级函数）都提供 8 个函数：

```
XxxFullVar / XxxConfigVar / XxxEnvVar / XxxVar      // 写入已定义变量
XxxFull    / XxxConfig    / XxxEnv    / Xxx          // 返回 *T
```

加上对应 8 个 `Hidden` 变体。覆盖类型：

| 基础类型 | 切片类型 |
|---|---|
| Bool / String | Strings |
| Int / Int64 | Ints / Int64s |
| Uint / Uint64 | Uints / Uint64s |
| Float64 | — |
| Duration | — |

## 🧪 测试

```bash
go test ./...
```

## 📄 License

MIT
