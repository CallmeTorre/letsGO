package app

import (
	"github.com/CallmeTorre/letsGO/api/log/option_b"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	err := godotenv.Load()
	if err != nil {
		option_b.Fatal(
			"Error loading .env file",
			err,
			option_b.Field("status", "fail"))
	}
	router = gin.Default()
	//router = gin.New() //Blank engine without middleware
}

func StartApp() {
	option_b.Info(
		"About to map URLS",
		option_b.Field("status", "pending"))
	mapUrls()
	option_b.Info(
		"URLS successfully mapped",
		option_b.Field("status", "success"))
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
