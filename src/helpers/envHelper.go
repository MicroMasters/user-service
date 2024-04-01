package helpers

import (
	"errors"
	"os"
	"strconv"
)

var envNotFound = errors.New("getENV : Environment variable not found")

func GetEnvStringVal(key string) (string, error) {

	keyVal := os.Getenv(key)

	if keyVal != "" {
		return keyVal, nil
	}

	return keyVal, envNotFound
}

func GetEnvIntVal(key string) (int, error) {
	keyVal, err := GetEnvStringVal(key)
	if err != nil {
		return 0, err
	}
	val, err := strconv.Atoi(keyVal)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func GetEnvBoolVal(key string) (bool, error) {
	keyVal, err := GetEnvStringVal(key)
	if err != nil {
		return false, err
	}
	val, err := strconv.ParseBool(keyVal)
	if err != nil {
		return false, err
	}
	return val, nil
}
