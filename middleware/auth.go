package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		var value string
		var err error

		tokenString, err := ctx.Request.Cookie("session_token")
		if err == nil {
			value = tokenString.Value
		}
		
		if err != nil && ctx.Request.URL.Path != "/" {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		
		if value == "" {
			ctx.String(http.StatusSeeOther, "Unauthorized")
			ctx.Abort()
			return
		}


		var Claims model.Claims

		_, err = jwt.ParseWithClaims(value, &Claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(model.JwtKey), nil
		})

		if err != nil {
			ctx.String(http.StatusBadRequest, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Set("email", Claims.Email)
		
		ctx.Next()
	})
}
