# vscode 常用快捷键

| 操作          | 快捷键                  | 说明                |
| ----------- | -------------------- | ----------------- |
| 快速打开文件      | `Cmd + P`            | 模糊文件名查找           |
| 全局查找文本      | `Cmd + Shift + F`    | 搜索整个工作区的内容        |
| 当前文件查找      | `Cmd + F`            | 搜索当前文件中的文本        |
| 当前文件替换      | `Cmd + Option + F`   | 当前文件内替换           |
| 全局替换        | `Cmd + Shift + H`    | 全局文本替换            |
| 转到符号（函数/变量） | `Cmd + Shift + O`    | 当前文件内按函数名跳转       |
| 转到行号        | `Cmd + P` → 输入 `:行号` | 例：`:42` 跳转到第 42 行 |



# 全局默认使用 UTF-8 编码
全局设置 VSCode 默认编码为 UTF-8

1、打开 VSCode。

2、使用快捷键：Cmd + ,（Mac） 或 Ctrl + ,（Windows） 打开 设置。

3、搜索：encoding

4、找到并设置以下选项：

* Files: Encoding
* 设置为：utf8（推荐，等价于 utf8 或 utf8bom，但不要选带 BOM 的除非明确需要）

如果你喜欢手动编辑配置文件，可以这样做：

打开设置文件（Command Palette -> 输入 Preferences: Open Settings (JSON)），添加：

```
{
"files.encoding": "utf8",
"files.autoGuessEncoding": true
}
```

解释：


* "files.encoding": "utf8"：默认以 UTF-8 编码打开和保存文件。

* "files.autoGuessEncoding": true：启用自动识别文件编码，比如 GBK、ISO-8859-1 等。

