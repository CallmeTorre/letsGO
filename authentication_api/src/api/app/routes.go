package app

import (
	"github.com/CallmeTorre/letsGO/api/controllers/healthcheck"
	"github.com/CallmeTorre/letsGO/authentication_api/src/api/controllers/oauth"
)

func mapUrls() {
	router.GET("/healthcheck", healthcheck.HealthCheck)
	router.POST("/oauth/token", oauth.CreateAccessToken)
	router.GET("/oauth/token/:token_id", oauth.GetAccessToken)
}
