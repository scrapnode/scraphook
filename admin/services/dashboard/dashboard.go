package dashboard

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"regexp"
	"strings"
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

	public := []*regexp.Regexp{
		regexp.MustCompile("^/scraphook\\.admin\\.dashboard\\.v1\\.Account/.*"),
	}
	verify := application.NewAccountVerify(service.app)
	service.server = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				for _, r := range public {
					if r.MatchString(info.FullMethod) {
						service.app.Logger.Debugw("ignore public method", "method", info.FullMethod)
						return handler(ctx, req)
					}
				}
				meta, ok := metadata.FromIncomingContext(ctx)
				if !ok {
					return nil, status.Error(codes.Unauthenticated, "no given headers")
				}

				// authentication
				header, ok := meta["authorization"]
				if !ok || len(header) == 0 {
					return nil, status.Error(codes.Unauthenticated, "no given authorization header")
				}
				segments := strings.Split(header[0], " ")
				token := segments[0]
				// bearer token
				if len(segments) == 2 {
					token = segments[1]
				}
				request := &application.AccountVerifyReq{AccessToken: token}
				if types := meta.Get("X-ScrapNode-Token-Type"); len(types) > 0 {
					request.Type = types[0]
				}
				verifyCtx := context.WithValue(ctx, pipeline.CTXKEY_REQ, request)
				verifyCtx, err = verify(verifyCtx)
				if err != nil {
					return nil, status.Error(codes.Unauthenticated, "could not verify account")
				}
				response := verifyCtx.Value(pipeline.CTXKEY_RES).(*application.AccountVerifyRes)
				ctx = context.WithValue(ctx, pipeline.CTXKEY_ACC, response.Account)

				// workspace validation
				if header := meta.Get("X-ScrapNode-Workspace-Id"); len(header) > 0 {
					workspace, err := service.app.Repo.Workspace.GetById(header[0])
					if err != nil {
						service.app.Logger.Errorw("could not verify workspace", "error", err.Error())
						return nil, status.Error(codes.Unauthenticated, "workspace is not found")
					}
					//	@TODO: validate workspace is exist
					ctx = context.WithValue(ctx, pipeline.CTXKEY_WS, workspace)
				}

				return handler(ctx, req)
			},
		)),
	)
	protos.RegisterAccountServer(service.server, NewAccountServer(service.app))
	protos.RegisterWebhookServer(service.server, NewWebhookServer(service.app))
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
	service.logger.Debugw("running", "listen_address", service.app.Configs.GRPC.ListenAddress)

	listener, err := net.Listen("tcp", service.app.Configs.GRPC.ListenAddress)
	if err != nil {
		return err
	}

	return service.server.Serve(listener)
}
