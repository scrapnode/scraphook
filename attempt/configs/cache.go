package configs

import (
	"github.com/scrapnode/scrapcore/xcache"
	"github.com/spf13/viper"
)

type Cache struct {
	Dsn           string `json:"dsn" mapstructure:"SCRAPHOOK_ATTEMPT_CACHE_DSN"`
	SecondsToLive int64  `json:"seconds_to_live" mapstructure:"SCRAPHOOK_ATTEMPT_CACHE_SECONDS_TO_LIVE"`
}

func (cfg *Configs) useCache(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_ATTEMPT_CACHE_DSN", "bigcache://localhost")
	provider.SetDefault("SCRAPHOOK_ATTEMPT_CACHE_SECONDS_TO_LIVE", 3600)

	var configs Cache
	if err := provider.Unmarshal(&configs); err != nil {
		return err
	}

	cfg.Cache = &xcache.Configs{
		Dsn:           configs.Dsn,
		SecondsToLive: configs.SecondsToLive,
	}
	return nil
}
