package controllers

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"user-service/src/connection/db"
	"user-service/src/helpers"
	"user-service/src/jwt"
	"user-service/src/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	collection := db.OpenCollection(db.Client, "MongoUsersRepository")

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	mongoUsersRepository.ID = uuid.New().String()

	plainPassword := mongoUsersRepository.Password
	mongoUsersRepository.Password = base64.StdEncoding.EncodeToString([]byte(plainPassword)) // Encode the plain text password

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

	// print mongoUsersRepository
	fmt.Println(mongoUsersRepository)

	JWT_SECRET, err := helpers.GetEnvStringVal("JWT_SECRET")

	if err != nil {
		log.Error("Failed to load environment variable : JWT_SECRET")
		println("Failed to load environment variable : JWT_SECRET " + err.Error())
		log.Debug(err.Error())
		os.Exit(1)
	}

	JWT_ISSUER, err := helpers.GetEnvStringVal("JWT_ISSUER")

	if err != nil {
		log.Error("Failed to load environment variable : JWT_ISSUER")
		println("Failed to load environment variable : JWT_ISSUER " + err.Error())
		log.Debug(err.Error())
		os.Exit(1)
	}

	JWT_EXPIRED, err := helpers.GetEnvIntVal("JWT_EXPIRED")

	if err != nil {
		log.Error("Failed to load environment variable : JWT_EXPIRED")
		println("Failed to load environment variable : JWT_EXPIRED " + err.Error())
		log.Debug(err.Error())
		os.Exit(1)
	}

	jwtService := jwt.NewJWTService(JWT_SECRET, JWT_ISSUER, JWT_EXPIRED)

	var token string

	switch mongoUsersRepository.Role {
	case "admin":
		token, err = jwtService.GenerateToken(mongoUsersRepository.ID, true, true, true, mongoUsersRepository.Email)
	case "buyer":
		token, err = jwtService.GenerateToken(mongoUsersRepository.ID, false, true, false, mongoUsersRepository.Email)
	case "supplier":
		token, err = jwtService.GenerateToken(mongoUsersRepository.ID, false, false, true, mongoUsersRepository.Email)
	}

	if err != nil {
		// Handle token generation error
		fmt.Println("Error generating token:", err)
	} else {
		fmt.Println("Generated token:", token)
		mongoUsersRepository.Token = token
		mongoUsersRepository.RefreshToken = token
	}

	// Insert Data
	_, err = collection.InsertOne(ctx, mongoUsersRepository)

	// print err
	fmt.Println(err)

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

	// log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Backend User Get All Function Called.")

	var backendUser models.MongoUsersRepository
	var backendUsers []models.MongoUsersRepository

	collection := db.OpenCollection(db.Client, "MongoUsersRepository")

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	findOptions := options.Find()
	// Generate Filtering Conditions
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		c.JSON(helpers.GetHTTPError("Failed while fetching Backend users.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(err.Error())
		return
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		cursor.Decode(&backendUser)
		backendUsers = append(backendUsers, backendUser)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(helpers.GetHTTPError("Failed while decoding Backend Users.", http.StatusInternalServerError, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(err.Error())
		return
	}

	if backendUsers == nil {
		c.JSON(http.StatusOK, []string{})
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("No Documents Found.")
		return
	}

	c.JSON(http.StatusOK, backendUsers)
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("All Backend Users details responsed.")
}

func GetUserByPhone(c *gin.Context) {
	log := helpers.GetLogger()

	// log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Backend User Get All by Phone Function Called.")

	phoneNumber := c.Param("phone_number")

	// Validate phone number
	if phoneNumber == "" {
		c.JSON(helpers.GetHTTPError("Invalid request parameters.", http.StatusBadRequest, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug("Invalid request parameters.")
		return
	}

	// Open users collection
	collection := db.OpenCollection(db.Client, "MongoUsersRepository")

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	// Execute a Database Query to check this phone number is exist.
	var userBackendUser models.MongoUsersRepository
	findResult := collection.FindOne(ctx, bson.M{"phone_number": phoneNumber}).Decode(&userBackendUser)

	if findResult != nil {
		// phone number Not Exists.
		c.JSON(helpers.GetHTTPError("phone number '"+phoneNumber+"' not found.", http.StatusNotFound, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug(findResult.Error())
		return
	}

	// Send created reponse with Status 201.
	c.JSON(http.StatusCreated, userBackendUser)
	log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("All Backend users details by phone responsed.")
}

func GetUserByID(c *gin.Context) {
	log := helpers.GetLogger()

	// log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("User Backend User Get All by ID Function Called.")

	userID := c.Param("id")

	println(userID)

	// Validate phone number
	if userID == "" {
		c.JSON(helpers.GetHTTPError("Invalid request parameters.", http.StatusBadRequest, c.FullPath()))
		// log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Debug("Invalid request parameters.")
		return
	}

	// Open users collection
	collection := db.OpenCollection(db.Client, "MongoUsersRepository")

	// Create a Background Context with Timeout Value configured as Environment Variable.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(API_CONFIG_REQUEST_TIMEOUT)*time.Second)

	// Execute a Database Query to check this phone number is exist.
	var userBackendUser models.MongoUsersRepository
	findResult := collection.FindOne(ctx, bson.M{"id": userID}).Decode(&userBackendUser)

	if findResult != nil {
		// phone number Not Exists.
		c.JSON(helpers.GetHTTPError("user with ID '"+userID+"' not found.", http.StatusNotFound, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error(findResult.Error())
		return
	}

	// Send created reponse with Status 201.
	c.JSON(http.StatusCreated, userBackendUser)
	// log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Info("All Backend User details by ID responsed.")
}


