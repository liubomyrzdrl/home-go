package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liubomyrzdrl/home-go/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.services.User.GetAllUsers()
	logrus.Println("Users", users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}


}

func (h *Handler) CreateUser(c *gin.Context) {
	user:= models.User{}

	c.BindJSON(&user)
	err := h.services.User.CreateUser(user)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	intId, error := strconv.ParseInt(id[3:], 10, 64)
	if error != nil {
		panic(error)
	}

	err := h.services.User.GetUserById(user, intId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	intId, error := strconv.ParseInt(id[3:], 10, 64)
	if error != nil {
		panic(error)
	}
	err := h.services.User.DeleteUser(user, intId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id[3:]: "is deleted"})
	}
}