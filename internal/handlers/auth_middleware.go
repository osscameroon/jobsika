package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/elhmn/camerdevs/internal/server"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//AuthMiddleware checks that the user is authorized to hit that endpoint
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Middleware is being called")
		auth := c.GetHeader("Authorization")
		auth = strings.TrimSpace(auth)

		//if there is no Authorization header set
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not found"})
			c.Abort()
			return
		}

		//Get token
		w := "Bearer"
		i := strings.Index(auth, w)
		if i < 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token wrong format"})
			c.Abort()
			return
		}
		token := auth[i+len(w):]
		token = strings.TrimSpace(token)

		//Verify token
		_, err := verifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func verifyToken(token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(server.GetDefaultConfig().JWTKey), nil
	})

	if err != nil {
		log.Error(err)
		if err == jwt.ErrSignatureInvalid {
			return claims, errors.New("token wrong signature")
		}
		return claims, errors.New("token wrong format")
	}

	if !tkn.Valid {
		return claims, errors.New("token has expired")
	}
	return claims, nil
}
