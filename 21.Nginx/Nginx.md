# Nginx 504 Gateway Time-out

| 配置项                     | 默认值   | 作用说明                 |
| ----------------------- | ----- | -------------------- |
| `proxy_connect_timeout` | `60s` | 与后端服务器建立连接的超时时间      |
| `proxy_send_timeout`    | `60s` | 向后端发送请求的超时时间         |
| `proxy_read_timeout`    | `60s` | 等待后端响应的超时时间          |
| `send_timeout`          | `60s` | 向客户端发送响应的超时时间（不活跃连接） |


链路一、前端 Nginx 配置超时时间太短

```
   location /api/xxx/ {
        proxy_connect_timeout 120s;
        proxy_send_timeout 120s;
        proxy_read_timeout 120s;
        send_timeout 120s;
        proxy_pass http://testxxx.xxx.xxx/kjj/;
    }
```
链路二、ingress 配置

如果你使用的是 Nginx Ingress Controller，可以直接在 Ingress 资源中添加如下 Annotation：

```
kind: Ingress
apiVersion: networking.k8s.io/v1
metadata:
  name: mkt-test
  namespace: xxx-mkt
  creationTimestamp: null
  labels:
    app: mkt-test
    stark-app: mkt
    stark-ns: xxx-mkt
  annotations:
    nginx.ingress.kubernetes.io/proxy-body-size: 64m
    nginx.ingress.kubernetes.io/proxy-connect-timeout: '300'
    nginx.ingress.kubernetes.io/proxy-read-timeout: '300'
    nginx.ingress.kubernetes.io/proxy-send-timeout: '300'
spec:
  ingressClassName: nginx
  rules:
```


链路三、后端服务 如GO的配置

`*.yaml`文件中的配置：

```
http:
  addr: 0.0.0.0:8000
  timeout: 300
```

```
var Set = wire.NewSet(
	NewServer,
)

func NewServer(
	demo *demo.Service,
	hotlist *hotlist.Service,
	proxy *proxy.Service,
	assistant *assistant.Service,
) *http.Server {
	opts := setOpt()
	srv := http.NewServer(opts...)
	engine := gin.New()
	registerPprof(srv)
	registerHealth(srv)
	registerDoc(srv)
	registerProxy(srv)
	engine.Use(middlewares())

	//服务注册
	mktNoLogin := engine.Group("/xxxx")
	mkt := engine.Group("/xxx")
	mkt.Use(kgin.Middlewares(auth.Auth()))
	demo.Register(mktNoLogin)
	hotlist.Register(mkt, mktNoLogin)
	proxy.Register(engine)
	assistant.Register(engine.Group("/assistant"))

	srv.HandlePrefix("/", engine)
	return srv
}

func setOpt() []http.ServerOption {
	var opts []http.ServerOption
	if network := config.GetString("http.network"); network != "" {
		opts = append(opts, http.Network(network))
	}
	if addr := config.GetString("http.addr"); addr != "" {
		opts = append(opts, http.Address(addr))
	}
	if timeout := config.GetInt("http.timeout"); timeout > 0 {
		opts = append(opts, http.Timeout(time.Duration(timeout)*time.Second))
	}

	opts = append(opts, http.ErrorEncoder(reply.ErrorEncoder))
	opts = append(opts, http.ResponseEncoder(reply.ResponseEncoder))

	return opts
}
```

标准库中的超时，通过 http.Server 来设置超时时间：

```
srv := &http.Server{
Addr:         ":8080",
Handler:      myHandler, // 可以是 mux、Gin、Echo 等
ReadTimeout:  120 * time.Second, // 读取请求头 + body 的最大时间
WriteTimeout: 120 * time.Second, // 写响应给客户端的最大时间
IdleTimeout:  120 * time.Second, // keep-alive连接最大空闲时间
}

log.Fatal(srv.ListenAndServe())
```

如果是调用下游（如外部 API）也要设置超时：

如果服务是网关或中间层，还要设置 http.Client 的超时，例如：

```
client := &http.Client{
    Timeout: 30 * time.Second, // 整体请求超时
}
resp, err := client.Get("http://example.com/api")
```

