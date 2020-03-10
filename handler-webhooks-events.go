// handler-main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showWebhookEventPOST(c *gin.Context) {
	err := saveContextToFile(c)
	if err != nil {
		fmt.Println("Error en showWebhookEventPOST: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"response": "ERROR"})
	}
	c.JSON(http.StatusOK, gin.H{"response": "OK"})
}
