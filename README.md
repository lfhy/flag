# flag

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

一个增强版的 Go 命令行参数解析库，兼容标准库 `flag` 的用法，并在此基础上扩展了配置文件、环境变量、隐藏参数、子命令、别名、字符串/数值切片等能力。

> 📖 [English Documentation](./README.en.md)

## ✨ 特性

- **多来源优先级**：命令行参数 > 环境变量 > 配置文件 > 默认值
- **🌟 统一配置入口 `-c`**：`Parse()` 自动注册隐藏参数 `-c`；建议先用 `ConfigFile` 启用默认路径和严格文件读错，再按命令行/环境变量/配置文件优先级合并
- **配置文件集成**：基于 [viper](https://github.com/spf13/viper)，支持 JSON / TOML / YAML 等格式，参数变更可回写配置
- **隐藏参数**：支持注册不在 `--help` 中显示的参数，但仍可被解析
- **子命令 & 别名**：支持注册子命令（如 `git` 风格），并为参数/子命令设置别名（如 `-v` ↔ `--verbose`）
- **位置参数与选项可交错**：`FlagSet.Parse` 支持选项写在位置参数前后；布尔选项不会吞掉路径，`--` 后内容始终按字面位置参数保留。顶层子命令分发可使用 `ParseStandard` 在首个位置参数处停止
- **切片类型**：内置 `Strings / Ints / Int64s / Uints / Uint64s` 等切片类型，逗号分隔解析、多次传参自动累加
- **反射通用入口 `Var()`**：一个 API 覆盖受支持的基础类型及字符串/整数切片
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
    "os"

    "github.com/lfhy/flag"
)

func main() {
    fs := flag.NewFlagSet("app", flag.ContinueOnError)
    var port int
    var tags []string
    fs.Var(&flag.FlagVar{Value: &port, Name: "port", DefaultValue: 8080, Usage: "服务端口"})
    fs.Var(&flag.FlagVar{Value: &tags, Name: "tags", Usage: "标签，可重复"})
    host := fs.New("host").Default("127.0.0.1").Usage("监听地址").String()
    debug := fs.New("debug").Default(false).Usage("调试模式").Bool()

    if err := fs.Parse(os.Args[1:]); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(2)
    }
    fmt.Printf("Port=%d Host=%s Tags=%v Debug=%v\n", port, *host, tags, *debug)
}
```

运行：

```bash
$ go run main.go -port 9000 -tags a,b -tags c
Port=9000 Host=127.0.0.1 Tags=[a b c] Debug=false
```

## 📚 使用指南

### 1. 推荐的参数定义方式

优先在独立 `FlagSet` 上用 `fs.Var(&flag.FlagVar{...})` 定义配置丰富的参数（名称、别名、默认值、环境变量、配置节/键、帮助与隐藏状态），或用 `fs.New(...).Default(...).Config(...).Env(...).Usage(...).Int()` 等链式写法；已有变量也可在链末 `.Var(&value)`。两者都在运行期注册并由 `fs.Parse(os.Args[1:])` 解析。包级 `flag.Var` / `flag.New` 使用默认 `FlagSet`，适合确实需要共享全局状态的程序。

#### 链式定义（Fluent API）

`flag.New`（或 `fs.New`）可链式设置别名、默认值、配置键、环境变量和帮助文字。以类型方法结尾会立即注册参数并返回指针；`Parse` 会把最终值写入该指针：

```go
fs := flag.NewFlagSet("app", flag.ContinueOnError)
follow := fs.New("follow").Alias("f").Default(false).
    Config("title.key").Env("APP_FOLLOW").Usage("跟随更新").Bool()
