package configs

import (
	"github.com/spf13/viper"
)

type AuthRoot struct {
	AccessKeyId     string `json:"access_key_id" mapstructure:"SCRAPHOOK_ADMIN_ACCESS_KEY_ID"`
	AccessKeySecret string `json:"access_key_secret" mapstructure:"SCRAPHOOK_ADMIN_ACCESS_KEY_SECRET"`
}

func (cfg *Configs) useAuthRoot(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_ADMIN_ACCESS_KEY_ID", "aki_root")
	provider.SetDefault("SCRAPHOOK_ADMIN_ACCESS_KEY_SECRET", "aks_ashortsecretthatyoushouldchange")

	if err := provider.Unmarshal(&cfg.AuthRoot); err != nil {
		return err
	}
	return nil
}
