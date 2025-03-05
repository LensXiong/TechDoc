# API 调用
## 编译 pb.go 文件
```
protoc --plugin=protoc-gen-guardian=/Users/xxxx/go/bin/protoc-gen-guardian --guardian_out=plugins=grpc:. xx.xxx.proto
```
命令解释：
`protoc`: 这是 `Protocol Buffers` 编译器的命令。`Protocol Buffers` 是一种用于序列化结构化数据的语言和平台无关的技术。

`--plugin=protoc-gen-guardian=/Users/xxxx/go/bin/protoc-gen-guardian`: 这是指定编译器插件的选项。
`--plugin` 后面是插件的名称和路径。在这里，`protoc-gen-guardian` 是一个编译器插件，
它的执行文件位于当前工作目录（`/Users/xxxx/go/bin/protoc-gen-guardian`）。

`--guardian_out=plugins=grpc:.`: 这是指定编译器插件的参数。`--guardian_out` 后面的内容告诉编译器如何使用 `protoc-gen-guardian` 插件。
在这里，我们使用 grpc 插件，并且输出目录设置为当前工作目录（.）。

`xx.xxx.proto`: 这是要编译的 `Protocol Buffers` 文件的路径和名称。在这个例子中，我们要编译 `xx.xxx.proto` 这个文件。

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
            "xxx_data_tags": {
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
详细讲解：

这是一个示例的`Elasticsearch（ES）`索引模板配置：

`index_patterns`: 这个配置指定了哪些索引名称符合该模板的规则。 索引名称必须以 "-xxx-tags" 结尾，例如 "my-index-xxx-tags" 或 "another-index-xxx-tags" 等。

`settings`: 这是索引级别的设置部分，用于配置索引的一些基本参数。

`number_of_shards`: 定义索引被分成多少个分片。在这个示例中，索引将会有3个主分片。

`number_of_replicas`: 定义每个分片的副本数量。在这个示例中，每个主分片将会有1个副本，总共会有3个主分片和3个副本。

`refresh_interval`: 定义索引的刷新间隔，即将内存中的数据刷新到磁盘的频率。在这里设置为"1s"，表示每秒刷新一次，以便实时性更高的搜索结果。

`translog.flush_threshold_size`: 定义事务日志的大小阈值，当事务日志大小达到1GB时，将触发强制刷新操作，将数据持久化到磁盘上的索引文件中。

`max_result_window`: 定义一个搜索请求可以从该索引中获取的最大文档数量。在这个示例中，设置为2147483647，允许获取非常大的搜索结果集。

`max_inner_result_window`: 定义嵌套聚合结果窗口的大小。在这个示例中，设置为100000000，适用于处理大量嵌套聚合结果。

`mappings`: 这是索引的映射（mapping）配置部分，用于定义索引中字段的类型和属性。

`xxx_tags` 和 `xxx_data_tags`: 这两个字段分别被定义为 "object" 类型，这表示它们是嵌套对象。 每个对象有两个子字段：leaf 和 all。

`leaf`: 这些子字段被定义为 "keyword" 类型，这表示它们是精确值类型，不会被分词，用于精确匹配搜索。

`all`: 同样，这些子字段也被定义为 "keyword" 类型，它们用于存储一组关键字，同样不会被分词。

总结：
以上的配置示例是为了创建一个名为 "xxx-tags" 的索引模板。该模板用于匹配所有以 "-xxx-tags" 结尾的索引名称。
每个匹配的索引将有3个主分片和1个副本，索引将会每秒刷新一次，并设置了事务日志大小的阈值。
索引中包含两个嵌套对象字段 `xxx_tags` 和 `xxx_data_tags`，每个对象包含 leaf 和 all 两个关键字字段。
根据实际需求，可以根据这个模板创建多个匹配的索引，并且它们都会应用这个模板定义的配置和映射结构。

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

### 重建索引
要更改 Elasticsearch (ES) 索引结构并增加 `ctime` 和 `utime` 字段，你需要执行以下步骤：

1. 创建新的索引映射（mapping）：
首先，你需要创建一个新的索引，包含你想要的新映射。这是因为 Elasticsearch 不允许在现有索引上直接修改映射结构。你可以通过 `PUT` 请求来创建新索引。

```
PUT /your_new_index
{
  "mappings": {
    "properties": {
      "ctime": {
        "type": "date",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      "utime": {
        "type": "date",
        "format": "yyyy-MM-dd HH:mm:ss"
      },
      // 其他字段...
    }
  }
}
```

2. 将数据从旧索引迁移到新索引：
创建好新的索引后，你需要将数据从旧索引迁移到新索引。你可以使用 `_reindex` API 来实现。

```
POST /_reindex
{
  "source": {
    "index": "your_old_index"
  },
  "dest": {
    "index": "your_new_index"
  }
}
```

3. 删除旧索引（可选）：
如果你不再需要旧的索引，可以选择删除它。

```
DELETE /your_old_index
```

4. 重命名索引（可选）：
如果你希望新索引拥有与旧索引相同的名称，可以删除旧索引并使用 alias（别名）进行重命名。

```
POST /_aliases
{
  "actions": [
    {
      "remove": {
        "index": "your_old_index",
        "alias": "your_index_alias"
      }
    },
    {
      "add": {
        "index": "your_new_index",
        "alias": "your_index_alias"
      }
    }
  ]
}
```

通过这些步骤，你可以成功地更改 Elasticsearch 索引结构并添加新的 `ctime` 和 `utime` 字段。

### 新增索引
更新映射（Mapping）

```
PUT /your_index/_mapping
{
    "properties": {
      "ctime": { "type": "date", "format": "yyyy-MM-dd HH:mm:ss"},
      "utime": { "type": "date", "format": "yyyy-MM-dd HH:mm:ss"}
    }
}
```

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

查询不到数据：
原代码：
```
	//组装DSL
	must := []interface{}{}
	//榜单类别
	if params.HotlistName != "" {
		must = append(must, map[string]interface{}{
			"term": map[string]interface{}{
				"hotlist_names.keyword": map[string]interface{}{
					"value": params.HotlistName,
				},
			},
		})
	}
```

解决：
```
hotlist_names.keyword 改成 hotlist_names
```
原因：hotlist_names本身是keyword类型。

✅ keyword类型的字段适合做精确匹配查询，也就是term或terms。

✅ 不需要加.keyword，直接查hotlist_names就可以！

keyword是Elasticsearch里最适合存储"标签"、"分类"、"榜单名"这种的字段类型。 特点：

* 它是非分析的（not analyzed），存进去什么，查询就得拿一模一样的值去查。
* 它的用途是做精确匹配。
* 典型用法就是用来做过滤，比如： 查“榜单名等于‘爆款视频榜’” ；查“分类等于‘美妆’”

text是用来存长文本（description、文章、评论）的，它的特点：

* 存进去的文本会被分词，比如爆款视频榜可能会被拆成爆款、视频、榜。
* 适合做全文搜索（match查询），比如： 搜“爆款”能搜到“爆款视频榜”

hotlist_names本身就是keyword类型，它天生支持term和terms，可以直接查。

.keyword一般是用在text+keyword双字段的情况，比如：
```
{
    "hotlist_names": {
        "type": "text",
        "fields": {
            "keyword": {
                "type": "keyword"
            }
        }
    }
}
```

这种时候：

* hotlist_names是text，可以做分词搜索。
* hotlist_names.keyword是keyword，可以做精确匹配。

结论总结:

| 字段类型       | 查询字段名              | 查询方法               |
|------------|----------------|------------------|
| keyword    | hotlist_names  | ✅ term or terms |
| text          | hotlist_names  | ❌ 不适合term，要用match |
| text+keyword | hotlist_names.keyword | ✅ term or terms |


* keyword	存储标签、分类、榜单名等，用来做过滤
* text	存储长文本，用于全文搜索
* text+keyword	既能全文搜索，又能精确过滤，两者兼得