package memory

import (
	"errors"
	"fmt"
	"strconv"
	"url-shortener/internal/core/domain"
)

type UrlRepository struct {
}

func NewUrlRepository() (*UrlRepository, error) {
	return &UrlRepository{}, nil
}

var urls []domain.Url

func (repository *UrlRepository) Add(url domain.Url) error {
	urls = append(urls, url)
	fmt.Printf("URL '%s' added with code '%s'\n", url.Url, url.Code)
	return nil
}

func (repository *UrlRepository) FindByCode(code string) (*domain.Url, error) {
	for _, url := range urls {
		if url.Code == code {
			return &url, nil
		}
	}

	return nil, errors.New("url with code " + code + " not found")
}

func (repository *UrlRepository) FindById(id int) (*domain.Url, error) {
	for _, url := range urls {
		if url.Id == id {
			return &url, nil
		}
	}

	return nil, errors.New("url with id " + strconv.Itoa(id) + " not found")
}
