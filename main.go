package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

func main() {
	http.HandleFunc("/", urlHandler)
	http.ListenAndServe(":9000", nil)
}

type UrlData struct {
	Url      string `json:"url"`
	ShortUrl string `json:"short_url"`
	Host     string
}

//post   http://localhost:9000/ body : {url: https://github.com/devstackq/ozon-task }
//get http://localhost:9000/?url=gb.com/FvRIStQgV8

var UrlStruct UrlData

func urlHandler(w http.ResponseWriter, r *http.Request) {
	//domain name -> get query
	switch r.Method {

	case "GET":
		//get my server host url, r.URL.Query() - get after /, value(short link)
		//get from redis data key:value, key - shor link - value link

		keys, ok := r.URL.Query()["url"]

		if !ok || len(keys[0]) < 1 {
			log.Println("url param 'url' is missing or incorrect")
			return
		}

		if err, shortUrl := UrlStruct.IsUniqUrl(); err == nil && shortUrl != "" {
			//return full url from db || redis
			w.WriteHeader(200)
			w.Write([]byte(shortUrl))
			return
		}

		w.WriteHeader(400)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return

		//get json : url -> validUrl,  generate random short link, check isUniq in Db ? -> save Redis & Db
	case "POST":

		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			log.Println(err)
		}

		//check valid url ?
		if err, host := UrlStruct.IsValidUrl(UrlStruct.Url); err == nil {
			// check isUrlExist ?
			if err, shortUrl := UrlStruct.IsUniqUrl(); err == nil && shortUrl == "" {
				if err := UrlStruct.CreateShortHost(host); err == nil {
					UrlStruct.Randomaizer()

					if UrlStruct.ShortUrl != "" {
						fmt.Println(d, "create short url")
						//save redis, then db, return json. new Url
						w.WriteHeader(200)
						w.Write([]byte(UrlStruct.ShortUrl))
						return
					}
				}
			} else {
				w.WriteHeader(200)
				//get from db || redis by d.Url - shortUrl
				//return generated url by key
				w.Write([]byte(shortUrl))
			}
			w.WriteHeader(500)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}

//set new short host
func (d *UrlData) CreateShortHost(host string) error {

	re, err := regexp.Compile(`[.]`)
	if err != nil {
		return err
	}

	hostSplit := re.Split(host, -1)

	if len(hostSplit) == 2 {
		d.Host = string(hostSplit[0][0]) + string(hostSplit[0][len(hostSplit[0])-1]) + "." + hostSplit[1] + "/"
	} else if len(hostSplit) == 3 {
		d.Host = string(hostSplit[0]) + "." + string(hostSplit[1][0]) + string(hostSplit[1][len(hostSplit[1])-1]) + "." + hostSplit[2] + "/"
	}
	return nil
}

//62 elem, 450penta variance
func (d *UrlData) Randomaizer() {

	items := [62]int{}

	for i := 0; i < 26; i++ {
		items[i] = i + 65
	}
	for i := 0; i < 26; i++ {
		items[i+26] = i + 97
	}
	for i := 0; i < 10; i++ {
		items[i+52] = i + 48
	}

	//create short rand url
	var short string
	//uniq index
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < 10; i++ {
		short += string(items[r1.Intn(62)])
	}

	if len(short) == 10 && short != "" {
		d.ShortUrl = d.Host + short
	}
}

func (d *UrlData) IsUniqUrl() (error, string) {
	//check redis, then db
	return nil, ""
}

func (d *UrlData) IsValidUrl(host string) (error, string) {
	u, err := url.ParseRequestURI(d.Url)
	if err != nil {
		return err, ""
	}
	return nil, u.Hostname()
}
