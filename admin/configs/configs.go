package configs

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scrapcore/xcache"
	"github.com/scrapnode/scrapcore/xconfig"
	"github.com/scrapnode/scrapcore/xmonitor"
	"github.com/spf13/viper"
)

type Configs struct {
	*xconfig.Configs

	BucketTemplate string `json:"bucket_template" mapstructure:"SCRAPHOOK_BUCKET_TEMPLATE"`

	GRPC     *transport.Configs `json:"grpc"`
	MsgBus   *msgbus.Configs    `json:"msg_bus"`
	Database *database.Configs  `json:"database"`
	Cache    *xcache.Configs    `json:"xcache"`
	Monitor  *xmonitor.Configs  `json:"monitor"`
}

func New(provider *viper.Viper) (*Configs, error) {
	provider.SetDefault("SCRAPHOOK_BUCKET_TEMPLATE", "2006010215")

	cfg := &Configs{Configs: &xconfig.Configs{}}
	if err := cfg.Configs.Unmarshal(provider); err != nil {
		return nil, err
	}
	if err := provider.Unmarshal(cfg); err != nil {
		return nil, err
	}
	if err := cfg.useGRPC(provider); err != nil {
		return nil, err
	}
	if err := cfg.useMsgBus(provider); err != nil {
		return nil, err
	}
	if err := cfg.useDatabase(provider); err != nil {
		return nil, err
	}
	if err := cfg.useCache(provider); err != nil {
		return nil, err
	}
	if err := cfg.useMonitor(provider); err != nil {
		return nil, err
	}

	return cfg, nil
}
