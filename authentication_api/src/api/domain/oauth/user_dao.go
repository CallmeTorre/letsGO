package oauth

import (
	"github.com/CallmeTorre/letsGO/api/utils/errors"
)

const queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username = ? AND password = ?;"

var users map[string]*User = map[string]*User{
	"alexis": &User{
		ID:       1,
		Username: "alexis",
	},
}

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundApiError("No user found")
	}
	return user, nil
}
