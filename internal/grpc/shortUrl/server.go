package shorturl

import (
	"context"
	"net/url"

	shortUrl "github.com/dennnniska/shortUrl/internal/grpc/shortUrl/protofile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const lenSortURL = 10

type serverAPI struct {
	shortUrl.UnimplementedShortUrlServer
}

func RegisterGRPCServer(gRPC *grpc.Server) {
	shortUrl.RegisterShortUrlServer(gRPC, &serverAPI{})
}

func (s *serverAPI) GET(ctx context.Context, req *shortUrl.GetRequest) (*shortUrl.GetResponse, error) {
	if req.GetShortUrl() == "" {
		return nil, status.Error(codes.InvalidArgument, "shortUrl is required")
	}

	if len(req.GetShortUrl()) != lenSortURL {
		return nil, status.Error(codes.InvalidArgument, "length of shortUrl is not equal to 10")
	}
	return &shortUrl.GetResponse{
		Url: req.GetShortUrl(),
	}, nil
}
func (s *serverAPI) Post(ctx context.Context, req *shortUrl.PostRequest) (*shortUrl.PostResponse, error) {
	if req.GetUrl() == "" {
		return nil, status.Error(codes.InvalidArgument, "URL is required")
	}
	if _, err := url.ParseRequestURI(req.GetUrl()); err != nil {
		return nil, status.Error(codes.InvalidArgument, "URL is required")
	}

	return &shortUrl.PostResponse{
		ShortUrl: req.GetUrl(),
	}, nil
}
