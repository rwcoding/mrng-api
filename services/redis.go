package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/rwcoding/mrng/config"
)

var redisPools []*redisClient

func NewRedis() *redisClient {
	if len(redisPools) == 0 {
		return nil
	}
	return redisPools[0]
}

type redisClient struct {
	client *redis.Client
	ctx    context.Context
}

func (r *redisClient) Set(k string, v interface{}) {
	r.client.Set(r.ctx, k, v, 0)
}

func (r *redisClient) Del(k string) {
	r.client.Del(r.ctx, k)
}

func (r *redisClient) Get(k string) *redis.StringCmd {
	return r.client.Get(r.ctx, k)
}

func InitRedis() {
	redisList := config.GetRedis()
	for _, v := range redisList {
		if v.Addr == "" {
			continue
		}
		rdb := redis.NewClient(&redis.Options{
			Addr:     v.Addr,
			Password: v.Password,
			PoolSize: v.Pool,
			DB:       0,
		})

		redisPools = append(redisPools, &redisClient{
			client: rdb,
			ctx:    context.Background(),
		})
	}
}
