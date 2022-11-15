
# 容器中 mysql 执行sql数据
```
docker cp  ./xxx.sql  xxx_mysql:/home
mysql -uroot -p
use xxx
source xxx

```

# 查看容器映射目录
```
[root@v merged]# docker inspect skygo_mysql | grep Merge
                "MergedDir": "/opt/skygo/data/docker/overlay2/ba607d1a14f0ea9e6ae1734a87543e494fb0f431520e2b889f75965166a5c5f6/merged",
[root@ merged]# cd /opt/skygo/data/docker/overlay2/ba607d1a14f0ea9e6ae1734a87543e494fb0f431520e2b889f75965166a5c5f6/merged
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