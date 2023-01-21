package repositories

import (
	"github.com/scrapnode/scrapcore/database"
)

type Repo struct {
	Database database.Database
	Message  MessageRepo
	Request  RequestRepo
	Response ResponseRepo
	Endpoint EndpointRepo
}
