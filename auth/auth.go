package auth

import (
	"book_inventory/models"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func HomeHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "login.html")
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"content": ""})
}

func LoginPostHandler(c *gin.Context) {
	var credential models.Credentials

	err := c.Bind(&credential)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "binding error",
		})
	}

	if credential.Username != os.Getenv("DB_USER") || credential.Password != os.Getenv("DB_PASS") {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "username or password invalid",
		})
	} else {
		claim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "issue",
			IssuedAt:  time.Now().Unix(),
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

		secret := os.Getenv("TOKEN_SECRET")

		token, err := sign.SignedString([]byte(secret))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"content": "token signing error",
			})
			c.Abort()
		}

		q := url.Values{}

		q.Set("auth", token)

		location := url.URL{Path: "/books", RawQuery: q.Encode()}

		c.Redirect(http.StatusMovedPermanently, location.RequestURI())
	}
}
