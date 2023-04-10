package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osscameroon/jobsika/internal/server"
	log "github.com/sirupsen/logrus"
)

// GetJobTitles return a list of job titles
func GetJobTitles(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find job titles"})
		return
	}

	titles, err := db.GetJobTitles()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find job titles"})
		return
	}

	c.JSON(http.StatusOK, titles)
}
