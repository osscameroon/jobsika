package handlers

import (
	"net/http"
	"strconv"

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
		c.JSON(http.StatusNotFound,
			gin.H{"error": "could not find company ratings"})
		return
	}

	c.JSON(http.StatusOK, ratings)
}

//GetCompanyRatingsByID return company ratings by id
func GetCompanyRatingsByID(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find company ratings"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "failed to parse parameters"})
	}

	ratings, err := db.GetCompanyRatingsByID(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound,
			gin.H{"error": "could not find company ratings"})
		return
	}

	c.JSON(http.StatusOK, ratings)
}
