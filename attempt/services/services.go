package services

import (
	"context"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scraphook/attempt/application"
	"github.com/scrapnode/scraphook/attempt/configs"
	"github.com/scrapnode/scraphook/attempt/services/trigger"
)

func New(ctx context.Context, name string) (transport.Transport, error) {
	cfg := configs.FromContext(ctx)
	app, err := application.New(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// by default, we will serve trigger
	return trigger.New(ctx, app), nil
}
