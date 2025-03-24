# Simple RPC framework from scratch
- 学习RPC framework并理解RPC框架的重点

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

## 学习笔记
### [学习记录](https://soongao.github.io/posts/rpcframe/)
- 项目参考[geektutu](https://geektutu.com/post/geerpc.html)
