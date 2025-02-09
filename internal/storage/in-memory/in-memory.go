package inmemory

import "github.com/dennnniska/shortUrl/internal/storage"

type Storage struct {
	urls      map[string]string
	shortUrls map[string]string
}

func New() storage.Storage {
	urls := make(map[string]string)
	shortUrls := make(map[string]string)
	return &Storage{
		urls:      urls,
		shortUrls: shortUrls,
	}
}

func (s *Storage) FoundUrl(shortUrl string) (string, bool, error) {
	url, ok := s.shortUrls[shortUrl]
	return url, ok, nil
}

func (s *Storage) FoundShortUrl(url string) (string, bool, error) {
	shortURL, ok := s.urls[url]
	return shortURL, ok, nil
}

func (s *Storage) SaveUrl(url, shortUrl string) error {
	s.urls[url] = shortUrl
	s.shortUrls[shortUrl] = url
	return nil
}
