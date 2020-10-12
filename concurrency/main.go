package concurrency

import (
	"bufio"
	"os"
	"sync"

	"github.com/CallmeTorre/letsGO/api/domain/repositories"
	"github.com/CallmeTorre/letsGO/api/services"
	"github.com/CallmeTorre/letsGO/api/utils/errors"
)

var (
	success map[string]string
	fail    map[string]errors.ApiError
)

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result  *repositories.CreateRepoResponse
	Error   errors.ApiError
}

func getRequests() []repositories.CreateRepoRequest {
	result := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("/path/to/file.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		request := repositories.CreateRepoRequest{
			Name: line,
		}
		result = append(result, request)
	}
	file.Close()
	return result
}

func main() {
	requests := getRequests()
	input := make(chan createRepoResult)
	//This limits the amount of concurrent events
	buffer := make(chan bool, 10)
	var wg sync.WaitGroup

	go handleResults(input, &wg)

	for _, request := range requests {
		buffer <- true
		wg.Add(1)
		go createRepo(request, input, buffer)
	}
	wg.Wait()
	close(input)
}

func handleResults(input chan createRepoResult, wg *sync.WaitGroup) {
	for result := range input {
		if result.Error != nil {
			fail[result.Request.Name] = result.Error
		} else {
			success[result.Request.Name] = result.Result.Name
		}
		wg.Done()
	}
}

func createRepo(request repositories.CreateRepoRequest, output chan createRepoResult, buffer chan bool) {
	result, err := services.RepositoryService.CreateRepo(request)
	output <- createRepoResult{
		Request: request,
		Result:  result,
		Error:   err,
	}
	<-buffer
}
