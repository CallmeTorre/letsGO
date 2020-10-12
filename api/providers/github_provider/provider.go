package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/CallmeTorre/letsGO/api/clients/rest_client"
	"github.com/CallmeTorre/letsGO/api/domain/github"
)

const (
	urlCreateRepo             string = "https://api.github.com/user/repos"
	headerAuthorization       string = "Authorization"
	headerAuthorizationFormat string = "token %s"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := rest_client.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println(fmt.Sprintf("Error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid Response Body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errorResponse github.GithubErrorResponse
		if err := json.Unmarshal(jsonBytes, &errorResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid JSON Response Body",
			}
		}
		errorResponse.StatusCode = response.StatusCode
		return nil, &errorResponse
	}
	var result github.CreateRepoResponse
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		log.Println(fmt.Sprintf("Error when trying to unmarshal create repo successful response: %s", err.Error()))
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when trying to unmarshal github create repo response",
		}
	}
	return &result, nil
}
