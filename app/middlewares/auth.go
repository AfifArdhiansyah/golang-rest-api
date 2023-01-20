package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// generate jwt token
func GenerateToken(username string, password string) (tokenString string, err error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte("secret"))
	return
}
