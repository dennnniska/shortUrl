package postgres

import "github.com/dennnniska/shortUrl/internal/storage"

type Storage struct {
	urls      map[string]string
	shortUrls map[string]string
}

func New() (storage.Storage, error) {
	urls := make(map[string]string)
	shortUrls := make(map[string]string)
	return &Storage{
		urls:      urls,
		shortUrls: shortUrls,
	}, nil
}

func (s *Storage) FoundUrl(shortUrl string) (string, bool) {
	url, ok := s.shortUrls[shortUrl]
	return url, ok
}

func (s *Storage) FoundShortUrl(url string) (string, bool) {
	shortURL, ok := s.urls[url]
	return shortURL, ok
}

func (s *Storage) SaveUrl(url, shortUrl string) {
	s.urls[url] = shortUrl
	s.shortUrls[shortUrl] = url
}
