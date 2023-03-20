# Spring 的注解式开发

摘要：本篇文章主要是对 `Spring` 的几个基本注解式开发的介绍。其中包括实例化相关注解`@Component`和其子类注解`@Repository`、`@Service`、`@Controller`进行介绍 ；控制单列和多例的注解`@Scope`；对注入注解`@Autowired`和`@Resource`进行对照和区别；最后对事务的相关注解`@Transactional`进行了重点介绍，对事务`propagation`传播属性和事务`isolation`隔离级别也进行了讲解。

开启注解扫描配置：

```xml
	<!--开启注解扫描-->
  <context:component-scan base-package="com"/>
```

## 实例化相关注解

###  @Component

>  格式：@Component(value="conversionImpl")

作用：通用的创建实例的注解，用来创建当前这个类的实例。

```xml
<bean id="conversionImpl" class="com.wangxiong.xxx"></bean>
```

细节:	`value` 属性用来指定创建的对象在工厂中的唯一标识，如果不指定，默认创建对象在工厂中的标识为类名首字母小写。该注释的修饰范围只能作用在类上。

示例：

```java
@Component("conversionImpl")
// 默认的 spring 中的 Bean id 为 conversionImpl(首字母小写)
public class ConversionImpl implements Conversion {
    @Autowired
    private RedisClient redisClient;
}
```

### @Repository 

> @Repository(value="conversionImpl")

作用:   `@component`的子类注解专用于 DAO 组件的创建，通常加在 DAO 组件上。

```xml
<!--创建DAO对象-->
<bean class="org.mybatis.spring.mapper.MapperScannerConfigurer">
  <property name="sqlSessionFactoryBeanName" value="sqlSessionFactory"/>
  <property name="basePackage" value="com.wangxiong.dao"/>
</bean>
```

细节：`value` 属性用来指定创建的对象在工厂中的唯一标识，如果不指定，默认创建对象在工厂中的标识为类名首字母小写。该注释的修饰范围只能作用在类上。

示例：

```java
@Repository
// 使用 @Repository 将 DAO 类声明为 Bean 
public class UserDaoImpl implements UserDao{ …… } 
```

### @Service

> @Service(value="conversionSerImpl")

作用:   `@component`的子类注解专用于 Service 组件的创建，通常加在 Service 组件上。

```xml
<!--管理Service组件对象-->
<bean id="userService" class="com.wangxiong.service.UserServiceImpl">
  <property name="userDAO" ref="userDAO"/>
</bean>
```

细节：`value` 属性用来指定创建的对象在工厂中的唯一标识，如果不指定，默认创建对象在工厂中的标识为类名首字母小写。该注释的修饰范围只能作用在类上。

示例：

```java
@Service
// 使用 @Service 将 Service 类声明为 Bean 
public class SmsHomePopServiceImpl implements SmsHomePopService { …… } 
```

### @Controller 

> @Controller(value="conversionController")

作用:   `@component`的子类注解专用于 Action 组件的创建,通常加在 Action 组件上。

```xml
<!--管理Action scope="03.prototype"-->
<bean id="userAction" class="com.wangxiong.action.UserAction" scope="prototype">
  <property name="userService" ref="userService"/>
</bean>
```

细节：`value` 属性用来指定创建的对象在工厂中的唯一标识，如果不指定，默认创建对象在工厂中的标识为类名首字母小写。该注释的修饰范围只能作用在类上。

示例：

```java
@Controller
// 使用 @Controller 将 Controller 类声明为 Bean
public class SmsHomePopController { …… } 
```



其中，`DAO`层、`Service`层、`Controller`层有各自细分的注解，除了此三层外的注解都用 `@Component`通用注解。

## 单例和多例的注解

### @Scope

> @Scope(value="singleton|prototype")

作用:   `@Scope`用来控制这个实例在工厂中的创建次数。
细节：`value` 属性`singleton`为单例，`prototype`为多例，默认单例。

示例：

```java
@Scope("prototype")
// 使用 Scope 将 UserServiceImpl 类实例化为多例
public class UserServiceImpl implements UserService { …… } 
```

