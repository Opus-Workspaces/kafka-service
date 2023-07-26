package redis

import (
	config "kafka-service/app/config/database"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg *config.DatabaseType) *redis.Client {
	redisCfg := cfg.Redis

	addr := redisCfg.Host + redisCfg.Port

	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     redisCfg.Password,
		DB:           redisCfg.DB,
		PoolSize:     redisCfg.PoolSize,
		PoolTimeout:  time.Duration(redisCfg.IdleTimeout) * time.Second,
		MinIdleConns: redisCfg.MinIdleCons,
	})

	return client
}
