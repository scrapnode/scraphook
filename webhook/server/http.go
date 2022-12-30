package server

import (
	"context"
	jsoniter "github.com/json-iterator/go"
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

func New(infra *infrastructure.Infra) *Http {
	logger := infra.Logger.With("pkg", "server.http")
	return &Http{infra: infra, logger: logger}
}

func (server *Http) Start(ctx context.Context) error {
	if err := server.infra.Connect(ctx); err != nil {
		return err
	}

	server.server = &http.Server{
		Addr: server.infra.Configs.Http.ServerListenAddress,
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")

		data, err := json.Marshal(map[string]interface{}{
			"version": server.infra.Configs.Version,
			"env":     server.infra.Configs.Env,
		})
		i, err := writer.Write(data)
		log.Println(i, err)
	})

	server.logger.Debug("connected")
	return nil
}

func (server *Http) Stop(ctx context.Context) error {
	if err := server.server.Shutdown(ctx); err != nil {
		server.logger.Errorw("shutdown http server got error", "error", err.Error())
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
