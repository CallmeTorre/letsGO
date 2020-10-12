package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/CallmeTorre/letsGO/api/clients/rest_client"
	"github.com/CallmeTorre/letsGO/api/domain/repositories"
	"github.com/CallmeTorre/letsGO/api/utils/errors"
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

func TestCreateRepoConcurrentInvalidRequest(T *testing.T) {
	request := repositories.CreateRepoRequest{}
	output := make(chan repositories.CreateReposResult)
	service := reposService{}
	go service.createRepoConcurrent(request, output)
	result := <-output
	assert.NotNil(T, result)
	assert.Nil(T, result.Response)
	assert.NotNil(T, result.Error)
	assert.EqualValues(T, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(T, "Invalid Repository Name", result.Error.Message())
}

func TestCreateRepoConcurrentGithubError(T *testing.T) {
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

	output := make(chan repositories.CreateReposResult)
	service := reposService{}
	go service.createRepoConcurrent(request, output)
	result := <-output
	assert.NotNil(T, result)
	assert.Nil(T, result.Response)
	assert.NotNil(T, result.Error)
	assert.EqualValues(T, http.StatusUnauthorized, result.Error.Status())
	assert.EqualValues(T, "Requires authenthication", result.Error.Message())
}

func TestCreateRepoConcurrentNoError(T *testing.T) {
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

	output := make(chan repositories.CreateReposResult)
	service := reposService{}
	go service.createRepoConcurrent(request, output)
	result := <-output
	assert.NotNil(T, result)
	assert.NotNil(T, result.Response)
	assert.Nil(T, result.Error)
	assert.EqualValues(T, 123, result.Response.ID)
	assert.EqualValues(T, "test", result.Response.Name)
}

func TestHandleResults(T *testing.T) {
	var wg sync.WaitGroup
	inputChannel := make(chan repositories.CreateReposResult)
	outputChannel := make(chan repositories.CreateReposResponse)
	service := reposService{}
	go service.handleRepoResults(inputChannel, outputChannel, &wg)
	wg.Add(1)
	go func() {
		inputChannel <- repositories.CreateReposResult{
			Error: errors.NewBadRequestError("Invalid Repository Name"),
		}
	}()

	wg.Wait()
	close(inputChannel)

	result := <-outputChannel
	assert.NotNil(T, result)
	assert.EqualValues(T, 0, result.StatusCode)
	assert.EqualValues(T, 1, len(result.Results))
	assert.EqualValues(T, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(T, "Invalid Repository Name", result.Results[0].Error.Message())
}

func TestCreateReposInvalidRequests(T *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "   "},
	}
	result := RepositoryService.CreateRepos((requests))
	assert.NotNil(T, result)
	assert.EqualValues(T, 2, len(result.Results))
	assert.EqualValues(T, http.StatusBadRequest, result.StatusCode)
	assert.Nil(T, result.Results[0].Response)
	assert.EqualValues(T, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(T, "Invalid Repository Name", result.Results[0].Error.Message())
	assert.Nil(T, result.Results[1].Response)
	assert.EqualValues(T, http.StatusBadRequest, result.Results[1].Error.Status())
	assert.EqualValues(T, "Invalid Repository Name", result.Results[1].Error.Message())
}

func TestCreateReposOneSuccessOneFail(T *testing.T) {
	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123, "name":"test"}`)),
		},
	})
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "testing"},
	}
	result := RepositoryService.CreateRepos((requests))
	assert.NotNil(T, result)
	assert.EqualValues(T, 2, len(result.Results))
	assert.EqualValues(T, http.StatusPartialContent, result.StatusCode)
	for _, value := range result.Results {
		if value.Error != nil {
			assert.EqualValues(T, http.StatusBadRequest, value.Error.Status())
			assert.EqualValues(T, "Invalid Repository Name", value.Error.Message())
		} else {
			assert.EqualValues(T, 123, value.Response.ID)
			assert.EqualValues(T, "test", value.Response.Name)
		}
	}
}

func TestCreateReposNoError(T *testing.T) {
	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123, "name":"test"}`)),
		},
	})
	requests := []repositories.CreateRepoRequest{
		{Name: "testing"},
		{Name: "testing"},
	}
	result := RepositoryService.CreateRepos((requests))
	assert.NotNil(T, result)
	assert.EqualValues(T, 2, len(result.Results))
	assert.EqualValues(T, http.StatusCreated, result.StatusCode)
	assert.Nil(T, result.Results[0].Error)
	assert.EqualValues(T, 123, result.Results[0].Response.ID)
	assert.EqualValues(T, "test", result.Results[0].Response.Name)
	assert.Nil(T, result.Results[1].Error)
	assert.EqualValues(T, 123, result.Results[1].Response.ID)
	assert.EqualValues(T, "test", result.Results[1].Response.Name)
}
