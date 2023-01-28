package repositories

import (
	"github.com/scrapnode/scrapcore/database"
)

type SqlEndpoint struct {
	db *database.SQL
}
