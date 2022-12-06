package di

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/controller/http"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
	"github.com/tuxoo/smart-loader/facade-service/internal/server"
	"github.com/tuxoo/smart-loader/facade-service/internal/service"
	"go.uber.org/fx"
)

var configModule = fx.Options(
	fx.Provide(config.NewHTTPConfig),
	fx.Provide(config.NewPostgresConfig),
	fx.Provide(config.NewNatsConfig),
)

var repositoryModule = fx.Options(
	fx.Provide(repository.NewPostgresPool),
	fx.Provide(repository.NewRepositories),
)

var App = fx.New(
	configModule,
	repositoryModule,
	fx.Provide(service.NewServices),
	fx.Provide(http.NewHandler),
	fx.Provide(http.NewRouter),
	fx.Provide(server.NewHTTPServer),
	fx.Invoke(registerServerHooks),
	fx.Invoke(registerPostgresHooks),
)

func registerServerHooks(lifecycle fx.Lifecycle, s *server.HTTPServer) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				// TODO: return error
				go func() {
					if err := s.Run(); err != nil {
						logrus.Errorf("error occurred while running http server: %s\n", err.Error())
					}
				}()
				logrus.Printf("SMART LOADER facade application has been started [%s]", s.HttpServer.Addr)
				return nil
			},
			OnStop: func(context.Context) error {
				// TODO: return error
				logrus.Print("SMART LOADER application is shutting down")
				if err := s.Shutdown(context.Background()); err != nil {
					logrus.Errorf("error occured on http server shutting down: %s", err.Error())
				}
				return nil
			},
		},
	)
}

func registerPostgresHooks(lifecycle fx.Lifecycle, pool *pgxpool.Pool) {
	lifecycle.Append(
		fx.Hook{
			OnStop: func(context.Context) error {
				// TODO: return error
				pool.Close()
				return nil
			},
		},
	)
}
