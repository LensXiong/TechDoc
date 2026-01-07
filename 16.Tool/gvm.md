




# GVM（Go Version Manager）概览
## 什么是 GVM
GVM 是一个 Go 语言版本管理工具，类似于 Python 的 pyenv 或 Node 的 nvm，主要用于：

* 在同一台机器上 安装和管理多个 Go 版本
* 灵活切换当前使用的 Go 版本
* 测试项目在不同 Go 版本下的兼容性

官网文档：https://github.com/moovweb/gvm

## 核心功能

| 功能     | 说明                                           |
| ------ | -------------------------------------------- |
| 安装不同版本 | 支持安装稳定版或 RC 版 Go，例如 `gvm install go1.23.2`   |
| 版本切换   | `gvm use go1.23.2` 可临时切换，`--default` 可设置默认版本 |
| 沙箱隔离   | 每个 Go 版本独立管理 `GOROOT`，避免冲突                   |
| 兼容旧项目  | 对旧版本依赖（如 Go 1.12、1.15）非常方便                   |

## 安装 GVM
使用官方脚本（macOS / Linux）：
```
bash < <(curl -sSL https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

安装完成后，按照提示执行：

```
source ~/.gvm/scripts/gvm
```

⚠️ 建议把 source ~/.gvm/scripts/gvm 添加到 ~/.zshrc 或 ~/.bashrc，保证每次终端打开自动加载 GVM。

## GVM 常用命令

列出已安装版本:
```
gvm list
```

示例输出：
```
gvm gos (installed)
   go1.20.14
=> go1.20.14
```
=> 表示当前正在使用的版本。

安装新版本:
```
gvm install go1.23.2
```
> GVM 会自动下载、编译并安装 Go 版本到 ~/.gvm/gos/go1.23.2。

使用指定版本:
临时切换：
```
gvm use go1.23.2
```
永久默认切换：
```
gvm use go1.23.2 --default
```
切换后，执行：
```
go version
```
即可看到当前版本。
卸载某个版本:
```
gvm uninstall go1.20.14
```

## GVM 原理
GVM 本质上是通过 修改 PATH 环境变量 来实现不同 Go 版本切换的：
* 每个 Go 版本都有独立的 GOROOT（安装目录）
* 当执行 gvm use 时，会把指定版本的 bin 目录放到 PATH 最前面
* 这样终端里执行 go 命令时，就使用指定版本
⚠️ 如果系统上存在其他 Go 版本（Homebrew 或官方安装），需要注意 PATH 优先级，否则可能还是调用旧版本。