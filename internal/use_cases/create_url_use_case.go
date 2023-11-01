package use_cases

import "url-shortener/internal/core/domain"

type CreateUrlUseCase struct{}

func (useCase *CreateUrlUseCase) CreateUrl(url string) domain.Url {
	return domain.Url{
		Url:  url,
		Code: domain.GenerateShortCode(),
	}
}
