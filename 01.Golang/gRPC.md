
# gRPC 调用流程图
> 从客户端请求到服务端处理再到响应返回的完整过程。


## 方式一：普通 gRPC	调用流程

```mermaid
sequenceDiagram
    participant Client
    participant Stub as Client Stub (由 protoc-gen-go-grpc 生成)
    participant Network as HTTP/2 通信
    participant ServerStub as Server Stub (由 protoc-gen-go-grpc 生成)
    participant ServiceImpl as FundServiceImpl (你自己实现)
    
    Client->>Stub: 调用 client.GetAccount(req)
    Stub->>Network: 序列化 request (protobuf)
    Network->>ServerStub: 通过 HTTP/2 发送数据流
    ServerStub->>ServiceImpl: 反序列化为 req, 调用 GetAccount(ctx, req)
    ServiceImpl-->>ServerStub: 返回 response 数据
    ServerStub-->>Network: 序列化 response (protobuf)
    Network-->>Stub: 返回数据流
    Stub-->>Client: 反序列化为 response 结构体
```

![grpc_01.png](grpc/grpc_01.png)


## 方式二：grpc-gateway	调用

```mermaid
sequenceDiagram
    participant RESTClient as HTTP Client（浏览器/前端）
    participant Gateway as grpc-gateway（反向代理层）
    participant ClientStub as gRPC 客户端 Stub
    participant Network as HTTP/2
    participant ServerStub as gRPC 服务端 Stub
    participant ServiceImpl as FundServiceImpl

    RESTClient->>Gateway: 发送 HTTP 请求（如 GET /v1/account?id=123）
    Gateway->>ClientStub: 转换为 gRPC 请求对象（使用 Protobuf）
    ClientStub->>Network: 发起 gRPC 请求（HTTP/2 + Protobuf）
    Network->>ServerStub: 到达 gRPC Server
    ServerStub->>ServiceImpl: 调用 GetAccount(ctx, req)
    ServiceImpl-->>ServerStub: 返回业务响应
    ServerStub-->>Network: 响应序列化
    Network-->>ClientStub: 返回 gRPC 响应
    ClientStub-->>Gateway: 反序列化为 JSON
    Gateway-->>RESTClient: 返回 HTTP JSON 响应
```

![grpc_02.png](grpc/grpc_02.png)

```mermaid
sequenceDiagram
    participant Client as 客户端（浏览器/curl/前端）
    participant Gateway as grpc-gateway（HTTP 反向代理）
    participant JSONParser as JSON 解析器
    participant ProtoBinder as Protobuf 构造器
    participant GRPCStub as gRPC 客户端 Stub
    participant ServerStub as gRPC 服务端 Stub
    participant Business as 业务逻辑实现（Go/Java/...）

    Client->>Gateway: 发起 HTTP 请求（POST /v1/account）
    Gateway->>JSONParser: 解析 JSON Body
    JSONParser-->>Gateway: 得到结构化数据（map[string]interface{}）

    Gateway->>ProtoBinder: 将 JSON 映射到 Protobuf 消息结构
    ProtoBinder->>ProtoBinder: 字段名校验 + 类型转换
    ProtoBinder->>ProtoBinder: 检查 required/optional/oneof
    ProtoBinder-->>Gateway: 得到 Protobuf 请求对象（如 *GetAccountRequest）

    Gateway->>GRPCStub: 调用 gRPC 方法（stub 发送请求）
    GRPCStub->>ServerStub: 使用 HTTP/2 发送 Protobuf 数据
    ServerStub->>Business: 反序列化请求，调用 GetAccount(ctx, req)
    Business-->>ServerStub: 返回业务处理结果
    ServerStub-->>GRPCStub: 序列化 Protobuf 响应
    GRPCStub-->>Gateway: 返回 gRPC 响应
    Gateway->>Gateway: Protobuf 响应 转换为 JSON
    Gateway-->>Client: 返回 HTTP JSON 响应
```


![grpc_gateway_01.png](grpc/grpc_gateway_01.png)

说明：
* grpc-gateway 充当“HTTP → gRPC 的中间层”
* 前端无需了解 gRPC，只通过 RESTful API 调用
* 适用于服务既提供给 Web App，也用于微服务间高效通信的场景


## 方式三：双向流

