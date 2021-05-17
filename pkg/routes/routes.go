package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/liubomyrzdrl/home-go/pkg/handlers"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	userRoute := r.Group("/user-api")
	{
		userRoute.GET("/users", handlers.GetUsers)
		userRoute.POST("/user", handlers.CreateUser)
		userRoute.GET("/user/:id", handlers.GetUserByID)
		userRoute.DELETE("user/:id", handlers.DeleteUser)
	}
	return r
}
