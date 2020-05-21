// handler-main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showApiSamplesGET(c *gin.Context) {
	err := saveContextToFile(c)
	if err != nil {
		fmt.Println("Error en showRootGET: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"response": "ERROR"})
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "OK"})
	}
}

func showApiSamplesPOST(c *gin.Context) {
	err := saveContextToFile(c)
	if err != nil {
		fmt.Println("Error en showRootPOST: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"response": "ERROR"})
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "OK"})
	}
}