```mermaid
sequenceDiagram
    participant Client
    participant Stub as Client Stub
    participant Network as HTTP/2 Stream
    participant ServerStub
    participant ServiceImpl as ChatServiceImpl

    Client->>Stub: 调用 stream := client.ChatStream()
    Stub->>Network: 建立 gRPC Stream（双向）

    loop 多轮交互
        Client->>Stub: stream.Send(req1)
        Stub->>Network: 发送 req1
        Network->>ServerStub: 接收 req1
        ServerStub->>ServiceImpl: 调用 stream.Recv()
        ServiceImpl-->>ServerStub: stream.Send(resp1)
        ServerStub-->>Network: 发送 resp1
        Network-->>Stub: 接收 resp1
        Stub-->>Client: 返回 resp1
    end

    Client->>Stub: stream.CloseSend()
    Stub-->>Network: 关闭发送通道
    Network-->>ServerStub: 通知流关闭
```

![grpc_03.png](grpc/grpc_03.png)

应用场景：
* 聊天系统
* 实时语音识别（客户端不断上传音频，服务端实时返回文本）
* 股票行情推送

总结对比：

| 类型           | 客户端          | 服务端         | 特点                           |
| ------------ | ------------ | ----------- | ---------------------------- |
| 普通 gRPC      | 单请求单响应       | 单请求单响应      | 类似 REST 但用 Protobuf 和 HTTP/2 |
| grpc-gateway | HTTP JSON 请求 | 自动转 gRPC 调用 | 前端友好，可双协议共存                  |
| 双向流          | 多请求多响应       | 持续收发数据流     | 实时、复杂交互                      |


## 方式四：Lua grpcaller 动态调用

```mermaid
sequenceDiagram
    participant Go调用者
    participant LVM
    participant Lua执行器
    participant grpcaller模块
    participant gRPC客户端
    participant gRPC服务端

    Go调用者->>LVM: Call(ctx, "scripts/grpc/grpc.lua", "call", params)
    LVM->>Lua执行器: 执行grpc.lua中的 call 函数
    Lua执行器->>grpcaller模块: require("grpcaller")
    Note right of grpcaller模块: 加载 Lua C 扩展 / 动态库

    Lua执行器->>grpcaller模块: grpcaller.open(params.addr)
    grpcaller模块->>gRPC客户端: 创建 gRPC channel（连接池、负载均衡）
    grpcaller模块-->>Lua执行器: 返回 client 实例

    Lua执行器->>gRPC客户端: client.call(method, metadata, params, timeout)
    gRPC客户端->>gRPC服务端: 发起 gRPC 请求（携带 metadata 和参数）
    gRPC服务端-->>gRPC客户端: 返回响应数据
    gRPC客户端-->>Lua执行器: 返回 rep 数据

    Lua执行器-->>LVM: rep
    LVM-->>Go调用者: 返回结果
```

```mermaid
sequenceDiagram
    participant Lua执行器
    participant grpcaller接口
    participant Method解析器
    participant ProtoDesc加载器
    participant 参数编码器
    participant Metadata构造器
    participant gRPC Stub
    participant gRPC网络IO
    participant gRPC服务端

    Lua执行器->>grpcaller接口: client.call(method, metadata, params, timeout)

    grpcaller接口->>Method解析器: 拆解 "Service/Method" 名称
    Method解析器-->>grpcaller接口: 返回 service_name, method_name

    grpcaller接口->>ProtoDesc加载器: 加载 proto 描述（本地 JSON or 反射）
    alt 使用预编译 Proto JSON
        ProtoDesc加载器-->>grpcaller接口: 从本地缓存加载 proto desc
    else 服务端支持反射
        ProtoDesc加载器->>gRPC服务端: 使用 grpc reflection 获取 method schema
        gRPC服务端-->>ProtoDesc加载器: 返回 method descriptor
    end

    grpcaller接口->>参数编码器: 使用 method descriptor，将 Lua table 编码为 protobuf
    参数编码器-->>grpcaller接口: 返回请求二进制数据

    grpcaller接口->>Metadata构造器: 转换 metadata 到 gRPC headers
    Metadata构造器-->>grpcaller接口: 返回 headers

    grpcaller接口->>gRPC Stub: 发起请求 (headers + request + timeout)

    gRPC Stub->>gRPC网络IO: 发送数据 (HTTP/2)
    gRPC网络IO->>gRPC服务端: 请求 protobuf
    gRPC服务端-->>gRPC网络IO: 返回响应
    gRPC网络IO-->>gRPC Stub: 响应数据帧
    gRPC Stub-->>grpcaller接口: 返回 protobuf 响应数据

    grpcaller接口->>ProtoDesc加载器: 获取响应 message 描述
    grpcaller接口->>参数编码器: 解码 protobuf 响应为 Lua table
    参数编码器-->>grpcaller接口: 返回 Lua 格式响应

    grpcaller接口-->>Lua执行器: 返回 rep, nil
```

