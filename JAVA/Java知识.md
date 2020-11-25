

待解决问题：

```java
SecurityContextHolder.getContext().getAuthentication() == null
UserDetails userDetails = this.userDetailsService.loadUserByUsername(username)  
```





# 封装通用返回对象

`api`返回码和返回信息接口：

```java
public interface IErrorCode {
    long getCode();

    String getMessage();
}
```

`api`返回码和返回信息接口的实现：

```java
public enum ResultCode implements IErrorCode {
    // 操作成功
    SUCCESS(200, "success"),
    // 操作失败
    FAILED(500, "failed"),
    // 参数检验失败
    VALIDATE_FAILED(404, "parameter check failed"),
    // 暂未登录或token已经过期
    UNAUTHORIZED(401, "not logged in yet or token expired"),
    // 没有相关权限
    FORBIDDEN(403, "no relevant authority");
    private long code;
    private String message;

    private ResultCode(long code, String message) {
        this.code = code;
        this.message = message;
    }

    public long getCode() {
        return code;
    }

    public String getMessage() {
        return message;
    }
}
```



通用返回对象封装：

```java
public class CommonResult<T> {
    /**
     * 状态码
     */
    private long code;
    /**
     * 提示信息
     */
    private String message;
    /**
     * 数据封装
     */
    private T data;

    protected CommonResult() {
    }

    protected CommonResult(long code, String message, T data) {
        this.code = code;
        this.message = message;
        this.data = data;
    }

    /**
     * 成功返回结果
     *
     * @param data 获取的数据
     */
    public static <T> CommonResult<T> success(T data) {
        return new CommonResult<T>(ResultCode.SUCCESS.getCode(), ResultCode.SUCCESS.getMessage(), data);
    }

    /**
     * 成功返回结果
     *
     * @param data 获取的数据
     * @param  message 提示信息
     */
    public static <T> CommonResult<T> success(T data, String message) {
        return new CommonResult<T>(ResultCode.SUCCESS.getCode(), message, data);
    }

    /**
     * 失败返回结果
     * @param errorCode 错误码
     */
    public static <T> CommonResult<T> failed(IErrorCode errorCode) {
        return new CommonResult<T>(errorCode.getCode(), errorCode.getMessage(), null);
    }

    /**
     * 失败返回结果
     * @param errorCode 错误码
     * @param message 错误信息
     */
    public static <T> CommonResult<T> failed(IErrorCode errorCode,String message) {
        return new CommonResult<T>(errorCode.getCode(), message, null);
    }

    /**
     * 失败返回结果
     * @param message 提示信息
     */
    public static <T> CommonResult<T> failed(String message) {
        return new CommonResult<T>(ResultCode.FAILED.getCode(), message, null);
    }

    /**
     * 失败返回结果
     */
    public static <T> CommonResult<T> failed() {
        return failed(ResultCode.FAILED);
    }

    /**
     * 参数验证失败返回结果
     */
    public static <T> CommonResult<T> validateFailed() {
        return failed(ResultCode.VALIDATE_FAILED);
    }

    /**
     * 参数验证失败返回结果
     * @param message 提示信息
     */
    public static <T> CommonResult<T> validateFailed(String message) {
        return new CommonResult<T>(ResultCode.VALIDATE_FAILED.getCode(), message, null);
    }

    /**
     * 未登录返回结果
     */
    public static <T> CommonResult<T> unauthorized(T data) {
        return new CommonResult<T>(ResultCode.UNAUTHORIZED.getCode(), ResultCode.UNAUTHORIZED.getMessage(), data);
    }

    /**
     * 未授权返回结果
     */
    public static <T> CommonResult<T> forbidden(T data) {
        return new CommonResult<T>(ResultCode.FORBIDDEN.getCode(), ResultCode.FORBIDDEN.getMessage(), data);
    }

    public long getCode() {
        return code;
    }

    public void setCode(long code) {
        this.code = code;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public T getData() {
        return data;
    }

    public void setData(T data) {
        this.data = data;
    }
}
```



分页数据封装：

```java
public class CommonPage<T> {
    /**
     * 当前页码
     */
    private Integer pageNum;
    /**
     * 每页数量
     */
    private Integer pageSize;
    /**
     * 总页数
     */
    private Integer totalPage;
    /**
     * 总条数
     */
    private Long total;
    /**
     * 分页数据
     */
    private List<T> list;

    /**
     * 将PageHelper分页后的list转为分页信息
     */
    public static <T> CommonPage<T> restPage(List<T> list) {
        CommonPage<T> result = new CommonPage<T>();
        PageInfo<T> pageInfo = new PageInfo<>(list);
        result.setTotalPage(pageInfo.getPages());
        result.setPageNum(pageInfo.getPageNum());
        result.setPageSize(pageInfo.getPageSize());
        result.setTotal(pageInfo.getTotal());
        result.setList(pageInfo.getList());
        return result;
    }

    /**
     * 将SpringData分页后的list转为分页信息
     */
    public static <T> CommonPage<T> restPage(Page<T> pageInfo) {
        CommonPage<T> result = new CommonPage<T>();
        result.setTotalPage(pageInfo.getTotalPages());
        result.setPageNum(pageInfo.getNumber());
        result.setPageSize(pageInfo.getSize());
        result.setTotal(pageInfo.getTotalElements());
        result.setList(pageInfo.getContent());
        return result;
    }

    public Integer getPageNum() {
        return pageNum;
    }

    public void setPageNum(Integer pageNum) {
        this.pageNum = pageNum;
    }

    public Integer getPageSize() {
        return pageSize;
    }

    public void setPageSize(Integer pageSize) {
        this.pageSize = pageSize;
    }

    public Integer getTotalPage() {
        return totalPage;
    }

    public void setTotalPage(Integer totalPage) {
        this.totalPage = totalPage;
    }

    public List<T> getList() {
        return list;
    }

    public void setList(List<T> list) {
        this.list = list;
    }

    public Long getTotal() {
        return total;
    }

    public void setTotal(Long total) {
        this.total = total;
    }
}
```



使用示例：

```java
@ApiOperation("弹窗列表")
@RequestMapping(value = "/list", method = RequestMethod.GET)
@ResponseBody
public CommonResult<CommonPage<SmsHomePop>> list(SmsHomePopSearchParam params,
       @RequestParam(value = "pageSize", defaultValue = "15") Integer pageSize,
       @RequestParam(value = "pageNum", defaultValue = "1") Integer pageNum) 
{
       List<SmsHomePop> data = smsHomePopService.search(params, pageNum, pageSize);
       return CommonResult.success(CommonPage.restPage(data));
}
```





# 基础知识

## Java String

### getBytes()

	byte[] getBytes()

 使用平台的默认字符集将此 String 编码为 byte 序列，并将结果存储到一个新的 byte 数组中。

	byte[] getBytes(String charsetName)

