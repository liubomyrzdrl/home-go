package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/liubomyrzdrl/home-go/models"
	"net/http"
)

func GetRoles(c *gin.Context) {
	var roles []models.Role
	err := models.GetRoles(&roles)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, roles)
	}
}

