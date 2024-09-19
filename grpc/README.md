## gRPC

### 1. 概念
gRPC是一种现代化、开源的远程过程调用（Remote Procedure Call，RPC）框架，它由Google开发并基于HTTP/2协议进行通信。gRPC的关键特性包括跨语言支持、基于Protocol Buffers（protobuf）的序列化方式、支持双向流式通信等。

### 2. gRPC基本原理
gRPC允许客户端调用远程服务器上定义的方法，就像调用本地对象一样。它基于一下几个重要的概念：
- 客户端和服务器端：客户端请求调用方法，服务器端实现这些方法。
- Protocol Buffers：gRPC使用 Protocol Buffers(protobuf)作为接口定义语言（IDL）。
- HTTP/2：gRPC使用HTTP/2协议进行底层传输，者提供了低延迟和双向流的支持。


### 3. Protocol Buffers(protobuf)
Protocol Buffers 是一种高效的结构化数据序列化机制， protobuf 使用的是二进制格式，在表示数字和布尔值等基本类型时更加紧凑，序列化和反序列化更加快速。在gRPC中，服务和消息都用protobuf定义，常见文件格式是.proto文件。

### 4. gRPC 通信模式
gRPC 支持四种主要的通信模式：
- Unary RPC: 客户端发送一个请求，服务端返回一个响应。
- Server Streaming RPC: 客户端发送一个请求，服务端返回一系列响应。
- Client Streaming RPC: 客户端发送一系列请求，客户端返回一个响应。
- Bidirectional Stream RPC: 客户端和服务器可以在一个连接中互相流式传输信息。

### 5. 常用的gRPC特性
- 负载均衡：gRPC支持多种负载均衡策略，如客户端负载均衡和服务器端负载均衡。
- 拦截器：可以通过拦截器在RPC调用前后进行一些处理逻辑，比如日志记录、鉴权。
- 超时与重试：gRPC 支持设置调用的超时与重试策略。
- 安全性：gRPC 支持使用 TLS 加密通信。
