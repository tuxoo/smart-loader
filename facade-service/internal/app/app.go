package app

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/tuxoo/smart-loader/facade-service/internal/config"
	"github.com/tuxoo/smart-loader/facade-service/internal/controller/http"
	"github.com/tuxoo/smart-loader/facade-service/internal/dependency"
	"github.com/tuxoo/smart-loader/facade-service/internal/repository"
	"github.com/tuxoo/smart-loader/facade-service/internal/server"
	"github.com/tuxoo/smart-loader/facade-service/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	fmt.Println(`
 ####~~##~~~#~~####~~#####~~######~~~~##~~~~~~####~~~####~~#####~~#####~#####
##~~~~~###~##~##~~##~##~~##~~~##~~~~~~##~~~~~##~~##~##~~##~##~~##~##~~~~##~~##
 ####~~##~#~#~######~#####~~~~##~~~~~~##~~~~~##~~##~######~##~~##~####~~#####
~~~~##~##~~~#~##~~##~##~~##~~~##~~~~~~##~~~~~##~~##~##~~##~##~~##~##~~~~##~~##
 ####~~##~~~#~##~~##~##~~##~~~##~~~~~~######~~####~~##~~##~#####~~#####~##~~##
	`)

	cfg, err := dependency.InitConfig()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := config.NewPostgresPool(config.PostgresConfig{
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
	defer db.Close()

	nc, err := nats.Connect("nats://host.docker.internal:4222")
	if err != nil {
		logrus.Fatalf("error initializing nats: %s", err.Error())
	}

	err = nc.Publish("foo", []byte("Hello World"))

	repositories := repository.NewRepositories(db)

	services := service.NewServices(service.ServicesDeps{
		Repositories: repositories,
	})

	httpHandlers := http.NewHandler(services.JobService)
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
