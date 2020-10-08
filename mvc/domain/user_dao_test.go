package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(T *testing.T) {
	user, err := GetUser(-1)

	assert.Nil(T, user, "We are not expecting an user with id -1")
	assert.NotNil(T, err)
	assert.EqualValues(T, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(T, "User -1 not found", err.Message)
	assert.EqualValues(T, "not_found", err.Code)

	/*if user != nil {
		T.Errorf("We are not expecting an user with id -1")
	}

	if err == nil {
		T.Errorf("We are expecting an error when user id -1")
	}

	if err.StatusCode != http.StatusNotFound {
		T.Errorf("We are expecting 404 error")
	}*/
}

func TestGetUserFound(T *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(T, user)
	assert.Nil(T, err)
	assert.EqualValues(T, 123, user.ID)
	assert.EqualValues(T, "Alexis", user.FirstName)
	assert.EqualValues(T, "Torreblanca", user.LastName)
	assert.EqualValues(T, "alexis.torreblanca@wizeline.com", user.Email)
}
