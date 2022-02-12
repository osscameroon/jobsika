package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/elhmn/camerdevs/internal/server"
	"github.com/elhmn/camerdevs/pkg/models/v1beta"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//GetRatings handles user sign in
func GetRatings(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find ratings"})
		return
	}

	//Get parameters
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "20")
	jobtitle := c.Query("jobtitle")
	company := c.Query("company")

	ratings, err := db.GetRatings(page, limit, jobtitle, company)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find ratings"})
		return
	}

	c.JSON(http.StatusOK, ratings)
}

//GetRatingByID returns a rating by id
func GetRatingByID(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find ratings"})
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

	rating, err := db.GetRatingByID(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound,
			gin.H{"error": "could not find rating"})
		return
	}

	c.JSON(http.StatusOK, rating)
}

//GetAverageRating handles /average-rating GET endpoint
func GetAverageRating(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find ratings"})
		return
	}

	//Get parameters
	jobtitle := c.Query("jobtitle")
	company := c.Query("company")

	rating, err := db.GetAverageRating(jobtitle, company)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find average rating"})
		return
	}

	c.JSON(http.StatusOK, rating)
}

//PostRatings handles /ratings POST endpoint
func PostRatings(c *gin.Context) {
	if c.Request.Header.Get("Content-Type") != "application/json" {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post rating"})
		return
	}

	query := v1beta.RatingPostQuery{}
	if err := c.ShouldBind(&query); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post rating"})
	}

	fmt.Printf("query: %+v\n", query) // Debug
	c.JSON(http.StatusCreated, "ok")
}
