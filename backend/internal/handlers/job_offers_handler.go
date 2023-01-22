package handlers

import (
	"errors"
	"github.com/elhmn/jobsika/internal/server"
	"github.com/elhmn/jobsika/pkg/models/v1beta"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

// PostJobOffer handles /offers POST endpoint
func PostJobOffer(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		log.Error(errors.New("Wrong contentType"))
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post job offer"})
		return
	}

	query := v1beta.OfferPostQuery{}
	if err := c.ShouldBind(&query); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post job offer"})
	}

	err := query.Validate()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not post job offer"})
		return
	}

	err = db.PostJobOffer(query)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not post job offer"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}
