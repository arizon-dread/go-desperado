package main

import (
	"example/ws-test/businesslayer"
	"example/ws-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/ping", getPing)
	router.POST("/desperado", getAsDesperado)
	router.Run("localhost:8080")
}

func getPing(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Pong")
}
func getAsDesperado(c *gin.Context) {
	jsonData := models.Desperado{}

	if err := c.BindJSON(&jsonData); err == nil {
		response := businesslayer.GetTextAsDesperado(jsonData)
		c.IndentedJSON(http.StatusOK, response)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}

}
