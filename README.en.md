# flag

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

An enhanced command-line argument parsing library for Go. It stays compatible with the standard library `flag` package while adding configuration files, environment variables, hidden flags, subcommands, aliases, and string / numeric slice support.

> 📖 [中文文档](./README.md)

## ✨ Features

- **Multi-source priority**: CLI args > environment variables > config file > default value
- **🌟 Unified config entry `-c`**: `Parse()` auto-registers a hidden `-c` flag; passing a config file merges CLI/env/file by priority in a single pass — no manual viper calls needed
- **Config file integration**: powered by [viper](https://github.com/spf13/viper); supports JSON / TOML / YAML, and changes can be written back to the config
- **Hidden flags**: register flags that don't appear in `--help` but are still parsed
- **Subcommands & aliases**: register subcommands (à la `git`) and set aliases for flags/subcommands (e.g. `-v` ↔ `--verbose`)
- **Slice types**: built-in `Strings / Ints / Int64s / Uints / Uint64s` with comma-separated parsing and automatic accumulation across repeated flags
- **Generic reflection entry `Var()`**: a single API covering all primitive and slice types
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
    "github.com/lfhy/flag"
)

var (
    Port  int
    Host  string
    Tags  []string
    Debug bool
)

func init() {
    flag.IntVar(&Port, "port", 8080, "service port")
    flag.StringVar(&Host, "host", "127.0.0.1", "listen address")
    flag.StringsVar(&Tags, "tags", "go,linux", "tags (comma-separated)")
    flag.BoolVar(&Debug, "debug", false, "debug mode")
    flag.Parse()
}

func main() {
    fmt.Printf("Port=%d Host=%s Tags=%v Debug=%v\n", Port, Host, Tags, Debug)
}
```

Run it:

```bash
$ go run main.go -port 9000 -tags a,b -tags c
Port=9000 Host=127.0.0.1 Tags=[a b c] Debug=false
```

## 📚 Guide

### 1. Defining flags

The package level provides `XxxVar` / `Xxx` functions consistent with the standard library style; `FlagSet` can also be used standalone within a subcommand.

| Function | Description |
|---|---|
| `flag.IntVar(p, "n", 1, "...")` | Primitive type, CLI + default only |
| `flag.IntEnvVar(p, "n", "APP_N", 1, "...")` | Also binds an environment variable |
| `flag.IntConfigVar(p, "n", "server", "port", 1, "...")` | Also binds the config file |
| `flag.IntFullVar(p, "n", "server", "port", "APP_N", 1, "...")` | All sources |
| `flag.Int("n", 1, "...")` | Returns `*int`, no pre-declared variable needed |
| Append `Hidden` to the name (e.g. `IntHiddenVar`) | Registers as a hidden flag |

> All types (Bool / String / Int / Int64 / Uint / Uint64 / Float64 / Duration, plus the slice variants Strings/Ints/...) follow the same naming convention.

### 2. Slice types

Slice types parse comma-separated values and **accumulate across repeated flags** (consistent with `pflag`'s `StringSlice`):

```go
var ids []int
flag.IntsVar(&ids, "ids", "1,2,3", "ID list")
// -ids=4,5 -ids=6  →  [1 2 3 4 5 6]
```

Supported: `StringsVar` / `IntsVar` / `Int64sVar` / `UintsVar` / `Uint64sVar`.

### 3. Generic reflection entry Var()

`Var()` takes a `FlagVar` struct and uses reflection to auto-detect primitive and slice types, so you don't need a dedicated API per type:

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

Types supported by `Var()`: `bool / string / int* / uint* / float64 / time.Duration`, plus their slice variants.

### 4. Environment variables & config files (manual binding)

```go
// Bind an env var: read from it when not specified on the CLI
flag.StringEnvVar(&host, "host", "APP_HOST", "0.0.0.0", "listen address")

// Bind a config file: read from title.key
flag.IntConfigVar(&port, "port", "server", "port", 8080, "service port")

// Read/write the config
cfg := flag.GetConfig()
cfg.WriteToConfig("server", "port", 9000)
```

Priority: **CLI > environment variables > config file > default value**.

### 5. 🌟 Unified config entry: `-c` to specify a config file (signature feature)

This is the core capability that sets the library apart from the standard `flag`. When `Parse()` runs it auto-registers a hidden `-c` flag; passing a config file path makes **CLI args, the config file, and environment variables merge in a single parse according to priority** — no manual viper calls required.

```go
// app.go
var (
    Host string
    Port int
)

func init() {
    flag.StringVar(&Host, "host", "127.0.0.1", "listen address")
    flag.IntConfigVar(&Port, "port", "server", "port", 8080, "service port")
    flag.Parse()
}
```

Suppose there's a `config.toml`:

```toml
[server]
port = 9000
```

Three sources merged:

```bash
# 1. Use the config file only
$ go run app.go -c config.toml
# Host=127.0.0.1 (default)  Port=9000 (from config file)

# 2. CLI overrides the config file
$ go run app.go -c config.toml -port 8888
# Port=8888 (CLI > config file)

# 3. Env vars also participate (effective when absent from CLI/config)
$ APP_HOST=0.0.0.0 go run app.go -c config.toml
# Host=0.0.0.0 (env > default)
```

Notes:

- `-c` is an auto-registered hidden flag, no need to define it manually; file parsing is skipped when not specified
- The flag name can be customized via `fs.SetConfigFlagName("conf")` (per-FlagSet, independent; defaults to `"c"`)
- The config file is read by `ConfigTitle.ConfigKey` (the 2nd and 3rd args of `IntConfigVar`)
- The default flag name can be changed via `flag.DefaultConfigFlagName` (default `"c"`)
- After parsing, use `flag.GetConfig()` to obtain `*Config`, which supports reading/writing back to the file

### 6. Aliases

```go
flag.String("verbose", "false", "verbose output")
flag.Alias("verbose", "v")     // -v is equivalent to --verbose
```

### 7. Hidden flags

```go
flag.StringHiddenVar(&token, "token", "", "internal auth token")
// Won't appear in --help output, but -token=xxx still takes effect
```

### 8. Subcommands

Implement the `Cmd` interface (`Name / Init / Run / Help`) and register it:

```go
type ServeCmd struct { *flag.FlagSet }

func (*ServeCmd) Name() string { return "serve" }
func (s *ServeCmd) Init(args ...string) error {
    s.FlagSet = flag.NewFlagSet("serve", flag.ContinueOnError)
    // ... define subcommand flags
    return s.Parse(args)
}
func (*ServeCmd) Run(args ...string) error { /* ... */ return nil }
func (*ServeCmd) Help() string             { return "start the service" }

flag.RegisterCommand(&ServeCmd{})
flag.Parse()
```

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
