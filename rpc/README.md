## RPC(Remote Procedure Call)
### 1. RPC概念
- **RPC定义**：RPC允许一台计算机（客户端）调用另一台计算机（服务器）上的程序或函数，就像调用本地函数一样。
- **核心思想**：隐藏网络通信的复杂性，使得远程调用和本地调用差不多。
- **工作原理**：

  1. 客户端调用本地的存**根（stub）函数；
  2. 存根函数负责打包（序列化）参数并通过网络发送给远程服务器；
  3. 服务器接收请求，解包参数，执行目标函数，并返回结果；
  4. 客户端收到结果后解包并返回给调**用者。

### 2. RPC的组成部分
- **客户端和服务器端**：客户端发送请求，服务器处理请求并返回响应。
- **存根（stub）**：客户端和服务器端的代理函数。客户端的存根负责发送请求，服务器存根负责处理请求。
- **编解码**：参数的序列化与反序列化过程，使得数据能够通过网络传输。

### 3. RPC框架
不同语言中有很多成熟的RPC框架可以使用，它们提供了简化的开发流程，自动生成代码、处理网络细节等。常见的RPC:
  - **gRPC**（Google的RPC框架）：基于HTTP/2和Protocol Buffers，非常适合高性能、跨语言服务调用。
  - **Thrift**：Apache基金会的跨语言RPC框架，支持多种协议和传输层。
  - **Dubbo**：阿里巴巴开源的分布式服务框架，广泛应用微服务框架。