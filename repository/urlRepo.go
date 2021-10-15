package repository

import "github.com/devstackq/ozon/entity"

type UrlRepository interface {
	Create(url *entity.UrlData) (string, error)
	Get(url *entity.UrlData) (string, error)
}
