package configs

import (
	databaseconfigs "github.com/scrapnode/scrapcore/database/configs"
	msgbusconfigs "github.com/scrapnode/scrapcore/msgbus/configs"
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/spf13/viper"
)

var EVENT_TYPE_MESSAGE = "webhook.message"

type Configs struct {
	*xconfig.Configs

	Http      *Http
	Validator *Validator
	MsgBus    *msgbusconfigs.Configs
	Database  *databaseconfigs.Configs
}

func New(provider *viper.Viper) (*Configs, error) {
	cfg := &Configs{Configs: &xconfig.Configs{}}
	if err := cfg.Configs.Unmarshal(provider); err != nil {
		return nil, err
	}
	if err := provider.Unmarshal(cfg); err != nil {
		return nil, err
	}
	if err := cfg.useDatabase(provider); err != nil {
		return nil, err
	}
	if err := cfg.useHttp(provider); err != nil {
		return nil, err
	}
	if err := cfg.useMsgBus(provider); err != nil {
		return nil, err
	}
	if err := cfg.useValidator(provider); err != nil {
		return nil, err
	}

	return cfg, nil
}
