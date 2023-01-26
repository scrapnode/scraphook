package configs

import "github.com/spf13/viper"

type Examiner struct {
	MaxCount int64 `json:"max_count" mapstructure:"SCRAPHOOK_ATTEMPT_EXAMINER_MAX_COUNT"`
}

func (cfg *Configs) useExaminer(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_ATTEMPT_EXAMINER_MAX_COUNT", 3)

	return provider.Unmarshal(&cfg.Examiner)
}