## 注入相关的注解

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

## 事务的相关注解

> @Transactional(isolation = Isolation.DEFAULT,propagation = Propagation.REQUIRED)

作用：主要作用一是进行事务控制或事务细粒度配置，二是简化事务切面配置。

```xml
<!--创建事务管理器-->
<bean id="transactionManager"     		class="org.springframework.jdbc.datasource.DataSourceTransactionManager">
  <property name="dataSource" ref="dataSource"/>
</bean>

<!--基于事务管理器创建事务通知对象并配置事务细粒度控制-->
<tx:advice id="txAdvice" transaction-manager="transactionManager">
  <tx:attributes>
    <tx:method name="save*"/>
    <tx:method name="update*"/>
    <tx:method name="delete*"/>
    <tx:method name="find*" propagation="SUPPORTS"/>
  </tx:attributes>
</tx:advice>

<!--配置事务切面-->
<aop:config>
  <aop:pointcut id="pc" expression="within(com.wangxiong.service.*ServiceImpl)"/>
  <aop:advisor advice-ref="txAdvice" pointcut-ref="pc"/>
</aop:config>
```

细节：修饰范围用在类上（主要用在业务层组件类上）或者是方法上。当加在类上的时候，用来给类中所有的方法加入事务控制；当加在方法上时，代表当前方法加入事务控制；当类上和方法上同时存在该注解时，局部（方法）优先。

注解属性：

| 属性名称       | 说明                                                         |
| -------------- | ------------------------------------------------------------ |
| propagation    | 用来控制传播属性。                                           |
| isolation      | 用来控制隔离级别。                                           |
| timeout        | 用来设置超时性。-1  永不超时；大于0，表示超时时间（秒）。    |
| rollback-for   | 出现什么类型异常不会滚。 java.lang.RuntimeException          |
| norollback-for | 出现什么类型异常回滚 。默认出现 RuntimeException 及其子类异常回滚。 |
| readonly       | 用来设置事务读写性。true 代表只读，不能执行增删改操作。      |

如果要使用事务注解，需要在配置文件中开启事务注解生效：

```xml
<tx:annotation-driven transaction-manager="transactionManager"></tx:annotation-driven>
```

事务`propagation`传播属性：

| 属性名称      | 说明                                                         |
| ------------- | ------------------------------------------------------------ |
| REQUIRED      | 需要事务。如果外层没有事务则开启新的事务  ；如果外层存在事务，则融入当前事务。 |
| SUPPORTS      | 支持事务。如果外层没有事务不会开启新的事务；如果外层存在事务，则融入当前事务。 |
| REQUIRES_NEW  | 每次开启新的事务 。如果外层存在事务，外层事务挂起，自己开启新的事务执行，执行完成，恢复外层事务继续执行。 |
| NOT_SUPPORTED | 不支持事务。如果外层存在事务，外层事务挂起，自己以非事务方式执行，执行完成恢复外部事务执行。 |
| NEVER         | 不能有事务，存在事务报错。                                   |
| MANDATORY     | 强制事务， 没有事务报错。                                    |
| NESTED        | 嵌套事务 。事务之间可以嵌套运行，Oracle、MySQL数据库不支持。 |

事务`isolation`隔离级别：

| 属性名称                      | 说明                                                         |
| ----------------------------- | ------------------------------------------------------------ |
| DEFAULT                       | 使用数据库默认的隔离级别。                                   |
| READ_UNCOMMITTED（读未提交 ） | 一个客户端读到了另一个客户端没有提交的数据 ，会出现脏读现象。 |
| READ_COMMITTED（读提交）      | 一个客户端只能读到另一个客户端提交的数据  ，避免脏读现象。Oracle 默认隔离级别。 |
| REPEATABLE_READ（可重复读）   | 主要是用来避免不可重复读现象出现 ，行锁。 MySQL 默认隔离级别。 |
| SERIALIZABLE（序列化读）      | 主要是用来避免幻影读现象出现 ，表锁。                        |