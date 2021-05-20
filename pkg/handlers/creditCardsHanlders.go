package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liubomyrzdrl/home-go/models"
	"net/http"
)

func GetCreditCards(c *gin.Context) {
	var cards []models.CreditCard
	err := models.GetAllCreditCards(&cards)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cards)
	}
}

func CreateCreditCard(c *gin.Context) {
	var card models.CreditCard
	c.BindJSON(&card)
	err := models.CreateCreditCard(&card)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, card)
	}
}