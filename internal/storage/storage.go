package storage

type Storage interface {
	FoundUrl(shortUrl string) (string, bool, error)
	FoundShortUrl(url string) (string, bool, error)
	SaveUrl(url, shortUrl string) error
}
