package cache

type UrlCache interface {
	SaveRedis( string,string) error 
	GetRedis( string) string
}
