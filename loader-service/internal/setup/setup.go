package setup

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/loader-service/internal/client"
	"github.com/tuxoo/smart-loader/loader-service/internal/config"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/loader-service/internal/handler"
	"go.uber.org/fx"
)

var configModule = fx.Options(
	fx.Provide(config.NewAppConfig),
	fx.Provide(config.NewPostgresConfig),
	fx.Provide(config.NewNatsConfig),
)

var repositoryModule = fx.Options(
	fx.Provide(repository.NewPostgresDB),
	fx.Provide(provideJobRepository),
	fx.Provide(provideJobStageRepository),
	fx.Provide(provideDownloadRepository),
	fx.Provide(provideLockRepository),
)

var serviceModule = fx.Options(
	fx.Provide(provideJobService),
	fx.Provide(provideJobStageService),
	fx.Provide(provideDownloadService),
	fx.Provide(provideLockService),
)

var handlerModule = fx.Options(
	fx.Provide(provideJobHandler),
)

var utilModule = fx.Options(
	fx.Provide(provideHasher),
	fx.Provide(provideDownloader),
)

var App = fx.New(
	configModule,
	repositoryModule,
	serviceModule,
	handlerModule,
	utilModule,
	fx.Provide(client.NewNatsClient),
	fx.Invoke(registerNatsHooks),
	fx.Invoke(registerPostgresHooks),
)

func registerPostgresHooks(lifecycle fx.Lifecycle, pool *repository.PostgresDB) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				err := pool.Connect(ctx)
				if err != nil {
					return err
				}

				return nil
			},
			OnStop: func(_ context.Context) error {
				pool.Disconnect()
				return nil
			},
		},
	)
}

func registerNatsHooks(lifecycle fx.Lifecycle, client *client.NatsClient, handler *handler.JobHandler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				if err := client.Connect(); err != nil {
					logrus.Fatalf("error connecting nats: %s", err.Error())
					return err
				}

				err := handler.Handle()
				if err != nil {
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
