package repository

import (
	"database/sql"
	"log"

	"github.com/devstackq/ozon/entity"
	_ "github.com/lib/pq"
)

type pSqlRepo struct{ db *sql.DB }

var ()

func NewPSqlRepository() UrlRepository {
	configDb := "postgres://postgres:password@localhost:5432/ozondb?sslmode=disable"
	db, err := sql.Open("postgres", configDb)
	if err != nil {
		log.Println(err)
	}
	if err = db.Ping(); err != nil {
		log.Println(err)
	}

	linksTable, err := db.Prepare(`CREATE TABLE IF NOT EXISTS links  (
		id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
		url varchar(255) NOT NULL UNIQUE,
		short varchar(255) NOT NULL,
		createdtime timestamp
	)`)
	if err != nil {
		log.Println(err)
	}
	linksTable.Exec()

	return &pSqlRepo{db: db}
}

//db gloabal
func (p *pSqlRepo) Create(url *entity.UrlData) error {
	_, err := p.db.Exec("INSERT INTO links (url, short) VALUES ($1,$2)", url.Url, url.ShortUrl)
	if err != nil {
		return err
	}
	log.Println(url, "insert db", url.ShortUrl, err)

	return nil
}

func (p *pSqlRepo) GetByUrl(url *entity.UrlData) (result string, err error) {
	p.db.QueryRow(`SELECT short FROM where url=$1`, url.Url).Scan(&result)
	log.Println("get from db by url", result)
	return result, nil
}

func (p *pSqlRepo) GetByShort(url *entity.UrlData) (result string, err error) {
	p.db.QueryRow(`SELECT url FROM where short=$1`, url.ShortUrl).Scan(&result)
	log.Println("get from db by short", result)
	
	return result, nil
}
