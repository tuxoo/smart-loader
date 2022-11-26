package app

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	fmt.Println(`
 ####~~##~~~#~~####~~#####~~######~~~~##~~~~~~####~~~####~~#####~~#####~#####
##~~~~~###~##~##~~##~##~~##~~~##~~~~~~##~~~~~##~~##~##~~##~##~~##~##~~~~##~~##
 ####~~##~#~#~######~#####~~~~##~~~~~~##~~~~~##~~##~######~##~~##~####~~#####
~~~~##~##~~~#~##~~##~##~~##~~~##~~~~~~##~~~~~##~~##~##~~##~##~~##~##~~~~##~~##
 ####~~##~~~#~##~~##~##~~##~~~##~~~~~~######~~####~~##~~##~#####~~#####~##~~##
	`)

	cfg, err := config.NewConfig(configPath)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	postgresDB, err := config.NewPostgresPool(config.PostgresConfig{
		Host:            cfg.Postgres.Host,
		Port:            cfg.Postgres.Port,
		DB:              cfg.Postgres.DB,
		User:            cfg.Postgres.User,
		Password:        cfg.Postgres.Password,
		MaxConns:        cfg.Postgres.MaxConns,
		MinConns:        cfg.Postgres.MinConns,
		MaxConnLifetime: cfg.Postgres.MaxConnLifetime,
		MaxConnIdleTime: cfg.Postgres.MaxConnIdleTime,
	})
	if err != nil {
		logrus.Fatalf("error initializing postgres: %s", err.Error())
	}
	defer postgresDB.Close()

	httpHandlers := InitHa
	httpServer := server.NewHTTPServer(cfg, httpHandlers.Init(cfg.HTTP))

	go func() {
		if err := httpServer.Run(); err != nil {
			logrus.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logrus.Printf("SMART LOADER facade application has been started on :%s port", cfg.HTTP.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("IDLER facade application shutting down")

	if err := httpServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on http server shutting down: %s", err.Error())
	}
}
