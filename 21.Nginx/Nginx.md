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