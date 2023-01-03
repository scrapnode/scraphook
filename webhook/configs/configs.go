package configs

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/spf13/viper"
)

var (
	EVENT_TYPE_MESSAGE     = "webhook.message"
	EVENT_TYPE_FORWARD_REQ = "webhook.forward.request"
	EVENT_TYPE_FORWARD_RES = "webhook.forward.response"
)

type Configs struct {
	*xconfig.Configs

	Http      *transport.Configs
	Validator *Validator
	MsgBus    *msgbus.Configs
	Database  *database.Configs
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
