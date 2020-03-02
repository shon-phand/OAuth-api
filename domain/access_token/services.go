package access_token

import (
	"strings"

	"github.com/shon-phand/OAuth-api/utils/errors"
)

type Service interface{
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Repository interface{
	GetById(string) (*AccessToken, *errors.RestErr)
}


type service struct{
	repository Repository
}

func NewService(repo Repository) Service{

	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenID string) (*AccessToken, *errors.RestErr){
	accessTokenID = strings.TrimSpace(accessTokenID)
	if len(accessTokenID) == 0{
		return nil,errors.StatusBadRequestError("invalid access token id")
	}
	accesstoken,err:= s.repository.GetById(accessTokenID)	
	if err !=nil{
		return nil,err
	}
	return accesstoken,nil
}
