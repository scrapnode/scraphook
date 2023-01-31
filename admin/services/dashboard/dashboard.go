package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

func New(ctx context.Context, app *application.App) *Dashboard {
	logger := xlogger.FromContext(ctx).With("service", "dashboard")
	return &Dashboard{app: app, logger: logger}
}

type Dashboard struct {
	app    *application.App
	logger *zap.SugaredLogger
	server *grpc.Server
}

func (service *Dashboard) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	service.server = grpc.NewServer()
	protos.RegisterAccountServer(service.server, &AccountServer{app: service.app})
	reflection.Register(service.server)

	service.logger.Debug("connected")
	return nil
}

func (service *Dashboard) Stop(ctx context.Context) error {
	if service.server != nil {
		if service.app.Configs.Debug() {
			service.server.Stop()
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			go func() {
				// GracefulStop stops the gRPC server gracefully. It stops the server from
				// accepting new connections and RPCs and blocks until all the pending RPCs are
				// finished. So we need to call .Stop after a couple of seconds
				service.server.GracefulStop()
			}()
			<-ctx.Done()
			service.server.Stop()
		}
	}

	if err := service.app.Disconnect(ctx); err != nil {
		service.logger.Errorw("disconnect app got error", "error", err.Error())
	}

	service.logger.Debug("disconnected")
	return nil
}

func (service *Dashboard) Run(ctx context.Context) error {
	service.logger.Debugw("running")

	listener, err := net.Listen("tcp", service.app.Configs.GRPC.ListenAddress)
	if err != nil {
		return err
	}

	return service.server.Serve(listener)
}
