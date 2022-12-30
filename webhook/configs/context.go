package configs

import "context"

type ctxkey string

const CTXKEY ctxkey = "webhook.configs"

func WithContext(ctx context.Context, cfg *Configs) context.Context {
	return context.WithValue(ctx, CTXKEY, cfg)
}

func FromContext(ctx context.Context) *Configs {
	return ctx.Value(CTXKEY).(*Configs)
}
