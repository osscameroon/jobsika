package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/elhmn/jobsika/internal/server"
	"github.com/elhmn/jobsika/pkg/models/v1beta"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func PostSubscribers(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		log.Error(errors.New("wrong contentType"))
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post subscriber"})
		return
	}

	var query v1beta.SubscribersPostQuery
	if err := c.ShouldBind(&query); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post subscriber"})
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
			gin.H{"error": "could not post subscriber"})
		return
	}

	subscriber, err := db.PostSubscribers(query)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not post subscriber"})
		return
	}

	c.JSON(http.StatusCreated, subscriber)
}
