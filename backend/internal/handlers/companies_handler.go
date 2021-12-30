package handlers

import (
	"net/http"

	"github.com/elhmn/camerdevs/internal/server"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//GetCompanies return a list of companies
func GetCompanies(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find companies"})
		return
	}

	companies, err := db.GetCompanies()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find companies"})
		return
	}

	c.JSON(http.StatusOK, companies)
}
