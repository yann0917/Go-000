# Lesson 01 Go 架构实践

> * 微服务概览
> * 微服务设计
> * gRPC & 服务发现
> * 多集群 & 多租户

## 微服务概览

## 微服务设计

## gRPC & 服务发现

## 多集群 & 多租户


## Documents

* [HTTP2]()
* [K8S中文社区](https://www.kubernetes.org.cn/)
* [Kubernetes教程](https://www.kuboard.cn/learning/)
* [Prometheus](https://github.com/prometheus/prometheus)
* [grafana](https://github.com/grafana/grafana)
* [gRPC-go](https://github.com/grpc/grpc-go)
* [Golang gRPC实践 连载](https://segmentfault.com/a/1190000007880647)
* [Consul](https://github.com/hashicorp/consul)
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

---
