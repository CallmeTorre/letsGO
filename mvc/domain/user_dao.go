package domain

import (
	"fmt"
	"net/http"

	"github.com/CallmeTorre/letsGO/mvc/utils"
)

var users = map[int64]*User{
	123: {ID: 123, FirstName: "Alexis", LastName: "Torreblanca", Email: "alexis.torreblanca@wizeline.com"},
}

func GetUser(userID int64) (*User, *utils.ApplicationError) {
	user := users[userID]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("User %v not found", userID),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
