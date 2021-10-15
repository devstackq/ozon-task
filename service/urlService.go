package service

import (
	"github.com/devstackq/ozon/entity"
	"github.com/devstackq/ozon/repository"
)

type UrlService interface {
	Randomaizer() (error, string)
	CreateShortHost(*entity.UrlData) (error, string)
	IsValidUrl(*entity.UrlData) (error, string)
	IsUniqUrl() bool
	SaveUrlDB(*entity.UrlData) error
	GetUrlDB(*entity.UrlData) string
}

var (
	repo repository.UrlRepository
)

type service struct{}

func NewUrlService(repository repository.UrlRepository) UrlService {
	repo = repository
	return &service{}
}

func (*service) Randomaizer() (error, string) {
	return nil, ""
}

func (*service) CreateShortHost(url *entity.UrlData) (error, string) {
	return nil, ""
}

func (*service) IsValidUrl(url *entity.UrlData) (error, string) {
	return nil, ""
	// repo.Save()
}

func (*service) IsUniqUrl() bool {
	return true
}

func (*service) SaveUrlDB(url *entity.UrlData) error {
	return nil
}

func (*service) GetUrlDB(url *entity.UrlData) string {
	return ""
}
