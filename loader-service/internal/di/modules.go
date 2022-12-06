package di

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/loader-service/internal/client"
	"github.com/tuxoo/smart-loader/loader-service/internal/config"
	"github.com/tuxoo/smart-loader/loader-service/internal/repository"
	"go.uber.org/fx"
)

var configModule = fx.Options(
	fx.Provide(config.NewPostgresConfig),
	fx.Provide(config.NewNatsConfig),
)

var repositoryModule = fx.Options(
	fx.Provide(repository.NewPostgresDB),
	fx.Provide(repository.NewRepositories),
)

var App = fx.New(
	configModule,
	repositoryModule,
	//fx.Provide(service.NewServices),
	fx.Provide(client.NewNatsClient),
	fx.Invoke(registerPostgresHooks),
	fx.Invoke(registerNatsHooks),
)

func registerPostgresHooks(lifecycle fx.Lifecycle, pool *repository.PostgresDB) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				if err := pool.Connect(); err != nil {
					logrus.Fatalf("error initializing postgres: %s", err.Error())
					return err
				}
				return nil
			},
			OnStop: func(context.Context) error {
				pool.Disconnect()
				return nil
			},
		},
	)
}

func registerNatsHooks(lifecycle fx.Lifecycle, client *client.NatsClient) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				if err := client.Connect(); err != nil {
					logrus.Fatalf("error connecting nats: %s", err.Error())
					return err
				}
				return nil
			},
			OnStop: func(context.Context) error {
				client.Disconnect()
				return nil
			},
		},
	)
}
