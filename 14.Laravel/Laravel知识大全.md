# 问题列表

1、[简单说一下Laravel的执行过程 ？阐述`Laravel`框架一次请求的完整生命周期。](#01)







# 解答列表

1、<span id="01">简单说一下Laravel的执行过程 ？阐述`Laravel`框架一次请求的完整生命周期。</span>

当`Web`服务器（`Apache/Nginx`）的请求被导向单入口文件 `public/index.php`时， `index.php`中极简的代码出色的完成整个`HTTP`的请求，逻辑组织既严谨又完美，以下便是其优雅的六行代码：

```php
<?php
// 加载项目依赖
require __DIR__.'/../bootstrap/autoload.php';
// 创建服务容器实例并完成内核的绑定（包含`Http`内核和`Console`内核）和注册异常处理
$app = require_once __DIR__.'/../bootstrap/app.php';
//  服务容器解析内核实例
$kernel = $app->make(Illuminate\Contracts\Http\Kernel::class);
//  接收请求并返回响应
$response = $kernel->handle(
    $request = Illuminate\Http\Request::capture()
);
// 发送响应到客户端
$response->send();
// 终止程序，清理中间件
$kernel->terminate($request, $response);
```

上面这段单入口文件代码完成了一次`Http`请求的完整生命周期，具体过程如下：

> ① 首先，所有的请求都会被`web`服务器导向单入口文件`public/index.php`；
> ② 第一步注册`composer`自动加载，以完成项目中所需依赖的加载；
> ③ 第二步在创建服务容器实例时，不仅包括项目基础服务、项目服务提供者别名、目录路径等在内的一系列注册工作，还会绑定 `HTTP` 内核及 `Console` 内核到 `APP`服务容器，绑定异常处理到服务容器；
> ④ 第三步通过服务容器去解析内核实例， 将`HTTP` 内核定义的中间件组注册到路由器，注册完后就可以在实际处理 `HTTP` 请求前调用这些中间件实现过滤请求的目的；
> ⑤ 第四步将`HTTP`请求实例注册到服务容器，并且通过启动引导程序来加载环境变量、全局配置、异常处理、注册门面、注册服务提供者、启动服务等，随后请求被分发到匹配的路由，在路由中执行中间件以过滤不满足校验规则的请求，只有通过中间件处理的请求才最终处理实际的控制器或匿名函数生成响应结果；
> ⑥ 第五步获取到响应结果后发送给请求的客户端；
> ⑦ 第六步运行程序`$kernel->terminate()`做一些善后清理工作，并最终退出脚本。