package grpcApp

import (
	"fmt"
	"log/slog"
	"net"

	shortUrlGRPC "github.com/dennnniska/shortUrl/internal/grpc/shortUrl"
	"github.com/dennnniska/shortUrl/internal/lib/service"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int, service service.ServiceShortUrl) *App {
	gRPCServer := grpc.NewServer()

	shortUrlGRPC.RegisterGRPCServer(gRPCServer, service, log)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcApp.Run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Start gRPC server", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	const op = "grpcApp.Stop"

	a.log.With(slog.String("op", op)).
		Info("stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
