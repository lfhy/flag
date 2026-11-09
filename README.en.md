# flag

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

An enhanced command-line argument parsing library for Go. It stays compatible with the standard library `flag` package while adding configuration files, environment variables, hidden flags, subcommands, aliases, and string / numeric slice support.

> 📖 [中文文档](./README.md)

## ✨ Features

- **Multi-source priority**: CLI args > environment variables > config file > default value
- **🌟 Unified config entry `-c`**: `Parse()` auto-registers a hidden `-c`; opt into a default path and strict file read errors with `ConfigFile`, then merge CLI/env/file by priority
- **Config file integration**: powered by [viper](https://github.com/spf13/viper); supports JSON / TOML / YAML, and changes can be written back to the config
- **Hidden flags**: register flags that don't appear in `--help` but are still parsed
- **Subcommands & aliases**: register subcommands (à la `git`) and set aliases for flags/subcommands (e.g. `-v` ↔ `--verbose`)
- **Interspersed positional arguments and flags**: `FlagSet.Parse` accepts flags before or after positional arguments; boolean flags do not consume paths, and tokens after `--` remain literal positional arguments. Top-level subcommand dispatch can use `ParseStandard` to stop at the first positional argument
- **Slice types**: built-in `Strings / Ints / Int64s / Uints / Uint64s` with comma-separated parsing and automatic accumulation across repeated flags
- **Generic reflection entry `Var()`**: one API for supported scalar types and string/integer slices
- **Colored help output**: built-in ANSI color printing; `PrintAll` lists all flags as a table

## 📦 Installation

```bash
go get github.com/lfhy/flag
```

## 🚀 Quick Start

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
    fs.Var(&flag.FlagVar{Value: &port, Name: "port", DefaultValue: 8080, Usage: "service port"})
    fs.Var(&flag.FlagVar{Value: &tags, Name: "tags", Usage: "tags (repeatable)"})
    host := fs.New("host").Default("127.0.0.1").Usage("listen address").String()
    debug := fs.New("debug").Default(false).Usage("debug mode").Bool()

    if err := fs.Parse(os.Args[1:]); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(2)
    }
    fmt.Printf("Port=%d Host=%s Tags=%v Debug=%v\n", port, *host, tags, *debug)
}
```

Run it:

```bash
$ go run main.go -port 9000 -tags a,b -tags c
Port=9000 Host=127.0.0.1 Tags=[a b c] Debug=false
```

## 📚 Guide

### 1. Recommended ways to define flags

Prefer `fs.Var(&flag.FlagVar{...})` on an independent `FlagSet` for richer settings (name, aliases, default, environment variable, config section/key, usage, and hidden status), or use fluent definitions such as `fs.New(...).Default(...).Config(...).Env(...).Usage(...).Int()`. Use `.Var(&value)` at the end of a chain to bind an existing variable. Register at runtime and call `fs.Parse(os.Args[1:])`. Package-level `flag.Var` / `flag.New` use the shared default `FlagSet` when global state is intentional.

#### Fluent definitions

`flag.New` (or `fs.New`) chains aliases, defaults, config keys, environment variables, and help text. A terminal type method registers the flag immediately and returns a pointer; `Parse` updates that pointer with the resolved value:

```go
fs := flag.NewFlagSet("app", flag.ContinueOnError)
follow := fs.New("follow").Alias("f").Default(false).
    Config("title.key").Env("APP_FOLLOW").Usage("follow updates").Bool()
