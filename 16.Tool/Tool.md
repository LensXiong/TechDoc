
# Excel转Json


# JSON文本处理工具jq

[Linux下JSON文本处理工具jq](https://gitbook.curiouser.top/origin/linux-jq.html)

Mac安装jq工具：
```
brew install jq
```
使用示例：
```
curl 'https://xx.xx.cn/xxx-api/apiz/xxz/cloud?key=xxx' \
  -H 'Accept: */*' \
  -H 'Accept-Language: zh-CN,zh;q=0.9,en;q=0.8' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Cookie: .' \
  -H 'Pragma: no-cache' \
  -H 'Referer: https://xxx.xxx.cn/chat/xxx' \
  -H 'Sec-Fetch-Dest: empty' \
  -H 'Sec-Fetch-Mode: cors' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36' \
  -H 'access-token;' \
  -H 'device-platform: Web' \
  -H 'func-ver: 1' \
  -H 'sec-ch-ua: "Chromium";v="122", "Not(A:Brand";v="24", "Google Chrome";v="122"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'timestamp: 2024-03-19T17:05:22+08:00' \
  -H 'zm-token: xxxx' \
  -H 'zm-ua: xxxxx' \
  -H 'zm-ver: 1.2' | jq '.data.brain_spec_robots.list[] | select(.index_id!="") |{"target","index_id"}'
```
示例1：
```
| jq '.data.brain_spec_robots.list[] | select(.index_id!="") |{"target","index_id"}'  
```
从输入的 `JSON` 数据中遍历 `data.brain_spec_robots.list` 数组，筛选出那些 `index_id` 不为空字符串的元素，
并为这些元素构造一个新的 JSON 对象，该对象仅包含每个元素的 `target` 和 `index_id` 字段。
如果在 `{"target","index_id"}` 这一部分有误解，正常情况应该是提取具体的值，
比如 `{"target": .target, "index_id": .index_id}`。

```
| jq '.data.brain_spec_robots.list[] | select(.index_id!="") | "\(.target),\(.index_id)"'
```
从输入的 `JSON` 数据中，遍历 `data.brain_spec_robots.list` 数组，筛选出 `index_id` 字段不为空字符串的元素。
对于每个筛选后的元素，它会输出一个由该元素的 `target` 字段和 `index_id` 字段的值组成的字符串，这两个值之间以逗号分隔。
这种输出格式可能用于需要这种特定格式的场景，如生成 `CSV` 文件的一部分。


# adb

adb（Android Debug Bridge），安卓平台调试桥，是连接Android手机与PC端的桥梁。
通过adb可以管理、操作模拟器和设备，如安装软件、查看设备软硬件参数、系统升级、运行shell命令等。

相关日志：
```
brew install --cask android-platform-tools
adb version
adb devices
adb shell
adb root
# 查看日志
adb logcat
setenforce 0
```

蓝牙抓包：
```
// ubuntu 22.04 上安装 ubertooth 
sudo apt install ubertooth
// 
adb pull PATH
// ubertooth-rx 是抓取命令 -t 20 表示抓 20s 然后退出 -r <file> 表示抓到的包在本地存储的文件名是什么。
ubertooth-rx -t 20 -r traffic.pcapng
```

# frp

[frp官方文档](https://gofrp.org/zh-cn/)  

[frp Github](https://github.com/fatedier/frp/blob/dev/README_zh.md)  

[版本下载](https://github.com/fatedier/frp/releases)

[conf/frps_full_example.toml](https://github.com/fatedier/frp/blob/dev/conf/frps_full_example.toml)

[conf/frpc_full_example.toml](https://github.com/fatedier/frp/blob/dev/conf/frpc_full_example.toml)

[frp 内网穿透 实现MAC远程桌面](https://yuqiangcoder.com/2019/11/22/frp-%E5%86%85%E7%BD%91%E7%A9%BF%E9%80%8F-%E5%AE%9E%E7%8E%B0MAC%E8%BF%9C%E7%A8%8B%E6%A1%8C%E9%9D%A2.html)

frp 采用 C/S 模式，将服务端部署在具有公网 IP 的机器上，客户端部署在内网或防火墙内的机器上，通过访问暴露在服务器上的端口，反向代理到处于内网的服务。
在此基础上，frp 支持 TCP, UDP, HTTP, HTTPS 等多种协议，提供了加密、压缩，身份认证，代理限速，负载均衡等众多能力。
此外，还可以通过 xtcp 实现 P2P 通信。

## Darwin、Freebsd、Linux

关于 xxx_darwin_amd64.tar.gz和 xxx_darwin_arm64.tar.gz和xxx_freebsd_amd64.tar.gz 的区别.

Darwin 是一种类 Unix 作业系统，包含开放原始码的 XNU 内核， 其以微核心为基础的核心架构来实作 Mach，
而作业系统的服务和使用者空间工具则以 BSD 为基础。


* xxx_darwin_amd64.tar.gz: 这个文件适用于在 macOS（Darwin）操作系统上运行的 64 位 x86 架构的计算机。
* xxx_darwin_arm64.tar.gz: 这个文件适用于在 macOS（Darwin）操作系统上运行的 64 位 ARM 架构的计算机。
通常，这适用于使用 M1 芯片或其他 ARM 架构的 Mac 设备。
* xxx_freebsd_amd64.tar.gz: 这个文件适用于在 FreeBSD 操作系统上运行的 64 位 x86 架构的计算机。

mac系统：
```
uname -a
Darwin xxx-Pro.local 22.3.0 Darwin Kernel Version RELEASE_ARM64_T6000 arm64
```
centos系统：
```
[root@ opt]# uname -a
Linux 3.10.0-1062.12.1.el7.x86_64 #1 SMP  UTC 2020 x86_64 x86_64 x86_64 GNU/Linux
```

## amd64、arm64、mips、mips64、mips64le、mipsle 的区别是

这些都是不同的处理器架构，用于指代不同类型的计算机硬件。

1. **amd64（x86-64）**: 这是目前大多数个人计算机和服务器上使用的 64 位 x86 处理器架构。它是由 AMD 公司推出的，后来被 Intel 公司采用，因此也被称为 x86-64。
2. **arm64（AArch64）**: 这是一种用于移动设备（如智能手机和平板电脑）和服务器的 64 位 ARM 处理器架构。它由 ARM 公司设计，提供了更高的性能和能效比，因此在移动设备和云计算中得到广泛应用。
3. **mips**: MIPS 是一种早期的 RISC 处理器架构，广泛应用于嵌入式系统、网络设备和一些早期的个人计算机。然而，在通用计算领域，它的使用逐渐减少。
4. **mips64**: 这是 MIPS 处理器的 64 位版本，扩展了地址空间和寄存器位数，提供更高的性能和更大的内存支持。
5. **mipsle** 和 **mips64le**: 这两者都表示 Little-Endian 的 MIPS 架构。Little-Endian 和 Big-Endian 是数据存储的不同方式，区别在于字节的排列顺序。
在 Little-Endian 中，最低有效字节存储在最低地址，而在 Big-Endian 中，最高有效字节存储在最低地址。 
LE 版本通常用于 x86 等架构，而 BE 版本通常用于 MIPS 等。

端口设置：
```
webServer：7500、7400
serverPort：17000、16000
rdpServer: 17001
rdp：3389
vnc：5900
```
使用场景：

| 编号 | 类型  |描述|
|---|-----|-----|
|1|通过 SSH 访问内网机器	|简单配置 TCP 类型的代理让用户访问到内网的服务器|
|2	|自定义域名访问内网的 Web 服务|	简单配置 HTTP 类型的代理让用户访问到内网的 Web 服务|
|3	|转发 DNS 查询请求	|通过简单配置 UDP 类型的代理转发 DNS 查询请求|
|4	|转发 Unix 域套接字|	配置 Unix 域客户端插件来通过 TCP 端口访问内网的 Unix 域服务|
|5|对外提供简单的文件访问服务	|配置 static_file 客户端插件将本地文件暴露在公网上供其他人访问|
|6	|为本地 HTTP 服务启用 HTTPS	|https2http 插件可以让本地 HTTP 转换成 HTTPS 服务对外提供|
|7	|安全地暴露内网服务	|将会创建一个只有自己能访问到的 SSH 服务代理|
|8	|点对点内网穿透	|将会演示一种不通过服务器中转流量的方式来访问内网服务|

### 1、通过 SSH 访问内网机器

* 可以安全的暴露内网服务。对于某些服务来说如果直接暴露于公网上将会存在安全隐患。
* 使用 stcp(secret tcp) 类型的代理可以避免让任何人都能访问到要穿透的服务，但是访问者也需要运行另外一个 frpc 客户端。
* 在有公网IP的机器(已部署frps服务)，配置frps.toml文件。如果配置文件如下，注意对应服务器安全组规则放行17000、16000端口。


```
# 设置frp服务器用户接收客户端连接的端口
bindPort = 17000
auth.method = "token"
auth.token = "xxx"
```

配置好之后，运行相关命令，启动服务端
```
./frps -c ./frps.toml
```
在后台使用服务：
```
nohup ./frps -c ./frps.toml &
```
可以使用jobs查看当前运行的任务，查看进程信息：
```
[root@ frp_0.52.3_linux_amd64]# jobs
[1]+  Running                 nohup ./frps -c ./frps.toml &
[root@ frp_0.52.3_linux_amd64]# ps -ef | grep "frps"
root  927 479  0 16:11 pts/0 00:00:00 ./frps -c ./frps.toml
```
* 在内网IP的机器(部署frpc服务)，配置frpc.toml文件。
```
user = "xxx"
serverAddr = "xxx.xx.xx.xx"
serverPort = 17000

loginFailExit = true
auth.method = "token"
auth.token = "xxx"

[[proxies]]
name = "tcp-22"
type = "tcp"
localIP = "127.0.0.1"
localPort = 22
remotePort = 16000
```
如果要考虑安全的方式，可以配置以下信息：
```
[[proxies]]
name = "secret_tcp"
# If the type is secret tcp, remotePort is useless
# Who want to connect local port should deploy another frpc with stcp proxy and role is visitor
type = "stcp"
# secretKey is used for authentication for visitors
secretKey = "xxxx"
localIP = "127.0.0.1"
localPort = 22
# If not empty, only visitors from specified users can connect.
# Otherwise, visitors from same user can connect. '*' means allow all users.
allowUsers = ["*"]
```
配置好之后，运行相关命令，启动客户端
```
./frpc -c ./frpc.toml
```
* 通过SSH访问内网机器。
```
# frp会将请求x.x.x.x:16000的流量转发到内网机器的22端口
$ ssh -oPort=16000 test@x.x.x.x
```

如果要开启dashboard，则在`frps.toml`文件中配置如下信息：
```
webServer.addr = "0.0.0.0"
webServer.port = 7500
webServer.user = "xxx"
webServer.password = "xxx"
```
配置好后重启frps的服务端，访问以下地址并输入相应的用户名和密码即可访问：
```
serverAddr:webServer.port
```

### 2、通过自定义域名访问内网的 Web 服务
* `frps.toml` 配置：
```
vhostHTTPPort = 8080
vhostHTTPSPort = 443
vhostHTTPTimeout = 60
```
* `frpc.toml`配置
```
[[proxies]]
name = "web01"
type = "http"
localIP = "127.0.0.1"
localPort = 80
httpUser = "xxx"
httpPassword = "xxx"
```
###  frp 实现 Microsoft remote desktop 远程桌面连接

为什么使用 RDP？

RDP 是 Windows 自带的协议。相比市面上的远程桌面软件，如 Todesk、Anydesk、向日葵等，有以下优势：

* 兼容性更佳，根据设备自适应分辨率，可连接键鼠使用。
* 自由度高，不限制设备数量，也没有会员体系。
* 连接速度取决于电脑网速和服务器配置。

需要在 windows 的设置->系统->远程桌面->点击开启远程桌面。

配置好`frpc.toml` 信息：
```
[[proxies]]
name = "tcp-windows-rdp"
type = "tcp"
localIP = "127.0.0.1"
localPort = 3389
remotePort = 17001
```
配置好之后，运行相关命令，启动客户端：
```
./frpc -c ./frpc.toml
```
点击Microsoft remote desktop，在连接中心点击+号 ,点击Add PC。

```
PC name：填写公网服务IP:rdp的remotePort（上述示例为17001）
User account：PC机的账号
密码：PC机的密码。如果的PC机原来没有密码，请设置密码后，再远程桌面，安全第一。
```

### frp 实现MAC远程桌面

什么是VNV?

VNC（Virtual Network Computing），为一种使用RFB协议的屏幕画面分享及远程操作软件。
此软件借由网络，可发送键盘与鼠标的动作及即时的屏幕画面。

服务端配置同上。主要介绍关于`frpc`客户端的配置：

* localPort 为 VNC 的默认端口号 5900。
* remotePort 为屏幕共享时所输入的IP地址的端口号。

```
[[proxies]]
name = "tcp-mac-vnc"
type = "tcp"
localIP = "127.0.0.1"
localPort = 5900
remotePort = 17003
transport.useEncryption = true
transport.useCompression = true
```
配置好之后，运行相关命令，启动客户端：
```
./frpc -c ./frpc.toml
```

Mac 测试连接，需要开启mac电脑的屏幕共享（也就是5900端口，系统设置里面，搜索共享，打开屏幕共享）：

* 使用快捷键 command + k 唤出连接服务器页面。
* 输入 公网 IP 和 remotePort 指定的端口号。例如：vnc://xx.xx.xx.xx:17003
* 输入用户名和密码进行连接即可。


相对于远程桌面软件的优势：

* 不需要下载各种远程软件（Teamviewer、向日葵、toDesk）。但是需要下载包并写入配置文件。
* 可以限制端口和token，不让任意端口连接，如果服务端配置了token，需在客户端加入同样token才能连接到服务器。


# Go 清单
[Go 语言高性能编程](https://geektutu.com/post/high-performance-go.html)

[Go 语言设计与实现](https://draveness.me/golang/)

# Git 清单
[Git-中文版](https://git-scm.com/book/zh/v2)