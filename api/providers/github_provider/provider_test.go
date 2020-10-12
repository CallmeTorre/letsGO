package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/CallmeTorre/letsGO/api/clients/rest_client"
	"github.com/CallmeTorre/letsGO/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestMain(M *testing.M) {
	rest_client.StartMocks()
	os.Exit(M.Run())
}

func TestConstants(T *testing.T) {
	assert.EqualValues(T, "Authorization", headerAuthorization)
	assert.EqualValues(T, "token %s", headerAuthorizationFormat)
	assert.EqualValues(T, "https://api.github.com/user/repos", urlCreateRepo)

}

func TestGetAuthorizationHeader(T *testing.T) {
	header := getAuthorizationHeader("token-test")
	assert.EqualValues(T, "token token-test", header)
}

func TestCreateRepoErrorRestClient(T *testing.T) {
	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Error:      errors.New("Invalid Rest Client Response"),
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(T, response)
	assert.NotNil(T, err)
	assert.EqualValues(T, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(T, "Invalid Rest Client Response", err.Message)
}

func TestCreateRepoInvalidResponseBody(T *testing.T) {
	rest_client.FlushMock()
	invalidCloser, _ := os.Open("test")
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       invalidCloser,
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(T, response)
	assert.NotNil(T, err)
	assert.EqualValues(T, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(T, "Invalid Response Body", err.Message)
}

func TestCreateRepoInvalidErrorInterface(T *testing.T) {
	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":1}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(T, response)
	assert.NotNil(T, err)
	assert.EqualValues(T, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(T, "Invalid JSON Response Body", err.Message)
}

func TestCreateRepoInvalidUnauthorized(T *testing.T) {
	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires authenthication", "documentation_url":"https://developer.github.com/"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(T, response)
	assert.NotNil(T, err)
	assert.EqualValues(T, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(T, "Requires authenthication", err.Message)
}

func TestCreateRepoInvalidInvalidSuccessResponse(T *testing.T) {
	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":"123", "name":-1}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(T, response)
	assert.NotNil(T, err)
	assert.EqualValues(T, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(T, "Error when trying to unmarshal github create repo response", err.Message)
}

func TestCreateRepoNoError(T *testing.T) {
	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123, "name":"test"}`)),
		},
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.NotNil(T, response)
	assert.Nil(T, err)
	assert.EqualValues(T, 123, response.ID)
	assert.EqualValues(T, "test", response.Name)

}
