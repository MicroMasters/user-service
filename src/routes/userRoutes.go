package routes

import (
	"user-service/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetAllUsers)
	// router.GET("/users/:id", controllers.GetUserByID)
	router.POST("/users/create", controllers.CreateMongoUser)
	// router.PUT("/users/:id", controllers.UpdateUser)
	// router.DELETE("/users/:id", controllers.DeleteUser)
}
