package services

import (
	"context"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/configs"
	"github.com/scrapnode/scraphook/admin/services/dashboard"
)

func New(ctx context.Context, name string) (transport.Transport, error) {
	cfg := configs.FromContext(ctx)
	app, err := application.New(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// by default, we will serve dashboard service
	return dashboard.New(ctx, app), nil
}
