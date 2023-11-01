package ports

import "url-shortener/internal/core/domain"

type UrlRepository interface {
	Add(url domain.Url)
	FindById(id uint64) (*domain.Url, error)
	FindByCode(code string) (*domain.Url, error)
}

type CreateUrlUseCase struct{}

func (useCase *CreateUrlUseCase) CreateUrl(url string) domain.Url {
	return domain.Url{
		Url:  url,
		Code: domain.GenerateShortCode(),
	}
}
