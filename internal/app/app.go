package app

import (
	"log/slog"

	grpcApp "github.com/dennnniska/shortUrl/internal/app/grpc"
)

type App struct {
	GRPCSrv *grpcApp.App
}

func New(log *slog.Logger, grpcPort int) *App {
	grpcApp := grpcApp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
