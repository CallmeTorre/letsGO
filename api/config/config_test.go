package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(T *testing.T) {
	assert.EqualValues(T, "SECRET_GITHUB_ACCESS_TOKEN", apiGithubAccessToken)
}

func TestGetGithubAccessToken(T *testing.T) {
	assert.EqualValues(T, "", GetGithubAccessToken())
}
