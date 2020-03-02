package db

import (
	"github.com/shon-phand/OAuth-api/domain/access_token"
	"github.com/shon-phand/OAuth-api/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.StatusInternalServerError("database connection not implemented")
}
