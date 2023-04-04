

# frpc






# WAN口和LAN口

`WAN` 口和 `LAN` 口是计算机网络中常用的术语，用于描述网络设备的物理连接。

`WAN（Wide Area Network）`口是指连接到广域网（`Internet`）的接口，通常由`Internet`服务提供商（`ISP`）
提供并连接到调制解调器或路由器上。

`WAN`口的主要作用是将本地网络连接到互联网，以便能够访问远程网络和互联网上的资源。

`LAN（Local Area Network）`口是指连接到局域网的接口，通常由交换机或路由器提供并连接到本地设备上。
`LAN`口的主要作用是将多个本地设备连接到同一个网络中，以便这些设备之间可以进行数据交换和共享资源。

通常情况下，一个路由器会同时提供一个`WAN`口和多个`LAN`口。`WAN`口用于连接到互联网，而`LAN`口用于将多个本地设备连接到同一个网络中。
这使得多个设备可以共享互联网连接，并且可以相互访问和交换数据。

# 去除 goland 代码中的波浪线(黄色警告线)
搜索关键词：Weak Warning
![](file/weak_warning.png)

# mac 系统终端 sudo 免输入密码技能

[mac系统终端sudo免输入密码技能get](https://www.cnblogs.com/princesong/p/10293885.html)
`/etc/sudoers` 这个文件的权限是`r/r/n`，配置之前需要加写权限。
```
sudo chmod u-w /etc/sudoers
```
添加 `NOPASSWD`：
```
sudo vi /etc/sudoers
// 将 #%admin ALL=(ALL) ALL
// 替换为 %admin ALL=(ALL) NOPASSWD: ALL
```
修改完后配置为原来的读权限：
```
sudo chmod u-w /etc/sudoers
```

# Homebrew 镜像助手

[Homebrew 镜像助手](https://brew.idayer.com/guide/change-source/)
针对首次安装换源:

```
export HOMEBREW_BREW_GIT_REMOTE="https://mirrors.ustc.edu.cn/brew.git"
export HOMEBREW_CORE_GIT_REMOTE="https://mirrors.ustc.edu.cn/homebrew-core.git"
export HOMEBREW_API_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles/api"
export HOMEBREW_BOTTLE_DOMAIN="https://mirrors.ustc.edu.cn/homebrew-bottles/bottles"
```

```
/bin/bash -c "$(curl -fsSL https://gitee.com/ineo6/homebrew-install/raw/master/install.sh)"
```

# SSH 客户端无法识别主机密钥类型

问题：运行 `ssh -p 2xxx xxxx@jumper.xxxx.net` 时出现 `Unable to negotiate with 10.xxx.xxx.xx port 22xx: no matching host
key type found. Their offer: ssh-rsa`。

原因：该错误表示 `SSH` 客户端无法识别主机密钥类型，是由于 `SSH` 服务器配置了旧的加密算法或不受支持的密钥类型导致的。

解决：
① `SSH` 目录的权限
用户目录下的 `~/Users/xxxx/.ssh`目录以及下面的文件需要特别小心的管理其权限，

整个 `~/.ssh` 目录需要设置 700 `(drwx------)` 权限
`public key` 需要设置 `644(-rw-r--r--)`
客户端的私钥需要设置 600 `(-rw-------)`
需要保证该目录不会被其他 `group` 的用户读取和修改。
使用 `vim ~/Users/xxxx/.ssh` 输入以下内容：

```
 # man ssh_config
 # http://einverne.github.io/post/2017/05/ssh-keep-alive.html
 # http://einverne.github.io/post/2020/07/sync-ssh-config.html
 Host *
     User git
     PubkeyAcceptedAlgorithms +ssh-rsa
     HostkeyAlgorithms +ssh-rsa
     ForwardAgent yes
     ServerAliveInterval 30
     ServerAliveCountMax 10
     TCPKeepAlive no
     ControlMaster auto
     ControlPath ~/.ssh/conn-%r@%h:%p
     ControlPersist 120h
     Compression yes
```

如上是关于 SSH 客户端的配置文件，用于定制 SSH 连接的行为和选项。以下是对每个配置参数的解释：

```
Host *: 这是一个通配符，表示适用于所有的主机。
User git: 这个参数指定连接到远程主机时使用的用户名。
PubkeyAcceptedAlgorithms +ssh-rsa 和 HostkeyAlgorithms +ssh-rsa: 这两个参数指定使用 ssh-rsa 算法进行密钥认证和主机密钥验证。
ForwardAgent yes: 这个参数启用SSH代理转发功能，可以在连接到远程主机后让 SSH 客户端继续代理您的认证信息。
ServerAliveInterval 30: 这个参数指定发送保活消息的时间间隔，以确保连接保持活动状态。
ServerAliveCountMax 10: 这个参数指定发送保活消息后，如果没有收到任何响应，则尝试重新连接的次数。
TCPKeepAlive no: 这个参数指定是否启用 TCP keepalive 功能。
在SSH连接中，由于 SSH 协议本身已经包含了保活功能，因此建议禁用 TCP keepalive 功能，以避免不必要的流量和 CPU 开销。
ControlMaster auto: 这个参数指定是否使用 SSH 主控模式，可以让 SSH 客户端复用现有的连接，从而加速 SSH 连接的建立和执行。
ControlPath ~/.ssh/conn-%r@%h:%p: 这个参数指定主控模式下，用于保存控制连接的 Unix socket文件路径。
ControlPersist 120h: 这个参数指定主控模式下，控制连接的持续时间。在指定时间内，如果没有任何活动，则保持连接打开状态。
Compression yes: 这个参数指定是否启用数据压缩功能。在网络带宽受限的情况下，启用数据压缩功能可以提高 SSH 连接的效率。
```

配置 `iterm session`:

![](file/iterm_session.png)

# 去掉 vim 每行结尾的 ^M

有时候用 `vim` 打开文件，每行结尾都有一个灰色的`^M`。
这个原因是该文件在 `windows` 或 `mac` 系统上被创建：

`windows` 的换行符是`\n\r`；
`mac` 的换行符是 `\r`；
`unix` 下的文本换行只需要`\n`。
在`vim`下，这个多余的`\r`就被显示为`^M`，虽然显示为两个字符，但其实是一个字符。

大多数情况下，打开这种文件，`vim`的状态栏会显示文件格式：`utf-8[dos]`或者`utf-8[mac]`，此时只需要下面命令可转为`unix`
格式，即可删除或转换所有的`^M`:

```
:set ff=unix
```

如果`vim`显示文件格式已经是`utf-8[unix]`，这时候上述命令就不管用了，说明 `vim` 识别类型错误，可以先把它纠正（即用`dos`
或`mac`格式打开当前文件），再变更类型：

```
:e ++ff=dos
# or
# :e ++ff=mac
:set ff=unix
```

还有一种方法是字符串替换：

```
:%s/\r//g
```

注意这里是用`\r`而不是`^M`，这也是很多人不会删除`^M`的原因。

# ChatGPT

[国外接码平台 sms-activate.org 官网]https://sms-activate.org/getNumber

[VPN-ClashX](https://github.com/yichengchen/clashX)

[VPN-ClashX 下载](https://github.com/yichengchen/clashX/releases)

[DuangCloud](https://portal.dc-site3.com/#/login)

登录`ChatGPT`报不可用 `country`:

```
OpenAI's services are not available in your country. (error=unsupported_country)
```

核心要点：

① 出站模式（全局）：全局连接

② 勾选设置为系统代理。

③ GLOBAL 海外专线。

④ 配置导入相关`DuangCloud`配置。

⑤ 浏览器系统设置代理：切换为系统代理（海外）。直接连接（国内）。

⑥ 设置为系统代理来回切换几次，还是不行就重启`ClashX`。

# MAC 技巧

## 三指拖移

新版 `macOS Catalina/Big Sur`系统里面三指拖移已经被苹果从触控板中移除并把这个功能放入了辅助功能里面。
这个位置还不太好找，`mac`老手在这里也会频频翻车，这个功能开启也非常简单。

新版 mac ：系统设置 -> 辅助功能 -> 指针控制 -> 触控板选项 -> 拖移样式 -> 三指拖移

![img_1.png](file/img_1.png)

# iterm2-zmodem

解决：在 mac 下，实现与服务器进行便捷的文件上传和下载操作。

[iterm2-zmodem 文档](https://github.com/aikuyun/iterm2-zmodem)
主要过程：

```
x. rz, sz
https://github.com/aikuyun/iterm2-zmodem

> sz 下载功能
服务端send

> rz 上传功能
服务端receive

服务器:
yum -y install lrzsz

客户端:
brew install lrzsz
```

主要命令：

```
git clone git@github.com:aikuyun/iterm2-zmodem.git

cd iterm2-zmodem && chmod +x iterm2-*  && sudo mkdir -p /usr/local/bin && \n
sudo cp iterm2-* /usr/local/bin && sudo ln -s /opt/homebrew/bin/rz /usr/local/bin && \n
sudo ln -s /opt/homebrew/bin/sz /usr/local/bin
```

设置 `Iterm2`的`Tirgger`特性，`profiles->default->editProfiles->Advanced`中的`Tirgger`：
添加两条`trigger`，分别设置 `Regular expression，Action，Parameters，Instant`如下：

① 第一条

```
Regular expression: rz waiting to receive.\*\*B0100
Action: Run Silent Coprocess
Parameters: /usr/local/bin/iterm2-send-zmodem.sh
Instant: checked
```

② 第二条

```
Regular expression: \*\*B00000000000000
Action: Run Silent Coprocess
Parameters: /usr/local/bin/iterm2-recv-zmodem.sh
Instant: checked
```

![img.png](file/img.png)

# Github Add SSH keys

问题：Git 当前项目设置 用户名、邮箱。

```
➜  git config --list
➜  git config user.name "xxxx"
➜  git config user.email "xxx@gmail.com"
```

问题：`Github`上添加了`ssh key`，但是在`push`代码的时候还会要求我们输入用户名密码?

其中一种可能的原因是，在克隆代码的时候，使用了`https`协议，而不是`git`协议。

`https`协议要求你每次都需要输入用户名以及密码，自由`git`协议才可以使用`ssh-key`文件。

```
➜ git remote -v
origin  https://github.com/LensXiong/TechDoc.git (fetch)
origin  https://github.com/LensXiong/TechDoc.git (push)
➜ git remote set-url origin git@github.com:LensXiong/TechDoc.git
➜ git remote -v                                                 
origin  git@github.com:LensXiong/TechDoc.git (fetch)
origin  git@github.com:LensXiong/TechDoc.git (push)
```

# macOS Zsh 使用 oh-my-zsh 打造高效便捷的 shell 环境

[macOS Zsh 使用 oh-my-zsh 打造高效便捷的 shell 环境](https://sysin.org/blog/macos-zsh/)

## zsh-autosuggestions

作用是根据历史输入命令的记录即时的提示（建议补全），然后按 → 键即可补全。

```
git clone --depth=1 https://github.com/zsh-users/zsh-autosuggestions.git ${ZSH_CUSTOM:-${ZSH:-~/.oh-my-zsh}/custom}/plugins/zsh-autosuggestions
```

编辑 `~/.zshrc`，找到 `plugins=(git)` 这一行，修改为：

```
plugins=(
    git
    # other plugins...
    zsh-autosuggestions
)
```

# docker-compose network_mode=“host”

问题：`docker-compose` 中 `network_mode=“host”` 外网访问不了。

原因：`docker` 默认的 `network`是`bridge`，这个默认会把映射的端口加到宿主机防火墙。而`host`模式是不会主动加入防火墙的，所以需要添加端口。

解决：

```
# 开放指定端口
firewall-cmd --zone=public --add-port=9203/tcp --permanent
# 重启防火墙
firewall-cmd --reload
```