你没有自定义 http.Server，所以使用的是 Go 标准库的默认超时配置:

| 配置项            | 默认值       |
| -------------- | --------- |
| `ReadTimeout`  | **0**（无限） |
| `WriteTimeout` | **0**（无限） |
| `IdleTimeout`  | **0**（无限） |


链路四：后端服务崩溃或无响应

* 服务没有监听端口，或已挂掉。

* 有时数据库连接池已耗尽或 Redis 卡死。

链路五：负载过高 / 网络不通

* 服务器 CPU / 内存使用率过高，响应变慢。

* 网络断开、DNS 配置错误等

配置说明：

```
nginx.ingress.kubernetes.io/proxy-read-timeout: "120"
```

* 作用：指定 Nginx 等待后端服务器响应的最大时间（单位为秒）。

* 对应 Nginx 配置项：proxy_read_timeout

* 典型用途：处理一些耗时较长的接口，比如视频转码、AI 推理、慢查询等。

```
nginx.ingress.kubernetes.io/proxy-send-timeout: "120"
```

当客户端发起请求，Ingress 会将请求转发到后端服务（你的 Go 服务）。 

如果请求体较大、网络慢，或者是 POST/PUT 等带有 body 的请求，Ingress 需要一定时间将请求数据发送给后端。

proxy-send-timeout 就是限制这个“发送过程”的最长时间，如果超时，就会中断连接并返回错误。



# Nginx 413 Request Entity Too Large

前端nginx配置、Ingress配置、服务端代码配置 三者保持一致。

方法一、前端Nginx配置文件中：`/etc/nginx/nginx.conf`

解决步骤： 打开 Nginx 配置文件，通常在：

```
/etc/nginx/nginx.conf
```

或者某个 `server` 块中（比如 `/etc/nginx/sites-enabled/default`）

增加或修改如下配置：

```
http {
...
client_max_body_size 50M;  # 或者更大，根据你的上传大小设置
}
```

或者在 server 块里：

```
server {
...
client_max_body_size 50M;
}
```

重启 Nginx：

```
sudo nginx -s reload
```

方法二： 后端服务（如 `PHP / Node.js / Go`）配置限制
如果你已经修改了 Nginx，但问题仍存在，还需要检查后端服务本身的限制。

PHP：php.ini 文件中
```
post_max_size = 50M
upload_max_filesize = 50M
```

Express (Node.js)：
```
app.use(express.json({ limit: '50mb' }));
app.use(express.urlencoded({ limit: '50mb', extended: true }));
```

Gin (Go)：你需要手动读取 Body 并限制最大 size：

```
r.MaxMultipartMemory = 50 << 20 // 50 MB
```

方法三：Ingress 默认的 NGINX 配置中，`client_max_body_size` 默认值是 1MB 或 20MB 或 8MB，上传大文件时很容易触发 413 错误。

需要在 Ingress 的 annotations 中增加配置：

```
metadata:
annotations:
nginx.ingress.kubernetes.io/proxy-body-size: "100m"
```

在容器的ingress配置与编辑中，找到便签与注解，键 `nginx.ingress.kubernetes.io/proxy-body-size` 值 64m。

| 参数                      | 控制阶段              | 默认值 | 典型用途             |
| ----------------------- | ----------------- | --- | ---------------- |
| `proxy_connect_timeout` | 连接到后端的最大时间（TCP握手） | 60s | 后端服务繁忙或启动缓慢时     |
| `proxy_send_timeout`    | 向后端发送请求数据的超时      | 60s | 请求体较大时（POST/PUT） |
| `proxy_read_timeout`    | 从后端读取响应数据的超时      | 60s | 后端处理慢或响应大时       |
| `send_timeout`          | 向客户端发送响应的超时时间     | 60s | 客户端接收慢时          |



