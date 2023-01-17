package services

import (
	"context"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scraphook/capture/application"
	"github.com/scrapnode/scraphook/capture/configs"
	"github.com/scrapnode/scraphook/capture/services/message"
)

func New(ctx context.Context, name string) (transport.Transport, error) {
	cfg := configs.FromContext(ctx)
	app, err := application.New(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// by default, we will serve trigger
	return message.New(ctx, app), nil
}
