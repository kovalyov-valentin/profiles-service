package main

import (
	"context"
	"github.com/kovalyov-valentin/profiles-service/internal"
	"github.com/kovalyov-valentin/profiles-service/internal/config"
	"github.com/kovalyov-valentin/profiles-service/internal/handler"
	"github.com/kovalyov-valentin/profiles-service/internal/repository"
	"github.com/kovalyov-valentin/profiles-service/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.CtxTimeout)
	defer cancel()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("DEBUG messages are enabled")

	db, err := repository.NewPostgresDB(cfg.DB)
	if err != nil {
		logrus.Fatalf("failes to connect postgres db: %s", err.Error())
	}
	defer db.Close()

	repos := repository.NewRepository(db)
	services := service.NewService(repos, cfg)
	handlers := handler.NewHandler(services, cfg)

	srv := new(internal.Server)

	serverErrors := make(chan error, 1)
	go func() {
		logrus.Printf("Start listen http service on %s at %s\n", cfg.Address, time.Now().Format(time.DateTime))
		err := srv.Run(cfg.HTTPServer, handlers.InitRoutes())
		if err != nil {
			logrus.Printf("shutting down the server: %s\n", cfg.Address)
		}
		serverErrors <- err
	}()

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGTERM, syscall.SIGINT)
	select {
	case err := <-serverErrors:
		logrus.Printf("error starting server: %v\n", err)
	case <-osSignal:
		logrus.Println("start shutdown...")
		if err := srv.Shutdown(ctx); err != nil {
			logrus.Printf("graceful shutdown error: %v\n", err)
			os.Exit(1)
		}
	}
	logrus.Info("server stopped")
}
