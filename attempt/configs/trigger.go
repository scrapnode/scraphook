package configs

import "github.com/spf13/viper"

type Trigger struct {
	CronPattern string `json:"cron_pattern" mapstructure:"SCRAPHOOK_ATTEMPT_TRIGGER_CRON_PATTERN"`
}

func (cfg *Configs) useTrigger(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_ATTEMPT_TRIGGER_CRON_PATTERN", "*/3 * * * *")

	return provider.Unmarshal(&cfg.Trigger)
}
