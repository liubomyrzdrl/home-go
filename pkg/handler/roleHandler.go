package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetRoles(c *gin.Context) {
	 roles, err :=  h.services.Role.GetRoles()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, roles)
	}
}
