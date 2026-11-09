# 项目协作指南

本仓库是 Go module `github.com/lfhy/flag`（Go 版本见 `go.mod`），提供与标准库 `flag` 相近的命令行解析，以及配置文件、环境变量、别名、隐藏参数和子命令。使用说明以 [中文 README](README.md) 与 [English README](README.en.md) 为入口；更新公开行为时同步两种语言，不让示例依赖仓库外文件或本机目录。

## 代码入口与职责

- `flag_set.go` 拥有 `FlagSet`、参数注册/解析、来源优先级和配置路径参数；`flag_args.go` 是默认参数集的包级入口。
- `flag_var.go` 定义通用 `FlagVar` 绑定及注册时错误处理；`flag_reflect_value.go` 负责反射值与切片的支持边界；`flag_builder.go` 提供 `New` 链式定义和读取。
- `flag_config_file.go` 定义显式 `ConfigFile` 及文件读取装配；`flag_config.go` 提供配置读写；`flag_cmd.go` 拥有 `Cmd`、命令注册和分发。`flag_bool.go` 等按类型命名的文件保留 `XxxVar` 兼容 API，`flag_value.go` 等实现类型值。
- 根目录 `*_test.go` 验证解析/来源/链式定义/子命令等行为；`example/` 含使用示例。变动前先读对应实现与测试，不以 README 代替真实 API。

## 安装与验证

在仓库根目录、已安装与 `go.mod` 兼容的 Go 工具链后运行：

```sh
go mod download       # 下载锁定依赖
go test ./...         # 全部可运行包测试
go vet ./...          # 静态检查
gofmt -l .            # 列出未格式化的 Go 文件；改源码时先对目标文件执行 gofmt -w
git diff --check      # 检查差异中的空白问题
```

公开 API 示例应保留完整导入、变量和错误处理，优先用独立 `flag.NewFlagSet(..., flag.ContinueOnError)`、`fs.Var(&flag.FlagVar{...})` 或 `fs.New(...).Bool()/Var(...)`；把 README Go 示例提取到本模块下的临时目录编译、运行并核对所列输出。单纯文档改动不必格式化无关源码；若更改行为还需覆盖相应测试并同步两个 README。

## 编写与安全边界

保持修改聚焦，复用既有结构，不为示例做不必要的源码重构。区分显式 `ConfigFile` 的读错承诺与不启用时的旧路径，区分解析错误和注册阶段错误；子命令由顶层 `Run` 分发，避免顶层 `Parse` 抢解析。隐藏选项不是安全边界，不在示例、测试或文档中写入密钥、凭据及个人路径；新建/更新文件使用仓库内相对引用，确保可独立检出。
