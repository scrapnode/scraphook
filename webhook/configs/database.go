package configs

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/spf13/viper"
)

type Database struct {
	Dsn        string `json:"dsn" mapstructure:"SCRAPHOOK_WEBHOOK_DATABASE_DSN"`
	MigrateDir string `json:"migrate_dir" mapstructure:"SCRAPHOOK_WEBHOOK_DATABASE_MIGRATE_DIR"`
}

func (cfg *Configs) useDatabase(provider *viper.Viper) error {
	provider.SetDefault("SCRAPHOOK_WEBHOOK_DATABASE_DSN", "sqlite3:///tmp/scraphook.sqlite")
	provider.SetDefault("SCRAPHOOK_WEBHOOK_DATABASE_MIGRATE_DIR", "./db/migrations")

	var db Database
	if err := provider.Unmarshal(&db); err != nil {
		return err
	}

	cfg.Database = &database.Configs{
		Dsn:        db.Dsn,
		MigrateDir: db.MigrateDir,
	}
	return nil
}
