package configs

import (
	"github.com/scrapnode/scrapcore/xmonitor"
	"github.com/spf13/viper"
)

type Monitor struct {
	Provider       string `json:"provider" mapstructure:"SCRAPHOOK_CAPTURE_MONITOR_PROVIDER"`
	ServiceName    string `json:"service_name" mapstructure:"SCRAPHOOK_CAPTURE_MONITOR_SERVICE_NAME"`
	ServiceVersion string `json:"version" mapstructure:"SCRAPHOOK_CAPTURE_MONITOR_SERVICE_VERSION"`
}

func (cfg *Configs) useMonitor(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_CAPTURE_MONITOR_PROVIDER", "noop")
	provider.SetDefault("SCRAPHOOK_CAPTURE_MONITOR_SERVICE_NAME", "")
	provider.Set("SCRAPHOOK_CAPTURE_MONITOR_VERSION", cfg.Version)

	var configs Monitor
	if err := provider.Unmarshal(&configs); err != nil {
		return err
	}

	cfg.Monitor = &xmonitor.Configs{
		Provider:       configs.Provider,
		ServiceName:    configs.ServiceName,
		ServiceVersion: configs.ServiceVersion,
	}
	return nil
}