调用代码：
```go
func (this *Lvm) GrpcCall(ctx context.Context, addr string, method string, params map[string]interface{}, timeout int) (interface{}, error) {
	m := map[string]interface{}{
		"addr":   addr,
		"method": method,
		"metadata": map[string]interface{}{
			"trace_id":   global.GetTraceId(ctx),
			"user_id":    xxxx,
			"User-Agent": "mkt-xxx",
		},
		"params":  params,
		"timeout": timeout,
	}
	reply, err := this.Call(ctx, "scripts/grpc/grpc.lua", "call", m)
	if err != nil {
		return nil, err
	}
	if reply.Code != 0 {
		return nil, fmt.Errorf("[%d][%s]%s", reply.Code, reply.Msg, reply.Reason)
	}
	return reply.Data, nil
}
````

```lua
local grpcaller = require("grpcaller")

function call(params)
    local client,err = grpcaller.open(params.addr)
    if(err~=nil) then
        return nil,err
    end
    local rep,err = client.call(params.method,params.metadata,params.params,params.timeout)
    if(err~=nil) then
        return nil,err
    end
    return rep,nil
end
```

Go 通过 Lua 执行器调用 gRPC:

![grpc_lua_01.png](grpc/grpc_lua_01.png)

Lua 执行器通过 grpcaller 调用 gPRC:

![grpc_lua_02.png](grpc/grpc_lua_02.png)

概览对比表：四种调用方式的核心区别

| 特性 / 模式             | Lua 脚本动态调用       | Go 强类型调用            | gRPC-Gateway（REST 转 gRPC） | gRPC 双向流            |
| ------------------- | ---------------- | ------------------- | ------------------------- | ------------------- |
| **语言支持**            | Lua（通过 grpcaller 封装） | Go（静态语言）            | 前端 HTTP（任意语言）             | 主要 Go、Java、Python 等 |
| **类型安全**            | ❌ 弱类型（动态传 map）   | ✅ 强类型（编译时检查）        | ❌ 弱类型（JSON → Protobuf）    | ✅ 强类型               |
| **是否依赖 proto 编译**   | ⚠️ 可选，用反射或 ProtoDesc | ✅ 必须预编译 .proto      | ✅ 必须预编译 .proto + 注解       | ✅ 必须预编译             |
| **是否用 Protobuf 协议** | ✅ 是              | ✅ 是                 | ✅ 后端是                     | ✅ 是                 |
| **使用协议**            | HTTP/2 + Protobuf | HTTP/2 + Protobuf   | HTTP/1.1 or HTTP/2 + JSON | HTTP/2 + Protobuf   |
| **接口调用方式**          | 反射式、方法名 + 参数 map 调用 | Stub 接口直接调用         | REST API（POST/GET）        | 数据流持续读写             |
| **接入方便性**           | ✅ 灵活嵌入脚本、动态性强    | ⚠️ 需要 proto 编译、静态绑定 | ✅ 对前端极友好                  | ❌ 不适合浏览器            |
| **适用场景**            | 动态中转、脚本调度、热更     | 内部服务强约束通信           | 对外提供 REST API             | 流式传输，如聊天/实时日志       |
| **性能**              | 中等               | 高性能                 | 中等（JSON 序列化开销）            | 高（流式可复用连接）          |
| **错误处理能力**          | 手动处理错误码、反射结果     | gRPC Code + Typed Error | HTTP Code + JSON 错误信息     | gRPC Code，错误流通知     |
| **连接复用/连接池支持**      | 视 `grpcaller` 实现而定 | ✅ 支持连接复用            | 网关内部转发连接                  | ✅ 支持长连接             |




# gRPC 使用

## 工具安装
```
.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
```

## 编译命令
```
 protoc \
  -I . \
  -I ./third_party \
  --go_out . --go_opt paths=source_relative \
  --go-grpc_out . --go-grpc_opt paths=source_relative \
  --grpc-gateway_out . --grpc-gateway_opt paths=source_relative \
  ./api/fund/fund.proto
```