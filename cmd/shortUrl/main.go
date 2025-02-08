package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/dennnniska/shortUrl/internal/app"
	"github.com/dennnniska/shortUrl/internal/config"
	"github.com/dennnniska/shortUrl/internal/storage"
	inmemory "github.com/dennnniska/shortUrl/internal/storage/in-memory"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting", slog.String("env", cfg.Env))

	application := app.New(log, cfg)

	go application.HTTPSrv.MustRun()

	go application.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application GRPC", slog.String("signal", sign.String()))

	application.GRPCSrv.Stop()

	log.Info("application stopped GRPC")

	log.Info("stopping server HTTP", slog.String("signal", sign.String()))

	application.HTTPSrv.Stop()

	log.Info("server stopped HTTP")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func NewStorage(inMemory bool) (storage.Storage, error) {
	return inmemory.New(), nil
}
