package oauth

import (
	"strings"

	"github.com/CallmeTorre/letsGO/api/utils/errors"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *TokenRequest) Validate() errors.ApiError {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return errors.NewBadRequestError("Invalid Username")
	}
	r.Password = strings.TrimSpace(r.Password)
	if r.Password == "" {
		return errors.NewBadRequestError("Invalid Password")
	}
	return nil
}