# 实战中前端Nginx配置模板
```
user                 nobody;
worker_processes     4;
worker_rlimit_nofile 65535;

error_log  /data/nginx/logs/error.log  notice;

events {
    use epoll;
    worker_connections  4096;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    full_log_format  full  '$remote_addr $request_length $body_bytes_sent $request_time[s] - - [$time_local] '
                           '"$request" $status $http_referer "-" "$http_user_agent" $http_host $server_addr '
                           '$http_x_forwarded_for $http_x_real_ip';
    full_access_log  /data/nginx/logs/allweb.log  full;

    log_format       combinedio  '$remote_addr - $remote_user [$time_local] '
                                 '"$request" $status $body_bytes_sent '
                                 '"$http_referer" "$http_user_agent" $request_length $request_time $upstream_response_time';
    access_log off;

    sendfile                     on;
    gzip                         on;
    tcp_nopush                   on;
    tcp_nodelay		         on;

    keepalive_timeout            0;
    client_body_timeout          10;
    client_header_timeout        10;

    client_header_buffer_size    1k;
    large_client_header_buffers  4  4k;
    output_buffers               1  32k;
    client_max_body_size	 64m;
    client_body_buffer_size      256k; 
 
    #lua_package_path "/usr/local/luajit/share/lua/5.1/ngx_metric/?.lua;;";
    #lua_shared_dict shared_dict 64M;
    #log_by_lua_file /usr/local/luajit/share/lua/5.1/ngx_metric/ngx_metric.lua;

    server {
        listen       80;
        server_name  localhost;

        location /server-status {
            stub_status  on;
            allow        127.0.0.1;
            deny         all;
        }

        location /status {
            include      fastcgi.conf;
            fastcgi_pass 127.0.0.1:9000;
            allow        127.0.0.1;
            deny         all;
        }

        #location /ngx/metric {
        #    content_by_lua_file /usr/local/luajit/share/lua/5.1/ngx_metric/ngx_metric_output.lua;
        #    allow 127.0.0.1;
        #    deny all;
        #}
    }

    include include/*/vhost.conf;
}


server {
    listen 80;
    listen 443 ssl;
    server_name test.x.xxx.xx;

    root /home/q/system/fe/xxx-xxx/;
    
    # 协商缓存
    add_header Cache-Control no-cache;
    add_header Cache-Control public;
    # 跨域处理
    add_header Access-Control-Allow-Origin *;
    add_header Access-Control-Allow-Methods 'GET, POST, OPTIONS';
    add_header Access-Control-Allow-Headers 'DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization';

    # 预检请求
    if ($request_method = 'OPTIONS') {
        return 204;
    }

    access_log /data/nginx/logs/test.xxx.xxx.xx/web/test.xxx.xxx.xx-access.log combinedio;
    error_log /data/nginx/logs/test.xxx.xxx.xx/web/test.xxx.xxx.xx-error.log;

    client_body_temp_path /data/nginx/client_body_temp/ 1 2;
    proxy_temp_path /data/nginx/proxy_temp/ 1 2;
    fastcgi_temp_path /data/nginx/fastcgi_temp/ 1 2;

    # The 'ENV' file in document root contains various environment information,
    # such as the path of document root, the location of log file, and so on.
    # For security reason, you CAN NOT remove this location!!
    location = /ENV {
        allow 127.0.0.1;
        deny all;
    }
    
     #  service worker
    location /xx-app/sw/ {
        alias /home/q/system/fe/quick-clip/sw/;
        try_files $uri $uri.js;
        add_header Service-Worker-Allowed /;
    }

    # 
    location /xx-app/ {
        alias /home/q/system/fe/quick-clip/;
        try_files $uri $uri.html $uri.htm;
    }
  
    # 新服务接口，适用于爆款视频等
    location /api/xxx/ {
        proxy_pass http://testxxx.xxx.xx/xxx/;
    }

    # 旧服务接口，适用于如user/info等
    location /api/brain/  {
        proxy_pass http://testxxxapi.xxx.xx/api/;
    }
    
    # 后端接口
    location /api/kxx/ {
        proxy_pass http://testkxx.xxx.xx/kxx/;
    }

    # 前端
    location / {
        root /home/q/system/fe/xxx-xxxx;
        try_files $uri $uri.html $uri/ /index.html;
    }
    
    #ssl  on;
    ssl_session_cache  shared:SSL:50m;
    ssl_session_timeout  300;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers HIGH:!aNULL:!MD5;
    ssl_prefer_server_ciphers  on;

    ssl_certificate /usr/local/nginx/ssl/test.x.xxx.xx.crt;
    ssl_certificate_key /usr/local/nginx/ssl/test.x.xxx.xx.key;

}
```

