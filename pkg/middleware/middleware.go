package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	database "github.com/onlinevoting/pkg/Database"
)

func Auth(c *gin.Context) {
	//get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECERET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		newcontroller, err := database.NewMyController()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to connect",
			})
		}
		// var user model.Users
		user, err := newcontroller.GetUserByEmail(claims["email"].(string))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func AuthVoter(c *gin.Context) {
	//get the cookie off req
	tokenString, err := c.Cookie("Authorisation")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECERET")), nil
	})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		newcontroller, err := database.NewMyController()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "failed to connect",
			})
		}
		// var user model.Users
		voter, err := newcontroller.GetVoterByEmail(claims["email"].(string))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("voter", voter)
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