使用指定的字符集将此 String 编码为 byte 序列，并将结果存储到一个新的 byte 数组中。

```java
String lock = "redis_lock" + 12345789;
lock.getBytes();

long expireAt = System.currentTimeMillis() + 8000 + 1;
String.valueOf(expireAt).getBytes();
```

 

### String.valueOf()

将基本数据类型转换成`String`类型：

* `String.valueOf(boolean b) `: 将 `boolean` 变量 b 转换成字符串 
* `String.valueOf(char c)` : 将 char 变量 c 转换成字符串 
* `String.valueOf(char[] data)` : 将 char 数组 data 转换成字符串 
* `String.valueOf(char[] data, int offset, int count)` : 将 `char` 数组 `data` 中 由 `data[offset]` 开始取 count 个元素 转换成字符串 
* `String.valueOf(double d)` : 将 double 变量 d 转换成字符串 
* `String.valueOf(float f)` : 将 float 变量 f 转换成字符串 
* `String.valueOf(int i) `: 将 int 变量 i 转换成字符串 
* `String.valueOf(long l) `: 将 long 变量 l 转换成字符串 
* `String.valueOf(Object obj)` : 将 obj 对象转换成 字符串, 等于 `obj.toString() `



字符串转基本数据类型：

* `byte` :`Byte.parseByte(String s)` : 将 s 转换成 `byte` 

* `double` : `Double.parseDouble(String s)` : 将 s 转换成 `double` 

* `float` : `Double.parseFloat(String s)` : 将 s 转换成 `float` 

* `int` : `Integer.parseInt(String s)` : 将 s 转换成 `int` 

* `long` : `Long.parseLong(String s)`：将 s 转换成 `long`



字符串格式化：

```java
public final static String FORMAT_KEY = "sms:templates:%s:%s:%s";
String captchaKey = String.format(FORMAT_KEY, template, mobile, "captcha");
```

判断对象是否为空：

```java
Objects.isNull(skuStock) 
```

判断`list`、`map`是否为空：

```java
list == null || list.isEmpty()
map == null || map.isEmpty()   
// list.size()和 list.isEmpty()在数据量小的时候没有什么区别,但是在数据量多的时候,isEmpty()的效率高。  
list == null || list.size() == 0  
```

使用`CollectionUtils`判断集合是否为空：

```java
import org.springframework.util.CollectionUtils;

if (!CollectionUtils.isEmpty(oneGoodsPlans)) {}
```

使用`Hutool`开源工具包简化代码和方法：

```java
import cn.hutool.core.collection.CollUtil;

if(CollUtil.isNotEmpty(roleList)){}
```



判断`map`是否为空：

```java
map != null && !map.isEmpty()
```



条件查询：

```java
private List<SmsHomeAdvertise> getHomeAdvertiseList() {
        SmsHomeAdvertiseExample example = new SmsHomeAdvertiseExample();
        SmsHomeAdvertiseExample.Criteria criteria = example.createCriteria();
        criteria.andTypeEqualTo(1).andStatusEqualTo(1);
        Date nTime = new Date();
        criteria.andStartTimeLessThanOrEqualTo(nTime);
        criteria.andEndTimeGreaterThanOrEqualTo(nTime);
			  if (!StringUtils.isEmpty(keyword)) {
            criteria.andTitleLike("%" + keyword + "%");
        }
        example.setOrderByClause("sort desc");
        return advertiseMapper.selectByExample(example);
    }
```



`Long.parseLong(String s)`方法：将字符串`s`解析成十进制参数表示的`long`。

```java
Long.parseLong(oneGoodsPlans.get(0).getGoodsId().toString());
```

判断库存：

```java
//判断购物车中商品是否都有库存
Integer realSkuStock = skuStock.getStock() - skuStock.getLockStock();
if (realSkuStock == null || realSkuStock <= 0 || realSkuStock < quantity) {
    Asserts.fail("库存不足，无法下单");
}
```



属性拷贝：

```java
BeanUtils.copyProperties(product, pmsProductExtra);
```



#  序列化和反序列化

>  序列化是将对象状态转换为可保持或传输的格式的过程。与序列化相对的是反序列化，它将流转换为对象。这两个过程结合起来，可以轻松地存储和传输数据。

* **把对象转换为字节序列的过程称为对象的序列化**。

* **把字节序列恢复为对象的过程称为对象的反序列化**。

  

