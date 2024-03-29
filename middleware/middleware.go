package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthValidator(c *gin.Context) {
	tokenString := c.Query("auth")

	if tokenString == "" {
		tokenString = c.PostForm("auth")
		if tokenString == "" {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{
				"content": "token not found",
			})
			c.Abort()
		}
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, valid := t.Method.(*jwt.SigningMethodHMAC); !valid {
			return nil, fmt.Errorf("invalid token: %s", t.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if token != nil && err == nil {
		fmt.Println("token verified")
		c.Next()
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"content": err.Error(),
		})
		c.Abort()
	}
}
