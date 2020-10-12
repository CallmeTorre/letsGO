package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router = gin.Default()
	//router = gin.New() //Blank engine without middleware
}

func StartApp() {
	mapUrls()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
