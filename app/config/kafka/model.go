package config

type KafkaType struct {
	BootstrapServers      string
	Topic                 string
	GroupID               string
	AutoOffsetReset       string
	SessionTimeoutMs      int
	EnableAutoOffsetStore bool
}
