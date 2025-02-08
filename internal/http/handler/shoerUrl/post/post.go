package post

import (
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	render "github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	URL string `json:"URL" validate:"required"`
}

type Response struct {
	Status   string `json:"status"`
	Error    string `json:"error,omitempty"`
	ShortURL string `json:"shortURL"`
}

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.shortUrl.get.New"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")

			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "empty request",
			})

			return
		}
		if err != nil {
			log.Error("failed to decode request body", slog.Attr{
				Key:   "error",
				Value: slog.StringValue(err.Error()),
			})

			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "failed to decode request",
			})

			return
		}
		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("invalid request", slog.Attr{
				Key:   "error",
				Value: slog.StringValue(err.Error()),
			})

			render.JSON(w, r, Response{
				Status: "Error",
				Error:  "invalid request",
			})

			return
		}

		render.JSON(w, r, Response{
			Status:   "OK",
			ShortURL: req.URL,
		})
	}
}
