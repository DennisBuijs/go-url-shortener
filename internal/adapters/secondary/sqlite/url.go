package sqlite

import (
	"database/sql"
	"errors"
	"url-shortener/internal/core/domain"
)

type UrlRepository struct {
	db *sql.DB
}

func NewUrlRepository() (*UrlRepository, error) {
	return &UrlRepository{}, nil
}

func (repository *UrlRepository) Add(url domain.Url) error {
	return errors.New("not implemented yet")
}

func (repository *UrlRepository) FindById(id int) (*domain.Url, error) {
	return nil, errors.New("not implemented yet")
}
func (repository *UrlRepository) FindByCode(code string) (*domain.Url, error) {
	return nil, errors.New("not implemented yet")
}
