# 基础命令

`mongodb` 查询：
```
db.getCollection("message_history").find().limit(1000).skip(0)
db.message_history.aggregate([
   { $limit: 1000 },
   { $skip: 0 }
])
db.message_history.aggregate([
{
      $match: {
         role_id: "xxx",
         create_time: { $gt: 1650211200 }
      }
   },
   {
      $group: {
         _id: null,
         total: { $sum: { $toInt: "$len" } }
      }
   }
])
```


连接 mongodb 命令：

```shell
[root@wangxiong /]# mongo
MongoDB shell version v4.4.1
connecting to: mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("eb16dec1-6642-aff3-d86bbb3d533e") }
MongoDB server version: 4.2.5
```

连接 mongodb 后，通过以下命令查看版本：

```shell
> db.version();
4.2.5
```

相对路径使用 mongo 命令查看版本：

```shell
[root@wangxiong /]# whereis mongo
mongo: /usr/bin/mongo /usr/share/man/man1/mongo.1
[root@wangxiong /]# /usr/bin/mongo --version
MongoDB shell version v4.4.1
Build Info: {
    "version": "4.4.1",
    "gitVersion": "ad91a93a5a31e175f5cbf8c69561e788bbc55ce1",
    "openSSLVersion": "OpenSSL 1.0.1e-fips 11 Feb 2013",
    "modules": [],
    "allocator": "tcmalloc",
    "environment": {
        "distmod": "rhel70",
        "distarch": "x86_64",
        "target_arch": "x86_64"
    }
}
```

启动mongo：

```powershell
[root@wangxiong bin]# /usr/bin/mongod -f  /etc/mongod.conf 
child process started successfully, parent exiting
```

关闭mongo:

```powershell
[root@wangxiong bin]# /usr/bin/mongod  --shutdown  --dbpath  /var/lib/mongo
killing process with pid: 23419
```



```shell
mongoimport --host 10.66.187.127:27017 -u mongouser -p thepasswordA1 --authenticationDatabase=admin --db=testdb --collection=testcollection2 --file=/data/export_testdb_testcollection.json

mongoimport  --db=yapi  --collection=testcollection2 --file=/export/servers/my-yapi/yapi.sql
```

