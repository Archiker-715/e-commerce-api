package redis

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

type client struct {
	client          *goredislib.Client
	redsyncInstance *redsync.Redsync
}

func NewRedisClient() *client {
	c := goredislib.NewClient(&goredislib.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pool := goredis.NewPool(c)
	rs := redsync.New(pool)

	return &client{
		client:          c,
		redsyncInstance: rs,
	}
}
