package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"url-shortener/internal/core/domain"
)

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) (*UrlRepository, error) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS url (id INTEGER PRIMARY KEY, url TEXT, code VARCHAR(36))")
	if err != nil {
		log.Fatal(err)
	}

	return &UrlRepository{db: db}, nil
}

func (repository *UrlRepository) Add(url domain.Url) error {
	query, err := repository.db.Prepare("INSERT INTO url (url, code) VALUES (?, ?)")
	if err != nil {
		return err
	}

	defer query.Close()

	_, err = query.Exec(url.Url, url.Code)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UrlRepository) FindById(id int) (*domain.Url, error) {
	return nil, errors.New("not implemented yet")
}

func (repository *UrlRepository) FindByCode(code string) (*domain.Url, error) {
	url := domain.Url{}
	err := repository.db.QueryRow("SELECT url FROM url WHERE code = ?", code).Scan(&url.Url)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No rows found.")
		} else {
			log.Fatal(err)
		}
	}

	return &url, nil
}
