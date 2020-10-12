package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CallmeTorre/letsGO/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(T *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/healthcheck", nil)
	c := test_utils.GetMockedContext(request, response)
	HealthCheck(c)
	assert.EqualValues(T, http.StatusOK, response.Code)
	assert.EqualValues(T, "OK", response.Body.String())

}
