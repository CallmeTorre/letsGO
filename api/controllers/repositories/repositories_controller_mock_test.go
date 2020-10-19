package repositories

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/CallmeTorre/letsGO/api/domain/repositories"
	"github.com/CallmeTorre/letsGO/api/services"
	"github.com/CallmeTorre/letsGO/api/utils/errors"
	"github.com/CallmeTorre/letsGO/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

var (
	funcCreateRepo  func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	funcCreateRepos func(request []repositories.CreateRepoRequest) *repositories.CreateReposResponse
)

type repoServiceMock struct {
}

func (s *repoServiceMock) CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	return funcCreateRepo(request)
}
func (s *repoServiceMock) CreateRepos(request []repositories.CreateRepoRequest) *repositories.CreateReposResponse {
	return funcCreateRepos(request)
}

func TestCreateRepoNoErrorMockingTheEntireService(T *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return &repositories.CreateRepoResponse{
			ID: 123,
		}, nil
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing"}`))
	c := test_utils.GetMockedContext(request, response)
	CreateRepo(c)

	assert.EqualValues(T, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(T, err)
	assert.NotNil(T, result)
	assert.EqualValues(T, 123, result.ID)
}

func TestCreateRepoGithubErrorMockingTheEntireService(T *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
		return nil, errors.NewBadRequestError("Invalid Repository Name")
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing"}`))
	c := test_utils.GetMockedContext(request, response)
	CreateRepo(c)

	assert.EqualValues(T, http.StatusBadRequest, response.Code)

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())
	assert.Nil(T, err)
	assert.NotNil(T, apiErr)
	assert.EqualValues(T, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(T, "Invalid Repository Name", apiErr.Message())
}
