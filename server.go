package main

import (
	"log"
	"runtime"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func main() {

	r := SetupRouter()
	log.Println("gin version: " + gin.Version)
	log.Println("go version: " + runtime.Version())
	r.Run() // listen and serve on 0.0.0.0:8080
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(location.New(location.Config{
		Scheme: "http",
		Host:   "localhost:9000",
		Base:   "/",
	}))
	// The "type" variable is not used in this simplified version.
	router.POST("/:type", NewOrganizationController().Create())
	return router
}
