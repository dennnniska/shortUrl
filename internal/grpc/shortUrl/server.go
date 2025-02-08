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

func (s *serverAPI) GET(context.Context, *shortUrl.GetRequest) (*shortUrl.GetResponse, error) {
	panic("implement me")
}
func (s *serverAPI) Post(context.Context, *shortUrl.PostRequest) (*shortUrl.PostResponse, error) {
	panic("implement me")
}
