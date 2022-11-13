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
		c.JSON(http.StatusOK, response)

	} else {
		fmt.Printf("err: %vn", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

}
func birthdaySeconds(c *gin.Context) {
	jsonData := models.BirthdaySeconds{}

	if err := c.BindJSON(&jsonData); err == nil {
		response, err := businesslayer.DateToSeconds(jsonData.Birthday)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, response)
		}

	} else {
		fmt.Printf("err: %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
