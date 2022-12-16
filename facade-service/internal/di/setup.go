package di

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/facade-service/internal/client"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/controller/http"
	"github.com/tuxoo/smart-loader/facade-service/internal/domain/repository"
	"github.com/tuxoo/smart-loader/facade-service/internal/server"
	"go.uber.org/fx"
)

var configModule = fx.Options(
	fx.Provide(config.NewHTTPConfig),
	fx.Provide(config.NewPostgresConfig),
	fx.Provide(config.NewNatsConfig),
	fx.Provide(config.NewAppConfig),
)

var repositoryModule = fx.Options(
	fx.Provide(repository.NewPostgresDB),
	fx.Provide(provideUserRepository),
	fx.Provide(provideJobRepository),
	fx.Provide(provideJobStageRepository),
)

var serviceModule = fx.Options(
	fx.Provide(provideUserService),
	fx.Provide(provideJobService),
	fx.Provide(provideJobStageService),
)

var utilModule = fx.Options(
	fx.Provide(provideHasher),
	fx.Provide(provideTokenManager),
)

var App = fx.New(
	configModule,
	repositoryModule,
	serviceModule,
	utilModule,
	fx.Provide(http.NewHandler),
	fx.Provide(server.NewHTTPServer),
	fx.Provide(client.NewNatsClient),
	fx.Invoke(registerPostgresHooks),
	fx.Invoke(registerServerHooks),
	fx.Invoke(registerNatsHooks),
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
