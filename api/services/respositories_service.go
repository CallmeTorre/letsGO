package services

import (
	"net/http"
	"sync"

	"github.com/CallmeTorre/letsGO/api/config"
	"github.com/CallmeTorre/letsGO/api/domain/github"
	"github.com/CallmeTorre/letsGO/api/domain/repositories"
	"github.com/CallmeTorre/letsGO/api/providers/github_provider"
	"github.com/CallmeTorre/letsGO/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(request []repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}

func (s *reposService) CreateRepos(request []repositories.CreateRepoRequest) *repositories.CreateReposResponse {
	channel := make(chan repositories.CreateReposResult)
	outputChannel := make(chan repositories.CreateReposResponse)
	defer close(outputChannel)
	var wg sync.WaitGroup
	go s.handleRepoResults(channel, outputChannel, &wg)
	for _, current := range request {
		wg.Add(1)
		go s.createRepoConcurrent(current, channel)
	}
	wg.Wait()
	close(channel)
	result := <-outputChannel

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations++
		}
	}
	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(request) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return &result
}

func (s *reposService) handleRepoResults(inputChannel chan repositories.CreateReposResult, outputChannel chan repositories.CreateReposResponse, wg *sync.WaitGroup) {
	var results repositories.CreateReposResponse
	for incomingEvent := range inputChannel {
		repoResult := repositories.CreateReposResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	outputChannel <- results
}

func (s *reposService) createRepoConcurrent(input repositories.CreateRepoRequest, channel chan repositories.CreateReposResult) {
	if err := input.Validate(); err != nil {
		channel <- repositories.CreateReposResult{
			Error: err,
		}
		return
	}

	result, err := s.CreateRepo(input)

	if err != nil {
		channel <- repositories.CreateReposResult{
			Error: err,
		}
		return
	}
	channel <- repositories.CreateReposResult{
		Response: result,
	}
}
