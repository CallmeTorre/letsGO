package repositories

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/CallmeTorre/letsGO/api/clients/rest_client"
	"github.com/CallmeTorre/letsGO/api/domain/repositories"
	"github.com/CallmeTorre/letsGO/api/utils/errors"
	"github.com/CallmeTorre/letsGO/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(M *testing.M) {
	rest_client.StartMocks()
	os.Exit(M.Run())
}

func TestCreateRepoInvalidJsonRequest(T *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))

	c := test_utils.GetMockedContext(request, response)

	CreateRepo(c)
	assert.EqualValues(T, http.StatusBadRequest, response.Code)
	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())
	assert.Nil(T, err)
	assert.NotNil(T, apiErr)
	assert.EqualValues(T, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(T, "Invalid JSON Body", apiErr.Message())
}

func TestCreateRepoGithubError(T *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing"}`))
	c := test_utils.GetMockedContext(request, response)

	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message":"Requires authenthication", "documentation_url":"https://developer.github.com/"}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(T, http.StatusUnauthorized, response.Code)

	apiErr, err := errors.NewApiErrorFromBytes(response.Body.Bytes())

	assert.Nil(T, err)
	assert.NotNil(T, apiErr)
	assert.EqualValues(T, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(T, "Requires authenthication", apiErr.Message())
}

func TestCreateRepoNoError(T *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing"}`))
	c := test_utils.GetMockedContext(request, response)
	rest_client.FlushMock()
	rest_client.AddMock(rest_client.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
		},
	})

	CreateRepo(c)

	assert.EqualValues(T, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(T, err)
	assert.NotNil(T, result)
	assert.EqualValues(T, 123, result.ID)
}
