package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func main() {
	http.HandleFunc("/", urlHandler)
	http.ListenAndServe(":9000", nil)
}

type Data struct {
	Url       string
	ShortLink string
}

// curl -H "Content-Type: application/json" --request POST  --data '{"url":"https://github.com/destackq/real-time-forum" }'  http://localhost:9000

func urlHandler(w http.ResponseWriter, r *http.Request) {
	//domain name -> get query
	switch r.Method {
	case "GET":
		//get my server host url, r.URL.Query() - get after /, value(short link)
		//get from redis data key:value, key - shor link - value link
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		//get json : url -> validUrl,  generate random short link, check isUniq in Db ? -> save Redis & Db
	case "POST":
		d := Data{}

		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			log.Println(err)
		}
		// fmt.Printf("%s\n url value from client", d.Url)

		//check valid url ?
		if ok := d.IsValidUrl(d.Url); ok {
			d.Randomaizer()
			log.Print(d, "data")
			if ok := d.IsUniqUrl(); ok {
				//save redis, then db
			}
		}

		w.WriteHeader(200)
		w.Write([]byte(d.ShortLink))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}


//62 elem, 450penta variance

func (d *Data) Randomaizer() {

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

	log.Println(short, "short link create")
	d.ShortLink = short
}

func (d *Data) IsUniqUrl() bool {
	//check redis, then db
	return true
}

func (d *Data) IsValidUrl(host string) bool {
	_, err := url.ParseRequestURI(d.Url)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
