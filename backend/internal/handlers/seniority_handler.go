package handlers

import (
	"net/http"

	"github.com/elhmn/camerdevs/pkg/models/v1beta"
	"github.com/gin-gonic/gin"
)

//GetSeniority handles user sign in
func GetSeniority(c *gin.Context) {
	seniority := []string{
		v1beta.SeniorityEntryLevel,
		v1beta.SeniorityMidLevel,
		v1beta.SenioritySenior,
		v1beta.SeniorityAboveSenior,
		v1beta.SeniorityExecutive,
	}

	c.JSON(http.StatusOK, seniority)
}
