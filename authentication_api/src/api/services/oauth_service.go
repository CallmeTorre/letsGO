package services

import (
	"time"

	"github.com/CallmeTorre/letsGO/api/utils/errors"
	"github.com/CallmeTorre/letsGO/authentication_api/src/api/domain/oauth"
)

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.TokenRequest) (*oauth.TokenRequest, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.Token, errors.ApiError)
}

var OauthService oauthServiceInterface

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.TokenRequest) (*oauth.Token, errors.ApiError) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	user, err := oauth.GetUserByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	token := oauth.Token{
		UserID:  user.ID,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}
	if err := token.Save(); err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.Token, errors.ApiError) {
	token, err := oauth.GetAccessToken(accessToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}
