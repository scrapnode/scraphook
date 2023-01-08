package configs

import (
	"github.com/scrapnode/scrapcore/msgbus"
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

	var configs MsgBus
	if err := provider.Unmarshal(&configs); err != nil {
		return err
	}

	cfg.MsgBus = &msgbus.Configs{
		Uri:      configs.Uri,
		Region:   configs.Region,
		Name:     configs.Name,
		MaxRetry: configs.MaxRetry,
	}
	return nil
}