if err := fs.Parse(os.Args[1:]); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(2) }
fmt.Println(*follow) // resolved from -follow/-f, APP_FOLLOW, the config file, or the default
```

`.Config("title.key")` maps to the `[title]` section and `key` in the config file. Without `ConfigFile`, pass `-c config.toml` to read the file, or use `fs.SetConfigFlagName("conf")` to select `-conf`. Environment variables are **inputs**, not destinations: parsing does not write them. Priority is **CLI > environment > config file > default**.

Use `.Var(&data)` to bind an existing variable; it registers the flag immediately and does not return a pointer:

```go
var data bool
fs.New("enabled").Default(false).Var(&data)
if err := fs.Parse(os.Args[1:]); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(2) }
fmt.Println(data)
```

You can also keep a definition without calling a terminal type method or `.Var`, then read it after parsing:

```go
fs := flag.NewFlagSet("app", flag.ContinueOnError)
definition := fs.New("follow").Default(false).Env("FOLLOW")
count := fs.New("count").Default("3") // no terminal type method: default is stored as text
if err := fs.Parse([]string{"-follow", "-count=5"}); err != nil { panic(err) }
fmt.Println(definition.Value().Bool()) // true (a bool value, not *bool)
fmt.Println(count.Value().Int(), count.Value().String()) // 5 5
```

`Parse` registers non-terminal definitions with string-backed storage. `.Value().Bool()/Int()/Int64()/String()` and the other accessors convert on read; numeric and boolean accessors **panic** if the text cannot be converted, while `.Value().String()` returns the text unchanged. Terminal `.Bool()/Int()/Int64()/String()` and the other type methods instead return typed pointers (`*bool`, `*int`, `*int64`, `*string`, etc.).

`FlagSet` retains parsed flag state. Create a new `FlagSet` to parse independently with a different environment or configuration.

### 2. Slice types

Slice types parse comma-separated values and **accumulate across repeated flags**. Non-empty slice defaults are also retained when values are appended; use an empty default if explicit values should start a fresh list:

```go
var ids []int
fs := flag.NewFlagSet("app", flag.ContinueOnError)
fs.Var(&flag.FlagVar{Value: &ids, Name: "ids", DefaultValue: []int{1, 2}, Usage: "ID list"})
if err := fs.Parse([]string{"-ids=4,5", "-ids=6"}); err != nil { panic(err) }
fmt.Println(ids) // [1 2 4 5 6]
```

The compatibility APIs `StringsVar` / `IntsVar` / `Int64sVar` / `UintsVar` / `Uint64sVar` take comma-separated strings for defaults and also append. Reflection-backed `FlagVar` supports only `[]string`, `[]int*`, and `[]uint*`, not arbitrary slices.

### 3. Generic reflection entry Var()

`Var()` accepts `FlagVar` to configure `Value`, `Name`, `Aliases`, `DefaultValue`, `Env`, `ConfigSection`, `ConfigKey`, `Usage`, and `Hidden` together and reflect over supported types. Prefer a non-nil variable pointer. It is more expressive than multiple `XxxVar` calls when a flag needs several sources or aliases:

```go
fs := flag.NewFlagSet("app", flag.ContinueOnError)
var port int
var ids []int64
fs.Var(&flag.FlagVar{
    Value: &port, Name: "port", Aliases: []string{"p"}, DefaultValue: 8080,
    ConfigSection: "server", ConfigKey: "port", Env: "APP_PORT", Usage: "service port",
})
fs.Var(&flag.FlagVar{Value: &ids, Name: "ids", DefaultValue: []int64{1, 2}})
if err := fs.Parse(os.Args[1:]); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(2) }
```

`Var()` supports `bool / string / int* / uint* / float32 / float64 / time.Duration` and string, signed-integer, or unsigned-integer slices. Omitting the section/key makes `FlagVar` infer a mapping: a simple name maps to `config.<name>`, but a hyphenated name can produce a surprising key (for example, `server-port` maps to `server.server-port`). Set `ConfigSection` and `ConfigKey` explicitly, or use fluent `.Config("server.port")`, for a stable mapping. With `ContinueOnError`, some registration errors are printed during registration and **are not guaranteed** to be returned by a later `Parse`; check definitions rather than treating a successful parse as proof of successful registration.

### 4. Environment variables & config files

Bind environment, config key, and CLI in the same `FlagVar` or fluent definition. There is no need to parse twice or read the file separately:

```go
var host string
fs.Var(&flag.FlagVar{
    Value: &host, Name: "host", Env: "APP_HOST", DefaultValue: "127.0.0.1",
    ConfigSection: "server", ConfigKey: "host", Usage: "listen address",
})
// To update the file after parsing: fs.GetConfig().WriteToConfig("server", "host", host)
```

Priority: **CLI > environment variables > config file > default value**. `XxxVar` / `XxxEnvVar` / `XxxConfigVar` / `XxxFullVar` and their pointer-returning `Xxx` counterparts remain as compatibility APIs. For example, `fs.IntFullVar(&port, "port", "server", "port", "APP_PORT", 8080, "service port")` binds the section, key, and environment variable before specifying default and usage.

### 5. 🌟 Unified config entry: a default file or `-c` to specify one (signature feature)

`FlagSet.Parse` auto-registers a hidden `-c`. **Opt in with `fs.ConfigFile(...)`** for a default config path and dependable read errors; CLI, environment variables, and file values then merge by priority without calling viper yourself. You can pass `-c` without calling `ConfigFile`, but this legacy path cannot be relied on for strict read errors; `GetConfig()` may also create its default file.

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
    fs.Var(&flag.FlagVar{Value: &host, Name: "host", Env: "APP_HOST", DefaultValue: "127.0.0.1", Usage: "listen address"})
    fs.New("port").Default(8080).Config("server.port").Var(&port)
    if err := fs.Parse(os.Args[1:]); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(2)
    }
    fmt.Printf("Host=%s Port=%d\n", host, port)
}
```

