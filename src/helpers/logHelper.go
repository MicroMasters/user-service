package helpers

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
)

func GetLogger() *logrus.Logger {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logFilePath := os.Getenv("LOG_FILE_PATH")
	logFileName := os.Getenv("LOG_FILE_NAME")
	logLevel := os.Getenv("LOG_LEVEL")

	fileName := path.Join(logFilePath, logFileName)

	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("err", err.Error())
	}

	logger := logrus.New()

	logger.Out = src

	loggersLogLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		loggersLogLevel = logrus.InfoLevel
	}

	logger.SetLevel(loggersLogLevel)

	logger.SetFormatter(&logrus.JSONFormatter{})

	return *&logger

}
