package postgres

import (
	"fmt"
	"os"

	"github.com/dennnniska/shortUrl/internal/config"
	"github.com/dennnniska/shortUrl/internal/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func New(cfg *config.Postgres) (storage.Storage, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Dbname, os.Getenv("DB_PASSWORD"), cfg.Sslmode))
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS shortUrl(
		id INTEGER PRIMARY KEY,
		url TEXT NOT NULL UNIQ
		shortUrl TEXT NOT NULL UNIQ
	)
	`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}
	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) FoundUrl(shortUrl string) (string, bool, error) {
	fmt.Print("fsafasf")
	stmt, err := s.db.Prepare("SELECT url FROM url WHERE shortUrl = ?")
	if err != nil {
		return "", false, err
	}
	fmt.Print("fsafasf")
	var resURL string

	err = stmt.QueryRow(shortUrl).Scan(&resURL)
	if err != nil {
		return "", false, err
	}
	if resURL == "" {
		return resURL, false, nil
	}
	return resURL, true, nil
}

func (s *Storage) FoundShortUrl(url string) (string, bool, error) {
	stmt, err := s.db.Prepare("SELECT shortUrl FROM url WHERE url = ?")
	if err != nil {
		return "", false, err
	}

	var resURL string

	err = stmt.QueryRow(url).Scan(&resURL)
	if err != nil {
		return "", false, err
	}
	if resURL == "" {
		return resURL, false, nil
	}
	return resURL, true, nil
}

func (s *Storage) SaveUrl(url, shortUrl string) error {
	stmt, err := s.db.Prepare("INSERT INTO shortUrl(url, shortUrl) VALUES(?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(url, shortUrl)
	if err != nil {
		return err
	}
	return nil
}