Suppose there's a `config.toml`:

```toml
[server]
port = 9000
```

After opting in with `ConfigFile`, the same program merges sources this way:

```bash
$ go run app.go
Host=127.0.0.1 Port=9000
$ go run app.go -port 8888
Host=127.0.0.1 Port=8888
$ APP_HOST=0.0.0.0 go run app.go
Host=0.0.0.0 Port=9000
```

With `fs.ConfigFile("config.toml")`, an existing default file is loaded when `-c` is absent; a missing default is skipped. An explicitly selected missing or unreadable file, or a malformed existing file, makes `Parse` return an error. Without `ConfigFile`, an explicitly missing `-c` path is skipped; an existing path goes through the legacy reader, which may ignore read or parse errors. A separate call to `fs.GetConfig()` can create its default `config.toml`; a missing `-c` file is **not** automatically created just because it was passed to `Parse`.

`go run app.go` uses `config.toml`; `go run app.go -c other.toml` uses the explicit file instead. Priority remains CLI > environment variables > config file > default value. The path flag can also be named fluently (choose one of these alternatives):

```go
fs := flag.NewFlagSet("app", flag.ContinueOnError)
fs.ConfigFile("config.toml").Alias("config") // accept -c and -config
// To use only -config instead: fs.ConfigFile("config.toml").Name("config")
```

`fs.SetConfigFlagName("conf")` can still rename the flag before parsing; if used with `.Name(...)`, the last setting wins. Define config keys via `FlagVar.ConfigSection` / `ConfigKey` or fluent `.Config("server.port")`. After parsing, `fs.GetConfig()` provides the current configuration for optional writes; the example above does not write anything.

### 6. Aliases

```go
fs := flag.NewFlagSet("app", flag.ContinueOnError)
verbose := fs.New("verbose").Alias("v").Default(false).Bool()
if err := fs.Parse(os.Args[1:]); err != nil { fmt.Fprintln(os.Stderr, err); os.Exit(2) }
fmt.Println(*verbose) // -v and -verbose set the same bool flag
```

Alternatively, with `var verbose bool` already declared, use `fs.Var(&flag.FlagVar{Value: &verbose, Name: "verbose", Aliases: []string{"v"}, DefaultValue: false})`. Choose one of the two registration styles; do not define the same flag twice.

### 7. Hidden flags

```go
var quiet bool
fs.New("quiet").Default(false).Hidden().Var(&quiet)
// -quiet still works; hidden only removes it from help, not access or disclosure
```

### 8. Subcommands

The minimal `Cmd` interface requires only `Name() string` and `Run(args ...string) error`. `Init(args ...string) error` is optional for parsing subcommand flags before `Run`; `Help() string` is optional help text. Complete standalone example (save as `main.go` and run within this module):

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
func (*ServeCmd) Help() string { return "start the service" } // optional

func main() {
    root := flag.NewFlagSet("app", flag.ContinueOnError)
    root.RegisterCommand(&ServeCmd{})
    if err := root.Run(os.Args[1:]...); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(2)
    }
}
```

`go run main.go serve -port=9000` prints `Port: 9000`. Top-level `Run` stops parsing at the command name and passes later options to the subcommand's `Init`. Do not call `root.Parse(os.Args[1:])` before `root.Run(...)`: top-level parsing could consume or reject subcommand options.

## 📖 API Overview

Each type provides 8 functions at each level (`FlagSet` methods / `ArgsFlag` methods / package-level functions):

```
XxxFullVar / XxxConfigVar / XxxEnvVar / XxxVar      // writes into a pre-declared variable
XxxFull    / XxxConfig    / XxxEnv    / Xxx          // returns *T
```

Plus the 8 `Hidden` variants. Covered types:

| Primitive | Slice |
|---|---|
| Bool / String | Strings |
| Int / Int64 | Ints / Int64s |
| Uint / Uint64 | Uints / Uint64s |
| Float64 | — |
| Duration | — |

## 🧪 Testing

```bash
go test ./...
```

## 📄 License

MIT
