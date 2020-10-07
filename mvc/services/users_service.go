package services

import (
	"github.com/CallmeTorre/letsGO/mvc/domain"
	"github.com/CallmeTorre/letsGO/mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
