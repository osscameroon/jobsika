package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/osscameroon/jobsika/internal/server"
	log "github.com/sirupsen/logrus"
)

// GetSalaries handles user sign in
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

// GetSalaryByID returns a salary by id
func GetSalaryByID(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find salaries"})
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

	salary, err := db.GetSalaryByID(id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound,
			gin.H{"error": "could not find salary"})
		return
	}

	c.JSON(http.StatusOK, salary)
}
