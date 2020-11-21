# Lesson 01 Go 架构实践

> * 微服务概览
> * 微服务设计
> * gRPC & 服务发现
> * 多集群 & 多租户

## 微服务概览

## 微服务设计

## gRPC & 服务发现

2个心跳周期内可以
服务发现-客户端发现
服务发现-服务端发现
Servic Mesh 会变得很复杂
## 多集群 & 多租户

N+2 的节点来冗余节点

## Q&A

1. [第一课问题收集](https://shimo.im/docs/x8dxHkQRcdCHX8j3/read)
2. [第二课问题收集](https://shimo.im/docs/WxJp66WCtjVwKDK3)

## Documents

* [Google SRE](Google SRE)
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
* [bilibili/discovery](https://github.com/bilibili/discovery)
* [canal](https://github.com/alibaba/canal)
* [canal_mysql_nosql_sync](https://github.com/liukelin/canal_mysql_nosql_sync)

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

### [CAP 理论]()

* 一致性（Consistency）
* 可用性（Availability）
* 分区容错性（Partition Tolerance）

> CAP 不可能三角说的是对于一个分布式系统而言，一致性（Consistency）、可用性（Availability）、分区容错性（Partition Tolerance）3 个指标不可兼得，只能在 3 个指标中选择 2 个。

* CA 模型，在分布式系统中不存在。因为舍弃 P，意味着舍弃分布式系统，就比如单机版关系型数据库 MySQL，如果 MySQL 要考虑主备或集群部署时，它必须考虑 P。
* CP 模型，采用 CP 模型的分布式系统，舍弃了可用性，一定会读到最新数据，不会读到旧数据。一旦因为消息丢失、延迟过高发生了网络分区，就影响用户的体验和业务的可用性（比如基于 Raft 的强一致性系统，此时可能无法执行读操作和写操作）。典型的应用是 Etcd，Consul 和 Hbase。
* AP 模型，采用 AP 模型的分布式系统，舍弃了一致性，实现了服务的高可用。用户访问系统的时候，都能得到响应数据，不会出现响应错误，但会读到旧数据。典型应用就比如 Cassandra 和 DynamoDB。


---
