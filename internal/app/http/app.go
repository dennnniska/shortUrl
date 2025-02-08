package httpApp

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	get "github.com/dennnniska/shortUrl/internal/http/handler/shoerUrl/Get"
	"github.com/dennnniska/shortUrl/internal/http/handler/shoerUrl/post"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	log        *slog.Logger
	hTTPServer *http.Server
	address    string
}

func New(log *slog.Logger, address string, timeout, idleTimeout time.Duration) *App {
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Post("/post", post.New(log))
	router.Get("/get", get.New(log))

	HTTPServer := &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  idleTimeout,
	}
	return &App{
		log:        log,
		hTTPServer: HTTPServer,
		address:    address,
	}
}

func (a *App) MustRun() {
	a.log.Info("starting HTTP server", slog.String("address", a.address))
	if err := a.hTTPServer.ListenAndServe(); err != nil {
		a.log.Error("failed to start HTTP server")
	}
}

func (a *App) Stop() {
	const op = "HTTPApp.Stop"
	a.log.With(slog.String("op", op)).
		Info("stopping HTTP server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.hTTPServer.Shutdown(ctx); err != nil {
		a.log.Error("failed to stop HTTP server", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(err.Error()),
		})
		return
	}
}
