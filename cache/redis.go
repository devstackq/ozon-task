package cache

import "time"

type Cache struct {
	host string
	db int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) *Cache {
	return &Cache{
		host : host,
		db: db,
		expires: expires,
	}
}

func (c *Cache) getClient() *Cache {
return &Cache{}
}

func Set(key string){}