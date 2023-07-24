## 安装
```
$ go get github.com/davveo/lemonShop
```

## 如何运行

### 必须

- Mysql
- Redis

### 准备

创建一个 `lemon` 数据库，并且导入建表的 [SQL](https://github.com/davveo/lemonShop/blob/master/docs/sql/lemonShop.sql)

### 配置

你应该修改 `conf/app.ini` 配置文件

```
[database]
Type = mysql
User = root
Password = rootroot
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```


### 运行
```
$ go install github.com/pilu/fresh

# 直接使用命令启动
$ active=dev go run main.go # conf-dev配置前置

# 或者使用fresh启动
$ fresh -c fresh.conf
```

项目的运行信息和已存在的 API's

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /auth                     --> github.com/davveo/lemonShop/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/davveo/lemonShop/vendor/github.com/swaggo/gin-swagger.WrapHandler.func1 (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/davveo/lemonShop/routers/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/davveo/lemonShop/routers/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/davveo/lemonShop/routers/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/davveo/lemonShop/routers/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/davveo/lemonShop/routers/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> github.com/davveo/lemonShop/routers/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/davveo/lemonShop/routers/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/davveo/lemonShop/routers/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/davveo/lemonShop/routers/api/v1.DeleteArticle (4 handlers)

Listening port is 8000
Actual pid is 4393
```

## 参考
- https://github.com/xxjwxc/gormt/blob/master/data/view/genfunc/genfunc_test.go


## 特性

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis

