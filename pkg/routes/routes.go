package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/liubomyrzdrl/home-go/pkg/controllers"
)

func InitRouter() *gin.Engine {
	//r:= gin.Default()
	r := gin.New()
	userRoute := r.Group("/user-api")
	{
		userRoute.GET("/users", controllers.GetUsers)
		userRoute.POST("/user", controllers.CreateUser)
		userRoute.GET("/user/:id", controllers.GetUserByID)
		userRoute.DELETE("user/:id", controllers.DeleteUser)
	}
	return r
}