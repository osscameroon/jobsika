package handlers

import (
	"net/http"

	"github.com/elhmn/camerdevs/internal/server"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//GetSalaries handles user sign in
func GetSalaries(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find salaries"})
		return
	}

	salaries, err := db.GetSalaries()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find salaries"})
		return
	}

	c.JSON(http.StatusOK, salaries)
}
