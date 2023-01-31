package configs

import (
	"github.com/scrapnode/scrapcore/transport"
	"github.com/spf13/viper"
)

type Http struct {
	ListenAddress string `json:"listen_address" mapstructure:"SCRAPHOOK_WEBHOOK_HTTP_LISTEN_ADDRESS"`
}

func (cfg *Configs) useServiceHttp(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_WEBHOOK_HTTP_LISTEN_ADDRESS", ":8080")

	var configs Http
	if err := provider.Unmarshal(&configs); err != nil {
		return err
	}

	cfg.Http = &transport.Configs{ListenAddress: configs.ListenAddress}
	return nil
}
