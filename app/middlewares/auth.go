package middlewares

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

// verify token
func VerifyToken(tokenString string) (data *jwt.Token, err error) {
	data, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte("secret"), nil
	})
	return
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		data, err := VerifyToken(strings.Split(tokenString, " ")[1])
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Request.Header.Set("username", data.Claims.(jwt.MapClaims)["username"].(string))
		c.Next()
	}
}
