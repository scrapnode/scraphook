package configs

import "github.com/spf13/viper"

type Http struct {
	ServerListenAddress string `json:"server_listen_address" mapstructure:"SCRAPHOOK_WEBHOOK_HTTP_SERVER_LISTEN_ADDRESS"`
}

func (cfg *Configs) useHttp(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_WEBHOOK_HTTP_SERVER_LISTEN_ADDRESS", ":8080")

	return provider.Unmarshal(&cfg.Http)
}
