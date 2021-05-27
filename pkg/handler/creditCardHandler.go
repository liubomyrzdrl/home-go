package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liubomyrzdrl/home-go/models"
	"net/http"
)

func (h *Handler) GetCreditCards(c *gin.Context) {

	cards,err := h.services.CreditCard.GetAllCreditCards()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, cards)
	}
}

func (h *Handler) CreateCreditCard(c *gin.Context) {
	var card models.CreditCard
	c.BindJSON(&card)
	err :=  h.services.CreateCreditCard(card)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, card)
	}
}
