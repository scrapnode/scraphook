package configs

import (
	"github.com/scrapnode/scrapcore/transport"
	"github.com/spf13/viper"
)

type GRPC struct {
	ListenAddress string `json:"listen_address" mapstructure:"SCRAPHOOK_ADMIN_GRPC_LISTEN_ADDRESS"`
}

func (cfg *Configs) useGRPC(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_ADMIN_GRPC_LISTEN_ADDRESS", ":8081")

	var configs GRPC
	if err := provider.Unmarshal(&configs); err != nil {
		return err
	}

	cfg.GRPC = &transport.Configs{ListenAddress: configs.ListenAddress}
	return nil
}
