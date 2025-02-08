package app

import (
	"log/slog"

	grpcApp "github.com/dennnniska/shortUrl/internal/app/grpc"
	httpApp "github.com/dennnniska/shortUrl/internal/app/http"
	"github.com/dennnniska/shortUrl/internal/config"
	"github.com/dennnniska/shortUrl/internal/lib/service"
	"github.com/dennnniska/shortUrl/internal/storage"
	inmemory "github.com/dennnniska/shortUrl/internal/storage/in-memory"
	"github.com/dennnniska/shortUrl/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcApp.App
	HTTPSrv *httpApp.App
}

func New(log *slog.Logger, cfg *config.Config) *App {
	storage, err := NewStorage(cfg.InMemoryStorage)
	if err != nil {
		log.Error("failed to start storage")
	}

	service := &service.ShortUrl{
		Storage: storage,
	}

	grpcApp := grpcApp.New(log, cfg.GRPC.Port, service)
	httpApp := httpApp.New(log, cfg.Http.Address, cfg.Http.Timeout, cfg.Http.IdleTimeout, service)
	return &App{
		GRPCSrv: grpcApp,
		HTTPSrv: httpApp,
	}
}

func NewStorage(inMemory bool) (storage.Storage, error) {
	if inMemory {
		return inmemory.New(), nil
	}
	return postgres.New()
}
