# `/docs`
# 安装框架

1. `brew install go` 安装 `go`
2. `go env` 查看环境，如果需要，修改`GOPATH`
3. `brew install glide` 安装 `glide` 来管理依赖

## ~~使用 swag 生成 API 文档~~

> [swag-github](https://github.com/swaggo/swag)

1. 使用 `go get -u github.com/swaggo/swag/cmd/swag` 安装 `swag` 命令，将`$HOME/go/bin` 加入 PATH
1. 在 `controllers` 的方法上添加符合 `swagger` 要求的注释，详见 [swagger api 操作文档](https://swaggo.github.io/swaggo.io/declarative_comments_format/api_operation.html)
2. 在项目根目录里执行 `swag init -g cmd/main.go` 更新 `swagger.json`  
3. 访问 `http://127.0.0.1:8001/swagger/index.html` 打开 `swagger web ui` 可以看到由注释生成的 API ，执行后可查看响应数据


## glide 使用方法

> [glide-github](https://github.com/Masterminds/glide)

```shell
$ glide create                            # Start a new workspace
$ open glide.yaml                         # and edit away!
$ glide get github.com/Masterminds/cookoo # Get a package and add to glide.yaml
$ glide install                           # Install packages and dependencies
# work, work, work
$ go build                                # Go tools work normally
$ glide up                                # Update to newest versions of the package
$ glide rm                                # remove packages and dependencies
```

## testify 测试框架使用方法

[testify-github](https://github.com/stretchr/testify)

## go-sqlmock 测试框架使用方法

[sqlmock-github](https://github.com/DATA-DOG/go-sqlmock)

## 备注

* [glide](https://github.com/Masterminds/glide)
* [gin](https://github.com/gin-gonic/gin)
* [gorm](https://github.com/jinzhu/gorm)
* [go-redis](https://github.com/go-redis/redis)
* [gin 中文文档](https://learnku.com/docs/gin-gonic/2018)
* [GORM Guides](http://gorm.io/docs/)
* [GORM 中文文档](http://gorm.io/zh_CN/)
* [GORM 中文文档 gitbook](https://jasperxu.github.io/gorm-zh/)
* [go-redis DOC](https://godoc.org/github.com/go-redis/redis)
* [gin+gorm+router 快速搭建 crud restful API 接口](https://learnku.com/articles/23548/gingormrouter-quickly-build-crud-restful-api-interface)
* [教程：使用 go 的 gin 和 gorm 框架来构建 RESTful API 微服务](https://learnku.com/golang/t/24598)

---