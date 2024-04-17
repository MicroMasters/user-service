package main

import (
	"fmt"
	"os"
	"user-service/src/connection/db"
	"user-service/src/controllers"
	"user-service/src/helpers"
	logger "user-service/src/loggers"
	"user-service/src/middlewares"
	"user-service/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log := helpers.GetLogger()

	controllers.InitializeControllers()

	//initialize gin mode
	ginMode, err := helpers.GetEnvStringVal("GIN_MODE")
	if err != nil {
		log.Error("GIN_MODE not found in environment variables")
		os.Exit(1)
	}
	gin.SetMode(ginMode)

	//api config port
	port, err := helpers.GetEnvStringVal("API_CONFIG_PORT")
	if err != nil {
		log.Error("API_CONFIG_PORT not found in environment variables")
		os.Exit(1)
	}

	if port != "8080" {
		port = "8090"
	}

	//initialize jwt
	// jwtSecret, err := helpers.GetEnvStringVal("JWT_SECRET")
	// if err != nil {
	// 	log.Error("JWT_SECRET not found in environment variables")
	// 	os.Exit(1)
	// }

	// jwtIssuer, err := helpers.GetEnvStringVal("JWT_ISSUER")
	// if err != nil {
	// 	log.Error("JWT_ISSUER not found in environment variables")
	// 	os.Exit(1)
	// }

	// jwtExpiry, err := helpers.GetEnvIntVal("JWT_EXPIRED")
	// if err != nil {
	// 	log.Error("JWT_EXPIRED not found in environment variables")
	// 	os.Exit(1)
	// }

	// jwt service
	// jwtService := jwt.NewJWTService(jwtSecret, jwtIssuer, jwtExpiry)

	//initialize database
	db.GetClientConnection()
	// rediss.GetRedisClientConnection()

	router := gin.New()

	//middleware
	router.Use(middlewares.CORSMiddleware())
	router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))
	router.Use(gin.Recovery())

	//initialize routes
	// routes.CorporateUserRoutes(router)
	api := router.Group("api")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to User Service",
		})
	})
	supplier := api.Group("supplier")
	// supplier.Use(middlewares.NewAuthMiddleware(jwtService, false, false, true))
	routes.UserRoutes(supplier)

	err = router.Run(":" + port)

	fmt.Println("API running on port : 🚀 @ http://localhost:" + port)

	if err != nil {
		log.Error("Failed to start server.")
		os.Exit(1)
	} else {
		fmt.Println("API running on port : 🚀 @ http://localhost:" + port)
		log.Info("API running on port : 🚀 @ http://localhost:" + port)
	}

}
