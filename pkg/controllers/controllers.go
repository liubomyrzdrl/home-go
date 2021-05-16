package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liubomyrzdrl/home-go/models"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	var user  []models.User
	err := models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

func CreateUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)
	err := models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	intId, error := strconv.ParseInt(id[3:], 10, 64)
	if error != nil {
		panic(error)
	}

	err := models.GetUserById(&user, intId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context)  {
	var user models.User
	id := c.Params.ByName("id")
	intId, error := strconv.ParseInt(id[3:], 10, 64)
	if error != nil {
		panic(error)
	}
	err := models.DeleteUser(&user, intId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id[3:]: "is deleted"})
	}
}