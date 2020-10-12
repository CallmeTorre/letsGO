package app

import (
	"github.com/CallmeTorre/letsGO/api/controllers/healthcheck"
	"github.com/CallmeTorre/letsGO/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/healthcheck", healthcheck.HealthCheck)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
