package service

import "github.com/devstackq/ozon/entity"

type UrlService interface {
	Randomaizer() string
	CreateShortHost(*entity.UrlData) (error, string)
	IsValidUrl(*entity.UrlData) (error, string)
	IsUniqUrl() bool
	SaveUrlDB(*entity.UrlData) error
	GetUrlDB(*entity.UrlData) (string, error)
	GetShortDB(*entity.UrlData) (string, error)
}
