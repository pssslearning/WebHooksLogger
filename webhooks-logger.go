// main.go

package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	logfile, _ := os.Create("./log/GIN-" + getTimeStampAsString() + "-request.log")
	gin.DefaultWriter = logfile

	errlogfile, _ := os.Create("./log/GIN-" + getTimeStampAsString() + "-error.log")
	gin.DefaultErrorWriter = errlogfile

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()

}
