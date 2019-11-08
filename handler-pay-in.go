// handler-main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showPayINWebhookPOST(c *gin.Context) {
	err := saveContextToFile(c)
	if err != nil {
		fmt.Println("Error en showRootPOST: ", err)
	}
	c.JSON(http.StatusOK, gin.H{"par1": "par1", "par2": "par2"})
}
