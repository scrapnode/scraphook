package configs

import "github.com/spf13/viper"

type Validator struct {
	ChallengeQueryName   string `json:"challenge_query_name" mapstructure:"SCRAPHOOK_WEBHOOK_VALIDATOR_CHALLENGE_QUERY_NAME"`
	VerifyTokenQueryName string `json:"verify_token_query_name" mapstructure:"SCRAPHOOK_WEBHOOK_VALIDATOR_VERIFY_TOKEN_QUERY_NAME"`
}

func (cfg *Configs) useValidator(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_WEBHOOK_VALIDATOR_CHALLENGE_QUERY_NAME", "wh.challenge")
	provider.SetDefault("SCRAPHOOK_WEBHOOK_VALIDATOR_VERIFY_TOKEN_QUERY_NAME", "wh.verify_token")

	return provider.Unmarshal(&cfg.Validator)
}
