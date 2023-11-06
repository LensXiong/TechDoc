





# 使用 docker-compose 启动 kafka 报错

报错信息：
```
ERROR Fatal error during KafkaServer startup. Prepare to shutdown (kafka.server.KafkaServer)
kafka.common.InconsistentClusterIdException: The Cluster ID ZfQZirUQRua6RnLVYiz_rA doesn't match stored clusterId Some(iFvDhNp5TP2jnCpEai461Q) in meta.properties. The broker is trying to join the wrong cluster. Configured zookeeper.connect may be wrong.
	at kafka.server.KafkaServer.startup(KafkaServer.scala:223)
	at kafka.server.KafkaServerStartable.startup(KafkaServerStartable.scala:44)
	at kafka.Kafka$.main(Kafka.scala:82)
	at kafka.Kafka.main(Kafka.scala)
```
报错原因：

Kafka Broker 尝试加入了一个错误的集群，原因是其集群ID（Cluster ID）与之前存储在meta.properties文件中的集群ID不匹配。这通常发生在以下情况下：
* Kafka配置文件更改：如果你更改了Kafka Broker的配置文件，尤其是broker.id或zookeeper.connect等配置项，可能导致集群ID不匹配。
* ZooKeeper连接配置错误：zookeeper.connect配置项指定了Kafka Broker连接ZooKeeper的信息。确保这个配置正确，并且Kafka Broker可以连接到正确的ZooKeeper集群。

解决办法：
* 检查Kafka配置文件：检查Kafka Broker的配置文件，特别是broker.id和zookeeper.connect的值是否正确。确保broker.id是唯一的，并且zookeeper.connect指向正确的ZooKeeper集群。
* 删除错误的Cluster ID：在Kafka数据目录中，有一个meta.properties文件，其中存储了Cluster ID。如果你确定配置正确，你可以尝试删除这个文件，然后重新启动Kafka Broker。Kafka将会重新生成正确的Cluster ID。
* 清除ZooKeeper数据：如果上述步骤没有解决问题，可能需要清除ZooKeeper中的一些数据。在做任何数据清除之前，务必备份数据以防万一。你可以尝试删除ZooKeeper数据目录中与Kafka相关的节点，然后重新启动Kafka Broker。

具体步骤：
```
rm -rf /opt/xxxx/data/db/kafka/data/meta.properties
docker-compose restart xxx_kafka
```

# 初始化mysql数据备份迁移到另一台mysql

docker 启动 mysql 容器时报错：
```
[System] [MY-013169] [Server] /usr/sbin/mysqld (mysqld 8.0.30) initializing of server in progress as process 81
[ERROR] [MY-010457] [Server] --initialize specified but the data directory has files in it. Aborting.
[ERROR] [MY-013236] [Server] The designated data directory /var/lib/mysql/ is unusable. You can remove all files that the server added to it.
[ERROR] [MY-010119] [Server] Aborting
```
原因：
```
xxx/data/db/mysql下应该是所需要的mysql初始化数据。包括表结构、数据库、日志数据等。
```
解决：
```
保证xxx/data/db/mysql下的初始化数据正常。之前遇到xxx/data/db/mysql/里面还有一层data数据，将data数据移出来即可。
```

# docker 导入和导出相关数据
## 导入数据
```
docker exec -i xxx_mysql sh -c 'mysql -uxxxx -pxxxxx -D database_name' </opt/xxxx/mysqldump_xxxxx.sql
docker exec -i xxx_mongo sh -c 'mongorestore -u xxx -p xxxx --authenticationDatabase admin -d xxxx --drop --archive' </opt/xxxx/mongodump_xxxx.archive
```

* `sh -c`：调用 `shell (sh)` 并允许在其中执行一个命令 `(-c)`。
* `--drop`：在还原之前删除目标数据库的现有数据。
* `--archive`：指定还原的数据来源为归档文件。
* `</opt/xxxx/mongodump_xxxx.archive`：这部分是输入重定向。它将 `mongodump_xxxx.archive` 文件的内容作为输入传递给容器内部的 `mongorestore` 命令.

## 导出数据

```
docker exec -i xxx_mysql sh -c 'mysqldump -uxxx -pxxxx --skip-add-locks --single-transaction -q xxxx' >/opt/mysqldump_$(date '+%Y%m%d').sql
docker exec -i xxx_mongo sh -c 'mongodump -u admin -p xxxx --authenticationDatabase admin -d xxxx --archive' >/opt/mongodump_$(date '+%Y%m%d').archive
```

