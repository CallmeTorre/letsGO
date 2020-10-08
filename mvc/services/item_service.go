package services

import (
	"net/http"

	"github.com/CallmeTorre/letsGO/mvc/domain"
	"github.com/CallmeTorre/letsGO/mvc/utils"
)

type itemService struct{}

var ItemService itemService

func (u *itemService) GetUser(itemID string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
