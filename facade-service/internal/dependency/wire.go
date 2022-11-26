//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tuxoo/smart-loader/facade-service/internal/controller/http"
)

func InitHandler() (*http.Handler, error) {
	wire.Build(
		http.NewHandler,
	)
	return &http.Handler{}, nil
}
