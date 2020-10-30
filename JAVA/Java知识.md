# 基础知识

判断对象是否为空：

```java
Objects.isNull(skuStock) 
```

判断库存：

```java
//判断购物车中商品是否都有库存
Integer realSkuStock = skuStock.getStock() - skuStock.getLockStock();
if (realSkuStock == null || realSkuStock <= 0 || realSkuStock < quantity) {
    Asserts.fail("库存不足，无法下单");
}
```

#  序列化和反序列化

* **把对象转换为字节序列的过程称为对象的序列化**。
* **把字节序列恢复为对象的过程称为对象的反序列化**。

`fastjson`组件是阿里巴巴开发的反序列化与序列化组件，具体细节可以参考[github文档](https://github.com/alibaba/fastjson/wiki/Quick-Start-CN)。

```java
JSONObject obj= JSON.parseObject(JSON.toJSONString(contentResult)); // 处理时间
```

```java
//序列化
String text = JSON.toJSONString(obj); 
//反序列化
VO vo = JSON.parse(); //解析为JSONObject类型或者JSONArray类型
VO vo = JSON.parseObject("{...}"); //JSON文本解析成JSONObject类型
VO vo = JSON.parseObject("{...}", VO.class); //JSON文本解析成VO.class类
```

# 注解相关

### @ResponseBody

`@responseBody`注解的作用是将`controller`的方法返回的对象通过适当的转换器转换为指定的格式之后，写入到`response`对象的`body`区，通常用来返回`JSON`数据或者是`XML`数据。



### @Autowired

> @Autowired(required = false) (Spring提供)

作用： `@Autowired`用来给类中成员变量赋值。

细节：该注解用在成员变量或成员变量的`GET/SET`方法上。注入原则是默认根据类型自动注入。该注释的修饰范围只能作用在类上。

示例：

```java
public class SecurityConfig extends WebSecurityConfigurerAdapter {
    @Autowired(required = false)
    // 根据 dynamicSecurityService 名称自动注入成员变量，如果找不到会报错
    private DynamicSecurityService dynamicSecurityService;
}
```

### @Resource

> @Resource() (JavaEE提供)

作用： 和`@Autowired`一样，`@Resource`也是用来给类中成员变量赋值。

细节：该注解用在成员变量或成员变量的`GET/SET`方法上。注入原则是默认根据名称自动注入名称，找不到根据类型自动注入。该注释的修饰范围只能作用在类上。

示例：

```java
public class SmsInterceptor extends HandlerInterceptorAdapter {
    @Resource
    // 根据 redisService 名称自动注入成员变量，如果找不到该名称则按照类型 RedisService 注入
    private RedisService redisService;
}
```



# maven

```java
mvn clean package -Dmaven.test.skip=true -P test
```

