package oauth

import (
	"fmt"

	"github.com/CallmeTorre/letsGO/api/utils/errors"
)

var tokens map[string]*Token = make(map[string]*Token, 0)

func (t *Token) Save() errors.NewApiError {
	t.Token = fmt.Sprintf("USR_%d", t.UserID)
	tokens[t.Token] = t
	return nil
}

func GetAccessToken(accessToken string) (*Token, errors.ApiError) {
	token := tokens[accessToken]
	if token == nil || token.IsExpired() {
		return nil, errors.NewNotFoundApiError("Token Not Found")
	}
	return token, nil
}
