package handlers

import (
	"github.com/elhmn/jobsika/internal/server"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetCities returns a list of cities
func GetCities(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find job titles"})
		return
	}

	cities, err := db.GetCities()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find job titles"})
		return
	}

	c.JSON(http.StatusOK, cities)
}
