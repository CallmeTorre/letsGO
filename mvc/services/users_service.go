package services

import (
	"github.com/CallmeTorre/letsGO/mvc/domain"
	"github.com/CallmeTorre/letsGO/mvc/utils"
)

type usersService struct{}

var UsersService usersService

func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userID)
}
