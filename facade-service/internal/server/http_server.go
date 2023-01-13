package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kavu/go_reuseport"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model/config"
	"net/http"
)

const protocol = "tcp"

type HTTPServer struct {
	HttpServer *http.Server
}

func NewHTTPServer(cfg *config.HTTPConfig, handler *gin.Engine) *HTTPServer {
	return &HTTPServer{
		HttpServer: &http.Server{
			Addr:           fmt.Sprintf(":%s", cfg.Port),
			Handler:        handler,
			MaxHeaderBytes: cfg.MaxHeaderMegabytes << 28,
		},
	}
}

func (s *HTTPServer) Run() error {
	listener, err := reuseport.NewReusablePortListener(protocol, s.HttpServer.Addr)
	if err != nil {
		return err
	}

	return s.HttpServer.Serve(listener)
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}
