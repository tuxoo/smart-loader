package setup

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/facade-service/internal/client"
	"github.com/tuxoo/smart-loader/facade-service/internal/controller/http"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/model/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/facade-service/internal/server"
	"go.uber.org/fx"
)

var configModule = fx.Options(
	fx.Provide(
		config.NewHTTPConfig,
		config.NewPostgresConfig,
		config.NewNatsConfig,
		config.NewMinioConfig,
		config.NewAppConfig,
	),
)

var repositoryModule = fx.Options(
	fx.Provide(repository.NewPostgresDB),
	fx.Provide(
		provideUserRepository,
		provideTokenRepository,
		provideJobRepository,
		provideJobStageRepository,
		provideDownloadRepository,
	),
)

var serviceModule = fx.Options(
	fx.Provide(
		provideUserService,
		provideTokenService,
		provideJobService,
		provideJobStageService,
		provideDownloadService,
		provideMinioService,
	),
)

var utilModule = fx.Options(
	fx.Provide(
		provideHasher,
		provideTokenManager,
	),
)

var App = fx.New(
	configModule,
	repositoryModule,
	serviceModule,
	utilModule,
	fx.Provide(http.NewHandler),
	fx.Provide(server.NewHTTPServer),
	fx.Provide(client.NewNatsClient),
	fx.Provide(client.NewMinioClient),
	fx.Invoke(
		registerPostgresHooks,
		registerServerHooks,
		registerNatsHooks,
		registerMinioHooks,
	),
)

func registerServerHooks(lifecycle fx.Lifecycle, s *server.HTTPServer) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				go func() {
					if err := s.Run(); err == nil {
						logrus.Errorf("error occurred while running http server: %s\n", err.Error())
					}
				}()

				logrus.Printf("SMART LOADER facade application has been started [%s]", s.HttpServer.Addr)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				logrus.Print("SMART LOADER application is shutting down")
				if err := s.Shutdown(ctx); err != nil {
					logrus.Errorf("error occured on http server shutting down: %s", err.Error())
					return err
				}
				return nil
			},
		},
	)
}

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

func registerNatsHooks(lifecycle fx.Lifecycle, client *client.NatsClient) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
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

func registerMinioHooks(lifecycle fx.Lifecycle, client *client.MinioClient) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				if err := client.Connect(ctx); err != nil {
					logrus.Fatalf("error connecting minio: %s", err.Error())
					return err
				}

				return nil
			},
		},
	)
}
