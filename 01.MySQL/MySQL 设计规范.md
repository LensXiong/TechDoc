# 数据库设计规范

以下所有规范会按照【高危】、【强制】、【推荐】三个级别进行标注，遵守优先级从高到低。

对于【高危】、【强制】两个级别的设计，必须强制执行。

## 一个规范的建表语句示例

```
CREATE TABLE user (
  `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint(11) NOT NULL COMMENT '用户id',
  `username` varchar(45) NOT NULL COMMENT '真实姓名',
  `email` varchar(30) NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `nickname` varchar(45) NOT NULL COMMENT '昵称',
  `avatar` varchar(255) NOT NULL COMMENT '头像',
  `sex` tinyint(4) DEFAULT '0' COMMENT '性别',
  `short_introduce` varchar(150) DEFAULT NULL COMMENT '一句话介绍自己，最多50个汉字',
  `user_resume` varchar(300) NOT NULL COMMENT '用户提交的简历存放地址',
  `user_register_ip` int NOT NULL COMMENT '用户注册时的源ip',
  `create_time` int NOT NULL COMMENT '用户记录创建的时间',
  `update_time` int NOT NULL COMMENT '用户资料修改的时间',
  `user_review_status` tinyint NOT NULL COMMENT '用户资料审核状态，1为通过，2为审核中，3为未通过，4为还未提交审核',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_user_id` (`user_id`),
  KEY `idx_username`(`username`),
  KEY `idx_create_time_user_review_status`(`create_time`,`user_review_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='网站用户基本信息';
```



## 字符集



## 库设计



## 表设计

1.【推荐】适当的反范式设计：字段允许适当冗余，以提高查询性能，但必须考虑数据一致。冗余字段应遵循:

- 不是频繁修改的字段。
- 不是 `varchar` 超长字段，更不能是 `text` 字段。

说明：

① 把经常需要`join`查询的字段，在其他表里冗余一份。如`user_name`属性在`user_account`，`user_login_log`等表里冗余一份，减少`join`查询。

② 商品类目名称使用频率高，字段长度短，名称基本一成不变，可在相关联的表中冗余存储类目名称，避免关联查询。



2.【推荐】关于主键的设计：

- `id` 类型没有特殊要求，强制要求主键为`id`，类型为`int`或`bigint`，且为`auto_increment`。
- 标识表里每一行主体的字段不要设为主键，建议设为其他字段如`user_id`，`order_id`等，并建立`unique key`索引。

说明：

① 如果设为主键且主键值为随机插入，则会导致`innodb`内部`page`分裂和大量随机`I/O`，性能下降。

| 类型     | 范围小                       | 范围大                       | 存储字节 |
| -------- | ---------------------------- | ---------------------------- | -------- |
| bigint   | -2^63 (-9223372036854775808) | 2^63-1 (9223372036854775807) | 8        |
| int      | -2^31 (-2,147,483,648)       | 2^31 – 1 (2,147,483,647)     | 4        |
| smallint | -2^15 (-32,768)              | 2^15 – 1 (32,767)            | 2        |
| tinyint  | 0                            | 255                          | 1        |



3. 【推荐】表中所有字段必须都是`NOT NULL`属性，业务可以根据需要定义`DEFAULT`值。

① 因为使用`NULL`值会存在每一行都会占用额外存储空间、数据迁移容易出错、聚合函数计算结果偏差等问题。

② 如字符型的默认值为一个空字符值串’’；数值型默认值为数值 0；逻辑型的默认值为数值 0。



## 列设计



## 索引设计



## 分库分表分区设计









