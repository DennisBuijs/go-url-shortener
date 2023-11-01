package ports

import (
	"url-shortener/internal/core/domain"
)

type UrlRepository interface {
	Add(url domain.Url) error
	FindById(id int) (*domain.Url, error)
	FindByCode(code string) (*domain.Url, error)
}
