package middlewares

import (
	"net/http"
	"strings"

	"user-service/src/constants"
	"user-service/src/helpers"
	"user-service/src/jwt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthMiddleware struct {
	jwtService jwt.JWTService
	isAdmin    bool
	isBuyer    bool
	isSupplier bool
}

func NewAuthMiddleware(jwtService jwt.JWTService, isAdmin bool) gin.HandlerFunc {
	return (&AuthMiddleware{
		jwtService: jwtService,
		isAdmin:    isAdmin,
	}).Handle
}

func (m *AuthMiddleware) Handle(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(helpers.GetHTTPError("Missing Authorization Header", http.StatusUnauthorized, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error("Missing Authorization Header")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.JSON(helpers.GetHTTPError("Invalid Header Format", http.StatusUnauthorized, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error("Invalid Header Format")
		return
	}

	if headerParts[0] != "Bearer" {
		c.JSON(helpers.GetHTTPError("Token must content bearer", http.StatusUnauthorized, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error("Token must content bearer")
		return
	}

	user, err := m.jwtService.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(helpers.GetHTTPError("Invalid Token", http.StatusUnauthorized, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error("Invalid Token")
		return
	}

	if user.IsAdmin != m.isAdmin && !user.IsAdmin {
		c.JSON(helpers.GetHTTPError("You don't have access for this action", http.StatusUnauthorized, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error("You don't have access for this action")
		return
	} else if user.IsBuyer != m.isBuyer && !user.isBuyer {
		c.JSON(helpers.GetHTTPError("You don't have access for this action", http.StatusUnauthorized, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error("You don't have access for this action")
		return
	} else if user.isSupplier != m.isSupplier && !user.isSupplier {
		c.JSON(helpers.GetHTTPError("You don't have access for this action", http.StatusUnauthorized, c.FullPath()))
		log.WithFields(logrus.Fields{"ID": c.MustGet("LogID")}).Error("You don't have access for this action")
		return
	}

	ctx.Set(constants.CtxAuthenticatedUserKey, user)
	ctx.Next()
}
