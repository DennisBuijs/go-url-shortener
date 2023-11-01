package use_cases

import (
	"url-shortener/internal/core/domain"
	"url-shortener/internal/ports"
)

type CreateUrlUseCase struct {
	UrlRepository ports.UrlRepository
}

func NewCreateUrlUseCase(urlRepository ports.UrlRepository) *CreateUrlUseCase {
	return &CreateUrlUseCase{UrlRepository: urlRepository}
}

func (useCase *CreateUrlUseCase) CreateUrl(url string) domain.Url {
	return domain.Url{
		Url:  url,
		Code: domain.GenerateShortCode(),
	}
}
