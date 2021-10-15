package repository

import "github.com/devstackq/ozon/entity"

type pSqlRepo struct{}

func NewPSqlRepository() UrlRepository {
	//connect db,
	//create table
	//create query methods
	return &pSqlRepo{}
}

func (*pSqlRepo) Create(url *entity.UrlData) (string, error) {
	return "", nil
}
func (*pSqlRepo) Get(url *entity.UrlData) (string, error) {
	return "", nil
}
