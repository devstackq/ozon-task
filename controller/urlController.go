package controller

import (
	"encoding/json"
	"errors"
	"log"
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

//rest : https://localhost:8000, json:{url : https://www.google.com/?search/dposak0932jdoisfojsa}

func (*controller) GenerateNewLink(res http.ResponseWriter, req *http.Request) {
	var url entity.UrlData
	err := json.NewDecoder(req.Body).Decode(&url)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.New("Error unmarshalling data"))
		return
	}

	if err, host := urlService.IsValidUrl(&url); err != nil || host == "" {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.New("Error: not valid url"))
		return
	} else {
		if urlFromRedis := urlCache.GetRedis(url.Url); urlFromRedis == "" {
			if urlFromDb, err := urlService.GetUrlDB(&url); urlFromDb == "" && err == nil {
				//create new short url, save redis & db
				log.Println(url.Host, host, 1)
				url.Host = host

				if err, shortHost := urlService.CreateShortHost(&url); err != nil {
					res.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(res).Encode(err.Error())
					return
				} else {
					if shortUrl := urlService.Randomaizer(); shortUrl == "" {
						res.WriteHeader(http.StatusInternalServerError)
						json.NewEncoder(res).Encode(err.Error())
						return
					} else {
						url.ShortUrl = shortHost + shortUrl

						urlCache.SaveRedis(url.ShortUrl, url.Url)
						urlService.SaveUrlDB(&url)
						res.WriteHeader(200)
						res.Write([]byte(url.ShortUrl))
					}
				}
			} else {
				res.WriteHeader(200)
				res.Write([]byte(urlFromDb))
			}
		} else {
			res.WriteHeader(200)
			res.Write([]byte(urlFromRedis))
		}
	}
}


//rest : http://localhost:8000/?short=ye.com/zk0Aef2lWP

func (*controller) GetLinkByShortLink(res http.ResponseWriter, req *http.Request) {
	
	var urlData entity.UrlData
	query := req.URL.Query()["short"]
	urlData.Url = query[0]

		if urlFromRedis := urlCache.GetRedis(urlData.Url); urlFromRedis == "" {
			if urlFromDb, err := urlService.GetShortDB(&urlData); urlFromDb == "" && err == nil {
				res.WriteHeader(400)
				res.Write([]byte("bad request1"))
				return
			} else {
				res.WriteHeader(200)
				res.Write([]byte(urlFromDb))
			}
		} else {
			res.WriteHeader(200)
			res.Write([]byte(urlFromRedis))
		}
}