# 刷新移动端页面报404

问题本质分析：

根据配置和现象，根本原因在于 **`if`指令和`root`路径的优先级冲突**，导致移动端刷新时`try_files`未能正确继承`root`路径。Nginx的`if`属于rewrite模块，其内部作用域中修改的`root`不会传递到外层`location`的`try_files`逻辑，导致路径解析混乱（移动端刷新时实际仍使用PC端的`root`路径）。

```
# 在 http{} 级别定义一个 map
map $http_user_agent $mobile_root {
    default "/home/q/system/web.xxx.xxx.cn/CloudDrive/";  # 默认 PC 端目录
    "~*android|mobile safari|openharmony|aphone|meego; nokian9|blackberry|rim tablet os|iphone|ipod|iemobile|opera mini|juc|iuc|opera mobi|avantgo|blazer|elaine|hiptop|palm|plucker|xiino|windows ce; (iemobile|ppc|smartphone)|windows phone os|acer|zte|lenovo|moto|samu|nokia|sony|kindle|240x320|mobile|mmp|ucweb|midp|pocket|psp|symbian|smartphone|treo|up.browser|up.link|vodafone|wap" "/home/q/system/web.xxx.xxx.cn/CloudDrive_m/";  # 移动端目录
}

server {
    server_name 10.xxx.xxx.xxx;
  
    charset utf-8;
    server_tokens off;

    root $mobile_root;
  
    index index.html index.htm index.php;
  
    error_log  /dev/stderr;

    location / {
        add_header Last-Modified $date_gmt;
        add_header Cache-Control 'no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0';
        if_modified_since off;
        expires off;
        etag off;
  
        try_files $uri $uri/ /index.html;
    }

    if ($http_user_agent ~ ^$) {
        return 403;
    }

    # 爬虫
    if ($http_user_agent ~* (Scrapy|Curl|HttpClient|HeadlessChrome)) {
        return 403;
    }

    location /status.html {
        #  ngx.say("ok");
    }

    location /status.php {
        try_files /not.exists /status.html;
    }
  
    location /psp_jump.html {
        root /home/q/system/web.xxx.xxx.cn/CloudDrive/;
    }
  
    gzip               on;
    gzip_min_length    1024;
    gzip_buffers       4 16k;
    gzip_comp_level    6;
    gzip_types         text/plain application/javascript application/x-javascript text/css application/xml;
    gzip_vary          on;
}
```

# 配置新域名

```
server {
    listen 443 ssl;
    server_name xxx.xxx.cn;
  
    add_header Last-Modified $date_gmt;
    add_header Cache-Control 'no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0';
    if_modified_since off;
    expires off;
    etag off;

    location / {
        proxy_pass https://xx.xxx.xxx.cn/managepage/;
    }

    location /managepage {
        proxy_pass https://xx.xxx.xxx.cn/;
    }
}

```

* add_header Last-Modified $date_gmt;: 添加Last-Modified头，值为当前的GMT时间。这告诉客户端资源的最后修改时间。
* add_header Cache-Control 'no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0';: 添加Cache-Control头，指示客户端和代理服务器不要缓存响应。
* if_modified_since off;: 禁用If-Modified-Since请求头的处理。
* expires off;: 禁用Expires头，以防止资源被缓存。
* etag off;: 禁用ETag头，这也是用于缓存控制的一个HTTP头。
* location / { ... }: 针对根路径（/）的请求，将它们反向代理到https://xx.xxx.xxx.cn/managepage/。即，所有对https://xxx.xxx.cn/的请求会被转发到https://xx.xxx.xxx.cn/managepage/。
* location /managepage { ... }: 针对/managepage路径的请求，将它们反向代理到https://xx.xxx.xxx.cn/。即，所有对https://xxx.xxx.cn/managepage的请求会被转发到https://xx.xxx.xxx.cn/。

