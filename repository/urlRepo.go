package repository

import "github.com/devstackq/ozon/entity"

type UrlRepository interface {
	Create(url *entity.UrlData) error
	GetByUrl(url *entity.UrlData) (string, error)
	GetByShort(url *entity.UrlData) (string, error)
}
