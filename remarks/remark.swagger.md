# 集成 swagger

[gin-swagger](https://github.com/swaggo/gin-swagger)

## 安装 swag cmd

```bash
go get -u github.com/swaggo/swag/cmd/swag

// 生成文档
swag init
```

## 下载 gin-swagger 并应用

```bash
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

设置 `docs.SwaggerInfo.BasePath = ""`

## swagger 访问

```bash
http://localhost:4001/swagger/index.html
```
