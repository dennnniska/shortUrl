package app

import (
	"log/slog"
	"time"

	grpcApp "github.com/dennnniska/shortUrl/internal/app/grpc"
	httpApp "github.com/dennnniska/shortUrl/internal/app/http"
)

type App struct {
	GRPCSrv *grpcApp.App
	HTTPSrv *httpApp.App
}

func New(log *slog.Logger, grpcPort int, address string, timeout, idleTimeout time.Duration) *App {
	grpcApp := grpcApp.New(log, grpcPort)
	httpApp := httpApp.New(log, address, timeout, idleTimeout)
	return &App{
		GRPCSrv: grpcApp,
		HTTPSrv: httpApp,
	}
}
