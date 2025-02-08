package storage

type Storage interface {
	FoundUrl(shortUrl string) (string, bool)
	FoundShortUrl(url string) (string, bool)
	SaveUrl(url, shortUrl string)
}
