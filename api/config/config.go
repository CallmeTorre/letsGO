package config

import (
	"os"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	LogLevel             = "info"
)

func GetGithubAccessToken() string {
	return os.Getenv(apiGithubAccessToken)
}
