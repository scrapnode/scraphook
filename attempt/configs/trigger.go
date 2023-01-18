package configs

import "github.com/spf13/viper"

type Trigger struct {
	CronPattern         string `json:"cron_pattern" mapstructure:"SCRAPHOOK_ATTEMPT_TRIGGER_CRON_PATTERN"`
	ScanSize            int    `json:"scan_size" mapstructure:"SCRAPHOOK_ATTEMPT_TRIGGER_BUCKET_SCAN_SIZE"`
	BucketCount         int    `json:"bucket_count" mapstructure:"SCRAPHOOK_ATTEMPT_TRIGGER_BUCKET_COUNT"`
	BucketSizeInMinutes int    `json:"bucket_size_in_minutes" mapstructure:"SCRAPHOOK_ATTEMPT_TRIGGER_BUCKET_SIZE_IN_MINUTES"`
}

func (cfg *Configs) useTrigger(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_ATTEMPT_TRIGGER_CRON_PATTERN", "*/3 * * * *")
	provider.SetDefault("SCRAPHOOK_ATTEMPT_TRIGGER_BUCKET_SCAN_SIZE", 100)
	provider.SetDefault("SCRAPHOOK_ATTEMPT_TRIGGER_BUCKET_COUNT", 3)
	provider.SetDefault("SCRAPHOOK_ATTEMPT_TRIGGER_BUCKET_SIZE_IN_MINUTES", 60)

	return provider.Unmarshal(&cfg.Trigger)
}
