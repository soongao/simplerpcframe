# Simple RPC framework learning from scratch
- 学习RPC framework并理解RPC框架的重点

## 架构
- ![流程图](/doc/流程.png)
- ![关系图](/doc/关系.png)

## 项目结构
```text
|   client.go
|   debug.go
|   server.go
|   service.go
|
+---cmd
|   +---httpdebug
|   |       main.go
|   |
|   \---registry
|           main.go
|
+---codec
|       codec.go
|       gob.go
|
+---registry
|       registry.go
|
\---xclient
        discovery.go
        discoveryx.go
        xclient.go
```

## 学习记录
### [学习记录](https://soongao.github.io/posts/rpcframe/)
- 项目参考geerpc
