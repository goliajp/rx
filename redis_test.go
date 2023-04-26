package gormx

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRedis(t *testing.T) {
	var (
		cfg = &RedisConfig{
			Host:     "root",
			Port:     6379,
			Password: "",
			Db:       0,
		}

		errCfg = &RedisConfig{
			Host:     "root",
			Port:     6666,
			Password: "",
			Db:       0,
		}
	)
	Convey("with config", t, func() {
		r := NewRedis(cfg)
		So(r, ShouldNotBeNil)
	})
	Convey("without config", t, func() {
		r := NewRedis(nil)
		So(r, ShouldNotBeNil)
	})
	Convey("not initialized", t, func() {
		r := &Redis{}
		cli := r.Open()
		So(cli, ShouldBeNil)
	})
	Convey("connect failed", t, func() {
		r := NewRedis(errCfg)
		cli := r.Open()
		So(r, ShouldNotBeNil)
		So(cli, ShouldBeNil)
	})
	Convey("normal", t, func() {
		r := NewRedis(nil)
		cli := r.Open()
		So(r, ShouldNotBeNil)
		So(cli, ShouldNotBeNil)
	})
	Convey("open with params", t, func() {
		r := NewRedis(nil)
		cli := r.Open("0", "localhost", "6379")
		So(r, ShouldNotBeNil)
		So(cli, ShouldNotBeNil)
	})
	Convey("open with invalid params", t, func() {
		r := NewRedis(nil)
		cli := r.Open("0", "localhost", "6379")
		cli1 := r.Open("0", "localhost")
		cli2 := r.Open("0", "localhost", "6379", "")
		cli3 := r.Open("0", "localhost", "6379", "", "foo")
		So(r, ShouldNotBeNil)
		So(cli, ShouldNotBeNil)
		So(cli1, ShouldBeNil)
		So(cli2, ShouldNotBeNil)
		So(cli3, ShouldBeNil)
	})
	Convey("get go-redis v8 singleton", t, func() {
		r := NewRedis(nil)
		cli := r.Client()
		So(r, ShouldNotBeNil)
		So(cli, ShouldNotBeNil)
	})
}
