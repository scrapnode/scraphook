package interfaces

import "context"

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Run(ctx context.Context) error
}
