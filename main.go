package main

import (
	"fmt"
	"net/http"

	"github.com/arizon-dread/go-desperado/businesslayer"
	"github.com/arizon-dread/go-desperado/models"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/ping", getPing)
	router.POST("/desperado", getAsDesperado)
	router.POST("/birthdaySeconds", birthdaySeconds)
	router.Run(":8080")
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
		fmt.Printf("err: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

}
func birthdaySeconds(c *gin.Context) {
	jsonData := models.BirthdaySeconds{}

	if err := c.BindJSON(&jsonData); err == nil {
		response := businesslayer.DateToSeconds(jsonData.Birthday)
		c.IndentedJSON(http.StatusOK, response)
	} else {
		fmt.Printf("err: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
