# Lesson 03 Go 语言实践 Concurrency

> * Corontine
> * Memory model
> * Package sync
> * chan
> * Package context

```txt
goroutine go关键字一定要关注:
1、go 生命周期（结束、终止）
2、go panic
3、把并行扔给调用者

内存模型：
1、搞清楚 原子性、可见性
2、go memory model（了解happen-before）
3、底层的 memory reordering（可以挖一挖 cpu cacline、锁总线、mesi、memory barrier）
```

## Goroutine

进程(Processes)和线程(Threads)

并行(Parallelism)和并发(Concurrency)

并发不是并行。并行是指两个或多个线程同时在不同的处理器执行代码。如果将运行时配置为使用多个逻辑处理器，则调度程序将在这些逻辑处理器之间分配 goroutine，这将导致 goroutine 在不同的操作系统线程上运行。但是，要获得真正的并行性，您需要在具有多个物理处理器的计算机上运行程序。否则，goroutines 将针对单个物理处理器并发运行，即使 Go 运行时使用多个逻辑处理器。

* 空的 `select` 语句在没有 `case` 可以运行的时候，将永远阻塞；
* 如果你的 goroutine 在从另一个 goroutine 获得结果之前无法取得进展，那么通常情况下，你自己去做这项工作比委托它( go func() )更简单。
* 如果函数启动 goroutine，则必须向调用方提供显式停止该goroutine 的方法。通常，将异步执行函数的决定权交给该函数的调用方通常更容易。
* Only use `log.Fatal` from main.main or init functions. log.Fatal 调用了 os.Exit，会无条件终止程序；defers 不会被调用到。

## Memory model

**happen-before** : Go程序中执行内存操作的偏序。如果事件`e1`发生在`e2`前，我们可以说`e2`发生在`e1`后。如果`e1`不发生在`e2`前也不发生在`e2`后，我们就说`e1`和`e2`是并发的。

内存重排(MemoryReordering)

* 在多核心场景下,没有办法轻易地判断两段程序是“等价”的。

