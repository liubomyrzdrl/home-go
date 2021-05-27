package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/liubomyrzdrl/home-go/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()
	userRoute := r.Group("/user-api")
	{
		userRoute.GET("/users", h.GetUsers)
		userRoute.POST("/user", h.CreateUser)
		userRoute.GET("/user/:id", h.GetUserByID)
		userRoute.DELETE("user/:id", h.DeleteUser)
	}
	creditCardRoute := r.Group("/credit-api")
	{
		creditCardRoute.GET("/cards", h.GetCreditCards)
		creditCardRoute.POST("/card", h.CreateCreditCard)
	}
	role := r.Group("/role-api")
	{
		role.GET("/roles", h.GetRoles)
	}
	return r
}
