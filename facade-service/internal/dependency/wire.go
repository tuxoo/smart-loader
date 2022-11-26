//go:build wireinject
// +build wireinject

package dependency

import (
	"github.com/google/wire"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
)

func InitConfig() (*config.Config, error) {
	wire.Build(
		config.NewConfig,
	)
	return &config.Config{}, nil
}
