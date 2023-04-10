package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/osscameroon/jobsika/pkg/models/v1beta"
)

// GetSeniority returns a list of seniority levels
func GetSeniority(c *gin.Context) {
	seniority := []string{
		v1beta.SeniorityIntern,
		v1beta.SeniorityEntryLevel,
		v1beta.SeniorityMidLevel,
		v1beta.SenioritySenior,
		v1beta.SeniorityAboveSenior,
		v1beta.SeniorityExecutive,
	}

	c.JSON(http.StatusOK, seniority)
}
