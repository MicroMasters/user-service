package helpers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetHTTPError(error_message string, http_error_code int, path string) (error_code int, response_message map[string]interface{}) {

	if http.StatusText(http_error_code) != "" {
		return http_error_code, gin.H{
			"timestamp": time.Now().Format(time.RFC3339Nano),
			"status":    http_error_code,
			"error":     http.StatusText(http_error_code),
			"message":   error_message,
			"path":      path,
		}
	} else {
		return 400, gin.H{
			"timestamp": time.Now().Format(time.RFC3339Nano),
			"status":    400,
			"error":     http.StatusText(400),
			"message":   error_message,
			"path":      path,
		}
	}

}

func IsArrayContains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}
