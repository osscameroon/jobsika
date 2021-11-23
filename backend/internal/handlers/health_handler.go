package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Health always returns 200 status code
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "all good sanix"})
}
