# RX - Golang Redis ユーティリティ

[![Go Report Card](https://goreportcard.com/badge/github.com/goliajp/rx)](https://goreportcard.com/report/github.com/goliajp/rx)
[![GoDoc](https://pkg.go.dev/badge/github.com/goliajp/rx)](https://pkg.go.dev/github.com/goliajp/rx)

---
[ENGLISH](README.md)
[简体中文](README_CN.md)

RXは、Golia シリーズの一部として Redis を操作するための Golang ヘルパーライブラリです。カスタム設定またはデフォルト設定を使用して、Redis 接続の作成と管理を簡略化できます

## インストール

```sh
go get -u github.com/goliajp/rx/v2
```

## 使い方
デフォルト設定またはカスタム設定を使用して、Redis 接続を作成できます。Open()から返された*redis.Clientは、go-redisライブラリと一緒に使用してさらなる操作ができます

### デフォルト設定を使用
```go
r := rx.NewRedis(nil) // nil はデフォルト設定を意味します
cli := r.Open() // cli は go-redis の *redis.Client です
```

### カスタム設定を使用
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
RX への貢献は大歓迎です！GitHub で問題を提出したり、機能のリクエストやプルリクエストを提出してください

## ライセンス
このプロジェクトはMITライセンスの下でライセンスされています。詳細については、 [LICENSE](LICENSE) ファイルを参照してください