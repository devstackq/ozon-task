package service

import (
	"math/rand"
	"net/url"
	"regexp"
	"time"

	"github.com/devstackq/ozon/entity"
	"github.com/devstackq/ozon/repository"
)

var (
	repo repository.UrlRepository
)

type service struct{}

func NewUrlService(repository repository.UrlRepository) UrlService {
	repo = repository
	return &service{}
}

//62 elem, 450penta variance
func (*service) Randomaizer() string {

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
		return short
	}
	return ""
}

func (*service) CreateShortHost(url *entity.UrlData) (error, string) {
	var result string
	re, err := regexp.Compile(`[.]`)
	if err != nil {
		return err, ""
	}

	hostSplit := re.Split(url.Host, -1)

	if len(hostSplit) == 2 {
		result = string(hostSplit[0][0]) + string(hostSplit[0][len(hostSplit[0])-1]) + "." + hostSplit[1] + "/"
	} else if len(hostSplit) == 3 {
		result = string(hostSplit[0]) + "." + string(hostSplit[1][0]) + string(hostSplit[1][len(hostSplit[1])-1]) + "." + hostSplit[2] + "/"
	}
	return nil, result
}

func (*service) IsValidUrl(data *entity.UrlData) (error, string) {
	u, err := url.ParseRequestURI(data.Url)
	if err != nil {
		return err, ""
	}
	return nil, u.Hostname()
}

func (*service) IsUniqUrl() bool {
	return true
}

func (*service) SaveUrlDB(url *entity.UrlData) (err error) {
	err = repo.Create(url)
	return err
}

func (s *service) GetUrlDB(url *entity.UrlData) (result string, err error) {
	result, err = repo.GetByUrl(url)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *service) GetShortDB(url *entity.UrlData) (result string, err error) {
	result, err = repo.GetByShort(url)
	if err != nil {
		return "", err
	}
	return result, nil
}