# 使用自己的域名访问第三方网站

例如：https://xx.xxx.xx/video

```
location ~* \.(js|png|jpg|css|mp4|svg|jpeg|html)$ {
        root /home/xxx/xxxx/;
        #如果是手机移动端访问内容
        if ($http_user_agent ~* "^((.*android.*)|(.*Mobile Safari.*)|(.*Aphone.*)|(.*MeeGo; NokiaN9*.)|(.*blackberry.*)|(.*rim tablet os.*)|(.*iphone.*)|(.*ipod.*)|(.*IEMobile*.)|(.*opera mini.*)|(.*JUC.*)|(.*IUC.*)|(.*opera mobi.*)|avantgo|blazer|elaine|hiptop|palm|plucker|xiino|(windows ce; (iemobile|ppc|smartphone))|(.*windows phone os.*)|acer|zte|lenovo|moto|samu|nokia|sony|kindle|240x320|mobile|mmp|ucweb|midp|pocket|psp|symbian|smartphone|treo|up.browser|up.link|vodafone|wap)") {
            root  /home/xxx/xxxx/react-m/;
        }
    
        if ($arg_h5 = "true") {
            root  /home/xxx/xxxx/react-m/;
        }

        if ($http_referer ~ '/privacy|/agreement|/a/|/i/|/feedback') {
            root /home/xxx/xxxx/;
        }

        if ($uri ~ '/secure.|/agree.|/feed_comment.') {
            root /home/xxx/xxxx/;
        }
    
        try_files $uri @static;
    }
  
location /video {
        proxy_pass http://xxx.xx.xxx.xx/;
        proxy_set_header Host xxx.xx.xxx.xx;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
    location /static {
        proxy_pass http://xxx.xx.xxx.xx;
    }
    location /api/ {
        proxy_pass http://xxx.xx.xxx.xx;
    }
    location /zhipu.json {
        proxy_pass http://xxx.xx.xxx.xx;
    }
  
    location @static {
        proxy_pass http://xxx.xx.xxx.xx;
    }
```

# Nginx 参数

基础参数：

```
user nginx;
worker_processes 6;
worker_rlimit_nofile 1024;
include /etc/nginx/modules-enabled/*.conf;
error_log /xxx/nginx/log/nginx_error.log error;

events {
    use epoll;
    worker_connections 1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;
    autoindex off;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3;
    ssl_prefer_server_ciphers on;
    charset utf-8;
    log_format main "$remote_addr\t$remote_user\t[$time_local]\t$request_method\t$host\t$request_uri\t"
                    "$request_time\t$status\t$body_bytes_sent\t'$http_referer'\t"
                    "'$http_user_agent'\t'$http_x_forwarded_for'\t$upstream_response_time\t$upstream_status";
    access_log /xxx/nginx/log/nginx_access.log main;

    sendfile        on;
    keepalive_timeout  65;
    server_names_hash_max_size 1024;
    server_names_hash_bucket_size 128;
    client_header_buffer_size 40k;
    large_client_header_buffers 8 256k;
    client_header_timeout  6m;
    client_body_timeout    6m;
    send_timeout           6m;
    tcp_nopush     on;
    tcp_nodelay    on;
    types_hash_max_size 4096;
    client_max_body_size 2048m;
    client_body_buffer_size 6m;
    proxy_connect_timeout 3600;
    proxy_send_timeout 3600;
    proxy_read_timeout 3600;
    proxy_buffer_size 4k;
    proxy_buffers 8 64k;
    proxy_busy_buffers_size 128k;
    proxy_temp_file_write_size 64k;
    proxy_intercept_errors  on;
    proxy_headers_hash_max_size 1024;
    proxy_headers_hash_bucket_size 256;
    variables_hash_max_size 512;
    variables_hash_bucket_size 128;

    gzip on;
    gzip_min_length 1k;
    gzip_buffers     8 32k;
    gzip_comp_level 2;
    gzip_http_version 1.1;
    gzip_types text/plain application/x-javascript text/css application/xml text/javascript image/gif image/png;
    gzip_vary on;
    map_hash_max_size 102400;
    map_hash_bucket_size  128;

    fastcgi_connect_timeout 300;
    fastcgi_send_timeout 300;
    fastcgi_read_timeout 300;
    fastcgi_buffer_size 64k;
    fastcgi_buffers 4 64k;
    fastcgi_busy_buffers_size 128k;
    fastcgi_temp_file_write_size 128k;

    server_tokens off;
    fastcgi_intercept_errors on;

    map $http_upgrade $connection_upgrade {
        default upgrade;
        '' close;
    }

    include /etc/nginx/conf.d/*.conf;
    include /etc/nginx/sites-enabled/*;
    include /xxx/nginx/conf/*.conf;
}
```

