package shorturl

import (
	"context"

	shortUrl "github.com/dennnniska/shortUrl/internal/grpc/shortUrl/protofile"
	"github.com/dennnniska/shortUrl/internal/lib/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const lenSortURL = 10

type serverAPI struct {
	shortUrl.UnimplementedShortUrlServer
	service service.ServiceShortUrl
}

func RegisterGRPCServer(gRPC *grpc.Server, service service.ServiceShortUrl) {
	shortUrl.RegisterShortUrlServer(gRPC, &serverAPI{service: service})
}

func (s *serverAPI) GET(ctx context.Context, req *shortUrl.GetRequest) (*shortUrl.GetResponse, error) {
	URL, code, err := s.service.Get(req.GetShortUrl())
	if err != nil {
		return nil, status.Error(code, err.Error())
	}
	return &shortUrl.GetResponse{
		Url: URL,
	}, nil
}
func (s *serverAPI) Post(ctx context.Context, req *shortUrl.PostRequest) (*shortUrl.PostResponse, error) {
	shortURL, code, err := s.service.Post(req.GetUrl())
	if err != nil {
		return nil, status.Error(code, err.Error())
	}

	return &shortUrl.PostResponse{
		ShortUrl: shortURL,
	}, nil
}
