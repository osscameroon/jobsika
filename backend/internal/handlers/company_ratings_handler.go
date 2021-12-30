package handlers

import (
	"net/http"

	"github.com/elhmn/camerdevs/internal/server"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//GetCompanyRatings return a list of company ratings
func GetCompanyRatings(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find company ratings"})
		return
	}

	ratings, err := db.GetCompanyRatings()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find company ratings"})
		return
	}

	c.JSON(http.StatusOK, ratings)
}
