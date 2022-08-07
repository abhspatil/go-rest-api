package middlewares

import (
	"fmt"
	"net/http"
	"users-api-with-orm-with-jwt-auth/auth"
	exception "users-api-with-orm-with-jwt-auth/exception"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func IsAuthorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		fmt.Println("token, ", token)

		if token == "" {
			var err exception.GenericError
			err.SetError("no token found")

			fmt.Print("no token found")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var tokenObj *jwt.Token
		tokenObj, err := auth.ValidateJWTToken(token)

		fmt.Print(tokenObj.Claims)

		if err != nil {
			fmt.Print("invalid token")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		fmt.Println("validated token")

		ctx.Next()
	}
}
