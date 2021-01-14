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

* 变更管理
  * 70% 的问题是由变更引起的，恢复可用代码并不总是坏事
* 避免过载
  * 过载保护、流量调度等
* 依赖管理
  * 任何依赖都可能故障，做`chaos monkey testing` 注入故障测试
* 优雅降级
  * 有损服务，避免核心链路依赖故障
* 重试退避
  * 退让算法，冻结时间，API retry detail 控制策略
* 超时控制
  * 进程内+服务间 超时控制
* 极限压测+故障演练
* 扩容+重启+消除有害流量

## Q&A

1. [第十一课问题收集](https://shimo.im/docs/dqrP3ChJGwJtdchg)
2. [第十二课问题收集](https://shimo.im/docs/xJHJXyKRgHyPXDhT)
3. [第十三课问题收集](https://shimo.im/docs/KXyxCdR9t86TDCyR))

## Documents

* [Google 软件测试之道](https://book.douban.com/subject/25742200/)
* [架构整洁之道 Clean Architecture](https://book.douban.com/subject/30333919/)
* [Google API Design Guide 中文版](https://www.bookstack.cn/books/API-design-guide)
* [API 设计指南](https://cloud.google.com/apis/design/)

## References

* [Sentinel: 分布式系统的流量防卫兵](https://github.com/alibaba/Sentinel)
* [http2-explained](https://github.com/bagder/http2-explained)
* [http3-explained](https://github.com/bagder/http3-explained)
* [Proposal: Monotonic Elapsed Time Measurements in Go](https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md)
* [大神讲解微服务治理的技术演进和架构实践](http://www.360doc.com/content/16/1124/21/31263000_609259745.shtml)
* [实施微服务，我们需要哪些基础框架？](http://www.infoq.com/cn/articles/basis-frameworkto-implement-micro-service/)
* [开源Linkerd项目庆祝一周年纪念日](http://www.infoq.com/cn/news/2017/04/linkerd-celebrates-one-year)
* [Rethinking Netflix’s Edge Load Balancing](https://medium.com/netflix-techblog/netflix-edge-load-balancing-695308b5548c)
* [亿级Web系统的容错性建设实践](https://mp.weixin.qq.com/s?__biz=MzAwNjQwNzU2NQ==&mid=402841629&idx=1&sn=f598fec9b370b8a6f2062233b31122e0)
* [阿里微服务之殇及分布式链路追踪技术原理](https://mp.weixin.qq.com/s?__biz=MzIzMzk2NDQyMw==&mid=2247486641&idx=1&sn=1660fb41b0c5b8d8d6eacdfc1b26b6a6)
* [Overload control for scaling WeChat microservices](https://blog.acolyer.org/2018/11/16/overload-control-for-scaling-wechat-microservices/)
* [socc18-final100](https://www.cs.columbia.edu/~ruigu/papers/socc18-final100.pdf)
* [alibaba/Sentinel](https://github.com/alibaba/Sentinel/wiki)
* [分布式系统常见负载均衡算法](https://blog.csdn.net/okiwilldoit/article/details/81738782)
* [“Predictive Load-Balancing: Unfair But Faster & More Robust” by Steve Gury](http://alex-ii.github.io/notes/2019/02/13/predictive_load_balancing.html)
* [深入解析TensorFlow中滑动平均模型与代码实现](https://blog.csdn.net/m0_38106113/article/details/81542863)

---
