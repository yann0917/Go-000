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

## Corontine

进程(Processes)和线程(Threads)

并行(Parallelism)和并发(Concurrency)

## Memory model

happen-before

内存重排

* [golang-ref-mem](https://golang.google.cn/ref/mem) ☆☆☆☆☆

## Package sync

## Package context

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

### TBD

[Goroutine Leaks - The Forgotten Sender](https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html)

https://www.ardanlabs.com/blog/2019/04/concurrency-trap-2-incomplete-work.html

https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html

https://dave.cheney.net/practical-go/presentations/qcon-china.html#_concurrency

https://golang.org/ref/mem

https://blog.csdn.net/caoshangpa/article/details/78853919

https://blog.csdn.net/qcrao/article/details/92759907

https://cch123.github.io/ooo/

https://blog.golang.org/codelab-share

https://dave.cheney.net/2018/01/06/if-aligned-memory-writes-are-atomic-why-do-we-need-the-sync-atomic-package

http://blog.golang.org/race-detector

https://dave.cheney.net/2014/06/27/ice-cream-makers-and-data-races

https://www.ardanlabs.com/blog/2014/06/ice-cream-makers-and-data-races-part-ii.html

https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549

https://medium.com/a-journey-with-go/go-discovery-of-the-trace-package-e5a821743c3c

https://medium.com/a-journey-with-go/go-mutex-and-starvation-3f4f4e75ad50

https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html

https://medium.com/a-journey-with-go/go-buffered-and-unbuffered-channels-29a107c00268

https://medium.com/a-journey-with-go/go-ordering-in-select-statements-fd0ff80fd8d6

https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html

https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html

https://www.ardanlabs.com/blog/2013/10/my-channel-select-bug.html

https://blog.golang.org/io2013-talk-concurrency

https://blog.golang.org/waza-talk

https://blog.golang.org/io2012-videos

https://blog.golang.org/concurrency-timeouts

https://blog.golang.org/pipelines

https://www.ardanlabs.com/blog/2014/02/running-queries-concurrently-against.html

https://blogtitle.github.io/go-advanced-concurrency-patterns-part-3-channels/

https://www.ardanlabs.com/blog/2013/05/thread-pooling-in-go-programming.html

https://www.ardanlabs.com/blog/2013/09/pool-go-routines-to-process-task.html

https://blogtitle.github.io/categories/concurrency/

https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c

https://blog.golang.org/context

https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html

https://golang.org/ref/spec#Channel_types

https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view

https://medium.com/a-journey-with-go/go-context-and-cancellation-by-propagation-7a808bbc889c

https://blog.golang.org/context

https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html

https://golang.org/doc/effective_go.html#concurrency


---
