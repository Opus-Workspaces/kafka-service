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
	fmt.Println("bootstrapServers: ", bootstrapServers)
	fmt.Println(reflect.TypeOf(bootstrapServers))
	return KafkaType{
		BootstrapServers: bootstrapServers,
		Topic:            topic,
		GroupID:          groupID,
		AutoOffsetReset:  autoOffsetReset,
	}
}
