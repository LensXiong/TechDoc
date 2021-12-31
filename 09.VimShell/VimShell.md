# Vim 操作

## 跳转

* 跳到到首行：两次`g`。
* 跳到最后一行：按`G`，即`shift+g`。
* 跳转到当前行的第一个字符：在当前行按`0`。
* 跳转到当前行的最后一个字符：在当前行按`$`。
* 跳转到文件第`n`行：`ngg/nG `。例如：25gg或者25G 跳转到第25行。



## 显示

设置 vim 显示行号：

* 临时**显示行号** ：如果只是临时**显示vim**的**行号**，只须按ESC键退出编辑内容模式，输入`：set number`后按回车键，就可以**显示行号**了。 **行号显示**只是暂时的，退出**vim**后再次打开**vim**就不**显示行号**了。
* 如果想让vim永久显示行号，则需要修改vim配置文件vimrc。如果没有此文件可以创建一个。在启动vim时，当前用户根目录下的vimrc文件会被自动读取，因此一般在当前用户的根目录下创建vimrc文件，即使用下面的命令：

```
vim ~/.vimrc
set number
```



配置全局变量`gopath`

```go
vim ~/zshrc

export GOPATH=/Users/wangxiong3/sdk/go1.14
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN

source ~/.zshrc
```







