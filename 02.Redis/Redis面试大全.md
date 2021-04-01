# 问题列表

1、[Redis 缓存设计和使用原则有哪些？](#01)

2、[Redis 的数据类型你都用了哪几种。 在什么样的场景下用到什么类型？](#02)

3、[ Redis 的过期失效是怎么失效的，怎么释放内存的？](#03)

4、[ Redis  “雪崩效应” 产生原因和解决方案？](#04)

5、[ 当redis做缓存时，如何确定缓存更新策略（当DB数据更新时）？](#05)

6、[ 一致性哈希的含义和应用场景及缺点？](#06)

7、[  当redis做db时如何保证数据落地（redis突然挂了）？](#07)

8、[ Redis主从结构了解么？](#08)

9、[ Redis 集群配置你觉得难点在那？](#09)

10、[ Redis数据一致性的底层原理？](#10)

11、[ Redis 和 Memcached 的区别?](#11)

12、[ Redis 的持久化的方式有几种，各自的优缺点是什么，区别是什么？](#12)



# 解答列表

## Redis 缓存设计和使用原则

1、<span id="01">Redis 缓存设计和使用原则有哪些？</span>

* ① 只应将热数据放在缓存中。

* ② 所有缓存信息都应设置过期时间， 缓存过期时间应当分散以避免集中过期。

* ③ 缓存key应具备可读性，应避免不同业务出现同名缓存key。

* ④ 可对key进行适当的缩写以节省内存空间。

* ⑤ 一个key对应的数据不应过大，对于string类型，一个key对应的value大小应控制在10K以内，1K左右更优；
  对于hash类型，不应超过5000行。

* ⑥ 读的顺序是先缓存，后数据库；写的顺序是先数据库，后缓存。

* ⑦ 避免缓存穿透，数据库中未查询到的数据，可在Redis中设置特殊标识，以避免因缓存中无数据而导致每次请求均达到数据库。

* ⑧ 可以进行适当的缓存预热，对于上线后可能会有大量读请求的应用，在上线之前可预先将数据写入缓存中。

* ⑨ 读的顺序是先缓存，后数据库；写的顺序是先数据库，后缓存。

* ⑩ 数据一致性问题，数据源发生变更时可能导致缓存中数据与数据源中数据不一致，应根据实际业务需求来选择适当的缓存更新策略。主动更新：在数据源发生变更时同步更新缓存数据或将缓存数据过期。一致性高，维护成本较高。被动删除：根据缓存设置的过期时间有Redis负责数据的过期删除。一致性较低，维护成本较低。

  

## Redis 的数据类型和应用场景

2、<span id="02">Redis 的数据类型你都用了哪几种。 在什么样的场景下用到什么类型？</span>

Redis 常用的5种基本类型和3种特殊类型。

5种基本类型：String（字符串）、List（列表）、Hash（字典）、Set（集合）、Sorted Set（有序集合）。

3种特殊类型：Hyperloglog（基数）、地理位置（Geospatia）和位图 （Bitmaps）。



应用场景：

① 缓存应用。String（字符串）---短信验证码；Hash（字典）---商品信息，新闻信息。

② 排行榜应用---Sorted Set（有序集合）。很多网站都有排行榜应用的，如京东的月度销量榜单、商品按时间的上新排行榜、sina 微博热门话题等。Redis 提供的有序集合数据类构能实现各种复杂的排行榜应用。

③ 社交网络应用---Set（集合）。点赞、踩、关注/被关注、共同好友等是社交网站的基本功能，社交网站的访问量通常来说比较大，而且传统的关系数据库类型不适合存储这种类型的数据，Redis提供的哈希、集合等数据结构能很方便的的实现这些功能。

④ 最新列表和数据应用---List（列表）。Redis列表结构，LPUSH可以在列表头部插入一个内容ID作为关键字，LTRIM可用来限制列表的数量，这样列表永远为N个ID，无需查询最新的列表，直接根据ID去到对应的内容页即可。因为list是有序的，比较适合存储一些有序且数据相对固定的数据。如省市区表、字典表等。

⑤ 消息队列应用---消息队列是大型网站必用中间件，如ActiveMQ、RabbitMQ、Kafka等流行的消息队列中间件，主要用于业务解耦、流量削峰及异步处理实时性低的业务。Redis提供了发布/订阅及阻塞队列功能，能实现一个简单的消息队列系统。另外，这个不能和专业的消息中间件相比。



## Redis 和 Memcached 的区别

11、<span id="11">Redis 和 Memcached 的区别? （数据类型、持久化、单线程和多线程）</span>

* ① Redis支持更多的复杂的数据类型和操作。Redis 不仅仅支持简单的`key-value`类型的数据，同时还提供`list`，`set`，`zset`，`hash`等数据结构的存储。 Memcached 只支持`key-value`这样简单的结构。

* ② Redis支持数据的备份，即master-slave(主—从)模式的数据备份。 
* ③ Redis支持数据的持久化，可以将内存中的数据保持在磁盘中，重启的时候可以再次加载进行使用。 Memcache不支持数据持久存储 。
* ④ 内存使用率方面，对于key-value这样简单的数据储存，memcache的内存使用率更高。如果采用hash结构，Redis 的内存使用率会更高。
* ⑤ Redis 是单线程，多路复用方式提高处理效率。Memcached 是多线程的，通过CPU线程切换来提高处理效率。
* ⑥ Redis 单个value的最大限制是1GB， memcached只能保存1MB的数据。

关于什么时候选择使用Redis，什么时候选择使用Memcached，可以参考 Redis 作者的原文：

```php
You should not care too much about performances. Redis is faster per core with small values, but memcached is able to use multiple cores with a single executable and TCP port without help from the client. Also memcached is faster with big values in the order of 100k. Redis recently improved a lot about big values (unstable branch) but still memcached is faster in this use case. The point here is: nor one or the other will likely going to be your bottleneck for the query-per-second they can deliver.

You should care about memory usage. For simple key-value pairs memcached is more memory efficient. If you use Redis hashes, Redis is more memory efficient. Depends on the use case.

You should care about persistence and replication, two features only available in Redis. Even if your goal is to build a cache it helps that after an upgrade or a reboot your data are still there.

You should care about the kind of operations you need. In Redis there are a lot of complex operations, even just considering the caching use case, you often can do a lot more in a single operation, without requiring data to be processed client side (a lot of I/O is sometimes needed). This operations are often as fast as plain GET and SET. So if you don’t need just GEt/SET but more complex things Redis can help a lot (think at timeline caching).
```

##Redis 持久化的方式 RDB 和 AOF

12、<span id="12">Redis 的持久化的方式有几种，各自的优缺点是什么，区别是什么？</span>

[Redis 持久化](https://segmentfault.com/a/1190000002906345)

> Redis 有两种持久化的方式：快照（RDB文件）和追加式文件（AOF文件）。

- RDB持久化方式会在一个特定的间隔保存那个时间点的一个数据快照。
- AOF持久化方式则会记录每一个服务器收到的写操作。在服务启动时，这些记录的操作会逐条执行从而重建出原来的数据。写操作命令记录的格式跟Redis协议一致，以追加的方式进行保存。
- Redis的持久化是可以禁用的，就是说你可以让数据的生命周期只存在于服务器的运行时间里。
- 两种方式的持久化是可以同时存在的，但是当Redis重启时，AOF文件会被优先用于重建数据。

RDB 的工作原理：

> Redis调用fork()，产生一个子进程。子进程把数据写到一个临时的RDB文件。当子进程写完新的RDB文件后，把旧的RDB文件替换掉。

RDB 的优点：

> RDB文件是一个很简洁的单文件，它保存了某个时间点的Redis数据，很适合用于做备份。你可以设定一个时间点对RDB文件进行归档，这样就能在需要的时候很轻易的把数据恢复到不同的版本。
>
> 基于上面所描述的特性，RDB很适合用于灾备。单文件很方便就能传输到远程的服务器上。
>
> RDB的性能很好，需要进行持久化时，主进程会fork一个子进程出来，然后把持久化的工作交给子进程，自己不会有相关的I/O操作。
>
> 比起AOF，在数据量比较大的情况下，RDB的启动速度更快。

RDB 的缺点：

> RDB容易造成数据的丢失。假设每5分钟保存一次快照，如果Redis因为某些原因不能正常工作，那么从上次产生快照到Redis出现问题这段时间的数据就会丢失了。
>
> RDB使用`fork()`产生子进程进行数据的持久化，如果数据比较大的话可能就会花费点时间，造成Redis停止服务几毫秒。如果数据量很大且CPU性能不是很好的时候，停止服务的时间甚至会到1秒。

AOF 的工作原理：

> 快照并不是很可靠。如果你的电脑突然宕机了，或者电源断了，又或者不小心杀掉了进程，那么最新的数据就会丢失。而AOF文件则提供了一种更为可靠的持久化方式。每当Redis接受到会修改数据集的命令时，就会把命令追加到AOF文件里，当你重启Redis时，AOF里的命令会被重新执行一次，重建数据。

AOF 的优点：

> 比RDB可靠。你可以制定不同的fsync策略：不进行fsync、每秒fsync一次和每次查询进行fsync。默认是每秒fsync一次。这意味着你最多丢失一秒钟的数据。
>
> AOF日志文件是一个纯追加的文件。就算是遇到突然停电的情况，也不会出现日志的定位或者损坏问题。甚至如果因为某些原因（例如磁盘满了）命令只写了一半到日志文件里，我们也可以用`redis-check-aof`这个工具很简单的进行修复。
>
> 当AOF文件太大时，Redis会自动在后台进行重写。重写很安全，因为重写是在一个新的文件上进行，同时Redis会继续往旧的文件追加数据。新文件上会写入能重建当前数据集的最小操作命令的集合。当新文件重写完，Redis会把新旧文件进行切换，然后开始把数据写到新文件上。
>
> AOF把操作命令以简单易懂的格式一条接一条的保存在文件里，很容易导出来用于恢复数据。例如我们不小心用`FLUSHALL`命令把所有数据刷掉了，只要文件没有被重写，我们可以把服务停掉，把最后那条命令删掉，然后重启服务，这样就能把被刷掉的数据恢复回来。

AOF 的缺点：

> 在相同的数据集下，AOF文件的大小一般会比RDB文件大。
>
> 在某些fsync策略下，AOF的速度会比RDB慢。通常fsync设置为每秒一次就能获得比较高的性能，而在禁止fsync的情况下速度可以达到RDB的水平。
>
> 在过去曾经发现一些很罕见的BUG导致使用AOF重建的数据跟原数据不一致的问题。