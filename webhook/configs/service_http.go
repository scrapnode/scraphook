package configs

import (
	"github.com/scrapnode/scrapcore/transport"
	"github.com/spf13/viper"
)

type Http struct {
	ListenAddress string `json:"listen_address" mapstructure:"SCRAPHOOK_WEBHOOK_HTTP_LISTEN_ADDRESS"`
}

func (cfg *Configs) useHttp(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_WEBHOOK_HTTP_LISTEN_ADDRESS", ":8080")

	var http Http
	if err := provider.Unmarshal(&http); err != nil {
		return err
	}

	cfg.Http = &transport.Configs{ListenAddress: http.ListenAddress}
	return nil
}