* `--skip-add-locks`：在导出期间跳过添加锁定语句。
* `--single-transaction`：使用单个事务导出数据，确保一致性。
* -q：使用"快速"模式导出数据，以减少导出时的负载。


# Docker 中与 MTU 相关的信息
## 什么是 MTU ?

MTU 指的是“最大传输单元”（Maximum Transmission Unit），是计算机网络中的一个术语，它是指在一个网络中，能够通过一次发送的数据包的最大大小。

网络中的每个设备和协议都有其自己的MTU限制。当数据包的大小超过某个设备或协议的MTU限制时，
该数据包将被分割成更小的数据包进行传输，这会增加网络的负载和延迟。

MTU 的大小通常以字节为单位进行计算。在以太网中，MTU 的默认值为1500字节，而在其他网络协议中可能会有不同的值。
管理员可以在网络设备中配置 MTU 大小，以满足特定的网络需求和性能要求。


## 如何理解 Docker 中的 MTU。
在 Docker 中，MTU 是指容器网络中的最大传输单元。与主机上的网络接口类似，Docker 容器也有其自己的网络接口和 MTU 值。
Docker 网络使用 bridge 驱动程序来提供容器网络连接。在此模式下，Docker 会创建一个虚拟网络桥接设备，该设备用于将多个容器连接到同一网络中。

在 Docker 中，容器的默认 MTU 大小为1500字节，与大多数以太网设备的默认 MTU 大小相同。
管理员可以使用 --mtu 选项在创建容器时指定容器的 MTU 值。通过调整 MTU 大小，管理员可以在容器网络中优化性能和避免网络分段。

例如，在高负载网络环境中，降低 MTU 大小可以减少分段和延迟，从而提高网络性能。
另外，一些特定的网络协议和设备可能需要特定的 MTU 值才能正常工作，管理员也可以根据需要调整 MTU 大小来支持这些协议和设备。

## Docker 中 python pip 网络的问题有哪些？
在 Docker 中使用 Python 和 pip 安装第三方包时可能会遇到网络问题，包括但不限于以下几种：

* 防火墙和代理问题：如果 Docker 主机上设置了防火墙或代理，可能会阻止 Docker 容器与外部网络进行通信，导致 pip 安装失败。

* DNS 问题：Docker 容器默认使用主机上的 DNS 服务器，如果主机上的 DNS 服务器出现问题，容器中的 pip 也无法正常解析域名。

* MTU 问题：在某些情况下，MTU 大小可能会影响 Docker 容器中 pip 的网络连接。如果默认的 MTU 值不适用于特定的网络环境，可能会导致网络连接失败。

* pip 源问题：默认的 pip 源可能会被屏蔽或限速，如果需要更快速、稳定的安装，可以选择其他的 pip 源，例如使用阿里云源、清华源等。

* 版本兼容性问题：有些 Python 包的不同版本可能存在不同的依赖关系和兼容性问题，可能会导致 pip 在 Docker 容器中无法正常安装或运行。

为了避免这些问题，可以尝试以下几种解决方法：

* 在 Dockerfile 中配置合适的防火墙和代理规则，确保容器可以与外部网络正常通信。

* 在 Docker 容器中设置 DNS 服务器或使用其他 DNS 解析方案，确保 pip 可以正常解析域名。

* 在 Docker 容器中调整 MTU 大小，以确保网络连接的稳定性和可靠性。

* 在 pip 安装时指定更快速、稳定的 pip 源，以提高安装速度和可靠性。

* 了解 Python 包的版本兼容性和依赖关系，尽可能使用与其他环境相同的版本，以确保在 Docker 容器中可以正常安装和运行。

## 在 Docker 容器中如何调整 MTU 大小？
在 Docker 容器中，可以通过以下几种方式来调整 MTU 大小：

