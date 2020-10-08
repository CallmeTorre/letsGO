package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CallmeTorre/letsGO/mvc/utils"
)

type userDao struct{}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

var users = map[int64]*User{
	123: {ID: 123, FirstName: "Alexis", LastName: "Torreblanca", Email: "alexis.torreblanca@wizeline.com"},
}

var UserDao userDaoInterface

func init() {
	UserDao = &userDao{}
}

func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("We are accessing the database")
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
