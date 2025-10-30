package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/PatrochR/disko/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		trimHeader := strings.TrimPrefix(header , "Bearer ")

		secretKey := os.Getenv("SECRET_KEY")
		token, err := jwt.Parse(trimHeader, func(t *jwt.Token) (any, error) {
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error ": "Unauthorized"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(user.CustomeClaim)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error ": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("userID" , claims.ID)
		ctx.Set("username" , claims.Username)

		ctx.Next()
	}
}
