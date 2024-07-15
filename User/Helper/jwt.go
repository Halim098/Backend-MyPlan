package Helper

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(username string) (string, error) {
	tokenTTL, err := strconv.Atoi(os.Getenv("JWT_Time"))
	if err != nil {
		return "", err
	}
	
	claims := jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Minute * time.Duration(tokenTTL)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}