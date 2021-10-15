package cache

import (
	"time"

	"github.com/go-redis/redis/v7"
)

type Cache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, expires time.Duration) UrlCache {
	return &Cache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (c *Cache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.host,
		Password: "",
		DB:       c.db,
	})
}

func (c *Cache) SaveRedis(key string, url string) error {
	client := c.getClient()
	//save like key url, value - short
	client.Set(key, url, c.expires)
	//case 2 save reverse, key - shor, value - url
	client.Set(url, key, c.expires)
	return nil
}

func  (c *Cache) GetRedis(key string) string{ return "" }
