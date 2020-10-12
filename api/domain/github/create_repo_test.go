package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(T *testing.T) {
	request := CreateRepoRequest{
		Name:        "Golang Test",
		Description: "Golang introduction repository",
		Homepage:    "https://github.com",
		Private:     false,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     false,
	}
	bytes, err := json.Marshal(request)
	assert.Nil(T, err)
	assert.NotNil(T, bytes)
	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(T, err)
	assert.EqualValues(T, request, target)
	assert.EqualValues(T, request.Name, target.Name)
	assert.EqualValues(T, request.Description, target.Description)
	assert.EqualValues(T, request.Homepage, target.Homepage)
	assert.EqualValues(T, request.Private, target.Private)
	assert.EqualValues(T, request.HasIssues, target.HasIssues)
	assert.EqualValues(T, request.HasProjects, target.HasProjects)
	assert.EqualValues(T, request.HasWiki, target.HasWiki)
}
