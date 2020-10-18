package app

import (
	"github.com/CallmeTorre/letsGO/api/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err, "status:fail")
	}
	router = gin.Default()
	//router = gin.New() //Blank engine without middleware
}

func StartApp() {
	log.Info("About to map URLS", "step:1", "status:pending")
	mapUrls()
	log.Info("URLS successfully mapped", "step:2", "status:success")
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
