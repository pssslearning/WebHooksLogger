// handler-main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showRootGET(c *gin.Context) {
	err := saveContextToFile(c)
	if err != nil {
		fmt.Println("Error en showRootGET: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"response": "ERROR"})
	}
	c.JSON(http.StatusOK, gin.H{"response": "OK"})

}

func showRootPOST(c *gin.Context) {
	err := saveContextToFile(c)
	if err != nil {
		fmt.Println("Error en showRootPOST: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"response": "ERROR"})
	}
	c.JSON(http.StatusOK, gin.H{"response": "OK"})

}