参数说明：

* `user nginx;`: 指定Nginx worker进程的运行用户，这里是"nginx"。
* `worker_processes 6;`: 指定Nginx启动时创建的worker进程的数量，这里是6个。通常，可以将其设置为机器的CPU核心数。
* `worker_rlimit_nofile 1024;`: 设置每个Nginx worker进程能够打开的最大文件描述符数量。这对于限制Nginx的资源使用是有用的。
* `include /etc/nginx/modules-enabled/*.conf;`: 包含指定目录下的所有以".conf"为后缀的文件。通常，这个目录用于启用或禁用Nginx模块。
* `error_log /xxx/nginx/log/nginx_error.log error;`: 配置错误日志的路径和级别，将错误日志记录到指定文件中。
* `events {...}`: 配置Nginx处理事件的模块，这里使用了epoll作为事件模型。
* `use epoll;` #参考事件模型，`use [ kqueue | rtsig | epoll | /dev/poll | select | poll ]`;
  `#epoll`模型是Linux 2.6以上版本内核中的高性能网络I/O模型，如果跑在FreeBSD上面，推荐使用kqueue模型。
* `worker_connections 1024`; 指定每个worker进程能够处理的最大连接数。
* `http {...}`: 定义HTTP模块的配置块，包含了Nginx的主要HTTP配置。
* `include mime.types;`: 包含MIME类型配置文件，用于指定文件扩展名和相应的MIME类型。
* `default_type application/octet-stream;`: 设置默认的MIME类型，如果无法从文件扩展名中确定。
* `autoindex off;`: 禁用目录列表显示功能。
* `ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3;`: 指定支持的SSL/TLS协议版本。
* `ssl_prefer_server_ciphers on;`: 设置Nginx使用服务器端的密码组合顺序。
* `charset utf-8;`: 设置字符集为UTF-8。
* `log_format main {...}`: 定义访问日志的格式。
* `access_log /xxx/nginx/log/nginx_access.log main;`: 配置访问日志的路径和格式。
* 下面是一系列的HTTP配置，包括文件传输、超时设置、缓冲区大小等，用于优化Nginx的性能。
* `fastcgi_*`:FastCGI 相关参数是为了改善网站的性能：减少资源占用，提高访问速度。
* `server_tokens off;`: 这个指令用于控制Nginx是否在HTTP响应头中发送服务器版本信息。关闭时增加安全性。
* `fastcgi_intercept_errors on;`: 这个指令用于控制当FastCGI后端返回HTTP错误时，是否将这些错误交由Nginx来处理。
* `gzip on;`: 启用Gzip压缩。
* `map {...}`: 定义一个变量映射，根据请求头中的Upgrade字段来设置连接升级。例如`websocket`。
* `include /etc/nginx/conf.d/*.conf;`: 包含所有在指定目录下以".conf"为后缀的文件。
* `include /etc/nginx/sites-enabled/*;`: 包含指定目录下的所有符号链接文件，通常用于包含虚拟主机配置。
* `include /xxx/nginx/conf/*.conf;`: 包含指定目录下的所有以".conf"为后缀的文件。

日志格式详解：

```
log_format main "$remote_addr\t$remote_user\t[$time_local]\t$request_method\t$host\t$request_uri\t"
                "$request_time\t$status\t$body_bytes_sent\t'$http_referer'\t"
                "'$http_user_agent'\t'$http_x_forwarded_for'\t$upstream_response_time\t$upstream_status";
```

