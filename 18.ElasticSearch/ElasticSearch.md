# API 调用

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
DELETE
http://xx.xx.xx.xx:9200/index
```

### 删除索引
```
{{kd_pre}}/kd-base-attack-cases-knowledge-tags
```

## 查询

### 查询索引数据

使用 Postman 示例(term)：
```
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

### 删除指定文档数据

使用 Postman 示例：
```
DELETE
http://xx.xx.xx.xx:9200/index/_doc/6fmI9YgBgpI27JNqrBV9
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

