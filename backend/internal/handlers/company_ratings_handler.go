package handlers

import (
	"net/http"
	"strconv"

	"github.com/elhmn/camerdevs/internal/server"
	"github.com/elhmn/camerdevs/pkg/models/v1beta"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	var query v1beta.CompanyRatingQuery
	if err := c.ShouldBindWith(&query, binding.Query); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "bad query parameters"})
		return
	}

	ratings, err := db.GetCompanyRatings(query)
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
		return
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
