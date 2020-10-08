package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/CallmeTorre/letsGO/mvc/domain"
	"github.com/CallmeTorre/letsGO/mvc/utils"
	"github.com/stretchr/testify/assert"
)

type usersDaoMock struct{}

var userDaoMock usersDaoMock

var getUserFunction func(userID int64) (*domain.User, *utils.ApplicationError)

func (u *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func init() {
	domain.UserDao = &usersDaoMock{}
}

func TestGetUserNotFoundDB(T *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    fmt.Sprintf("User %v not found", userID),
		}
	}
	user, err := UsersService.GetUser(-1)
	assert.Nil(T, user)
	assert.NotNil(T, err)
	assert.EqualValues(T, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(T, "User -1 not found", err.Message)
}

func TestGetUserNoError(T *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID: 123,
		}, nil
	}
	user, err := UsersService.GetUser(123)
	assert.Nil(T, err)
	assert.NotNil(T, user)
	assert.EqualValues(T, 123, user.ID)
}
