package server

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/infrastructure"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type Http struct {
	infra  *infrastructure.Infra
	logger *zap.SugaredLogger

	server *http.Server
}

func New(ctx context.Context, infra *infrastructure.Infra) *Http {
	logger := xlogger.FromContext(ctx).With("pkg", "server.http")
	return &Http{infra: infra, logger: logger}
}

func (server *Http) Start(ctx context.Context) error {
	if err := server.infra.Connect(ctx); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")

		data, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(map[string]interface{}{
			"version": server.infra.Configs.Version,
			"env":     server.infra.Configs.Env,
		})
		i, err := writer.Write(data)
		log.Println(i, err)
	})

	server.server = &http.Server{
		Addr:    server.infra.Configs.Http.ServerListenAddress,
		Handler: mux,
	}
	server.logger.Debug("connected")
	return nil
}

func (server *Http) Stop(ctx context.Context) error {
	if server.server != nil {
		if err := server.server.Shutdown(ctx); err != nil {
			server.logger.Errorw("shutdown http server got error", "error", err.Error())
		}
	}

	if err := server.infra.Disconnect(ctx); err != nil {
		return err
	}

	server.logger.Debug("disconnected")
	return nil
}

func (server *Http) Run() error {
	server.logger.Debugw("running", "address", server.server.Addr)

	if err := server.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}
