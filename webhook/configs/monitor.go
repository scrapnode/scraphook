package configs

import (
	"github.com/scrapnode/scrapcore/monitor"
	"github.com/spf13/viper"
)

type Monitor struct {
	Namespace string `json:"namespace" mapstructure:"SCRAPHOOK_WEBHOOK_MONITOR_NAMESPACE"`
	Version   string `json:"version" mapstructure:"SCRAPHOOK_WEBHOOK_MONITOR_VERSION"`
}

type MonitorTracer struct {
	Endpoint string  `json:"endpoint" mapstructure:"SCRAPHOOK_WEBHOOK_MONITOR_TRACER_ENDPOINT"`
	Ratio    float64 `json:"ratio" mapstructure:"SCRAPHOOK_WEBHOOK_MONITOR_TRACER_RATIO"`
}

func (cfg *Configs) useMonitor(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_WEBHOOK_MONITOR_NAMESPACE", "webhook")
	provider.Set("SCRAPHOOK_WEBHOOK_MONITOR_VERSION", cfg.Version)
	provider.SetDefault("SCRAPHOOK_WEBHOOK_MONITOR_TRACER_ENDPOINT", "0.0.0.0:4317")
	provider.SetDefault("SCRAPHOOK_WEBHOOK_MONITOR_TRACER_RATIO", 1)

	var configs Monitor
	if err := provider.Unmarshal(&configs); err != nil {
		return err
	}

	var tracerConfigs MonitorTracer
	if err := provider.Unmarshal(&tracerConfigs); err != nil {
		return err
	}

	cfg.Monitor = &monitor.Configs{
		Namespace: configs.Namespace,
		Version:   configs.Version,
		Tracer: &monitor.TracerConfigs{
			Endpoint: tracerConfigs.Endpoint,
			Ratio:    tracerConfigs.Ratio,
		},
	}
	return nil
}
