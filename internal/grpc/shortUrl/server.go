package shorturl

import (
	"context"
	"log/slog"

	shortUrl "github.com/dennnniska/shortUrl/internal/grpc/shortUrl/protofile"
	"github.com/dennnniska/shortUrl/internal/lib/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const lenSortURL = 10

type serverAPI struct {
	shortUrl.UnimplementedShortUrlServer
	service service.ServiceShortUrl
	log     *slog.Logger
}

func RegisterGRPCServer(gRPC *grpc.Server, service service.ServiceShortUrl, log *slog.Logger) {
	shortUrl.RegisterShortUrlServer(gRPC, &serverAPI{service: service, log: log})
}

func (s *serverAPI) GET(ctx context.Context, req *shortUrl.GetRequest) (*shortUrl.GetResponse, error) {
	URL, code, err := s.service.Get(req.GetShortUrl())
	if err != nil {
		s.log.Error("storage error", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(err.Error()),
		})
		return nil, status.Error(code, err.Error())
	}
	return &shortUrl.GetResponse{
		Url: URL,
	}, nil
}
func (s *serverAPI) Post(ctx context.Context, req *shortUrl.PostRequest) (*shortUrl.PostResponse, error) {
	shortURL, code, err := s.service.Post(req.GetUrl())
	if err != nil {
		s.log.Error("storage error", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(err.Error()),
		})
		return nil, status.Error(code, err.Error())
	}

	return &shortUrl.PostResponse{
		ShortUrl: shortURL,
	}, nil
}
