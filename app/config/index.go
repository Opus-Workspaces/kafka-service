package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/spf13/viper"

	cfg "kafka-service/app/config/consts"
	kafka "kafka-service/app/config/kafka"
	server "kafka-service/app/config/server"
)

var buildEnv string

func init() {
	var runDefaultEnv = cfg.DevEnv
	var getEnv = flag.String("env", runDefaultEnv, "Environment for app")

	flag.Parse()
	runEnv := string(*getEnv)
	path := cfg.PathConfig

	if buildEnv != "" {
		runEnv = buildEnv
	}

	var fullFileName = fmt.Sprintf("%s/%s.json", path, runEnv)
	viper.SetConfigFile(fullFileName)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(errors.New("config file not found"))
		}
		panic(errors.New("config file found but another error was produced"))
	}

	if viper.GetBool("DEBUG") {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	httpServer := viper.GetString(cfg.PrefixHttpServer+".address") + viper.GetString(cfg.PrefixHttpServer+".port")
	os.Setenv(cfg.PrefixHttpServer, httpServer)
	os.Setenv(cfg.Env, runEnv)
}

type Config struct {
	Server server.ServerType
	// DB     db.DatabaseType
	Kafka kafka.KafkaType
}

func LoadConfig() Config {

	cfgServer := server.ConfigServer()
	// cfgDB := db.ConfigDatabase()
	cfgKafka := kafka.ConfigKafka()

	return Config{
		Server: cfgServer,
		// DB:     cfgDB,
		Kafka: cfgKafka,
	}
}
