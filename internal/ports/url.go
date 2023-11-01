package ports

import (
	"url-shortener/internal/core/domain"
)

type UrlRepository interface {
	Add(url domain.Url)
	FindById(id uint64) (*domain.Url, error)
	FindByCode(code string) (*domain.Url, error)
}
