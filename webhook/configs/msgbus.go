package configs

import (
	msgbusconfigs "github.com/scrapnode/scrapcore/msgbus/configs"
	"github.com/spf13/viper"
)

type MsgBus struct {
	Uri    string `json:"uri" mapstructure:"SCRAPHOOK_WEBHOOK_MSGBUS_URI"`
	Region string `json:"region" mapstructure:"SCRAPHOOK_WEBHOOK_MSGBUS_REGION"`
	Name   string `json:"name" mapstructure:"SCRAPHOOK_WEBHOOK_MSGBUS_NAME"`

	MaxRetry int `json:"max_retry" mapstructure:"SCRAPHOOK_WEBHOOK_MSGBUS_MAX_RETRY"`
}

func (cfg *Configs) useMsgBus(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_WEBHOOK_MSGBUS_URI", "nats://127.0.0.1:4222")
	provider.SetDefault("SCRAPHOOK_WEBHOOK_MSGBUS_REGION", "earth")
	provider.SetDefault("SCRAPHOOK_WEBHOOK_MSGBUS_NAME", "scraphook")

	var bus MsgBus
	if err := provider.Unmarshal(&bus); err != nil {
		return err
	}

	cfg.MsgBus = &msgbusconfigs.Configs{
		Uri:      bus.Uri,
		Region:   bus.Region,
		Name:     bus.Name,
		MaxRetry: bus.MaxRetry,
	}
	return nil
}
