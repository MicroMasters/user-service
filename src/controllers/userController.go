package controllers

import (
	"net/http"
	"user-service/src/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers(c *gin.Context) {
	log := helpers.GetLogger()

	c.Set("LogID", uuid.New().String())
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Get All Function Called.")

	// return a json response with all drivers {id : 1, name : "driver1"}

	c.JSON(http.StatusOK, gin.H{
		"id":     1,
		"name":   "driver1",
		"status": http.StatusOK,
	})

}
