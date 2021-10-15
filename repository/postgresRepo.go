package repository

type pSqlRepo struct {}

func NewPSqlRepository() UrlRepository {
	return &pSqlRepo{}
}