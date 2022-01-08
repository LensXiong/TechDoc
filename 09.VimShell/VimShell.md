

# Linux 

## 2>&1 | tee

① 标准输出+标准错误

输出标准输出和标准错误，同时保存到文件`logfile`。

```shell
<command> 2>&1 | tee <logfile>
```

管道作用：把一个进程的标准输出作为另一个进程的标准输入。

`2>&1`作用：是把标准错误重定向到标准输出的副本一起输出。上面的命令，把标准输出和标准错误都输出作为`tee`命令的标准输入。

`tee`的作用：把标准输入的内容拷贝到文件，并输出。

② 忽略标准输出

若想忽略标准输出，只输出标准错误并保存到文件`logfile`。

```shell
<command> 2>&1 >/dev/null | tee logfile
```

`2>&1`： 将标准错误重定向到标准输出，注意，此时标准输出还没有被重定向。

`/dev/null` 等同于`1>/dev/null`。 `/dev/null`文件是一个空设备，类似于`windows`内的回收站，使用`>/dev/null`（等同于`1>/dev/null`）将标准输出重定向到`/dev/null`，即不显示标准输出的内容。所以这时的标准输出就仅变为重定向过来的标准错误了。

##  nohup |  &

```
nohup <command> > xxx.log &
```

`&`：后台运行程序，结果会输出到终端，使用`Ctrl + C`发送`SIGINT`信号，程序免疫，关闭`session`发送`SIGHUP`信号，程序关闭。

`nohup`：结果默认会输出到`nohup.out`，使用`Ctrl + C`发送`SIGINT`信号，程序关闭，关闭`session`发送`SIGHUP`信号，程序免疫。

注：使用`nohup`和`&`配合来启动程序`nohup ./test &`，同时免疫`SIGINT`和`SIGHUP`信号。

## 光标快捷键

| 命令     | 说明             | 作用                       |
| -------- | ---------------- | -------------------------- |
| ctrl + r | reverse-i-search | 反向搜索执行过的命令       |
| ctrl + a | head             | 光标跳到行首               |
| ctrl + e | end              | 光标跳到行尾               |
| ctrl + b | backward         | 光标后退一个字符           |
| ctrl + f | forward          | 光标前进一个字符           |
| ctrl+d   | delete           | 删除光标后一个字符         |
| ctrl+h   | head             | 删除光标前一个字符         |
| ctrl+k   |                  | 清除光标后至行尾的所有内容 |
| ctrl+u   |                  | 清除光标前至行首的所有内容 |
| ctrl + c | clear            | 杀死当前进程，另起一行     |
| ctrl + l | line             | 清屏                       |
| ctrl + w | words            | 移除光标前的一个单词       |
| ctrl + t |                  | 交换光标位置前的两个字符   |

## 复用命令

| 命令        | 作用                         |
| ----------- | ---------------------------- |
| !!          | 复用上一条命令               |
| !df         | 复用上一条以 “df” 开头的命令 |
| history     | 查看历史命令                 |
| history [n] | 列出最近的n条历史命令        |
| ![n]        | 复用命令历史中的 [n] 号命令  |

## 查看文件

| 命令                    | 作用                         |
| ----------------------- | ---------------------------- |
| tail -f /var/log/syslog | 实时显示日志文件中增加的内容 |
| tail -5 test.php        | 查看最后五行日志信息         |
| tail -n +10 test.php    | 从第10行开始显示文件         |

## 查询文件

| 命令                                  | 作用                                             |
| ------------------------------------- | ------------------------------------------------ |
| find / -name "php.ini"                | 查找php.ini文件的位置                            |
| find /home -iname "*.txt"             | 在/home目录下忽略大小写，查找以.txt结尾的文件名  |
| find . -type f -size +10M             | 搜索大于10M的文件                                |
| find . -type f -size -1k              | 搜索小于1k的文件                                 |
| find . -type f -size 10M              | 搜索等于10M的文件                                |
| find . -name "*.txt" -o -name "*.pdf" | 当前目录及子目录下查找所有以.txt和.pdf结尾的文件 |

# Vim 

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







