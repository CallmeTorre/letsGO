package rest_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	Error      error
}

var (
	enabledMocks bool             = false
	mocks        map[string]*Mock = make(map[string]*Mock)
)

func getMockId(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

func StartMocks() {
	enabledMocks = true
}

func StopMocks() {
	enabledMocks = false
}

func AddMock(mock Mock) {
	mocks[getMockId(mock.HTTPMethod, mock.URL)] = &mock
}

func FlushMock() {
	mocks = make(map[string]*Mock)
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		mock := mocks[getMockId(http.MethodPost, url)]
		if &mock == nil {
			return nil, errors.New("No mock found for given request")
		}
		return mock.Response, mock.Error
	}
	jsonBytes, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
