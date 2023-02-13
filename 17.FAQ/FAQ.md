
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