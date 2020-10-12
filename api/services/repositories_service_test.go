package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/CallmeTorre/letsGO/api/clients/rest_client"
	"github.com/CallmeTorre/letsGO/api/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestMain(M *testing.M) {
	rest_client.StartMocks()
	os.Exit(M.Run())
}

func TestCreateRepoInvalidInputName(T *testing.T) {
	request := repositories.CreateRepoRequest{}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(T, result)
	assert.NotNil(T, err)
	assert.EqualValues(T, http.StatusBadRequest, err.Status())
	assert.EqualValues(T, "Invalid Repository Name", err.Message())

}

func TestCreateRepoGithubError(T *testing.T) {
	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires authenthication", "documentation_url":"https://developer.github.com/"}`)),
		},
	})
	request := repositories.CreateRepoRequest{
		Name: "Test",
	}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(T, result)
	assert.NotNil(T, err)
	assert.EqualValues(T, http.StatusUnauthorized, err.Status())
	assert.EqualValues(T, "Requires authenthication", err.Message())
}

func TestCreateNoError(T *testing.T) {
	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123, "name":"test"}`)),
		},
	})
	request := repositories.CreateRepoRequest{
		Name: "Test",
	}
	result, err := RepositoryService.CreateRepo(request)
	assert.NotNil(T, result)
	assert.Nil(T, err)
	assert.EqualValues(T, 123, result.ID)
	assert.EqualValues(T, "test", result.Name)
}
