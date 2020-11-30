# Lesson 02 Go 语言实践 error

> * Error vs Exception
> * Error Type
> * Go 1.13 errore
> * Go 2 Error Inspection

## Error vs Exception

> 对于真正意外的情况，那些表示不可恢复的程序错误，例如索引越界、不可恢复的环境问题、栈溢出，我们才使用 panic。
> 对于其他的错误情况，我们应该是期望使用 error 来进行判定。

## Error Type

* 预定义的特定错误，我们叫为 `sentinel error`，这个名字来源于计算机编程中使用一个特定值来表示不可能进行进一步处理的做法。
* `Error type` 是实现了 error 接口的自定义类型。尽量避免使用 error types
* `Opaque errors` 不同吗错误处理——只需返回错误，而不假设其内容。

## Handling Error

* you should only handle errors once. Handling an error means inspecting the error value, and making a single decision.
* 日志记录与错误无关且对调试没有帮助的信息应被视为噪音，应予以质疑。记录的原因是因为某些东西失败了，而日志包含了答案。
* 在你的应用代码中，使用 errors.New 或者  errros.Errorf 返回错误。
* 如果调用其他的函数，通常简单的直接返回。
* 如果和其他库进行协作，考虑使用 errors.Wrap 或者 errors.Wrapf 保存堆栈信息。
* 直接返回错误，而不是每个错误产生的地方到处打日志。
* 在程序的顶部或者是工作的 goroutine 顶部(请求入口)，使用 `%+v` 把堆栈详情记录。
* 使用 errors.Cause 获取 root error，再进行和 sentinel error 判定。
* 选择 wrap error 是只有 applications 可以选择应用的策略。
* 如果函数/方法不打算处理错误，那么用足够的上下文 `wrap errors` 并将其返回到调用堆栈中。
* 一旦确定函数/方法将处理错误，错误就不再是错误。如果函数/方法仍然需要发出返回，则它不能返回错误值。

## Go 1.13 errore

* `errors.Is()`
* `errors.As()`
* `errors.Unwrap()`

## Go 2 Error Inspection

## Q&A

1. [第三课问题收集](https://shimo.im/docs/vr6yDVPxRxXGKRDd)
2. [第四课问题收集](https://shimo.im/docs/R6gP8qyvWqJrgRCk)

## Documents

* [Effective Go](https://golang.org/doc/effective_go.html) ☆☆☆☆☆

## References

* [errors](github.com/pkg/errors) Simple error handling primitives ☆☆☆☆☆
* [Why Go gets exceptions right](https://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right)
* [Errors and Exceptions, redux](https://dave.cheney.net/2015/01/26/errors-and-exceptions-redux)
* [Error handling vs. exceptions redux](https://dave.cheney.net/2014/11/04/error-handling-vs-exceptions-redux)
* [Why Go's Error Handling is Awesome](https://rauljordan.com/2020/07/06/why-go-error-handling-is-awesome.html)
* [Effective error handling in Go](https://morsmachine.dk/error-handling)
* [Error handling and Go](https://blog.golang.org/error-handling-and-go)
* [Error Handling In Go, Part I](https://www.ardanlabs.com/blog/2014/10/error-handling-in-go-part-i.html)
* [Error Handling In Go, Part II](https://www.ardanlabs.com/blog/2014/11/error-handling-in-go-part-ii.html)
* [Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)
* [Error handling in Upspin](https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html)
* [Errors are values](https://blog.golang.org/errors-are-values)
* [Stack traces and the errors package](https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package)
* [Design Philosophy On Logging](https://www.ardanlabs.com/blog/2017/05/design-philosophy-on-logging.html)
* [Go 1.13: xerrors](https://crawshaw.io/blog/xerrors)
* [Working with Errors in Go 1.13](https://blog.golang.org/go1.13-errors)
* [Error Handling in Go](https://medium.com/gett-engineering/error-handling-in-go-53b8a7112d04)
* [Error Handling in Go 1.13](https://medium.com/gett-engineering/error-handling-in-go-1-13-5ee6d1e0a55c)

## 作业

Q1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 `sql.ErrNoRows` 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

A1. 需要 wrap err, 抛给上层，[示例](main.go) 展示时，sql 如下，未工程化处理，配置文件硬编码在代码中，修改 `dao/mysql.go` 的配置信息即可，访问路由 `localhost:8000/user/1`

```sql
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

---
