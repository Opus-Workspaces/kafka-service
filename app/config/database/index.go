package config

import (
	cfg "kafka-service/app/config/consts"

	"github.com/spf13/viper"
)

func ConfigDatabase() DatabaseType {
	var redis = cfg.PrefixDB + ".redis"

	redisHost := viper.GetString(redis + ".host")
	redisPort := viper.GetString(redis + ".port")
	redisPassword := viper.GetString(redis + ".password")
	redisDB := viper.GetInt(redis + ".db")
	redisDefaultDB := viper.GetInt(redis + ".default_db")
	redisMinIdleCons := viper.GetInt(redis + ".min_idle_cons")
	redisPoolSize := viper.GetInt(redis + ".pool_size")
	redisIdleTimeout := viper.GetInt(redis + ".idle_timeout")

	return DatabaseType{
		Redis: &Redis{
			Host:        redisHost,
			Port:        redisPort,
			Password:    redisPassword,
			DB:          redisDB,
			DefaultDB:   redisDefaultDB,
			MinIdleCons: redisMinIdleCons,
			PoolSize:    redisPoolSize,
			IdleTimeout: redisIdleTimeout,
		},
	}
}
