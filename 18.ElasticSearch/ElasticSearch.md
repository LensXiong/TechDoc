# API 调用
## 备份

## 迁移

## 脚本

## FQA


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

