# Go 架构实践 -- 微服务可用性设计

## 隔离

隔离，本质上是对**系统或资源进行分割**，从而实现当系统发生故障时能限定传播范围和影响范围，即发生故障后只有出问题的服务不可用，保证其他服务仍然可用。

### 服务隔离

* 动静分离
* 读写分离

### 轻重隔离

* 核心
* 快慢
* 热点

### 物理隔离

* 线程
* 进程
* 进群
* 机房

## 超时控制

## 过载保护

## 限流

## 降级

## 重试

## 负载均衡

## 最佳实践

## Q&A

1. [第十一课问题收集](https://shimo.im/docs/dqrP3ChJGwJtdchg)
2. [第十二课问题收集](https://shimo.im/docs/xJHJXyKRgHyPXDhT)

## Documents

* [Google 软件测试之道](https://book.douban.com/subject/25742200/)
* [架构整洁之道 Clean Architecture](https://book.douban.com/subject/30333919/)
* [Google API Design Guide 中文版](https://www.bookstack.cn/books/API-design-guide)
* [API 设计指南](https://cloud.google.com/apis/design/)

## References

* [http2-explained](https://github.com/bagder/http2-explained)
* [http3-explained](https://github.com/bagder/http3-explained)
* [Proposal: Monotonic Elapsed Time Measurements in Go](https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md)

---
