package inmemory_test

import (
	"testing"

	inmemory "github.com/dennnniska/shortUrl/internal/storage/in-memory"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryStorage(t *testing.T) {
	storage := inmemory.New()

	url := "https://ya.ru/"
	shortUrl := "AAA_Aa"

	storage.SaveUrl(url, shortUrl)

	tests := []struct {
		name      string
		url       string
		shortUrl  string
		findShort bool
		findUrl   bool
	}{
		{
			name:      "Url and shortUrl Exist",
			url:       "https://ya.ru/",
			shortUrl:  "AAA_Aa",
			findShort: true,
			findUrl:   true,
		},
		{
			name:      "Url and shortUrl not Exist",
			url:       "ya.ru",
			shortUrl:  "1111",
			findShort: false,
			findUrl:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, ok1, _ := storage.FoundShortUrl(tt.url)
			_, ok2, _ := storage.FoundUrl(tt.shortUrl)

			assert.Equal(t, ok1, tt.findShort)
			assert.Equal(t, ok2, tt.findUrl)
		})
	}
}
