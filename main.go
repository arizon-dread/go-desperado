package main

import (
	"net/http"

	//"github.com/arizon-dread/go-desperado/businesslayer"
	//"github.com/arizon-dread/go-desperado/models"

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
	jsonData := Desperado{}

	if err := c.BindJSON(&jsonData); err == nil {
		response := GetTextAsDesperado(jsonData)
		c.IndentedJSON(http.StatusOK, response)
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}

}
