# API 调用
## 备份

## 迁移

## 脚本

## FQA
### 配置 JVM 的初始堆大小和最大堆大小
问题：导入大量数据的时候代码中报错如下信息。
```
err:elastic: Error 429 (Too Many Requests): [parent] Data too large, 
data for [<http_request>] would be [510313990/486.6mb],
which is larger than the limit of [510027366/486.3mb], real usage: [510313840/486.6mb], 
new bytes reserved: [150/150b], usages [request=49152/48kb, fielddata=0/0b, 
in_flight_requests=3035484/2.8mb, model_inference=0/0b, 
eql_sequence=0/0b, accounting=6200196/5.9mb] [type=circuit_breaking_exception]
```
原因：
```
jvm.options 文件是 Elasticsearch 中用于配置 Java 虚拟机 (JVM) 参数的文件。
在该文件中，-Xms1g 和 -Xmx1g 是用来配置 JVM 的初始堆大小和最大堆大小的参数。

-Xms1g: 这个参数设置 JVM 的初始堆大小为 1GB。初始堆大小是 JVM 启动时分配给堆内存的初始大小，也就是 Elasticsearch 在启动时会预留 1GB 的堆内存。
-Xmx1g: 这个参数设置 JVM 的最大堆大小为 1GB。最大堆大小是 JVM 堆内存的上限，JVM 在运行时会尝试将堆内存扩展到最大堆大小，但不会超过这个限制。
设置初始堆大小和最大堆大小的值可以根据你的 Elasticsearch 集群的实际情况进行调整。
如果你的集群处理大量数据或者有较多的查询负载，可能需要增大堆内存大小来提供更好的性能。当然，增大堆内存大小也会占用更多的系统资源，包括服务器的内存。
```
解决：
```
1、找到 Elasticsearch 的 JVM 配置文件： JVM 配置文件通常是 jvm.options，在 Elasticsearch 安装目录的 config 文件夹中。
2、修改堆内存设置： 打开 jvm.options 文件，查找以下两行。
-Xms1g
-Xmx1g
3、默认情况下，这两行表示初始堆大小 (-Xms) 和最大堆大小 (-Xmx) 都为 1GB。你可以根据实际情况将这两个值增大，例如：
-Xms4g
-Xmx4g
这样将初始堆和最大堆大小都设置为 4GB。请确保服务器有足够的物理内存支持你所设置的堆内存大小。
4、保存文件并重启 Elasticsearch： 保存 jvm.options 文件并重新启动 Elasticsearch 使配置生效。
注意： 在重启 Elasticsearch 之前，确保没有正在进行的重要操作，以免造成数据丢失或不可预料的情况。
5、监控 Elasticsearch 性能： 完成以上步骤后，监控 Elasticsearch 的性能，确保不再出现 "Too Many Requests" 错误。
同时，观察服务器的资源使用情况，确保服务器有足够的内存和其他资源来支持 Elasticsearch 运行。
```

### 模板示例
```
{
    "index_patterns": [
        "*-xxx-tags"
    ],
    "settings": {
        "index": {
            "number_of_shards": 3,
            "number_of_replicas": 1,
            "refresh_interval": "1s",
            "translog.flush_threshold_size": "1g",
            "max_result_window": 2147483647,
            "max_inner_result_window": 100000000
        }
    },
    "mappings": {
        "properties": {
            "xxx_tags": {
                "type": "object",
                "properties": {
                    "leaf": {
                        "type": "keyword"
                    },
                    "all": {
                        "type": "keyword"
                    }
                }
            },
            "xxx_tags": {
                "type": "object",
                "properties": {
                    "leaf": {
                        "type": "keyword"
                    },
                    "all": {
                        "type": "keyword"
                    }
                }
            }
        }
    }
}
```

