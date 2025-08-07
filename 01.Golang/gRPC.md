
# gRPC 调用流程图
> 从客户端请求到服务端处理再到响应返回的完整过程。
## 普通 gRPC	调用流程

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


## grpc-gateway	

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

说明：
* grpc-gateway 充当“HTTP → gRPC 的中间层”
* 前端无需了解 gRPC，只通过 RESTful API 调用
* 适用于服务既提供给 Web App，也用于微服务间高效通信的场景


## 双向流

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
