package shorturl

import (
	"context"

	shortUrl "github.com/dennnniska/shortUrl/internal/grpc/shortUrl/protofile"
	"google.golang.org/grpc"
)

type serverAPI struct {
	shortUrl.UnimplementedShortUrlServer
}

func RegisterGRPCServer(gRPC *grpc.Server) {
	shortUrl.RegisterShortUrlServer(gRPC, &serverAPI{})
}

func (s *serverAPI) GET(ctx context.Context, req *shortUrl.GetRequest) (*shortUrl.GetResponse, error) {
	return &shortUrl.GetResponse{
		Url: req.GetShortUrl(),
	}, nil
}
func (s *serverAPI) Post(ctx context.Context, req *shortUrl.PostRequest) (*shortUrl.PostResponse, error) {
	panic("implement me")
}
