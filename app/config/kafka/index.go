package config

import (
	"fmt"
	cfg "kafka-service/app/config/consts"
	"reflect"

	"github.com/spf13/viper"
)

func ConfigKafka() KafkaType {
	bootstrapServers := viper.GetString(cfg.PrefixKafka + ".bootstrap_servers")
	topic := viper.GetString(cfg.PrefixKafka + ".topic")
	groupID := viper.GetString(cfg.PrefixKafka + ".consumer_group_id")
	autoOffsetReset := viper.GetString(cfg.PrefixKafka + ".auto_offset_reset")
	sessionTimeout := viper.GetInt(cfg.PrefixKafka + ".session_timeout_ms")
	enableAutoOffsetStore := viper.GetBool(cfg.PrefixKafka + ".enable_auto_offset_store")
	fmt.Println("bootstrapServers: ", bootstrapServers)
	fmt.Println(reflect.TypeOf(bootstrapServers))

	return KafkaType{
		BootstrapServers:      bootstrapServers,
		Topic:                 topic,
		GroupID:               groupID,
		AutoOffsetReset:       autoOffsetReset,
		SessionTimeoutMs:      sessionTimeout,
		EnableAutoOffsetStore: enableAutoOffsetStore,
	}
}
