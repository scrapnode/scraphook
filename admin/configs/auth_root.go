package configs

import (
	"fmt"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/spf13/viper"
)

type AuthRoot struct {
	AccessKeyId     string `json:"access_key_id" mapstructure:"SCRAPHOOK_ADMIN_ACCESS_KEY_ID"`
	AccessKeySecret string `json:"access_key_secret" mapstructure:"SCRAPHOOK_ADMIN_ACCESS_KEY_SECRET"`
}

func (cfg *Configs) useAuthRoot(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_ADMIN_ACCESS_KEY_ID", fmt.Sprintf("%s_root", auth.ACCESS_KEY_ID_PREF))
	provider.SetDefault("SCRAPHOOK_ADMIN_ACCESS_KEY_SECRET", fmt.Sprintf("%s_ashortsecretthatyoushouldchange", auth.ACCESS_KEY_SECRET_PREF))

	if err := provider.Unmarshal(&cfg.AuthRoot); err != nil {
		return err
	}
	return nil
}
