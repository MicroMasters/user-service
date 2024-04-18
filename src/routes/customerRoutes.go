package routes

import (
	"os"
	"user-service/src/controllers"
	"user-service/src/helpers"
	"user-service/src/jwt"
	"user-service/src/middlewares"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(router *gin.RouterGroup) {
	log := helpers.GetLogger()

	router.POST("/users/create", controllers.CreateMongoUser)

	//initialize jwt
	jwtSecret, err := helpers.GetEnvStringVal("JWT_SECRET")
	if err != nil {
		log.Error("JWT_SECRET not found in environment variables")
		os.Exit(1)
	}

	jwtIssuer, err := helpers.GetEnvStringVal("JWT_ISSUER")
	if err != nil {
		log.Error("JWT_ISSUER not found in environment variables")
		os.Exit(1)
	}

	jwtExpiry, err := helpers.GetEnvIntVal("JWT_EXPIRED")
	if err != nil {
		log.Error("JWT_EXPIRED not found in environment variables")
		os.Exit(1)
	}

	// jwt service
	jwtService := jwt.NewJWTService(jwtSecret, jwtIssuer, jwtExpiry)

	router.Use(middlewares.NewAuthMiddleware(jwtService, false, true, false))
	{
		router.GET("/users/all", controllers.GetAllUsers)
		router.GET("/users/:id", controllers.GetUserByID)
	}
}
