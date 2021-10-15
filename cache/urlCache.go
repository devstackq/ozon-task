package cache

import "github.com/devstackq/ozon/entity"

type UrlCache interface {
	SaveRedis( string,string) error 
	GetRedis( string) *entity.UrlData
}
