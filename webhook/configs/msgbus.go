package configs

import (
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/spf13/viper"
)

type MsgBus struct {
	Dsn    string `json:"uri" mapstructure:"SCRAPHOOK_WEBHOOK_MSGBUS_DSN"`
	Region string `json:"region" mapstructure:"SCRAPHOOK_WEBHOOK_MSGBUS_REGION"`
	Name   string `json:"name" mapstructure:"SCRAPHOOK_WEBHOOK_MSGBUS_NAME"`

	MaxRetry int `json:"max_retry" mapstructure:"SCRAPHOOK_WEBHOOK_MSGBUS_MAX_RETRY"`
}

func (cfg *Configs) useMsgBus(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_WEBHOOK_MSGBUS_DSN", "nats://127.0.0.1:4222")
	provider.SetDefault("SCRAPHOOK_WEBHOOK_MSGBUS_REGION", "earth")
	provider.SetDefault("SCRAPHOOK_WEBHOOK_MSGBUS_NAME", "scraphook")
	provider.SetDefault("SCRAPHOOK_WEBHOOK_MSGBUS_MAX_RETRY", 1)

	var configs MsgBus
	if err := provider.Unmarshal(&configs); err != nil {
		return err
	}

	cfg.MsgBus = &msgbus.Configs{
		Dsn:      configs.Dsn,
		Region:   configs.Region,
		Name:     configs.Name,
		MaxRetry: configs.MaxRetry,
	}
	return nil
}
