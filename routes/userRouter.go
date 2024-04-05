package routes

import (
	controller "Restorant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUsers())
	incomingRoutes.POST("/users/sign-up",controller.SignUp())
	incomingRoutes.POST("/users/login",controller.Login())

}