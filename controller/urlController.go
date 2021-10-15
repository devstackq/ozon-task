package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/devstackq/ozon/cache"
	"github.com/devstackq/ozon/entity"
	"github.com/devstackq/ozon/service"
)

type controller struct{}

var (
	urlService service.UrlService
	urlCache   cache.UrlCache
)

type UrlController interface {
	GenerateNewLink(res http.ResponseWriter, req *http.Request)
	GetLinkByShortLink(res http.ResponseWriter, req *http.Request)
}

func NewUrlController(service service.UrlService, cache cache.UrlCache) UrlController {
	urlService = service
	urlCache = cache
	return &controller{}
}

func (*controller) GenerateNewLink(res http.ResponseWriter, req *http.Request) {
	var url entity.UrlData
	err := json.NewDecoder(req.Body).Decode(&url)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.New("Error unmarshalling data"))
		return
	}

	if err, host := urlService.IsValidUrl(&url); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.New("Error: not valid url"))
		return
	} else {
		if ok := urlService.IsUniqUrl(); !ok {
			//get Redis || db full url key, if not exist return true, else return by key redis || db short url

			// urlService.GetUrlRedis(&url)
			urlService.GetUrlDB(&url)
			return
		} else {
			url.Host = host
			//create new url
			if err, shortHost := urlService.CreateShortHost(&url); err != nil {
				res.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(res).Encode(err.Error())
				return
			} else {
				if err, shortUrl := urlService.Randomaizer(); err != nil {
					//save Cache, then Db
					res.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(res).Encode(err.Error())
					return
				} else {
					url.ShortUrl = shortHost + shortUrl
					//saveRedis(&url)
					urlCache.SaveRedis(url.ShortUrl, url.Url)
					urlService.SaveUrlDB(&url)
				}
			}
		}
	}
}

func (*controller) GetLinkByShortLink(res http.ResponseWriter, req *http.Request) {

}