* `$remote_addr`: 记录客户端的IP地址。
* `$remote_user`: 如果启用了基本身份验证，这将记录远程用户的用户名。
* `[$time_local]`: 记录访问时间，使用本地时间格式，包裹在方括号中。
* `$request_method`: 记录HTTP请求的方法，如GET、POST等。
* `$host`: 记录请求的主机名。
* `$request_uri`: 记录完整的请求URI，包括参数。
* `$request_time`: 记录从接收客户端请求到向客户端发送响应的总时间，以秒为单位。
* `$status`: 记录HTTP响应状态码。
* `$body_bytes_sent`: 记录发送给客户端的字节数。
* `'$http_referer'`: 记录HTTP Referer头，表示访问来源。
* `'$http_user_agent'`: 记录HTTP User-Agent头，表示客户端的用户代理信息。
* `'$http_x_forwarded_for'`: 记录X-Forwarded-For头，表示经过代理服务器时的客户端真实IP地址。
* `$upstream_response_time`: 记录与后端服务器建立连接、发送请求以及接收响应的总时间，以秒为单位。
* `$upstream_status`: 记录从上游服务器接收到的HTTP响应状态码。

请求示例：

```
xx.xx.xxx.xx - [24/Nov/2023:16:10:35 +0800]	GET	xxx.xx.xx.xx /api/v1/xxx/xxx
0.039 200 658 'https://xx.xx.xx.xx/' 
'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/xx.36 (KHTML, like Gecko) Chrome/11x.0.0.0 Safari/537.36' '-' 0.040	200
```

服务配置参考：

```
server {
        listen 80;
        server_name xx.xx.xx.xx;
        # 使用 rewrite 指令将所有的 HTTP 请求重定向到对应的 HTTPS 地址，并使用 permanent 永久重定向
        rewrite ^/(.*)$ https://xx.xx.xx.xx/$1 permanent;
    }

server {
        listen 443 ssl;
        server_name xx.xx.xx.xx;

        ssl_certificate     /xx/nginx/certs/server.crt;
        ssl_certificate_key /xx/nginx/certs/server.key;
        ssl_session_cache    shared:SSL:10m;
        ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
        ssl_protocols TLSv1.2;
        ssl_session_timeout  5m;

        client_header_buffer_size 1024k;
        client_body_buffer_size 1024k;
        large_client_header_buffers 4 32k;
        client_max_body_size 200m;
        autoindex off;
        index index.php index.html;
        location  / {
            proxy_pass http://xx.xx.xx.xx;
            proxy_set_header HTTP_X_FORWARDED_FOR "$remote_addr:$remote_port" ;
            proxy_set_header Host $host:$server_port;
            proxy_connect_timeout 60s;
            proxy_read_timeout 300s;
            proxy_send_timeout 60s;
        }
        location ~^/(info*) {
            proxy_pass http://xx.xx.xx.xx:xxxx;
            proxy_set_header Upgrade $http_upgrade;
            proxy_set_header Connection $connection_upgrade;
            proxy_set_header   X-Real-IP        $remote_addr;
            proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
            proxy_next_upstream error timeout invalid_header http_500 http_502 http_503 http_504;
            proxy_max_temp_file_size 0;
        }
     }
```

详细解释：

* `proxy_pass` 指令用于将请求代理到指定的后端服务器。
* `proxy_set_header` 用于设置向后端服务器发送的 HTTP 头信息。
* `proxy_connect_timeout`, `proxy_read_timeout`, `proxy_send_timeout` 分别设置代理连接、读取和发送的超时时间。
* `proxy_next_upstream` 用于定义当代理服务器出现错误时，Nginx 应该尝试连接的下一个服务器。
* `proxy_max_temp_file_size` 设置临时文件的最大大小。
* `location  / `的配置：主要用于将所有请求代理到指定的后端服务器，并设置一些请求头和超时时间。

`location  /`配置：主要用于将所有请求代理到指定的后端服务器，并设置一些请求头和超时时间。
其中，proxy_set_header用于自定义一些请求头信息，而proxy_connect_timeout、
proxy_read_timeout和proxy_send_timeout则用于设置连接、读取和发送的超时时间，以确保在合理的时间内完成这些操作。

