# Simple RPC framework learning from scratch
- 学习RPC framework并理解RPC框架的重点
- 根据net/rpc实现了服务端以及并发客户端, 拥有超时处理机制
- 支持TCP、HTTP等多种传输协议; 支持多种负载均衡模式, 并且实现了一个简易的服务注册和发现中心

## 技术点
- 自定义消息编码 [`Option`|`Header`|`Body`]
- 支持异步的并发客户端实现
  - 同步发送call请求后, 将call加入map, 等待返回结果
  - 异步接收服务端消息, 收到返回后, 通过call中的done chan通知客户端调用结束
- 超时处理机制
  - 使用time.After进行建立连接超时
  - 使用context.WithTimeout, 将控制权交给用户, 进行请求处理的超时
- 除tcp长连接外, 支持http传输协议
  - http进行tcp握手后Hijack这个tcp进行自己定义的消息格式传输
  - 这样做的目的支持添加不同路径提供不同服务, 支持一个address可以根据path(URL)进行不同的handle
- 实现了多种负载均衡算法
  - random
  - 轮转法RR
  - 一致性哈希
- 实现了一个简易的服务注册和发现中心

## 架构
- ![流程图](/doc/process.png)
- ![关系图](/doc/relation.png)

## 项目结构
```text
|   client.go // 并发客户端
|   server.go // 服务器
|   service.go // rpc服务抽象
|
+---codec
|       codec.go // 编解码器抽象
|       gob.go // gob编解码器实现
|
+---registry
|       registry.go // 心跳保活注册中心
|
\---xclient
        discovery.go // 服务列表由手工维护的服务发现
        discoveryx.go // 基于注册中心的服务发现
        xclient.go // 负载均衡客户端
```

## 学习记录
### [学习记录](https://soongao.github.io/posts/rpcframe/)
- 项目参考geerpc
