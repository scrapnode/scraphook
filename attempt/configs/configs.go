package configs

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/scrapnode/scrapcore/xmonitor"
	"github.com/spf13/viper"
)

type Configs struct {
	*xconfig.Configs

	BucketTemplate string   `json:"bucket_template" mapstructure:"SCRAPHOOK_BUCKET_TEMPLATE"`
	Trigger        *Trigger `json:"trigger"`

	MsgBus   *msgbus.Configs   `json:"msg_bus"`
	Database *database.Configs `json:"database"`
	Monitor  *xmonitor.Configs `json:"monitor"`
}

func New(provider *viper.Viper) (*Configs, error) {
	provider.SetDefault("SCRAPHOOK_BUCKET_TEMPLATE", "20060102")

	cfg := &Configs{Configs: &xconfig.Configs{}}
	if err := cfg.Configs.Unmarshal(provider); err != nil {
		return nil, err
	}
	if err := cfg.Unmarshal(provider); err != nil {
		return nil, err
	}
	if err := cfg.useTrigger(provider); err != nil {
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