```
location  / {
    proxy_pass http://xx.xx.xx.xx;
    proxy_set_header HTTP_X_FORWARDED_FOR "$remote_addr:$remote_port" ;
    proxy_set_header Host $host:$server_port;
    proxy_connect_timeout 60s;
    proxy_read_timeout 300s;
    proxy_send_timeout 60s;
}
```

* `proxy_pass http://xx.xx.xx.xx;`: 当请求匹配上述location块时，Nginx会将请求代理到指定的后端服务器。
  `http://xx.xx.xx.xx`是后端服务器的地址。这是一个标准的反向代理配置，将客户端的请求转发给后端服务器。
* `proxy_set_header HTTP_X_FORWARDED_FOR "$remote_addr:$remote_port";`: 设置请求头中的HTTP_X_FORWARDED_FOR字段，
  将客户端的IP地址和端口号添加到该字段中。这个字段通常用于标识客户端的原始IP地址，尤其是在经过多层代理时。
  在这里，它包含了客户端的IP地址和端口号，以提供更详细的信息。
* `proxy_set_header Host $host:$server_port;`: 设置请求头中的Host字段，将其值设置为`$host:$server_port`。
  这个字段表示客户端原始请求的Host头部值。在反向代理的情况下，确保将正确的Host信息传递给后端服务器是很重要的，以便后端服务器能够正确处理请求。
* `proxy_connect_timeout 60s;`: 设置连接到后端服务器的超时时间为60秒。如果在指定的时间内无法建立连接，则Nginx将认为连接超时。
* `proxy_read_timeout 300s;`: 设置从后端服务器读取响应的超时时间为300秒。如果在指定的时间内没有从后端服务器读取到响应，则Nginx将认为读取超时。
* `proxy_send_timeout 60s;`: 设置向后端服务器发送请求的超时时间为60秒。如果在指定的时间内无法将请求发送到后端服务器，则Nginx将认为发送超时

`location ~^/(info*)`配置：主要用于反向代理WebSocket请求到指定的后端服务器，并且通过设置一些请求头字段，
使得后端服务器能够获取到客户端的真实IP地址。这对于一些需要获取客户端IP的应用场景非常有用，例如在日志中记录客户端的真实IP。

```
location ~^/(info*) {
    proxy_pass http://xx.xx.xx.xx:xxxx;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection $connection_upgrade;
    proxy_set_header   X-Real-IP        $remote_addr;
    proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
    proxy_next_upstream error timeout invalid_header http_500 http_502 http_503 http_504;
    proxy_max_temp_file_size 0;
}
```

* `location ~^/(info*) {`: 这是一个`Nginx`的`location`块，使用正则表达式匹配以"`/info`"开头的URL路径。
  `~^`表示正则表达式要从字符串的开头匹配。`/info*`表示匹配以`"/info"`开头的路径，*表示匹配零个或多个's'字符。
* `proxy_set_header Upgrade $http_upgrade;`: 设置请求头中的Upgrade字段，用于实现WebSocket协议的升级。这对于支持WebSocket的应用程序很重要。
* `proxy_set_header Connection $connection_upgrade;`: 设置请求头中的`Connection`字段，也用于WebSocket协议的升级。
* `proxy_set_header X-Real-IP $remote_addr;`: 将客户端的真实IP地址添加到请求头中的X-Real-IP字段。这在后端服务器需要获取客户端真实IP时非常有用。
* `proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;`: 将客户端的原始IP地址添加到请求头中的X-Forwarded-For字段。这是通过代理时常用的做法，以便后端服务器能够获取到整个请求链上的所有客户端IP地址。
* `proxy_next_upstream error timeout invalid_header http_500 http_502 http_503 http_504;`: 指定在遇到特定错误时Nginx应该尝试连接到下一个后端服务器。具体来说，当发生错误、超时或者后端服务器返回的响应头不合法时，Nginx会尝试连接到下一个后端服务器。
* `proxy_max_temp_file_size 0;`: 禁用Nginx将大于指定大小的响应保存到临时文件的功能。将其设置为0表示禁用这个功能。
