package services

import (
	"context"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scraphook/webhook/application"
	"github.com/scrapnode/scraphook/webhook/configs"
	"github.com/scrapnode/scraphook/webhook/services/scheduler"
	"github.com/scrapnode/scraphook/webhook/services/sender"
	"github.com/scrapnode/scraphook/webhook/services/webserver"
)

func New(ctx context.Context, transport string) (transport.Transport, error) {
	cfg := configs.FromContext(ctx)
	app, err := application.New(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if transport == "scheduler" {
		return scheduler.New(ctx, app), nil
	}
	if transport == "sender" {
		return sender.New(ctx, app), nil
	}

	// by default, we will serve HTTP server
	return webserver.New(ctx, app), nil
}
