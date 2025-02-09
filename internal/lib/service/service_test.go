package service_test

import (
	"fmt"
	"testing"

	"github.com/dennnniska/shortUrl/internal/lib/service"
	mock_storage "github.com/dennnniska/shortUrl/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
)

func NewShortUrl(t *testing.T) *service.ShortUrl {
	t.Helper()

	storage := mock_storage.MockStorage{}

	var service service.ShortUrl
	service.AddStorage(&storage)

	return &service
}

func TestUrlIsRequired(t *testing.T) {
	service := NewShortUrl(t)

	url := ""
	_, code, err := service.Post(url)

	assert.Equal(t, code, codes.InvalidArgument)
	assert.Equal(t, err, fmt.Errorf("URL is required"))
}

func TestUrlInvalid(t *testing.T) {
	service := NewShortUrl(t)

	url := "asfafs"
	_, code, err := service.Post(url)

	assert.Equal(t, code, codes.InvalidArgument)
	assert.Equal(t, err, fmt.Errorf("invalid URL"))
}

func TestShortUrlIsRequired(t *testing.T) {
	service := NewShortUrl(t)

	shortUrl := ""
	_, code, err := service.Get(shortUrl)

	assert.Equal(t, code, codes.InvalidArgument)
	assert.Equal(t, err, fmt.Errorf("shortUrl is required"))
}

func TestShortUrlIsLenNotTen(t *testing.T) {
	service := NewShortUrl(t)

	shortUrl := "а"
	_, code, err := service.Get(shortUrl)

	assert.Equal(t, code, codes.InvalidArgument)
	assert.Equal(t, err, fmt.Errorf("length of shortUrl is not equal to 10"))
}
