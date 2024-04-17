package middlewares

import (
	"fmt"
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

func NewAuthMiddleware(jwtService jwt.JWTService, isAdmin bool, isBuyer bool, isSupplier bool) gin.HandlerFunc {
	return (&AuthMiddleware{
		jwtService: jwtService,
		isAdmin:    isAdmin,
		isBuyer:    isBuyer,
		isSupplier: isSupplier,
	}).Handle
}

func (m *AuthMiddleware) Handle(ctx *gin.Context) {
	log := helpers.GetLogger()
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(helpers.GetHTTPError("Missing Authorization Header", http.StatusUnauthorized, ctx.FullPath()))
		log.WithFields(logrus.Fields{"ID": ctx.MustGet("LogID")}).Error("Missing Authorization Header")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		ctx.JSON(helpers.GetHTTPError("Invalid Header Format", http.StatusUnauthorized, ctx.FullPath()))
		log.WithFields(logrus.Fields{"ID": ctx.MustGet("LogID")}).Error("Invalid Header Format")
		return
	}

	if headerParts[0] != "Bearer" {
		ctx.JSON(helpers.GetHTTPError("Token must content bearer", http.StatusUnauthorized, ctx.FullPath()))
		log.WithFields(logrus.Fields{"ID": ctx.MustGet("LogID")}).Error("Token must content bearer")
		return
	}

	user, err := m.jwtService.ParseToken(headerParts[1])
	// print user
	fmt.Println(user)
	fmt.Println(m.isAdmin, m.isBuyer, m.isSupplier)
	if err != nil {
		ctx.JSON(helpers.GetHTTPError("Invalid Token", http.StatusUnauthorized, ctx.FullPath()))
		log.WithFields(logrus.Fields{"ID": ctx.MustGet("LogID")}).Error("Invalid Token")
		return
	}

	if user.IsAdmin != m.isAdmin && !user.IsAdmin && m.isBuyer && !user.IsBuyer && m.isSupplier && !user.IsSupplier {
		ctx.JSON(helpers.GetHTTPError("You don't have access for this action as a admin", http.StatusUnauthorized, ctx.FullPath()))
		log.WithFields(logrus.Fields{"ID": ctx.MustGet("LogID")}).Error("You don't have access for this action as a admin")
		return
	} else if user.IsBuyer != m.isBuyer && !user.IsBuyer {
		ctx.JSON(helpers.GetHTTPError("You don't have access for this action as a buyer", http.StatusUnauthorized, ctx.FullPath()))
		log.WithFields(logrus.Fields{"ID": ctx.MustGet("LogID")}).Error("You don't have access for this action as a buyer")
		return
	} else if user.IsSupplier != m.isSupplier && !user.IsSupplier {
		ctx.JSON(helpers.GetHTTPError("You don't have access for this action as a supplier", http.StatusUnauthorized, ctx.FullPath()))
		log.WithFields(logrus.Fields{"ID": ctx.MustGet("LogID")}).Error("You don't have access for this action as a supplier")
		return
	}

	ctx.Set(constants.CtxAuthenticatedUserKey, user)
	ctx.Next()
}
