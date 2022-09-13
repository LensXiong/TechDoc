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