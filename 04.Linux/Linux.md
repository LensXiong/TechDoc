# 列表

* [查看服务器系统及硬件信息。](#lsb_cpu_free)
* [常用命令。](#common_command)
* [ Centos 静态IP设置。](#centos_static_ip)
* [ Ubuntu 静态IP设置。](#ubuntu_static_ip)
* [ Docker-compose 中两个桥接网络导致 Telnet 不通。](#bridge_gateway)
* [ 解决 CentOS8 查看网络管理服务配置，并设置开机自启。](#network_scripts_centos8)
* [ 解决 CentOS7 查看网络管理服务配置，并设置开机自启。](#network_scripts_centos7)

<span id="lsb_cpu_free">查看服务器系统及硬件信息。</span>
查看服务器发行版本、CPU型号、CPU核数、硬盘大小和内存大小：

```
xxx@xxxx:~# lsb_release -d | awk -F"\t" '{print "发行版本: "$2}';\
cat /proc/cpuinfo | grep "model name" | uniq | awk -F":" '{print "CPU型号: "$2}';\
cat /proc/cpuinfo | grep "cpu cores" | uniq | awk -F":" '{print "CPU核数: "$2}';\
sudo parted -l | grep "Disk /" | uniq | awk -F":" '{print "硬盘大小:" $2 " G",$2/1024 " TB"}';\
free -h | awk '/Mem:/{printf "内存大小: %s\n", $2}'
发行版本: Ubuntu 22.04.1 LTS
CPU型号:  Intel(R) Core(TM) i5-2400 CPU @ 3.10GHz
CPU核数:  4
硬盘大小: 1000GB G 0.976562 TB
内存大小: 8G
```
查看 `Ubuntu` 版本信息，包括发行版本、发行代号和描述信息：

```
root@xxx:~# lsb_release -a
No LSB modules are available.
Distributor ID:	Ubuntu
Description:	Ubuntu 22.04.1 LTS
Release:	22.04
Codename:	jammy
```
硬盘的分区信息：

```
root@xxx:~# sudo parted -l
Model: ATA ST1000LM048-2E71 (scsi)
Disk /dev/sda: 1000GB
Sector size (logical/physical): 512B/4096B
Partition Table: msdos
Disk Flags:

Number  Start   End     Size    Type     File system  Flags
1      1049kB  256MB   255MB   primary  fat32        boot, esp
2      256MB   1000GB  1000GB  primary  ext4
```
这是一个在`Ubuntu`系统中使用的命令，下面是每个组成部分的详细解释：

* `parted`: 这是一个`Linux`分区工具，可用于对硬盘进行分区、格式化、重命名等操作。
* `-l`: 这是`parted`命令的选项之一，用于显示当前系统上的所有硬盘和分区的详细信息。
* `Model`: `ATA ST10xxxxx-2E71 (scsi)`: 此行显示了硬盘的型号和接口类型。
* `Disk /dev/sda: 1000GB`此行显示了硬盘的设备文件名和总容量，这里的硬盘是`/dev/sda`，总容量是1000GB。
* `Sector size (logical/physical): 512B/4096B:` 此行显示了硬盘扇区的逻辑和物理大小，这里的逻辑大小是512B，物理大小是4096B。
* `Partition Table: msdos: `此行显示了硬盘分区表的类型，这里使用的是传统的`msdos`分区表。
* `Disk Flags:`此行显示了硬盘的标志，这里没有设置任何标志。

下面是硬盘的分区信息：
```
Number: 分区的编号。
Start: 分区在硬盘上的起始位置。
End: 分区在硬盘上的结束位置。
Size: 分区的大小。
Type: 分区的类型，这里的类型分别为primary（主分区）和extended（扩展分区）。
File system: 分区的文件系统类型，这里分别为fat32和ext4。
Flags: 分区的标志，这里分别为boot（引导分区）和esp（EFI系统分区）。
```

查看 CPU 核数和 CPU 信息：

```
root@xxx:~# cat /proc/cpuinfo
processor	: 0
vendor_id	: GenuineIntel
cpu family	: 6
model		: 42
model name	: Intel(R) Core(TM) i5-2x00 CPU @ 3.x0GHz
stepping	: 7
microcode	: 0x2f
cpu MHz		: 3288.426
cache size	: 6144 KB
physical id	: 0
siblings	: 4
core id		: 0
cpu cores	: 4
apicid		: 0
initial apicid	: 0
fpu		: yes
fpu_exception	: yes
cpuid level	: 13
wp		: yes
flags		: fpu vme 
bugs		: cpu_meltdown
bogomips	: 6185.79
clflush size	: 64
cache_alignment	: 64
address sizes	: 3x bits physical, 4x bits virtual
power management:
```

以G为单位，显示硬盘的使用情况：

```
root@codriver-desktop:~# df -Th
Filesystem     Type     Size  Used Avail Use% Mounted on
tmpfs          tmpfs    785M  2.8M  782M   1% /run
/dev/sda2      ext4     912G   33G  837G   4% /
tmpfs          tmpfs    3.9G     0  3.9G   0% /dev/shm
tmpfs          tmpfs    5.0M  4.0K  5.0M   1% /run/lock
/dev/sda1      vfat     240M  5.3M  235M   3% /boot/efi
tmpfs          tmpfs    785M   76K  785M   1% /run/user/127
tmpfs          tmpfs    785M   60K  785M   1% /run/user/0
overlay        overlay  916G   33G  837G   4% /
```

以G为单位，显示内存的使用情况：

```
root@codriver-desktop:~# free -h
total        used        free      shared  buff/cache   available
7.7Gi       1.9Gi       352Mi       9.0Mi       5.4Gi       5.4Gi
```


<span id="network_scripts_centos8">解决 CentOS8 查看网络管理服务配置，并设置开机自启。</span>

问题场景：`ssh`登录远程服务器时出现 `connect to host 192.16x.1.x port 22: Host is down`。

先检查静态 IP 配置信息，确认 ONBOOT 为 yes：
```
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=static
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
NAME=eno1
UUID=xxxxxxx
DEVICE=eno1
ONBOOT=yes
IPADDR=192.168.1.x
PREFIX=24
GATEWAY=192.168.1.1
```
设置网络服务的开机自启：
```
[root@localhost ~]# cat /etc/redhat-release
CentOS Linux release 8.4.2105
# 查看本机的网卡
[root@localhost ~]# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: eno1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 50:eb:f6:b7:12:44 brd ff:xxxf:xx:ff:ff:ff
    inet 192.xxx.1.3/24 brd 192.xxx.1.255 scope global noprefixroute eno1
       valid_lft forever preferred_lft forever
    inet6 fe80::52eb:f6ff:feb7:6244/64 scope link noprefixroute
       valid_lft forever preferred_lft forever
# 查看 eth0 的网卡信息
[root@localhost ~]# ifconfig eno1
eno1: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.16x.1.x  netmask 255.255.255.0  broadcast 192.16x.1.255
        inet6 fe80::52eb:fzxxx:feb7:6244  prefixlen 64  scopeid 0x20<link>
        ether 50:eb:f6:b7:62:44  txqueuelen 1000  (Ethernet)
        RX packets 20002  bytes 17065598 (16.2 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 21763  bytes 16970379 (16.1 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
        device interrupt 16  memory 0xa1200000-a1220000
[root@localhost ~]# ifconfig
br-8e439f7cc472: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.16.1.1  netmask 255.255.255.0  broadcast 172.16.1.255
        inet6 fe80::42:90ff:feb1:3ce9  prefixlen 64  scopeid 0x20<link>
        ether 02:42:90:b1:3c:e9  txqueuelen 0  (Ethernet)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 67  bytes 6905 (6.7 KiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
docker0: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255
        ether 02:42:e3:af:83:b6  txqueuelen 0  (Ethernet)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
# 查看网络服务管理状态并开启开机自启动
[root@localhost ~]# systemctl list-unit-files --type service | grep NetworkManager
NetworkManager-dispatcher.service          disabled
NetworkManager-wait-online.service         disabled
NetworkManager.service                     disabled
[root@localhost ~]# systemctl status NetworkManager.service  
# 开机启用 NetworkManager 服务
[root@localhost ~]# chkconfig NetworkManager on
Note: Forwarding request to 'systemctl enable NetworkManager.service'.
Created symlink /etc/systemd/system/multi-user.target.wants/NetworkManager.service → /usr/lib/systemd/system/NetworkManager.service.
Created symlink /etc/systemd/system/dbus-org.freedesktop.nm-dispatcher.service → /usr/lib/systemd/system/NetworkManager-dispatcher.service.
Created symlink /etc/systemd/system/network-online.target.wants/NetworkManager-wait-online.service → /usr/lib/systemd/system/NetworkManager-wait-online.service.
[root@localhost ~]# systemctl start NetworkManager.service 
[root@localhost ~]# systemctl status NetworkManager.service
● NetworkManager.service - Network Manager
   Loaded: loaded (/usr/lib/systemd/system/NetworkManager.service; enabled; vendor preset: enabled)
   Active: active (running) since Tue 20xx-xx-15 15:43:09 CST; 34min ago
     Docs: man:NetworkManager(8)
 Main PID: 1238 (NetworkManager)
    Tasks: 3 (limit: 202653)
   Memory: 9.4M
   CGroup: /system.slice/NetworkManager.service
           └─1238 /usr/sbin/NetworkManager --no-daemon  
[root@localhost ~]# systemctl list-unit-files --type service | grep NetworkManager
NetworkManager-dispatcher.service          enabled
NetworkManager-wait-online.service         enabled
NetworkManager.service                     enabled
```

<span id="network_scripts_centos7">解决 CentOS7 查看网络管理服务配置，并设置开机自启。</span>

问题场景：`ssh`登录远程服务器时出现 `connect to host 192.16x.1.x port 22: Host is down`。
```
# 查看系统版本
[root@localhost ~]# cat /etc/redhat-release
CentOS Linux release 7.4.1708 (Core)
# 查看网络服务管理的自启动配置状态
[root@localhost ~]# systemctl --type service list-unit-files | grep NetworkManager
NetworkManager-dispatcher.service          disabled
NetworkManager-wait-online.service         disabled
NetworkManager.service                     disabled
# 查看网络管理服务状态
[root@localhost ~]# systemctl status NetworkManager.service
● NetworkManager.service - Network Manager
   Loaded: loaded (/usr/lib/systemd/system/NetworkManager.service; disabled; vendor preset: enabled)
   Active: inactive (dead)
     Docs: man:NetworkManager(8)
# 如果未启动，设置服务开机自启动
[root@localhost ~]# systemctl enable NetworkManager.service
Created symlink from /etc/systemd/system/dbus-org.freedesktop.NetworkManager.service to /usr/lib/systemd/system/NetworkManager.service.
Created symlink from /etc/systemd/system/multi-user.target.wants/NetworkManager.service to /usr/lib/systemd/system/NetworkManager.service.
Created symlink from /etc/systemd/system/dbus-org.freedesktop.nm-dispatcher.service to /usr/lib/systemd/system/NetworkManager-dispatcher.service.
# 再次查看网络管理的自启动配置状态
[root@localhost ~]# systemctl list-unit-files --type service | grep NetworkManager
dbus-org.freedesktop.NetworkManager.service   enabled
NetworkManager-dispatcher.service             enabled
NetworkManager-wait-online.service            disabled
NetworkManager.service                        enabled
```

<span id="bridge_gateway">Centos 静态IP设置。</span>
问题：telnet 内网下另一台服务中某个端口无法访问。
```
networks:
skygo_net:
driver: bridge
enable_ipv6: false
ipam:
driver: default
config:
- subnet: 172.16.1.0/24
gateway: 172.16.1.1

[root@web-1 config]# ifconfig
br-783ac168832a: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.16.1.1  netmask 255.255.255.0  broadcast 172.16.1.255
        inet6 fe80::42:98ff:fe80:1648  prefixlen 64  scopeid 0x20<link>
        ether 02:42:98:80:16:48  txqueuelen 0  (Ethernet)
        RX packets 274992  bytes 209994168 (200.2 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 440790  bytes 358828865 (342.2 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
br-xxxxxxxxxxx: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.16.1.1  netmask 255.255.255.0  broadcast 172.16.1.255
        inet6 fe80::42:98ff:fe80:1648  prefixlen 64  scopeid 0x20<link>
        ether 02:42:98:80:16:48  txqueuelen 0  (Ethernet)
        RX packets 274992  bytes 209994168 (200.2 MiB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 440790  bytes 358828865 (342.2 MiB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```
解决办法：关闭多余的网卡，通过 ifconfig 查看到桥接 172.16.1.1 中有两个网络。
```
ifconfig br-xxxxxx down
```
注：`ifconfig eth0 up` 为启动网卡`eth0` ；`ifconfig eth0 down` 为关闭网卡 `eth0`。`ssh`登陆`linux`服务器操作要小心，关闭了就不能开启了，除非你有多网卡。


<span id="centos_static_ip">Centos 静态IP设置。</span>

<span id="common_command">Linux 常用命令</span>

```
ps -ef | grep php-fpm
ps aux|grep nginx
netstat -tunpl | grep 3306 # 查看端口占用情况（redis-6379，mysql-3306）
tar -czvf file.tar.gz file # 以 gzip 压缩打包file文件夹并命名为file.tar.gz（显示打包详细过程）。
tar -xzvf file.tar.gz file # 解压file.tar.gz包并命名为file（显示解压详细过程）
tail -f xxx.log # 循环读取日志文件的内容
tail -n +20 notes.log # 显示最后20行的日志文件内容
lsof -i :3306 # 列出谁在使用这个端口

```

```shell
[root@iZ2z ~]# ps -ef | grep php-fpm
root      9458     1  0 Mar05 ?        00:00:53 php-fpm: master process (/usr/local/php7/etc/php-fpm.conf)
www      18260  9458  0 Mar31 ?        00:04:21 php-fpm: pool www
www      18263  9458  0 Mar31 ?        00:04:23 php-fpm: pool www
www      18270  9458  0 Mar31 ?        00:03:54 php-fpm: pool www
root     19420 19396  0 13:12 pts/1    00:00:00 grep php-fpm
```



```shell
[root@iZ2 ~]# ps aux|grep nginx
root      9453  0.0  0.3  54880  6092 ?        Ss   Mar05   0:00 nginx: master process nginx -c /usr/local/nginx/conf/nginx.conf
www       9578  0.0  1.8  87860 36132 ?        S    Mar05   5:51 nginx: worker process
root     19422  0.0  0.0 103296   784 pts/1    S+   13:13   0:00 grep nginx
```



```shell
[root@iZ2z ~]# netstat -tunpl | grep 3306
tcp        0      0 0.0.0.0:3306                0.0.0.0:*                   LISTEN      1340/mysqld
```



```shell
[root@iZ2z ~]# lsof -i
COMMAND     PID      USER   FD   TYPE  DEVICE SIZE/OFF NODE NAME
memcached  1019 memcached   26u  IPv4    8983      0t0  TCP *:memcache (LISTEN)
memcached  1019 memcached   27u  IPv4    8992      0t0  UDP *:memcache
[root@iZ2z ~]# lsof -i :3306
COMMAND  PID  USER   FD   TYPE DEVICE SIZE/OFF NODE NAME
mysqld  1340 mysql   13u  IPv4   9791      0t0  TCP *:mysql (LISTEN)
```



2、

```shell
sed 's/7000/7001/g' redis7000/redis.conf > redis7001/redis.conf
```






