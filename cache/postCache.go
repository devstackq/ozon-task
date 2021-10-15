package cache

import "github.com/devstackq/ozon/entity"

type PostCache interface {
	Set(key string)  *entity.UrlData
	Get(key string)  *entity.UrlData
}