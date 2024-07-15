package Middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("JWT_SECRET"))

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenUser := c.Request.Header.Get("Authorization")
	
		tokenUser = strings.Replace(tokenUser, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenUser, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})


		if err != nil{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized 1"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("id", claims["id"])
			c.Set("username", claims["username"])
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized 2"})
		}
	}
}

