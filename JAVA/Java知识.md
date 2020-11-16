

待解决问题：

```java
SecurityContextHolder.getContext().getAuthentication() == null
UserDetails userDetails = this.userDetailsService.loadUserByUsername(username)  
```





# 基础知识

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

## 日志处理



使用场景：在实际的开发过程中，我们需要将接口的请求参数、返回数据甚至接口的消耗时间都以日志的形式打印出来以便排查问题，有些比较重要的接口甚至还需要将这些信息写入到数据库。像类似这种场景的代码相对来讲比较相似，为了提高代码的复用率，完全可以以 `AOP` 的方式将类似的代码封装起来。



* `@Aspect`：将当前类标识为一个切面类，`Spring`会将该类作为一个切面管理。
* `@Component`：将该类作为一个 `Spring` 组件。
* `@Order(1)`：主要用来控制配置类的加载顺序，`bean`加载的优先级，值越小，越先被加载。
* @Pointcut：定义一个切入点，Pointcut是植入Advice的触发条件。每个Pointcut的定义包括2部分，一是表达式，二是方法签名。方法签名必须是 public及void型。可以将Pointcut中的方法看作是一个被Advice引用的助记符，因为表达式不直观，因此我们可以通过方法签名的方式为此表达式命名。因此Pointcut中的方法只需要方法签名，而不需要在方法体内编写实际代码。
* @Before：在切入点开始处切入内容。标识一个前置增强方法，相当于BeforeAdvice的功能。
* @Around：环绕增强，在切入点前后切入内容，并自己控制何时执行切入点自身的内容。
* @After: final增强，不管是抛出异常或者正常退出都会执行。
* @AfterThrowing ：用来处理当切入内容部分抛出异常之后的处理逻辑。
* @AfterReturning：后置增强，相当于AfterReturningAdvice，方法正常退出时执行。



步骤一：添加 `AOP` 相关依赖。

步骤二：`Controller` 层的日志封装类`WebLog`。

步骤三：统一日志处理切面类`WebLogAspect`。



添加 `AOP` 相关依赖：

```java
<dependency>
  <groupId>org.springframework.boot</groupId>
  <artifactId>spring-boot-starter-aop</artifactId>
</dependency>
```

控制层的日志封装类：

```java
package com.wangxiong.common.domain;

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

`Page`接口：

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



分页参数`Model`：

```java
/**
 * 分页参数
 **/

public class PageVo implements Serializable{
    private static final long serialVersionUID = -1305720016123712695L;
        //当 前页
        private String pageNum;

        // 每页显示条数
        private String pageSize;

        // 查询参数
        private Map<String, Object> parameters = new HashMap<>(10);

        // 排序参数
        private Map<String, Object> sort = new HashMap<>(10);;

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



定义`ReturnInfo`信息：

```java
public class ReturnInfo {
    /**
     * 是否成功
     */
    private ReturnState status;

    /**
     * 返回的实体类对象
     */
    private Object returnData;


    private Map<String,Object> pageInfo;

    /**
     * 返回的响应描述
     */
    private String message;

    /**
     * 返回的错误码
     */
    private String errorCode;

    /**
     * 请求地址
     */
    private String url;

    public ReturnInfo() {
        this(ReturnState.SUCCESS, StringUtils.EMPTY);
    }

    public ReturnInfo(Object returnData) {
        this(ReturnState.SUCCESS, returnData);
    }


    public ReturnInfo(ReturnState status, Object returnInfo) {
        this.status = status;
        this.returnData = returnInfo;
    }

    public ReturnInfo(ReturnState status, Object returnData, String message) {
        this.status = status;
        this.returnData = returnData;
        this.message = message;
    }

    public ReturnInfo(ReturnState status, String errorCode, String message) {
        this.status = status;
        this.errorCode = errorCode;
        this.message = message;
    }

    public ReturnInfo(ReturnState status, Object returnData, String message, String url) {
        this.status = status;
        this.returnData = returnData;
        this.message = message;
        this.url = url;
    }

    public ReturnInfo(ReturnState status, Object returnData, Map<String,Object> pageInfo, String message) {
        this.status = status;
        this.returnData = returnData;
        this.pageInfo = pageInfo;
        this.message = message;
    }
    public ReturnState getStatus() {
        return this.status;
    }

    public void setStatus(ReturnState status) {
        this.status = status;
    }

    public Object getReturnData() {
        return returnData;
    }

    public void setReturnData(Object returnData) {
        this.returnData = returnData;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public Object getPageInfo() {
        return pageInfo;
    }

    public void setPageInfo(Map<String,Object> pageInfo) {
        this.pageInfo = pageInfo;
    }

    public String getErrorCode() {
        return errorCode;
    }

    public void setErrorCode(String errorCode) {
        this.errorCode = errorCode;
    }

    @Override
    public String toString() {
        return " status=" + this.status +" message=" + this.message+ " returnData=" + this.returnData;
    }
}

```

分页处理切面：

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





# 数据类型

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

`List<Object>`：

```java

```

`List<Map<String,Object>> filter = new ArrayList<>();`

```java
[
  {
    "id":1,
    "name":"wangxiong01"
  },
  {
    "id":2,
    "name":"wangxiong02"
  }
]
```

