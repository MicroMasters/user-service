package controllers

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"user-service/src/connection/db"
	"user-service/src/helpers"
	"user-service/src/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMongoUser(c *gin.Context) {
	log := helpers.GetLogger()

	c.Set("LogID", uuid.New().String())
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Create Function Called.")

	var mongoUsersRepository models.MongoUsersRepository
	json.NewDecoder(c.Request.Body).Decode(&mongoUsersRepository)

	validationErr := helpers.StructValidator(mongoUsersRepository)

	if validationErr != nil {
		c.JSON(helpers.GetHTTPError("Invalid Request Parameters", http.StatusBadRequest, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug("Invalid request parameters.")
		return
	}

	collection := db.OpenCollection(db.GetClientConnection(), "MongoUsersRepository")

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	mongoUsersRepository.ID = uuid.New().String()

	decPass, decErr := base64.StdEncoding.DecodeString(mongoUsersRepository.Password)

	if decErr != nil {
		// Password decoding error
		c.JSON(helpers.GetHTTPError("Failed while fetching password", http.StatusNotFound, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(decErr.Error())
		return
	}

	// Hash Password
	hashPass := md5.Sum(decPass)

	mongoUsersRepository.Password = fmt.Sprintf("%x", hashPass)

	mongoUsersRepository.CreatedTime = time.Now()

	// Insert Data
	_, err := collection.InsertOne(ctx, mongoUsersRepository)

	// Check errors
	if err != nil {

		// User Email Already Exists
		if mongo.IsDuplicateKeyError(err) {
			c.JSON(helpers.GetHTTPError("User Phone Number '"+mongoUsersRepository.PhoneNumber+"' already taken.", http.StatusConflict, c.FullPath()))
			log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug("User Phone Number '" + mongoUsersRepository.PhoneNumber + "' already taken.")
			return
		}

		// User creation failed.
		c.JSON(helpers.GetHTTPError("Failed to create  User '"+mongoUsersRepository.FirstName+"'.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(err.Error())
		return
	}

	// Send created reponse with Status 201.
	c.JSON(http.StatusCreated, mongoUsersRepository)
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Backend User Created successfully.")

}

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
