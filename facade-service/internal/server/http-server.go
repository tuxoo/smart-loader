package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kavu/go_reuseport"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"net/http"
)

const protocol = "tcp"

type HTTPServer struct {
	HttpServer *http.Server
}

func NewHTTPServer(cfg *config.HTTPConfig, mux *mux.Router) *HTTPServer {
	return &HTTPServer{
		HttpServer: &http.Server{
			Addr:           fmt.Sprintf(":%s", cfg.Port),
			Handler:        mux,
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
