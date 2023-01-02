package repositories

import (
	"context"
	"github.com/scrapnode/scraphook/webhook/repositories/interfaces"
	"github.com/scrapnode/scraphook/webhook/repositories/sql"
)

func New(ctx context.Context) (*interfaces.Repo, error) {
	return sql.New(ctx), nil
}
