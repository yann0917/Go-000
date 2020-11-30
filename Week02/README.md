# Lesson 02 Go 语言实践 error

> * Error vs Exception
> * Error Type
> * Go 1.13 errore
> * Go 2 Error Inspection

## Error vs Exception

## Error Type

## Go 1.13 errore

## Go 2 Error Inspection

## Q&A

1. [第三课问题收集](https://shimo.im/docs/vr6yDVPxRxXGKRDd)
2. [第四课问题收集](https://shimo.im/docs/R6gP8qyvWqJrgRCk)

## Documents

* [Effective Go](https://golang.org/doc/effective_go.html) ☆☆☆☆☆

## References

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
