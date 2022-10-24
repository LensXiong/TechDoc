# 列表

* [常用命令。](#common_command)
* [ Centos 静态IP设置。](#centos_static_ip)
* [ Ubuntu 静态IP设置。](#ubuntu_static_ip)
* [ Docker-compose 中两个桥接网络导致 Telnet 不通。](#bridge_gateway)


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






