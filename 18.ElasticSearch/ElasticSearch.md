# API 调用

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

