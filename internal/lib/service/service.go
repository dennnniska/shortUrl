package service

import (
	"fmt"
	"net/url"

	"github.com/dennnniska/shortUrl/internal/lib/random"
	"github.com/dennnniska/shortUrl/internal/storage"
	"google.golang.org/grpc/codes"
)

const lenSortURL = 10

type ServiceShortUrl interface {
	Post(URL string) (string, codes.Code, error)
	Get(shortURL string) (string, codes.Code, error)
}

type ShortUrl struct {
	storage storage.Storage
}

func (s *ShortUrl) AddStorage(storage storage.Storage) {
	s.storage = storage
}

func (s *ShortUrl) Post(URL string) (string, codes.Code, error) {
	if URL == "" {
		return "", codes.InvalidArgument, fmt.Errorf("URL is required")
	}
	if _, err := url.ParseRequestURI(URL); err != nil {
		return "", codes.InvalidArgument, fmt.Errorf("invalid URL")
	}

	shortURLExist := true
	var shortURL string

	for shortURLExist {
		shortURL = random.RandStringBytesMaskImpr(lenSortURL)
		_, shortURLExist, _ = s.storage.FoundUrl(shortURL)
	}
	err := s.storage.SaveUrl(URL, shortURL)
	if err != nil {
		return "", codes.Aborted, err
	}
	return shortURL, codes.OK, nil
}

func (s *ShortUrl) Get(shortURL string) (string, codes.Code, error) {
	if shortURL == "" {
		return "", codes.InvalidArgument, fmt.Errorf("shortUrl is required")
	}

	if len(shortURL) != lenSortURL {
		return "", codes.InvalidArgument, fmt.Errorf("length of shortUrl is not equal to 10")
	}
	URL, ok, err := s.storage.FoundUrl(shortURL)
	if err != nil {
		return "", codes.Aborted, err
	}
	if !ok {
		return "", codes.NotFound, fmt.Errorf("not found URL")
	}
	return URL, codes.OK, nil
}
