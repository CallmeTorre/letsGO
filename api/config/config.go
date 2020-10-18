package config

import (
	"os"
)

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	LogLevel             = "info"
	goEnvironment        = "ENVIRONMENT"
)

func GetGithubAccessToken() string {
	return os.Getenv(apiGithubAccessToken)
}

func IsProduction() bool {
	return os.Getenv(goEnvironment) == "production"
}
