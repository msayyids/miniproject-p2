package middleware

import (
	"miniproject/controller"
	"miniproject/helper"
	"miniproject/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	Authentication repo.Repo
}

func (a *Auth) AuthUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, controller.NewErrorResponse(402, "invalid request header"))
			ctx.Abort()
			return

		}

		claims, err := helper.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, controller.NewErrorResponse(401, "invalid token, please provide valid access token"))
			ctx.Abort()
			return
		}

		id := claims["id"].(float64)
		loggedInUser, err := a.Authentication.FindById(int(id))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, controller.NewErrorResponse(401, "missing id claims"))
		}
		ctx.Set("loggedInUser", loggedInUser.Id)
		ctx.Next()
	}
}
