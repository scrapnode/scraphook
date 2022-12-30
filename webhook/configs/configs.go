package configs

import (
	databaseconfigs "github.com/scrapnode/scrapcore/database/configs"
	msgbusconfigs "github.com/scrapnode/scrapcore/msgbus/configs"
	"github.com/scrapnode/scrapcore/xconfig"
)

var EVENT_TYPE_MESSAGE = "webhook.message"

type Configs struct {
	*xconfig.Configs

	Http      *Http
	Validator *Validator
	MsgBus    *msgbusconfigs.Configs
	Database  *databaseconfigs.Configs
}
