package ratelimit

import (
	"context"
	"labs-four/config"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisLimiter struct {
	client *redis.Client
	ctx    context.Context
	cfg    config.AppSettings
}

func NewRedisLimiter(cfg config.AppSettings) (*RedisLimiter, error) {
	db, err := strconv.Atoi(cfg.RedisDB)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       db,
	})

	return &RedisLimiter{
		client: rdb,
		cfg:    cfg,
		ctx:    context.Background(),
	}, nil
}

func (r *RedisLimiter) Rate(ip string, token string) (bool, error) {
	var key string
	var limit int

	if token != "" {
		key = "ratelimit:token:" + token
		val, err := r.client.Get(r.ctx, "ratelimit:tokenlimit:"+token).Result()
		if err == nil {
			if l, err := strconv.Atoi(val); err == nil {
				limit = l
			}
		}
		if limit == 0 {
			limit = r.cfg.DefaultTokenLimit
		}
	} else {
		key = "ratelimit:ip:" + ip
		limit = r.cfg.DefaultIPLimit
	}

	count, _ := r.client.Incr(r.ctx, key).Result()
	if count == 1 {
		r.client.Expire(r.ctx, key, time.Duration(r.cfg.RateLimitDuration)*time.Second)
	}

	if int(count) > limit {
		r.client.Expire(r.ctx, key, time.Duration(r.cfg.BlockDuration)*time.Second)
		return false, nil
	}

	return true, nil
}
