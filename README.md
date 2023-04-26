# RX - Golang Redis Helper Library

[![Go Report Card](https://goreportcard.com/badge/github.com/goliajp/rx)](https://goreportcard.com/report/github.com/goliajp/rx)
[![GoDoc](https://pkg.go.dev/badge/github.com/goliajp/rx)](https://pkg.go.dev/github.com/goliajp/rx)

---
[简体中文](README_CN.md)
[日本語](README_JP.md)

RX is a Golang helper library for working with Redis as part of the Golia series. It simplifies the process of creating and managing Redis connections with the ability to use custom or default configurations.

## Installation

```sh
go get -u github.com/goliajp/rx
```

## Usage
You can create a Redis connection with the default configuration or by providing a custom configuration. The returned *redis.Client from Open() can be used with the go-redis library for further operations.

### Using Default Configuration
```go
r := rx.NewRedis(nil) // nil for default configuration
cli := r.Open() // cli is a *redis.Client from go-redis
```

### Custom Connection Configuration
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

## Contributing
We welcome contributions to RX! Feel free to submit issues, feature requests, or pull requests on GitHub.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.