① 在创建容器时使用 `--mtu` 选项指定 MTU 大小：
```
docker run --mtu=1400 myimage
```
此命令将在创建名为 mycontainer 的容器时将其 MTU 大小设置为 1400 字节。
② 在 Docker Compose 中使用 networks 关键字指定 MTU 大小：
```yaml
networks:
  mynetwork:
    driver: 06.bridge
    driver_opts:
      com.docker.network.driver.mtu: "1400"
```
这个配置将在 `Docker Compose` 中使用 `mynetwork` 网络时将其 MTU 大小设置为 1400 字节。
③ 在 Docker 主机中设置默认的 MTU 大小，从而使所有容器使用相同的 MTU 大小。
可以通过编辑 `/etc/docker/daemon.json` 文件并添加以下内容来实现：
```
{
  "mtu": 1400
}
```
这个配置将使 Docker 主机上的所有容器的 MTU 大小设置为 1400 字节。

需要注意的是，MTU 大小应该与所连接的网络和设备相匹配，以避免网络分段和连接问题。
在调整 MTU 大小时，应该进行测试和调试，以确保网络连接的稳定性和可靠性。

## ip link show 和 ip addr 的区别？
`ip link show` 适合查看网络接口的基本信息，而 `ip addr` 则提供了更详细的网络配置信息。
`ip link show` 命令会列出所有的网络接口，并显示它们的状态、MAC 地址、MTU 等基本信息。
此命令的输出格式更加简洁，适合快速浏览和查找网络接口。例如，以下是 `ip link show` 的示例输出：
```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN mode DEFAULT group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP mode DEFAULT group default qlen 1000
    link/ether 08:00:27:52:3c:fa brd ff:ff:ff:ff:ff:ff
```
相比之下，ip addr 命令提供了更详细的网络信息，包括每个网络接口的 IP 地址、广播地址、子网掩码等。此命令的输出格式更加详细，适合进行网络故障排除和调试。例如，以下是 ip addr 的示例输出：
```
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever

2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 08:00:27:52:3c:fa brd ff:ff:ff:ff:ff:ff
    inet 192.168.0.2/24 brd 192.168.0.255 scope global dynamic enp0s3
       valid_lft 86396sec preferred_lft 86396sec
    inet6 fe80::a00:27ff:fe52:3cfa/64 scope link
       valid_lft forever preferred_lft forever
```
# 容器中 mysql 执行sql数据
```
docker cp  ./xxx.sql  xxx_mysql:/home
mysql -uroot -p
use xxx
source xxx

```

# 查看容器映射目录
```
[root@v merged]# docker inspect xxxx_mysql | grep Merge
                "MergedDir": "/opt/xxxx/data/docker/overlay2/ba607d1a14f0ea9e6ae1734a87543e494fb0f431520e2b889f75965166a5c5f6/merged",
[root@ merged]# cd /opt/xxxx/data/docker/overlay2/ba607d1a14f0ea9e6ae1734a87543e494fb0f431520e2b889f75965166a5c5f6/merged
[root@v merged]# ll etc/mysql
total 12K
drwxr-xr-x 1 root root   94 Dec 21  2021 ./
drwxr-xr-x 1 root root 4.0K Sep 22 10:52 ../
drwxr-xr-x 1 root root   62 Dec 21  2021 conf.d/
lrwxrwxrwx 1 root root   24 Dec 21  2021 my.cnf -> /etc/alternatives/my.cnf
-rw-r--r-- 1 root root  839 Jul 10  2016 my.cnf.fallback
-rw-r--r-- 1 root root 1.5K Jan  5  2021 mysql.cnf
drwxr-xr-x 1 root root   24 Dec 21  2021 mysql.conf.d/
[root@web-1 merged]# vim  etc/mysql/mysql.conf.d/mysqld.cnf
max_allowed_packet=16777216
innodb_log_file_size=256M
```

# 命令

```
//  将指定镜像保存成 tar 归档文件，将镜像 runoob/ubuntu:v3 生成 my_ubuntu_v3.tar 文档
docker save -o my_ubuntu_v3.tar runoob/ubuntu:v3

// 导入镜像
docker load < xxx.tar.gz
```


## cp
`docker cp` :用于容器与主机之间的数据拷贝。
从容器中拷贝至主机：
```
格式：docker cp [OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-
示例：docker cp xxx_server:/root/xx/xxx /root/xxx/
```
从主机拷贝至容器：
```
格式：docker cp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH
示例：docker cp  ./xxx.so  xxx_analyze:/xxxServer/src/plugins/
```

# 安装

## rpm安装

[在 CentOS 上安装 Docker 引擎](https://docs.docker.com/engine/install/centos/)

[centos rpm包](https://download.docker.com/linux/centos/8/x86_64/stable/Packages/)