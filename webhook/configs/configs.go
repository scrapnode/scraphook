package configs

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/monitor"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/spf13/viper"
)

var (
	EVENT_TYPE_MESSAGE      = "webhook.message"
	EVENT_TYPE_SCHEDULE_REQ = "webhook.schedule.request"
	EVENT_TYPE_SCHEDULE_RES = "webhook.schedule.response"
)

type Configs struct {
	*xconfig.Configs

	Validator *Validator
	Http      *transport.Configs
	MsgBus    *msgbus.Configs
	Database  *database.Configs
	Monitor   *monitor.Configs
}

func New(provider *viper.Viper) (*Configs, error) {
	cfg := &Configs{Configs: &xconfig.Configs{}}
	if err := cfg.Configs.Unmarshal(provider); err != nil {
		return nil, err
	}
	if err := cfg.useValidator(provider); err != nil {
		return nil, err
	}
	if err := cfg.useHttp(provider); err != nil {
		return nil, err
	}
	if err := cfg.useMsgBus(provider); err != nil {
		return nil, err
	}
	if err := cfg.useDatabase(provider); err != nil {
		return nil, err
	}
	if err := cfg.useMonitor(provider); err != nil {
		return nil, err
	}

	return cfg, nil
}