* [golang-ref-mem](https://golang.google.cn/ref/mem) ☆☆☆☆☆

## Package sync

* Do not communicate by sharing memory; instead, share memory by communicating.
* data race 是两个或多个 goroutine 访问同一个资源(如变量或数据结构)，并尝试对该资源进行读写而不考虑其他 goroutine。
* Copy-On-Write 思路在微服务降级或者 local cache 场景中经常使用。写时复制指的是，写操作时候复制全量老数据到一个新的对象中，携带上本次新写的数据，之后利用原子替换(atomic.Value)，更新调用者的变量。来完成无锁访问共享数据。

### Mutex

### RWMutex

### atomic

我们把一个复杂的任务，尤其是依赖多个微服务 rpc 需要聚合数据的任务，分解为依赖和并行，依赖的意思为: 需要上游 a 的数据才能访问下游 b 的数据进行组合。但是并行的意思为: 分解为多个小任务并行执行，最终等全部执行完毕。

`sync.Pool` 的场景是用来保存和复用临时对象，以减少内存分配，降低 GC 压力(Request-Driven 特别合适)。

* [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
* [kratos-errgroup](https://github.com/go-kratos/kratos/tree/master/pkg/sync/errgroup)

## Channel

> channels 是一种类型安全的消息队列，充当两个 goroutine 之间的管道，将通过它同步的进行任意资源的交换。chan 控制 goroutines 交互的能力从而创建了 Go 同步机制。当创建的 chan 没有容量时，称为无缓冲通道。反过来，使用容量创建的 chan 称为缓冲通道。

* 无缓冲 chan 没有容量，因此进行任何交换前需要两个 goroutine 同时准备好。当 goroutine 试图将一个资源发送到一个无缓冲的通道并且没有goroutine 等待接收该资源时，该通道将锁住发送 goroutine 并使其等待。当 goroutine 尝试从无缓冲通道接收，并且没有 goroutine 等待发送资源时，该通道将锁住接收 goroutine 并使其等待。

* 无缓冲信道的本质是保证同步。Receive 先于 Send 发生，100% 能保证收到，但是延迟时间未知。

* buffered channel 具有容量，因此其行为可能有点不同。当 goroutine 试图将资源发送到缓冲通道，而该通道已满时，该通道将锁住 goroutine并使其等待缓冲区可用。如果通道中有空间，发送可以立即进行，goroutine 可以继续。当goroutine 试图从缓冲通道接收数据，而缓冲通道为空时，该通道将锁住 goroutine 并使其等待资源被发送。

* 无缓冲通道，Send 先于 Receive 发生，延迟更小，但是不保证数据到达，越大的 buffer，越小的保障到达。buffer = 1 时，给你延迟一个消息的保障。

## Package context

Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named `ctx:`

* Incoming requests to a server should create a Context.
* Outgoing calls to servers should accept a Context.
* Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it.
* The chain of function calls between them must propagate the Context.
* Replace a Context using WithCancel, WithDeadline, WithTimeout, or WithValue.
* When a Context is canceled, all Contexts derived from it are also canceled.
* The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.
* Do not pass a nil Context, even if a function permits it. Pass a TODO context if you are unsure about which Context to use.
* Use context values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
* All blocking/long operations should be cancelable.
* Context.Value obscures your program’s flow.
* Context.Value should inform, not control.
* Try not to use context.Value.

## Q&A

1. [第五课问题收集](https://shimo.im/docs/8yxKjP8r8RjKghPd)
2. [第六课问题收集](https://shimo.im/docs/HYHRdYR6HDCRdvpP)

## Documents

* [Go语言高级编程](https://github.com/chai2010/advanced-go-programming-book)

## References

* [golang-ref-mem](https://golang.google.cn/ref/mem) ☆☆☆☆☆
* [Go的内存模型](https://www.jianshu.com/p/5e44168f47a3) ☆☆☆☆☆
* [golang-notes](https://github.com/cch123/golang-notes) ☆☆☆☆☆
* [Cache一致性和内存模型](https://wudaijun.com/2019/04/cpu-cache-and-memory-model/)
* [Goroutine Leaks - The Forgotten Sender](https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html)
* [Concurrency Trap #2: Incomplete Work](https://www.ardanlabs.com/blog/2019/04/concurrency-trap-2-incomplete-work.html)
* [Concurrency, Goroutines and GOMAXPROCS](https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html)
* [Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/qcon-china.html#_concurrency)
* [理解 Memory Barrier（内存屏障）](https://blog.csdn.net/caoshangpa/article/details/78853919)
* [曹大谈内存重排](https://blog.csdn.net/qcrao/article/details/92759907)
* [从 Memory Reordering 说起](https://cch123.github.io/ooo/)
* [Share Memory By Communicating](https://blog.golang.org/codelab-share)
* [If aligned memory writes are atomic, why do we need the sync/atomic package?](https://dave.cheney.net/2018/01/06/if-aligned-memory-writes-are-atomic-why-do-we-need-the-sync-atomic-package)
* [Introducing the Go Race Detector](http://blog.golang.org/race-detector)
* [Ice cream makers and data races](https://dave.cheney.net/2014/06/27/ice-cream-makers-and-data-races)
* [Ice Cream Makers and Data Races Part II](https://www.ardanlabs.com/blog/2014/06/ice-cream-makers-and-data-races-part-ii.html)
* [Go: How to Reduce Lock Contention with the Atomic Package](https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549)
* [Go: Discovery of the Trace Package](https://medium.com/a-journey-with-go/go-discovery-of-the-trace-package-e5a821743c3c)
* [Go: Mutex and Starvation](https://medium.com/a-journey-with-go/go-mutex-and-starvation-3f4f4e75ad50)
* [The Behavior Of Channels](https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html)
* [Go: Buffered and Unbuffered Channels](https://medium.com/a-journey-with-go/go-buffered-and-unbuffered-channels-29a107c00268)
* [Go: Ordering in Select Statements](https://medium.com/a-journey-with-go/go-ordering-in-select-statements-fd0ff80fd8d6)
* [The Nature Of Channels In Go](https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html)
* [My Channel Select Bug](https://www.ardanlabs.com/blog/2013/10/my-channel-select-bug.html)
* [Advanced Go Concurrency Patterns](https://blog.golang.org/io2013-talk-concurrency)
* [Concurrency is not parallelism](https://blog.golang.org/waza-talk)
* [Go videos from Google I/O 2012](https://blog.golang.org/io2012-videos)
* [Go Concurrency Patterns: Timing out, moving on](https://blog.golang.org/concurrency-timeouts)
* [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)
* [Running MongoDB Queries Concurrently With Go](https://www.ardanlabs.com/blog/2014/02/running-queries-concurrently-against.html)
* [Thread Pooling in Go Programming](https://www.ardanlabs.com/blog/2013/05/thread-pooling-in-go-programming.html)
* [Pool Go Routines To Process Task Oriented Wo](https://www.ardanlabs.com/blog/2013/09/pool-go-routines-to-process-task.html)
* [Go advanced concurrency patterns](https://blogtitle.github.io/categories/concurrency/)
* [Go: Context and Cancellation by Propagation](https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c)
* [Go Concurrency Patterns: Context](https://blog.golang.org/context)
* [Context Package Semantics In Go](https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html)
* [Channel types](https://golang.org/ref/spec#Channel_types)
* [Rethinking Classical Concurrency Patterns](https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view)
* [Go: Context and Cancellation by Propagation](https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c)
* [effective go concurrency](https://golang.org/doc/effective_go.html#concurrency)
* [Go Context的踩坑经历](https://zhuanlan.zhihu.com/p/34417106)
* [Cancelation, Context, and Plumbing](https://talks.golang.org/2014/gotham-context.slide#1)
* [How to correctly use context.Context in Go 1.7](https://medium.com/@cep21/how-to-correctly-use-context-context-in-go-1-7-8f2c0fafdf39)

## 作业

> 1.基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。

[作业代码](homework/main.go)

---