`fastjson`组件是阿里巴巴开发的反序列化与序列化组件，具体细节可以参考[github文档](https://github.com/alibaba/fastjson/wiki/Quick-Start-CN)。

```java
JSONObject obj= JSON.parseObject(JSON.toJSONString(contentResult)); // 处理时间
```

```java
// 序列化
String text = JSON.toJSONString(obj); 
// 反序列化
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



### @PathVariable

> `@PathVariable("xxx")`是`Spring3.0`的一个新功能：接收请求路径中占位符的值。
>
> 通过 `@PathVariable` 可以将`URL`中占位符参数`{xxx}`绑定到处理器类的方法形参中`@PathVariable(“xxx“) `

```java
@GetMapping("/detail/{id}")
@ResponseBody
public CommonResult detail(@PathVariable("id") Integer id) {
        Map<String, Object> res = oneGoodsSettingService.goodsDetail(id);
        JSONObject obj= JSON.parseObject(JSON.toJSONString(res)); // 处理时间
        return CommonResult.success(obj);
}
```

### @RequestParam

```java
@Target({ElementType.PARAMETER})
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface RequestParam {
    @AliasFor("name")
    String value() default "";

    @AliasFor("value")
    String name() default "";

    boolean required() default true;

    String defaultValue() default "\n\t\t\n\t\t\n\ue000\ue001\ue002\n\t\t\t\t\n";
}
```



使用示例：

```java
public CommonResult list(OmsOrderQueryParam queryParam,
                         @RequestParam(value = "pageSize", defaultValue = "5") Integer pageSize,
                         @RequestParam(value = "pageNum", defaultValue = "1") Integer pageNum) {
```



`@RequestParam`用来处理 `Content-Type` 为 `application/x-www-form-urlencoded` 编码的内容，`Content-Type`默认为该属性，也可以接收`application/json`。

### @RequestBody



>  @RequestParam和@RequestBody的区别：



注解`@RequestBody`接收的参数是来自`RequestBody`中，即请求体。一般用于处理非 `Content-Type: application/x-www-form-urlencoded`编码格式的数据，比如：`application/json`、`application/xml`等类型的数据。



`@RequestParam`和`@RequestBody`的区别：

* `content-type`角度：`form-data`、`x-www-form-urlencoded`：不可以用`@RequestBody`；可以用`@RequestParam`。`application/json`：`json`字符串部分可以用`@RequestBody`；`url`中的?后面参数可以用`@RequestParam`。
* 



# maven

```java
mvn clean package -Dmaven.test.skip=true -P test
```



`POM`文件：

`packaging`项目的打包类型：`pom`、`jar`、`war`。所有的父级项目的`packaging`都为`pom`，`packaging`默认类型`jar`类型，如果不做配置，`maven`会将该项目打成`jar`包。作为父级项目，还有一个重要的属性，那就是`modules`，通过`modules`标签将项目的所有子项目引用进来，在`build`父级项目时，会根据子模块的相互依赖关系整理一个`build`顺序，然后依次`build`。



```java
  <groupId>com.wangxiong.mall</groupId>
  <artifactId>mall</artifactId>
  <version>1.0-SNAPSHOT</version>
  <packaging>pom</packaging>

  <modules>
  <module>mall-common</module>
  <module>mall-mbg</module>
  <module>mall-security</module>
  <module>mall-demo</module>
  <module>mall-admin</module>
  <module>mall-search</module>
  <module>mall-portal</module>
  </modules>
```





# DAO和DTO

`DAO`: `data access object`数据访问对象，主要用来封装对数据库的访问，夹在业务逻辑与数据库资源中间。

`DTO`：`Data Transfer Object` 数据传输对象，主要用于远程调用等大量传输对象的地方。

`BO`：`Business Object` 业务对象层，

`PO`：`Persistant Object`持久对象，

`POJO`：`Plain Old Java Objects` 简单的`Java`对象，实际就是普通`JavaBeans`,使用`POJO`名称是为了避免和`EJB`混淆起来, 而且简称比较直接。

如何获取远程`IP`：

```java
ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
HttpServletRequest request = attributes.getRequest();
loginLog.setIp(request.getRemoteAddr());
```

如何获取远程`HttpServletRequest`:

```java
HttpServletRequest request = ((ServletRequestAttributes) RequestContextHolder.getRequestAttributes()).getRequest();
```



# AOP 编程

摘要：在实际的开发过程中，我们需要将接口的请求参数、返回数据甚至接口的消耗时间都以日志的形式打印出来以便排查问题，有些比较重要的接口甚至还需要将这些信息写入到数据库。像类似这种场景的代码相对来讲比较相似，为了提高代码的复用率，完全可以以 `AOP` 的方式将类似的代码封装起来。

## 日志处理

相关注解：

* `@Aspect`：将当前类标识为一个切面类，`Spring`会将该类作为一个切面管理。
* `@Component`：将该类作为一个 `Spring` 组件。
* `@Order(1)`：主要用来控制配置类的加载顺序，`bean`加载的优先级，值越小，越先被加载。
* `@Pointcut`：定义切点表达式，`Pointcut`是植入`Advice`的触发条件。每个`Pointcut`的定义包括2部分，一是表达式，二是方法签名。方法签名必须是 `public`及`void`型。可以将`Pointcut`中的方法看作是一个被`Advice`引用的助记符，因为表达式不直观，因此我们可以通过方法签名的方式为此表达式命名。因此`Pointcut`中的方法只需要方法签名，而不需要在方法体内编写实际代码。
* `@Before`：通知方法会在目标方法调用之前执行。

- `@After`：通知方法会在目标方法返回或抛出异常后执行。
- `@AfterReturning`：通知方法会在目标方法返回后执行。
- `@AfterThrowing`：通知方法会在目标方法抛出异常后执行。
- `@Around`：环绕增强，在切入点前后切入内容，并自己控制何时执行切入点自身的内容。



切点表达式：指定了通知被应用的范围，格式如下：

```java
execution(方法修饰符 返回类型 方法所属的包.类名.方法名称(方法参数)
```

```java
@Pointcut("@annotation(com.xxx.annotation.Page)")
@Pointcut("execution(public * com.xxx.controller.*.*(..)) ||
           execution(public * com.xxx.*.controller.*.*(..))")
```

日志切面使用步骤：

* 步骤一：添加 `AOP` 相关依赖。

* 步骤二：`Controller` 层的日志封装类`WebLog`。

* 步骤三：统一日志处理切面类`WebLogAspect`。

添加 `AOP` 相关依赖：

```java
<dependency>
  <groupId>org.springframework.boot</groupId>
  <artifactId>spring-boot-starter-aop</artifactId>
</dependency>
```

控制层的日志封装类：

```java
package com.xxx.common.domain;

import lombok.Data;
import lombok.EqualsAndHashCode;

/**
 * Controller 层的日志封装类
 */
@Data
@EqualsAndHashCode(callSuper = false)
public class WebLog {
    /**
     * 操作描述
     */
    private String description;

    /**
     * 操作用户
     */
    private String username;

    /**
     * 操作时间
     */
    private Long startTime;

    /**
     * 消耗时间
     */
    private Integer spendTime;

    /**
     * 根路径
     */
    private String basePath;

    /**
     * URI
     */
    private String uri;

    /**
     * URL
     */
    private String url;

    /**
     * 请求类型
     */
    private String method;

    /**
     * IP地址
     */
    private String ip;

    /**
     * 请求参数
     */
    private Object parameter;

    /**
     * 返回结果
     */
    private Object result;

}

```

统一日志处理切面：

```java

/**
 * 统一日志处理切面
 */
@Aspect
@Component
@Order(1)
public class WebLogAspect {
    private static final Logger LOGGER = LoggerFactory.getLogger(WebLogAspect.class);

    @Pointcut("execution(public * com.wangxiong.mall.controller.*.*(..))||execution(public * com.wangxiong.mall.*.controller.*.*(..))")
    public void webLog() {
    }

    @Before("webLog()")
    public void doBefore(JoinPoint joinPoint) throws Throwable {
        String traceLogId = String.valueOf(UUID.randomUUID());
        MDC.put("TRACE_LOG_ID", traceLogId);
    }

    @AfterReturning(value = "webLog()", returning = "ret")
    public void doAfterReturning(Object ret) throws Throwable {
        // 处理完请求  清除
        MDC.clear();
    }

    @Around("webLog()")
    public Object doAround(ProceedingJoinPoint joinPoint) throws Throwable {
        long startTime = System.currentTimeMillis();
        // 获取当前请求对象
        ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        HttpServletRequest request = attributes.getRequest();
        // 记录请求信息
        WebLog webLog = new WebLog();
        // 让目标方法执行
        Object result = joinPoint.proceed();
        // 获取封装署名信息的对象,在该对象中可以获取到目标方法名,所属类的 Class 等信息
        Signature signature = joinPoint.getSignature();
        MethodSignature methodSignature = (MethodSignature) signature;
        Method method = methodSignature.getMethod();

        // 设置注解描述, method 注释是否在 ApiOperation 上,如果在则返回 true ;不在则返回 false
        if (method.isAnnotationPresent(ApiOperation.class)) {
            ApiOperation log = method.getAnnotation(ApiOperation.class);
            webLog.setDescription(log.value());
        }
        long endTime = System.currentTimeMillis();
        String urlStr = request.getRequestURL().toString();
        webLog.setBasePath(StrUtil.removeSuffix(urlStr, URLUtil.url(urlStr).getPath()));
        // 获取当前缓存的用户
        webLog.setIp(request.getRemoteUser());
        webLog.setMethod(request.getMethod());
        // 获取传入目标方法的参数对象
        webLog.setParameter(getParameter(method, joinPoint.getArgs()));
        webLog.setResult(result);
        webLog.setSpendTime((int) (endTime - startTime));
        webLog.setStartTime(startTime);
        // URI：统一资源标识符 (Uniform Resource Identifier)
        webLog.setUri(request.getRequestURI());
        // URL: 统一资源定位符 (Uniform Resource Locator)
        webLog.setUrl(request.getRequestURL().toString());
        Map<String, Object> logMap = new HashMap<>();
        logMap.put("url", webLog.getUrl());
        logMap.put("method", webLog.getMethod());
        logMap.put("parameter", webLog.getParameter());
        logMap.put("spendTime", webLog.getSpendTime());
        logMap.put("description", webLog.getDescription());
        LOGGER.info(Markers.appendEntries(logMap), JSONUtil.parse(webLog).toString());
        return result;
    }


    /**
     * 根据方法和传入的参数获取请求参数
     */
    private Object getParameter(Method method, Object[] args) {
        List<Object> argList = new ArrayList<>();
        Parameter[] parameters = method.getParameters();
        for (int i = 0; i < parameters.length; i++) {
            // 将RequestBody注解修饰的参数作为请求参数
            RequestBody requestBody = parameters[i].getAnnotation(RequestBody.class);
            if (requestBody != null) {
                argList.add(args[i]);
            }
            // 将RequestParam注解修饰的参数作为请求参数
            RequestParam requestParam = parameters[i].getAnnotation(RequestParam.class);
            if (requestParam != null) {
                Map<String, Object> map = new HashMap<>();
                String key = parameters[i].getName();
                if (!StringUtils.isEmpty(requestParam.value())) {
                    key = requestParam.value();
                }
                map.put(key, args[i]);
                argList.add(map);
            }
        }
        if (argList.size() == 0) {
            return null;
        } else if (argList.size() == 1) {
            return argList.get(0);
        } else {
            return argList;
        }
    }
}
```

## 分页处理

摘要：本篇文章主要讲解的何使用`Spring boot AOP` +自定义注解+`PageHelper`来实现无侵入式的分页。传统的分页方式需要我们手动在每个接口中使用重复性的分页代码，这显然是不够明智的选择。相关`AOP`的内容可参考之前的文章，本文仅介绍使用自定义注解与`Spring boot AOP`完成分页处理的具体步骤。



## 常用分页问题

使用`PageHelper`分页工具的一般步骤如下：

* 编写一个查询`sql`，一般定义到`mapper`中；
* 编写一个分页查询方法，设置PageHelper的当前页和页大小；
* 执行查询语句；
* 查询完成后把`PageInfo`的数据填充到自定义的`PageBean`中。

示例代码如下：

```java
public PageBean<ReportTemplate> selectPage(PageBean<User> page) {
        // 通过PageHelper设置当前页和页大小
        PageHelper.startPage(page.getPageNo(), page.getPageSize());
        PageHelper.orderBy(page.getSortedField());
        List<User> users= userMapper.selectList(page.getKeyWords());
        PageInfo<User> pageInfo = new PageInfo<>(users);
        page.setCount(pageInfo.getTotal());
        page.setList(pageInfo.getList());
        return page;
}
```

分析以上步骤我们可以发现，除了第一步我们需要手动编写`SQL`查询语句，其他都是重复步骤。如果能将其封装成一个注解来使用，就会非常简单和优雅，因此需要结合`Spring boot AOP`来实现。

实现步骤总结如下：

* 引入相关依赖`pagehelper-spring-boot-starter`和`spring-boot-starter-aop`。
* 定义`@Page`注解；
* 定义`PageVo`分页参数；
* 定义`PageAspect`分页切面。

引入相关依赖`pagehelper-spring-boot-starter`和`spring-boot-starter-aop`：

```java
<!--MyBatis分页插件starter-->
<dependency>
  <groupId>com.github.pagehelper</groupId>
  <artifactId>pagehelper-spring-boot-starter</artifactId>
  <version>${pagehelper-starter.version}</version>
</dependency>
<dependency>
  <groupId>org.springframework.boot</groupId>
  <artifactId>spring-boot-starter-aop</artifactId>
</dependency>  
```

定义`@Page`注解：

```java
@Target({ElementType.TYPE, ElementType.METHOD})
@Retention(RetentionPolicy.RUNTIME)
@Documented
public @interface Page {
    /**
     * 用于绑定的请求参数名字
     */
    String value() default "";
}
```

定义`PageVo`分页参数：

```java
/**
 * 分页参数
 **/

public class PageVo implements Serializable{
    private static final long serialVersionUID = -1305720016123712695L;
        // 当前页
        private String pageNum;
        // 每页显示条数
        private String pageSize;
        // 查询参数
        private Map<String, Object> parameters = new HashMap<>(10);
        // 排序参数
        private Map<String, Object> sort = new HashMap<>(10);
  
        public String getPageNum() {
            return pageNum;
        }
        public void setPageNum(String pageNum) {
            this.pageNum = pageNum;
        }
        public String getPageSize() {
            return pageSize;
        }
        public void setPageSize(String pageSize) {
            this.pageSize = pageSize;
        }
        public Map<String, Object> getParameters() {
            return parameters;
        }
        public void setParameters(Map<String, Object> parameters) {
            this.parameters = parameters;
        }
        public Map<String, Object> getSort() {
            return sort;
        }
        public void setSort(Map<String, Object> sort) {
            this.sort = sort;
        }
}
```

定义`PageAspect`分页切面：

```java
/**
 * 分页处理切面
 **/
@Aspect
@Component
@Order(1)
public class PageAspect {

    private Logger logger = LoggerFactory.getLogger(this.getClass());

    @Pointcut("@annotation(com.xxx.annotation.Page)")
    public void page() {
    }

    @Around("page()")
    public Object processPage(ProceedingJoinPoint jp) throws java.lang.Throwable {
        // 获取目标方法原始的调用参数
        Object[] args = jp.getArgs();
        PageVo pageVo = new PageVo();
        if (args != null && args.length > 0 && args[0] instanceof PageVo) {
            // 修改目标方法的第一个参数
            pageVo = (PageVo) args[0];
            logger.info("当前页为：{},每页{}条数据", pageVo.getPageNum(), pageVo.getPageSize());
            logger.info("查询条件为：{}", pageVo.getParameters());
        }
        logger.info("执行查询===");
        // 以改变后的参数去执行目标方法，并保存目标方法执行后的返回值
        Object result = null;
        try {
            PageHelper.clearPage();
            if (pageVo != null && pageVo.getPageNum() != null) {
                PageHelper.startPage(Integer.parseInt(pageVo.getPageNum()), Integer.parseInt(pageVo.getPageSize())).setReasonable(false);
            }
            result = jp.proceed(args);
            logger.info("查询结束===");
            // 如果result的类型是list,并且参数类型为pageVo，将result初始化到分页中
            if (result != null && result instanceof List && args[0] instanceof PageVo) {
                ArrayList resultList = (ArrayList) result;
                logger.info("返回查询结果size={}", resultList.size());
                PageInfo<Object> pageInfo = new PageInfo<Object>(resultList);
                logger.info("pageInfo={},pageDataSize={}", pageInfo.getList(), pageInfo.getPageSize());
                // 将pageInfo中多余的参数去除掉
                ReturnInfo info = new ReturnInfo();
                info.setStatus(ReturnState.SUCCESS);
                info.setMessage("");
                Map<String, Object> page = Maps.newHashMap();
                page.put("totalCount", pageInfo.getTotal());
                page.put("pageSize", pageInfo.getPageSize());
                page.put("currentPage", pageInfo.getPageNum());
                page.put("totalPage", pageInfo.getPages());
                info.setReturnData(pageInfo.getList());
                info.setPageInfo(page);
                return info;
            }
            return result;
        } finally {
            logger.info("清除PageInfo的分页查询");
            PageHelper.clearPage();
        }
    }
}
```

使用示例：

```java
public Object listTransferMarket(PageVo pageVo) {
        logger.info("listTransferMarket para:{}", JSON.toJSONString(pageVo));
        return this.newProductTransferRecordManager.getTransferMarketListManager(pageVo);
    }

@Page
public Object getTransferMarketListManager(PageVo pageVo) {
  List<Map<String, Object>> list = this.newProductTransferRecordMapper.getTransferMarketList(pageVo);
  return list;
}

List<Map<String,Object>> getTransferMarketList(@Param("pageVo") PageVo pageVo);
```







# ELK

> `ELK`即`Elasticsearch`、`Logstash`、`Kibana`，组合起来可以搭建线上日志系统，可以使用`ELK`来收集`SpringBoot`应用产生的日志。

## ELK中各个服务的作用

* `Elasticsearch`：用于存储收集到的日志信息。
* `Logstash`：用于收集日志，`SpringBoot`应用整合了`Logstash`以后会把日志发送给`Logstash`，`Logstash`再把日志转发给`Elasticsearch`。
* `Kibana`：通过Web端的可视化界面来查看日志。



# LogBack

摘要：本篇文章主要是对`LogBack`的基础介绍，首先介绍了如何在`SpringBoot`应用集成`Logstash`，接着对一些常用的节点进行详细的描述，最后对`Logback`的五个日志级别和`log4j`8个级别的`log`进行介绍。其中最关键的是对`Logback`的相关节点的理解和使用。

## 应用集成

如果需要在`SpringBoot`应用集成`Logstash`，则需要在`pom.xml`中添加`logstash-logback-encoder`依赖：

```java
<properties>
  <logstash-logback.version>5.3</logstash-logback.version>
</properties>   
<!--集成logstash-->
<dependency>
    <groupId>net.logstash.logback</groupId>
    <artifactId>logstash-logback-encoder</artifactId>
    <version>${logstash-logback.version}</version>
</dependency>
```

添加配置文件`logback-spring.xml`让`logback`的日志输出到`logstash`：

```java
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE configuration>
<configuration>
    <!--输出到logstash的appender-->
    <appender name="LOGSTASH" class="net.logstash.logback.appender.LogstashTcpSocketAppender">
        <!--可以访问的logstash日志收集端口-->
        <destination>192.168.xx.xxx:4560</destination>
        <encoder charset="UTF-8" class="net.logstash.logback.encoder.LogstashEncoder"/>
    </appender>
    <root level="INFO">
        <appender-ref ref="CONSOLE"/>
        <appender-ref ref="FILE"/>
        <appender-ref ref="LOGSTASH"/>
    </root>
</configuration>
```

## 常用节点

常用节点：

* 根节点`<configuration>`。
* 子节点`<property>`，用来设置相关变量，通过`key-value`的方式配置，然后在后面的配置文件中通过 `${key}`来访问。
* 子节点`<appender>`，日志输出组件，主要负责日志的输出以及格式化日志。常用的属性有`name`和`class`。
* 子节点`<logger>`，`root`节点和`logger`节点其实都是表示`Logger`组件。
* 子节点`<root>`，`root`是最顶层的`logger`。
* 子节点`<contextName>`，设置日志上下文名称，后面输出格式中可以通过定义 `%contextName` 来打印日志上下文名称。
* 子节点`<timestamp>`，获取时间戳字符串。

### configuration

使用示例：

```java
<configuration scan="true" scanPeriod="60 seconds" debug="false"> 
　　  <!--其他配置省略--> 
</configuration>　
```

* `scan`：当此属性设置为`true`时，配置文件如果发生改变，将会被重新加载，默认值为`true`。
* `scanPeriod`: 设置监测配置文件是否有修改的时间间隔，如果没有给出时间单位，默认单位是毫秒。当`scan`为`true`时，此属性生效。默认的时间间隔为1分钟。
* `debug`: 当此属性设置为`true`时，将打印出`logback`内部日志信息，实时查看`logback`运行状态。默认值为`false`。

### contextName

`contextName`用来设置上下文名称，每个`logger`都关联到`logger`上下文，默认上下文名称为`default`。但可以使用`<contextName>`设置成其他名字，用于区分不同应用程序的记录，一旦设置，不能修改。

示例：

```java
<configuration scan="true" scanPeriod="60 seconds" debug="false"> 
     <contextName>appName</contextName> 
　　  <!--其他配置省略-->
</configuration>    
```

### property

`property`节点用来定义变量值，它有两个属性`name`和`value`，通过`<property>`定义的值会被插入到`logger`上下文中，可以使`${}`来使用变量。　

示例：

```java
<configuration scan="true" scanPeriod="60 seconds" debug="false"> 
     <property name="log.path" value="./logs/admin.log"/>
     <property name="AppName" value="myAppValue" /> 
　　　<contextName>${AppName}</contextName> 
　　　<!--其他配置省略--> 
</configuration>
```

### logger

子节点`<logger>`：用来设置某一个包或具体的某一个类的日志打印级别、以及指定`<appender>`。
`<logger>`仅有一个`name`属性，一个可选的`level`和一个可选的`addtivity`属性。

示例：

```java
<!-- project default level -->
<logger name="com.xxxx.xxx" level="INFO"/>
<logger name="org.apache.ibatis" level="INFO"/>
<logger name="java.sql" level="INFO"/>
<logger name="org.springframework" level="INFO"/>
<logger name="com.xxxx.xxx.mapper" level="ERROR"/>
<!--log4jdbc -->
<logger name="jdbc.sqltiming" level="INFO"/>
```

* `name`: 用来指定受此`loger`约束的某一个包或者具体的某一个类。
* `level` ：用来设置打印级别，与大小写无关：`TRACE`, `DEBUG`, `INFO`, `WARN`,` ERROR`,` ALL`和`OFF`，还有一个特殊值`INHERITED`或者同义词`NULL`，代表强制执行上级的级别。 如果未设置此属性，那么当前`loger`将会继承上级的级别。
* `addtivity`: 是否向上级`logger`传递打印信息。默认是`true`。可以包含零个或多个`<appender-ref>`元素，标识这个`appender`将会添加到这个`logger`。

### root

子节点`<root>`：它也是`<logger>`元素，但是它是根`loger`，是所有`<loger>`的上级。只有一个`level`属性，因为`name`已经被命名为`root`，且已经是最上级了。

同`<logger>`一样，可以包含零个或多个`<appender-ref>`元素，标识这个`appender`将会添加到这个`logger`。

示例：

```java
<root level="INFO">
    <appender-ref ref="console"/>
    <appender-ref ref="rollingFile"/>
</root>
```

### timestamp

两个属性:

* `key`： 标识此`<timestamp> `的名字；
* `datePattern`：设置将当前时间（解析配置文件的时间）转换为字符串的模式，遵循`java.txt.SimpleDateFormat`的格式。

示例：将解析配置文件的时间作为上下文名称

```java
<configuration scan="true" scanPeriod="60 seconds" debug="false">  
      <timestamp key="bySecond" datePattern="yyyyMMdd'T'HHmmss"/>   
      <contextName>${bySecond}</contextName>  
      <!-- 其他配置省略-->  
</configuration>
```

### appender

> `appender`意思是输出目的地，负责写日志的组件，它有两个必要属性`name`和`class`。`name`指定`appender`名称，`class`指定`appender`的全限定名。

示例：

```java
<appender name="console" class="ch.qos.logback.core.ConsoleAppender">
<appender name="rollingFile" class="ch.qos.logback.core.rolling.RollingFileAppender"> 
```

####  ConsoleAppender 

> `ConsoleAppender` 把日志输出到控制台，有以下子节点：

* `<encoder>`：对日志进行格式化。
* `<target>`：字符串`System.out(默认)`或者`System.err`。

示例：

```java
<appender name="console" class="ch.qos.logback.core.ConsoleAppender">
        <encoder>
  					<!-- 1格式化输出：%d表示日期，%thread表示线程名，%-5level：级别从左显示5个字符宽度%msg：日志消息，%n是换行符-->
            <pattern>%date{yyyy-MM-dd HH:mm:ss.SSS} [%thread] [%X{TRACE_LOG_ID}] %-5level %logger{36} --%mdc{client} - %msg%n
            </pattern>
        </encoder>
</appender>
```

`<pattern>`示例说明：

```java
xxxx-xx-xx 16:49:09.390 [http-nio-8081-exec-1] [c4f0f235-bb60-436c-b723-23c369981800] INFO  c.xxxx.xxx.common.log.WebLogAspect -- - {""method":"GET"".....}
```

#### FileAppender

> `FileAppender`把日志添加到文件，有以下节点：

* `<file>`：被写入的文件名，可以是相对目录，也可以是绝对目录，如果上级目录不存在会自动创建，没有默认值。
* `<append>`：如果是`true`，日志被追加到文件结尾，如果是 `false`，清空现存文件，默认是`true`。
* `<encoder>`：对记录事件进行格式化。
* `<prudent>`：如果是` true`，即使其他的`FileAppender`也在向此文件做写入操作，日志会被安全的写入文件，效率低，默认是` false`。

示例：将`DEBUG`级别及以上的日志都输出到`./logs/admin.log`。

```java
<configuration>
    <property name="log.path" value="./logs/admin.log"/>
    <appender name="rollingFile" class="ch.qos.logback.core.rolling.RollingFileAppender">
        <file>${log.path}</file>
        <append>true</append>
        <prudent>false</prudent>
        <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
            <fileNamePattern>
                ${log.path}.%d{yyyy-MM-dd}.log
            </fileNamePattern>
        </rollingPolicy>
        <encoder>
            <pattern>%date{yyyy-MM-dd HH:mm:ss.SSS} [%thread] [%X{TRACE_LOG_ID}] %-5level %logger{36} --%mdc{client} -
                %msg%n
            </pattern>
        </encoder>
    </appender>
    <root level="DEBUG">　
        <appender-ref ref="FILE"/>
    </root>
</configuration>
```

#### RollingFileAppender

>  `RollingFileAppender`滚动记录文件，先将日志记录到指定文件，当符合某个条件时，将日志记录到其他文件。有以下子节点：

* `<file>`：被写入的文件名，可以是相对目录，也可以是绝对目录，如果上级目录不存在会自动创建，没有默认值。
*  `<append>`：如果是 `true`，日志被追加到文件结尾，如果是 `false`，清空现存文件，默认是`true`。
* `<rollingPolicy>`：当发生滚动时，决定`RollingFileAppender`的行为，涉及文件移动和重命名。属性`class`定义具体的滚动策略。`TimeBasedRollingPolicy`为最常用的滚动策略，它根据时间来制定滚动策略，既负责滚动也负责触发滚动。`SizeBasedTriggeringPolicy`根据文件大小触发当前文件滚动。`FixedWindowRollingPolicy`根据固定窗口算法重命名文件的滚动策略。
* `<encoder>`：对记录事件进行格式化。

示例：

```java
<?xml version="1.0" encoding="UTF-8"?>
<configuration>
    <appender name="rollingFile" class="ch.qos.logback.core.rolling.RollingFileAppender">
        <file>/export/home/tomcat/domains/search.xxxx.com/server/logs/search-api.log
        </file>
        <rollingPolicy class="ch.qos.logback.core.rolling.TimeBasedRollingPolicy">
            <fileNamePattern>
                /export/home/tomcat/domains/search.xxxx.com/server/logs/search-api.log.%d{yyyy-MM-dd}.log
            </fileNamePattern>
        </rollingPolicy>
        <encoder>
            <pattern>%date{yyyy-MM-dd HH:mm:ss.SSS} [%thread] %-5level %logger{36} - %msg%n</pattern>
        </encoder>
    </appender>
    <root level="INFO">
        <appender-ref ref="console"/>
        <appender-ref ref="rollingFile"/>
    </root>
</configuration>
```

## 日志级别

####  logback

`Logback`分为五个日志级别，级别顺序由低到高分为：

> 优先级由高到低依次为：ERROR > WARN  > INFO  > DEBUG  >  TRACE 

- `TRACE` 级别最小，打印信息最为详细，一般不使用。
- `DEBUG`，需要在调试过程中输出的信息，但发布后是不需要的。
- `INFO`，需要持续输出的信息（无论调试还是发布状态）。
- `WARN`，警告级别的信息（不严重）。
- `ERROR`， 错误信息（较严重）。

####  log4j

`log4j`定义了8个级别的`log`（除去`OFF`和`ALL `分为6个级别）：

>  优先级从高到低依次为:OFF > FATAL > ERROR > WARN > INFO > DEBUG > TRACE > ALL

* `ALL `  ，最低等级，用于打开所有日志记录。
* `TRACE` ， 很低的日志级别， 一般不会使用。
* `DEBUG`  ，指出细粒度信息事件对调试应用程序是非常有帮助的，主要用于开发过程中打印一些运行信息。
* `INFO`   ，消息在粗粒度级别上突出强调应用程序的运行过程打印一些重要的信息，这个可以用于生产环境中输出程序运行的一些重要信息，但是不能滥用，避免打印过多的日志
* `WARN`   ，表明会出现潜在错误的情形，有些信息不是错误信息，但是也要给程序员的一些提示。
* `ERROR` ，指出虽然发生错误事件，但仍然不影响系统的继续运行，打印错误和异常信息，如果不想输出太多的日志，可以使用这个级别。
* `FATAL`  ，指出每个严重的错误事件将会导致应用程序的退出，这个级别比较高，重大错误，这种级别可以直接停止程序了。
* `OFF`  ， 最高等级的，用于关闭所有日志记录。

如果将`log` `level`设置在某一个级别上，那么比此级别优先级高的`log`都能打印出来。例如：

> 如果设置优先级为`WARN`, 那么`OFF`、 `FATAL`、 `ERROR`、 `WARN` 4个级别的`log`能正常输出而`INFO`、`DEBUG` 、`TRACE`、 `ALL`级别的`log`则会被忽略。



# 数据类型

## List

`List<String>`：

```java
List<String> list = new ArrayList<String>(); // [hello，javaee，world，java]
```

`List<Student> list`：

```java
List<Student> list = new ArrayList<Student>(); // [com.wangxiong.Student@e580929, com.wangxiong.Student@1cd072a9, com.wangxiong.Student@7c75222b]
```

`List<Object>`：

```java
[
    {
        "deliveryCompany": "",
        "deliverySn": "null",
        "orderId": 12121
    },
    {
        "deliveryCompany": "1122",
        "deliverySn": "",
        "orderId": 121211
    }
]
```

`List<Map<String,Object>>`：

```java
[
    {
        "id": 1,
        "params": {
            "method": "POST",
            "url": "xxx"
        },
        "name": "wangxiong"
    }
]
```




`Map`：

```java

```

`Object`：

```java

```

`Object[]`：

```java

```

` Map<String, Object> map = new HashMap<>()`：

```java

```

使用示例：

```java
 public CommonResult delivery(@RequestBody List<OmsOrderDeliveryParam> deliveryParamList) {
        for(OmsOrderDeliveryParam dpl:deliveryParamList) {
            String  deliverySn = dpl.getDeliverySn();
            String  deliveryCompany = dpl.getDeliveryCompany();
            if(deliverySn == null || deliverySn.isEmpty() ) {
               return  CommonResult.failed(dpl.getOrderId()+" deliverySn required!");
            }
            if(deliveryCompany == null || deliveryCompany.isEmpty() ) {
                return  CommonResult.failed(dpl.getOrderId()+" deliveryCompany required!");
            }
        }
 }
```



 filter = new ArrayList<>();`

```java

```



# String 和 StringBuilder 

`String` 类代表字符串，`Java`程序中的所有字符串文字（例如`"abc"`）都被实现为此类的实例。也就是说，`Java` 程序中所有的双引号字符串，都是 `String` 类的对象。在 `Java` 中字符串属于对象，`Java` 提供了 `String` 类来创建和操作字符串。`String` 类在 `java.lang` 包下，所以使用的时候不需要导包。

## String 类

### 创建字符串

`String` 类的特点：

* `String` 类是不可改变的，所以一旦创建了 `String` 对象，那它的值就无法改变，如果需要改变请选择使用[StringBuffer & StringBuilder 类](https://www.runoob.com/java/java-stringbuffer.html)。

* 虽然 `String` 的值是不可变的，但是它们可以被共享。

* 字符串效果上相当于字符数组( `char[]` )，但是底层原理是字节数组( `byte[]` )。

  

创建字符串对象的两种方式：

* 直接赋值创建，以""方式给出的字符串，只要字符序列相同（顺序和大小写），无论在程序代码中出现几次，`JVM` 都只会建立一 个 `String` 对象，并在字符串池中维护。
* 通过构造方法创建，通过 `new` 创建的字符串对象，每一次 `new` 都会申请一个内存空间，虽然内容相同，但是地址值不同。

直接赋值创建：

```java
host:
  mall:
   admin: http://localhost:8080
@Value("${host.mall.admin}")
private String HOST_MALL_ADMIN;
String url = HOST_MALL_ADMIN + "/brand/create";
```

构造方法创建：

```java
char[] nameArr = { 'w', 'a', 'n', 'g', 'x', 'i', 'o', 'n', 'g'};
String nameString = new String(nameArr);  


String lock = "LOCK_PREFIX" + "wangxiong";
byte[] value = lock.getBytes(); 
long expireTime = Long.parseLong(new String(value));
```

常用的构造方法：

| 方法名                    | 说明                                      |
| ------------------------- | ----------------------------------------- |
| public String()           | 创建一个空白字符串对象，不含有任何内容    |
| public String(char[] chs) | 根据字符数组的内容，来创建字符串对象      |
| public String(byte[] bys) | 根据字节数组的内容，来创建字符串对象      |
| String s = "abc";         | 直接赋值的方式创建字符串对象，内容就是abc |

示例代码：

```java
public class StringDemo01 {
		public static void main(String[] args) {
		  // public String():创建一个空白字符串对象，不含有任何内容
      String s1 = new String();
	    System.out.println("s1:" + s1); // s1:
      
      // public String(char[] chs):根据字符数组的内容，来创建字符串对象 
      char[] chs = {'a', 'b', 'c'};
      String s2 = new String(chs);
      System.out.println("s2:" + s2); // s2:abc
      
      //public String(byte[] bys):根据字节数组的内容，来创建字符串对象 
      byte[] bys = {97, 98, 99};
      String s3 = new String(bys);
      System.out.println("s3:" + s3); // s3:abc
      
      //String s = “abc”; 直接赋值的方式创建字符串对象，内容就是abc 
      String s4 = "abc";
      System.out.println("s4:" + s4); // s4:abc
		} 
}
```

### 字符串比较



* 使用`==`号，当比较基本数据类型时比较的是具体的值，当比较引用数据类型时，比较的是对象地址值。
* 使用`equals`方法，比较两个字符串内容是否相同、区分大小写。



示例代码：

```java
public class StringDemo02 {
    public static void main(String[] args) {
        // 构造方法的方式得到对象
        char[] chs = {'a', 'b', 'c'};
        String s1 = new String(chs);
        String s2 = new String(chs);

        // 直接赋值的方式得到对象
        String s3 = "abc";
        String s4 = "abc";

        // 比较字符串对象地址是否相同
        System.out.println(s1 == s2); // false
        System.out.println(s1 == s3); // false
        System.out.println(s3 == s4); // true
        System.out.println("--------");

        // 比较字符串内容是否相同
        System.out.println(s1.equals(s2)); // true
        System.out.println(s1.equals(s3)); // true
        System.out.println(s3.equals(s4)); // true
    }
}
```

> 字符串的内容比较， 用equals() 方法实现。

### 格式化字符串

`String` 类的静态方法 `format()` 能用来创建可复用的格式化字符串。

如下所示：

```java
String captchaKey = String.format("sms:templates:%s:%s:%s", template, mobile, "captcha");

String fs = String.format(
  								 "浮点型变量的值为 " +
                   "%f, 整型变量的值为 " +
                   " %d, 字符串变量的值为 " +
                   " %s", floatVar, intVar, stringVar);
```

### String 常用方法

| 方法名                                 | 说明                                                         |
| -------------------------------------- | ------------------------------------------------------------ |
| public boolean equals(Object anObject) | 比较字符串的内容，严格区分大小写(用户名和密码)               |
| public char charAt(int index)          | 返回指定索引处的 char 值                                     |
| public int length()                    | 返回此字符串的长度                                           |
| byte[\] getBytes()                     | 使用平台的默认字符集将此 String 编码为 byte 序列，并将结果存储到一个新的 byte 数组中。 |

示例代码：键盘录入一个字符串，统计该字符串中大写字母字符，小写字母字符，数字字符出现的次数(不考虑其他字符)。

```java
public class StringTest03 {
    public static void main(String[] args) {
        // 键盘录入一个字符串，用 Scanner 实现
        Scanner sc = new Scanner(System.in);

        System.out.println("请输入一个字符串：");
        String line = sc.nextLine();

        // 要统计三种类型的字符个数，需定义三个统计变量，初始值都为0
        int bigCount = 0;
        int smallCount = 0;
        int numberCount = 0;

        // 遍历字符串，得到每一个字符
        for(int i=0; i<line.length(); i++) {
            char ch = line.charAt(i);

            // 判断该字符属于哪种类型，然后对应类型的统计变量+1
            if(ch>='A' && ch<='Z') {
                bigCount++;
            } else if(ch>='a' && ch<='z') {
                smallCount++;
            } else if(ch>='0' && ch<='9') {
                numberCount++;
            }
        }

        // 输出三种类型的字符个数
        System.out.println("大写字母：" + bigCount + "个");
        System.out.println("小写字母：" + smallCount + "个");
        System.out.println("数字：" + numberCount + "个");

    }
}
```

示例：定义一个方法，实现字符串反转。

```java
public static String reverse(String s) {
        String ss = "";
        for(int i=s.length()-1; i>=0; i--) {
            ss += s.charAt(i);
        }
        return ss;
    }
```

## StringBuilder 类

当对字符串进行修改的时候，需要使用 `StringBuffer` 和 `StringBuilder` 类。和 `String` 类不同的是，`StringBuffer` 和 `StringBuilder` 类的对象能够被多次的修改，并且不产生新的未使用对象。



`StringBuilder` 类在 `Java 5` 中被提出，它和 `StringBuffer` 之间的最大不同在于 `StringBuilder` 的方法不是线程安全的（不能同步访问）。由于 `StringBuilder` 相较于 `StringBuffer` 有速度优势，所以多数情况下建议使用 `StringBuilder` 类。然而在应用程序要求线程安全的情况下，则必须使用 `StringBuffer` 类。



### 常用方法

| 方法名                           | 说明                                                |
| -------------------------------- | --------------------------------------------------- |
| public StringBuilder()           | 创建一个空白可变字符串对象，不含有任何内容          |
| public StringBuilder(String str) | 根据字符串的内容，来创建可变字符串对象              |
| public int length()              | 返回长度，实际存储值                                |
| public String toString()         | 通过toString()就可以实现StringBuilde转换为String |


示例代码：

```java
public class StringBuilderDemo01 {
    public static void main(String[] args) {
        // public StringBuilder()：创建一个空白可变字符串对象，不含有任何内容
        StringBuilder sb = new StringBuilder();
        System.out.println("sb:" + sb); // sb:
        System.out.println("sb.length():" + sb.length()); // sb.length():0

        // public StringBuilder(String str)：根据字符串的内容，来创建可变字符串对象
        StringBuilder sb2 = new StringBuilder("hello");
        System.out.println("sb2:" + sb2); // sb2:hello
        System.out.println("sb2.length():" + sb2.length()); // sb2.length():5
    }
}
```

### 添加和反转

| 方法名                                | 说明                     |
| ------------------------------------- | ------------------------ |
| public StringBuilder append(任意类型) | 添加数据，并返回对象本身 |
| public StringBuilder reverse()        | 返回相反的字符序列       |

示例代码：

```java
public class StringBuilderDemo01 {
    public static void main(String[] args) {
        // 创建对象
        StringBuilder sb = new StringBuilder();
        sb.append("hello").append("world").append("java").append(100); 
        System.out.println("sb:" + sb); // sb:helloworldjava100
        // 返回相反的字符序列
        sb.reverse();
        System.out.println("sb:" + sb); // sb:001avajdlrowolleh
    }
}
```

### 互相转换

* `StringBuilder`转换为`String`：通过 `toString() `就可以实现把 `StringBuilder` 转换为 `String` 。

* `String`转换为`StringBuilder`：通过构造方法就可以实现把 `String` 转换为 `StringBuilder`。

  

`StringBuilder` 转换为 `String` 示例：

```java
StringBuilder sb = new StringBuilder();
sb.append("hello");
// String s = sb; //这个是错误的做法
String s = sb.toString();
```

`String`转换为`StringBuilder` 示例：

```java
String s = "hello";
// StringBuilder sb = s; //这个是错误的做法
StringBuilder sb = new StringBuilder(s);
```

用`StringBuilder`实现字符串的反转，并把结果转成`String`返回：

```java
// String --- StringBuilder --- reverse() --- String
StringBuilder(String s).reverse().toString();
```



