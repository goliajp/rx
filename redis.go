package rx

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

type Redis struct {
	init bool
	cli  *redis.Client
	cfg  *RedisConfig
	once sync.Once
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	Db       int
}

var DefaultRedisConfig = &RedisConfig{
	Host:     "localhost",
	Port:     6379,
	Password: "",
	Db:       0,
}

func NewRedis(cfg *RedisConfig) *Redis {
	if cfg == nil {
		cfg = DefaultRedisConfig
	}
	return &Redis{
		init: true,
		cfg:  cfg,
	}
}

func (r *Redis) Client(args ...string) *redis.Client {
	r.once.Do(func() {
		r.cli = r.Open(args...)
	})
	if !TestConnection(r.cli) {
		r.cli = r.Open(args...)
	}
	return r.cli
}

func (r *Redis) Open(args ...string) *redis.Client {
	if !r.init {
		return nil
	}
	host := r.cfg.Host
	port := r.cfg.Port
	password := r.cfg.Password
	db := r.cfg.Db

	var err error
	switch len(args) {
	case 1: // db
		db, err = strconv.Atoi(args[0])
		if err != nil {
			log.Error("args[0] should be int")
			return nil
		}
	case 2:
		log.Error("num of args should be 0, 1 or 3")
		return nil
	case 3: // db, host, port
		db, err = strconv.Atoi(args[0])
		if err != nil {
			log.Error("args[0] should be int")
			return nil
		}
		host = args[1]
		port, err = strconv.Atoi(args[2])
		if err != nil {
			log.Error("args[2] should be int")
			return nil
		}
	case 4: // db, host, port, password
		db, err = strconv.Atoi(args[0])
		if err != nil {
			log.Error("args[0] should be int")
			return nil
		}
		host = args[1]
		port, err = strconv.Atoi(args[2])
		if err != nil {
			log.Error("args[2] should be int")
			return nil
		}
		password = args[3]
	default:
		if len(args) > 4 {
			log.Error("num of args should be 0, 1 or 3")
			return nil
		}
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	if r.cfg.Port == 0 {
		addr = r.cfg.Host
	}
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if !TestConnection(cli) {
		return nil
	}
	return cli
}

func TestConnection(cli *redis.Client) bool {
	if cli == nil {
		return false
	}
	err := cli.Ping(context.Background()).Err()
	return err == nil
}
