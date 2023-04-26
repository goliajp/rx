# RX - Golang Redis 实用工具库

[![Go Report Card](https://goreportcard.com/badge/github.com/goliajp/rx)](https://goreportcard.com/report/github.com/goliajp/rx)
[![GoDoc](https://pkg.go.dev/badge/github.com/goliajp/rx)](https://pkg.go.dev/github.com/goliajp/rx)

---
[ENGLISH](README.md)
[日本語](README_JP.md)

RX 是 Golia 系列中用于操作 Redis 的 Golang 辅助库。它简化了创建和管理 Redis 连接的过程，可以选择使用默认配置或自定义配置

## 安装

```sh
go get -u github.com/goliajp/rx
```

## 使用
您可以使用默认配置或自定义配置创建 Redis 连接。从 Open() 返回的 *redis.Client 可以与 go-redis 库一起使用以进行进一步操作

### 使用默认配置
```go
r := rx.NewRedis(nil) // nil 表示默认配置
cli := r.Open() // cli 是来自 go-redis 的 *redis.Client
```

### 使用自定义配置
```go
cfg := &rx.RedisConfig{
Host:     "root",
Port:     6379,
Password: "",
Db:       0,
}
r := rx.NewRedis(cfg)
cli := r.Open()
```

## 贡献
我们欢迎对 RX 的贡献！请随时在 GitHub 上提交问题、功能请求或拉取请求

## 许可
本项目采用MIT许可证 - 有关详细信息，请参阅 [LICENSE](LICENSE) 文件