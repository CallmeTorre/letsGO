package test_utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMockedContext(T *testing.T) {
	response := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/test", nil)
	assert.Nil(T, err)
	request.Header = http.Header{"X-Mock": {"true"}}
	c := GetMockedContext(request, response)
	assert.EqualValues(T, http.MethodPost, c.Request.Method)
	assert.EqualValues(T, "8080", c.Request.URL.Port())
	assert.EqualValues(T, "/test", c.Request.URL.Path)
	assert.EqualValues(T, "http", c.Request.URL.Scheme)
	assert.EqualValues(T, 1, len(c.Request.Header))
	assert.EqualValues(T, "true", c.GetHeader("X-Mock"))
	assert.EqualValues(T, "true", c.GetHeader("x-mock"))
}