if err := fs.Parse(os.Args[1:]); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(2) }
fmt.Println(*follow) // 根据 -follow/-f、APP_FOLLOW、配置文件或默认值确定
```

`.Config("title.key")` 映射配置文件的 `[title]` 节和 `key` 键；未设置 `ConfigFile` 时，要读取文件需在命令行传 `-c config.toml`（或通过 `fs.SetConfigFlagName("conf")` 改用 `-conf`）。环境变量是**输入来源**，解析不会将值写回环境变量。优先级为 **命令行 > 环境变量 > 配置文件 > 默认值**。

已有变量可用 `.Var(&data)` 绑定；它会立即注册参数，不返回指针：

```go
var data bool
fs.New("enabled").Default(false).Var(&data)
if err := fs.Parse(os.Args[1:]); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(2) }
fmt.Println(data)
```

也可以不调用类型方法或 `.Var`，直接保留定义并在解析后读取：

```go
fs := flag.NewFlagSet("app", flag.ContinueOnError)
definition := fs.New("follow").Default(false).Env("FOLLOW")
count := fs.New("count").Default("3") // 未以类型方法结尾：默认值保留为文本
if err := fs.Parse([]string{"-follow", "-count=5"}); err != nil { panic(err) }
fmt.Println(definition.Value().Bool()) // true（值类型为 bool，不是 *bool）
fmt.Println(count.Value().Int(), count.Value().String()) // 5 5
```

无类型结尾的定义由 `Parse` 注册为文本值，`.Value().Bool()/Int()/Int64()/String()` 等方法在读取时转换；文本不能转换时，布尔值与数值读取方法会 **panic**，而 `.Value().String()` 原样返回文本。类型结尾的 `.Bool()/Int()/Int64()/String()` 等方法则返回对应类型的指针（`*bool`、`*int`、`*int64`、`*string` 等）。

`FlagSet` 会保留已解析的参数状态；如果要用另一组环境变量或配置独立解析，请新建 `FlagSet`。

### 2. 切片类型

切片类型支持逗号分隔解析，**多次传参会累加**；非空默认切片也参与追加，因此需要显式值独占列表时请用空默认值：

```go
var ids []int
fs := flag.NewFlagSet("app", flag.ContinueOnError)
fs.Var(&flag.FlagVar{Value: &ids, Name: "ids", DefaultValue: []int{1, 2}, Usage: "ID 列表"})
if err := fs.Parse([]string{"-ids=4,5", "-ids=6"}); err != nil { panic(err) }
fmt.Println(ids) // [1 2 4 5 6]
```

兼容接口 `StringsVar` / `IntsVar` / `Int64sVar` / `UintsVar` / `Uint64sVar` 的默认值参数是逗号分隔字符串，同样会追加；反射 `FlagVar` 只支持 `[]string`、`[]int*`、`[]uint*`，不支持任意切片。

### 3. 通用反射入口 Var()

`Var()` 接收一个 `FlagVar` 结构体，集中设置 `Value`、`Name`、`Aliases`、`DefaultValue`、`Env`、`ConfigSection`、`ConfigKey`、`Usage`、`Hidden`，通过反射绑定受支持的类型。优先传入非 nil 变量指针；相较多个 `XxxVar` 调用，更适合需要多来源和别名的定义：

```go
fs := flag.NewFlagSet("app", flag.ContinueOnError)
var port int
var ids []int64
fs.Var(&flag.FlagVar{
    Value: &port, Name: "port", Aliases: []string{"p"}, DefaultValue: 8080,
    ConfigSection: "server", ConfigKey: "port", Env: "APP_PORT", Usage: "服务端口",
})
fs.Var(&flag.FlagVar{Value: &ids, Name: "ids", DefaultValue: []int64{1, 2}})
if err := fs.Parse(os.Args[1:]); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(2) }
```

`Var()` 支持 `bool / string / int* / uint* / float32 / float64 / time.Duration` 以及字符串、有符号或无符号整数切片。若省略配置节/键，`FlagVar` 会隐式推导：普通名称对应 `config.<名称>`，带连字符名称可能产生意外键（例如 `server-port` 对应 `server.server-port`）；需要稳定映射时显式指定 `ConfigSection` 和 `ConfigKey`，或用链式 `.Config("server.port")`。`ContinueOnError` 下部分注册错误只在注册时输出，**不保证**随后由 `Parse` 返回；检查参数定义本身，勿把解析成功当作注册已成功的证明。

### 4. 环境变量与配置文件

优先在同一 `FlagVar` 或链式定义上同时绑定环境变量、配置键和命令行选项，无需重复解析或单独读取配置：

```go
var host string
fs.Var(&flag.FlagVar{
    Value: &host, Name: "host", Env: "APP_HOST", DefaultValue: "127.0.0.1",
    ConfigSection: "server", ConfigKey: "host", Usage: "监听地址",
})
// 如果需要更新文件，解析后可通过 fs.GetConfig().WriteToConfig("server", "host", host) 回写
```

优先级：**命令行 > 环境变量 > 配置文件 > 默认值**。`XxxVar` / `XxxEnvVar` / `XxxConfigVar` / `XxxFullVar` 以及返回指针的对应 `Xxx` 系列仍供兼容调用；例如 `fs.IntFullVar(&port, "port", "server", "port", "APP_PORT", 8080, "服务端口")`，最后三个来源配置按节、键、环境变量依次绑定，随后才是默认值和帮助文本。

### 5. 🌟 统一配置入口：默认文件或 `-c` 指定配置文件（特有能力）

本库的 `FlagSet.Parse` 会自动注册隐藏的 `-c`。需要默认配置路径及可靠的读错误时，**显式调用 `fs.ConfigFile(...)` 启用文件加载**；命令行、环境变量与文件随后按优先级合并，不需要手动调用 viper。未调用 `ConfigFile` 时也可传 `-c`，但这是兼容行为，不能用于严格校验文件读错；调用 `GetConfig()` 还可能创建默认文件。

```go
// app.go
package main