### 索引示例
```
{
    "mappings": {
        "properties": {
            "xxx_date": {
                "type": "date",
                "format": "yyyy-MM-dd"
            },
            "xxxx_desc": {
                "type": "text"
            }
        }
    }
}
```
### 批量创建模板
es_template.sh
```
#!/bin/bash

url=$1
readDir=$2
user=$3
password=$4

auth=""
if [ -n "$user" ]; then
    auth="-u $user:$password"
fi

for df in `ls $readDir`
do
    if [ -f $readDir"/"$df ]; then
        indexName="${df/%.txt/}"
        echo ${indexName}
        filePath="@$readDir/$df"
        echo ${filePath}
        curl $auth -k -XPUT  ${url}/_template/${indexName} -H 'Content-Type: application/json'  --data $filePath
    fi
done
```
### 批量创建索引
```
#!/bin/bash

url=$1
readDir=$2
user=$3
password=$4

auth=""
if [ -n "$user" ]; then
    auth="-u $user:$password"
fi

for df in `ls $readDir`
do
    if [ -f $readDir"/"$df ]; then
        indexName="${df/%.txt/}"
        echo ${indexName}
        filePath="@$readDir/$df"
        echo ${filePath}
        curl $auth -k -XPUT  ${url}/${indexName} -H 'Content-Type: application/json'  --data $filePath
    fi
done 
```

## 模板

### 创建模板
```
PUT
http://xx.xx.xx.xx:9200/_template/name
{
    "index_patterns": [
        "xx-base-*"
    ],
    "settings": {
        "index": {
            "number_of_shards": 3,
            "number_of_replicas": 1,
            "refresh_interval": "1s",
            "translog.flush_threshold_size": "1g",
            "max_result_window": 2147483647,
            "max_inner_result_window": 100000000
        }
    },
    "mappings": {
        "properties": {
            "name": {
                "type": "keyword"
            },
            "xxx": {
                "type": "text"
            },
            "create_time": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss"
            },
            "update_time": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss"
            },
            "timestamp": {
                "type": "date",
                "format": "epoch_millis"
            }
        }
    }
}
```

### 获取所有模板名称
```
GET
http://xx.xx.xx.xx:9200/_cat/templates?v&h=name
```

### 获取具体模板
```
GET
http://xx.xx.xx.xx:9200/_template/name
```

### 删除模板
```
DELETE
http://xx.xx.xx.xx:9200/_template/name
```

## 索引

### 创建索引
```
PUT 
http://xx.xx.xx.xx:9200/index
{
    "mappings": {
        "properties": {
            "xxx_date": {
                "type": "date",
                "format": "yyyy-MM-dd"
            },
            "xxx_desc": {
                "type": "text"
            }
        }
    }
}
```

### 获取索引
```
GET
http://xx.xx.xx.xx:9200/index
```

### 删除索引
```
DELETE
http://xx.xx.xx.xx:9200/index
```

## 查询

### 查询索引数据

使用 Postman 示例(term)：
```
GET
http://xx.xx.xx.xx:9200/index/_search
{
  "query": {
    "term": {
      "name": {
        "value": "xxx"
      }
    }
  }
}
```
使用 term 查询： term 查询用于精确匹配某个字段的确切值，而不会对搜索词进行分析。

使用 Postman 示例(match)：
```
GET
http://xx.xx.xx.xx:9200/index/_search
{
    "query": {
        "match": {
            "name": "xxxxx"
        }
    },
    "size": 1
}
```
使用 match 查询： match 查询是一种灵活的文本查询，可以根据字段内容中的词项进行匹配。

使用 Postman 示例(match_all)：
```
GET
http://xx.xx.xx.xx:9200/index/_search
{
  "query": {
    "match_all": {}
  },
  "size": 1
}
```

使用 cURL 命令：
```
curl -XGET "http://xx.xx.xx.xx:9200/index/_search" -H 'Content-Type: application/json' -d'
{
  "query": {
    "match_all": {}
  }
}
'
```

## 删除

### 删除单条指定文档数据

使用 Postman 示例：
```
DELETE
http://xx.xx.xx.xx:9200/index/_doc/6fmI9YgBgpI27JNqrBV9
```

### 删除查询匹配的文档数据
使用 Postman 示例：
```
POST
http://xx.xx.xx.xx:9200/index/_delete_by_query
{
  "query": {
    "match": {
      "name": ""
    }
  }
}
```

### 删除全部文档但不删除索引结构

使用 Postman 示例：
```
POST
http://xx.xx.xx.xx:9200/index/_delete_by_query
{
  "query": {
    "match_all": {}
  }
}
```

使用 cURL 命令示例：
```
curl -XPOST "http://xx.xx.xx.xx:9200/index/_delete_by_query" -H 'Content-Type: application/json' -d'
{
  "query": {
    "match_all": {}
  }
}
'
```

