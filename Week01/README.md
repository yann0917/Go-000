# Lesson 01 Go 架构实践

> * 微服务概览
> * 微服务设计
> * gRPC & 服务发现
> * 多集群 & 多租户

## 微服务概览

| 架构       | 定义                                                                                                                                                                                                                                                              | 优点                                           | 缺点                       |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------|----------------------------|
| 单体架构   | 数据库，服务器，前端表示层部署在同一个应用中                                                                                                                                                                                                                      | 开发、测试、部署高效简单                       | 扩展困难、可靠性低         |
| 微服务架构 | 围绕业务功能构建的，服务关注单一业务，服务间 采用轻量级的通信机制，可以全自动独立部署，可 以使用不同的编程语言和数据存储技术。<br/> 微服务架 构通过业务拆分实现服务组件化，通过组件组合快 速开发系统，业务单一的服务组件又可以独立部署， 使得整个系统变得清晰灵活 | 原子服务、独立进程、隔离部署、去中心化服务治理 | 运维基础设施建设、复杂度高 |

* 组件服务化
  * kit:一个微服务基础库 kit
  * service: 业务逻辑+kit 依赖+第三方依赖组成的微服务
  * rpc+mq: 微服务间轻量级通讯
* 按业务组织服务『[康威定律](https://zh.wikipedia.org/wiki/%E5%BA%B7%E5%A8%81%E5%AE%9A%E5%BE%8B)』
  * 开发团队对软件在生产环境的运行负全部责任！
  * 开发团队对软件在生产环境的运行负全部责任！
  * 开发团队对软件在生产环境的运行负全部责任！
* 去中心化
  * 数据去中心化（每个服务独享数据存储设施，利于服务独立性）
  * 治理去中心化
  * 技术去中心化
* 基础设施自动化
  * CICD: gitlab+Gitlab Hook +k8s
  * Testing: 测试环境部署，单元测试、API 自动化测试 （推荐[Yapi](https://github.com/YMFE/yapi)）
  * 在线运行时: [k8s](https://www.kubernetes.org.cn/), [Prometheus](https://github.com/prometheus/prometheus)
* 可用性&兼容性设计『[伯斯塔尔法则](https://en.wikipedia.org/wiki/Robustness_principl)』
  * 隔离
  * 超时控制
  * 负载保护
  * 限流
  * 熔断
  * 降级
  * 重试
  * 负载均衡

## 微服务设计

### API-Gateway
### 微服务划分

### 微服务安全

## gRPC & 服务发现

> 不要过早关注性能问题，先标准化

### gRPC 简介

* [gRPC 中文官方文档](http://doc.oschina.net/grpc)
* [gRPC 健康检查协议](https://github.com/grpc/grpc/blob/v1.15.0/doc/health-checking.md)

### 服务发现

服务发现-客户端发现

服务发现-服务端发现

2个心跳周期内可以

Servic Mesh 会变得很复杂
## 多集群 & 多租户

N+2 的节点来冗余节点

## Q&A

1. [第一课问题收集](https://shimo.im/docs/x8dxHkQRcdCHX8j3/read)
2. [第二课问题收集](https://shimo.im/docs/WxJp66WCtjVwKDK3)

## Documents

* [SRE：Google运维解密](https://item.jd.com/11973579.html) ☆☆☆☆☆
* [UNIX环境高级编程第3版](https://item.jd.com/12720738.html) ☆☆☆☆☆
* [HTTP2]()
* [K8S中文社区](https://www.kubernetes.org.cn/)
* [Kubernetes教程](https://www.kuboard.cn/learning/)
* [Prometheus](https://github.com/prometheus/prometheus)
* [grafana](https://github.com/grafana/grafana)
* [gRPC 中文官方文档](http://doc.oschina.net/grpc)
* [gRPC-go](https://github.com/grpc/grpc-go)
* [Golang gRPC实践 连载](https://segmentfault.com/a/1190000007880647)
* [如何通过gRPC实现高效远程过程调用？](https://time.geekbang.org/column/article/247812)
* [跟煎鱼学 Go](https://eddycjy.com/go-categories/)
* [Consul](https://github.com/hashicorp/consul)
* [Eureka](https://github.com/Netflix/eureka/wiki)
* [nacos](https://nacos.io/zh-cn/docs/what-is-nacos.html)
* [bilibili/discovery](https://github.com/bilibili/discovery)
* [canal](https://github.com/alibaba/canal)
* [canal_mysql_nosql_sync](https://github.com/liukelin/canal_mysql_nosql_sync)
* [四层和七层负载均衡的区别](https://kb.cnblogs.com/page/188170/)

## References

### [康威定律(Conway’s law)](https://zh.wikipedia.org/wiki/%E5%BA%B7%E5%A8%81%E5%AE%9A%E5%BE%8B)

"设计系统的架构受制于产生这些设计的组织的沟通结构。"

> * 第一定律 组织沟通方式会通过系统设计表达出来
> * 第二定律 时间再多一件事情也不可能做的完美，但总有时间做完一件事情
> * 第三定律 线型系统和线型组织架构间有潜在的异质同态特性
> * 第四定律 大的系统组织总是比小系统更倾向于分解

### [伯斯塔尔法则(Postel’s Law)](https://en.wikipedia.org/wiki/Robustness_principle)

Be conservative in what you send, be liberal in what you accept。
“对发送的内容保持谨慎，对接收的内容保持自由”

> * 发送的数据要更保守， 意味着最小化的传送必要的信息
> * 接收时更开放意味着要最大限度的容忍冗余数据，保证兼容性

### [CAP 理论](http://www.ruanyifeng.com/blog/2018/07/cap.html)

* 一致性（Consistency）
* 可用性（Availability）
* 分区容错性（Partition Tolerance）

> CAP 不可能三角说的是对于一个分布式系统而言，一致性（Consistency）、可用性（Availability）、分区容错性（Partition Tolerance）3 个指标不可兼得，只能在 3 个指标中选择 2 个。

* CA 模型，在分布式系统中不存在。因为舍弃 P，意味着舍弃分布式系统，就比如单机版关系型数据库 MySQL，如果 MySQL 要考虑主备或集群部署时，它必须考虑 P。
* CP 模型，采用 CP 模型的分布式系统，舍弃了可用性，一定会读到最新数据，不会读到旧数据。一旦因为消息丢失、延迟过高发生了网络分区，就影响用户的体验和业务的可用性（比如基于 Raft 的强一致性系统，此时可能无法执行读操作和写操作）。典型的应用是 Etcd，Consul 和 Hbase。
* AP 模型，采用 AP 模型的分布式系统，舍弃了一致性，实现了服务的高可用。用户访问系统的时候，都能得到响应数据，不会出现响应错误，但会读到旧数据。典型应用就比如 Cassandra 和 DynamoDB。

* [CAP理论：怎样舍弃一致性去换取性能？](https://time.geekbang.org/column/article/251062)
* [CAP理论：分布式系统的PH试纸，用它来测酸碱度](https://time.geekbang.org/column/article/195675)
* [ACID理论：CAP的酸，追求一致性](https://time.geekbang.org/column/article/199663)

---