import (
    "fmt"
    "os"

    "github.com/lfhy/flag"
)

func main() {
    fs := flag.NewFlagSet("app", flag.ContinueOnError)
    fs.ConfigFile("config.toml")
    var host string
    var port int
    fs.Var(&flag.FlagVar{Value: &host, Name: "host", Env: "APP_HOST", DefaultValue: "127.0.0.1", Usage: "监听地址"})
    fs.New("port").Default(8080).Config("server.port").Var(&port)
    if err := fs.Parse(os.Args[1:]); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(2)
    }
    fmt.Printf("Host=%s Port=%d\n", host, port)
}
```

假设有配置文件 `config.toml`：

```toml
[server]
port = 9000
```

调用 `ConfigFile` 后，同一程序按以下方式合并来源：

```bash
$ go run app.go
Host=127.0.0.1 Port=9000
$ go run app.go -port 8888
Host=127.0.0.1 Port=8888
$ APP_HOST=0.0.0.0 go run app.go
Host=0.0.0.0 Port=9000
```

已调用 `fs.ConfigFile("config.toml")`：未传 `-c` 时读取存在的默认文件，默认文件不存在时跳过；显式 `-c` 指定的文件不存在、文件不可读或现有文件解析失败时，`Parse` 返回错误。未调用 `ConfigFile` 时，显式 `-c` 指向不存在的文件会跳过读取；若路径已存在则进入旧的读取实现，可能忽略读取或解析错误。单独调用 `fs.GetConfig()` 时旧实现还可能创建其默认的 `config.toml`，并非缺失的 `-c` 路径一定会被创建。

`go run app.go` 使用 `config.toml`；`go run app.go -c other.toml` 则使用显式指定的文件。命令行参数 > 环境变量 > 配置文件 > 默认值的优先级不变。配置路径参数还可链式命名（以下为分别选用的写法）：

```go
fs := flag.NewFlagSet("app", flag.ContinueOnError)
fs.ConfigFile("config.toml").Alias("config") // 接受 -c 和 -config
// 若只需要 -config：改用 fs.ConfigFile("config.toml").Name("config")
```

`fs.SetConfigFlagName("conf")` 仍可在解析前改名；与 `.Name(...)` 同时使用时以最后一次设置为准。配置键由 `FlagVar.ConfigSection` / `ConfigKey` 或链式 `.Config("server.port")` 指定。解析后 `fs.GetConfig()` 可取得当前配置并写回；本节配置示例不执行回写。

#### 多个配置文件按层合并

服务公共配置与实例配置可用 `ConfigFiles` 统一加载，不需要自行读写 TOML：

~~~go
package main

import (
    "fmt"
    "os"

    "github.com/lfhy/flag"
)

func main() {
    fs := flag.NewFlagSet("gateway", flag.ContinueOnError)
    fs.ConfigFiles("config.toml", "gateway.toml", "gateway-7420.toml").
        Name("config").Alias("c")
    port := fs.IntConfig("port", "gateway", "port", 7420, "监听端口")
    peers := fs.StringsConfig("peers", "gateway", "peers", "", "逗号分隔网关地址")
    if err := fs.Parse(os.Args[1:]); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(2)
    }
    fmt.Printf("Port=%d Peers=%v\n", *port, *peers)
}
~~~

文件按参数顺序从低到高合并，优先级为 **命令行 > 环境变量 > 显式 `-c` 文件 > 最后一个默认文件 > 前面的默认文件 > 参数默认值**。例如 `-c override.toml` 为上例追加最高配置层，未出现的键仍继承公共文件。`ConfigFile` 保持原有“显式 `-c` 替换默认文件”的行为；再次调用 `ConfigFiles` 则替换整组默认路径。

缺少的默认文件会跳过，已存在但不可读或损坏的默认文件、缺少或损坏的显式 `-c` 文件均返回错误。`-c=` 禁用全部文件读取。嵌套表按键合并，数组整体替换而不拼接；例如实例中的 `peers = []` 清除公共列表，`StringsConfig` 可直接读取原生 TOML 字符串数组，元素中的逗号保持原样。重复 CLI 切片参数仍按现有规则累加。

`fs.GetConfig()` 包含所有文件合并后的键值（包括没有绑定 flag 的键），`GetConfigPath()` 指向最后一个实际读取的文件。它保存文件来源的值，不把 CLI/环境覆盖值写回配置；写入合并配置会将继承的键一起写入该目标，按实例保存少量键时应自行选择写入内容。

### 6. 别名

```go
fs := flag.NewFlagSet("app", flag.ContinueOnError)
verbose := fs.New("verbose").Alias("v").Default(false).Bool()
if err := fs.Parse(os.Args[1:]); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(2) }
fmt.Println(*verbose) // -v 与 -verbose 都设置同一个 bool 参数
```

若已有 `var verbose bool`，也可改用 `fs.Var(&flag.FlagVar{Value: &verbose, Name: "verbose", Aliases: []string{"v"}, DefaultValue: false})`；这是与链式写法二选一的注册方式，不要同时定义同名参数。

### 7. 隐藏参数

```go
var quiet bool
fs.New("quiet").Default(false).Hidden().Var(&quiet)
// -quiet 仍可使用，只是不在帮助中显示；隐藏不是权限或保密措施
```

### 8. 子命令

`Cmd` 最小只需 `Name() string` 和 `Run(args ...string) error`；`Init(args ...string) error` 可选，用于在运行前解析子命令参数；`Help() string` 也是可选的帮助说明。完整独立示例（保存为 `main.go`，在本模块中运行）：

```go
package main

import (
    "fmt"
    "os"

    "github.com/lfhy/flag"
)

type ServeCmd struct{ port int }

func (*ServeCmd) Name() string { return "serve" }
func (s *ServeCmd) Init(args ...string) error {
    fs := flag.NewFlagSet("serve", flag.ContinueOnError)
    fs.New("port").Default(8080).Var(&s.port)
    return fs.Parse(args)
}
func (s *ServeCmd) Run(args ...string) error {
    fmt.Println("Port:", s.port)
    return nil
}
func (*ServeCmd) Help() string { return "启动服务" } // 可选

func main() {
    root := flag.NewFlagSet("app", flag.ContinueOnError)
    root.RegisterCommand(&ServeCmd{})
    if err := root.Run(os.Args[1:]...); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(2)
    }
}
```

`go run main.go serve -port=9000` 输出 `Port: 9000`。顶层 `Run` 会在命令名处停止解析并交给子命令的 `Init`；不要先用 `root.Parse(os.Args[1:])` 再 `root.Run(...)`，以免顶层抢先解析子命令选项。

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
