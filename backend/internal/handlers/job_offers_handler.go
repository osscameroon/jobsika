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

func GetJobOffers(c *gin.Context) {
	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find job offers"})
		return
	}

	//Get parameters
	var query v1beta.GetJobOffersQuery
	query.Page = c.DefaultQuery("page", "1")
	query.Limit = c.DefaultQuery("limit", "20")
	query.JobTitle = c.Query("jobtitle")
	query.Company = c.Query("company")
	query.City = c.Query("city")
	query.Country = c.Query("country")
	query.IsRemote = c.Query("isRemote")

	offers, err := db.GetJobOffers(query)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find job offers"})
		return
	}

	c.JSON(http.StatusOK, offers)
}

// PostJobOffer handles /offers POST endpoint
func PostJobOffer(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		log.Error(errors.New("wrong contentType"))
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

	offer, err := db.PostJobOffer(query)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not post job offer"})
		return
	}

	c.JSON(http.StatusCreated, v1beta.JobOfferPresenter{
		ID:                offer.ID,
		CreatedAt:         offer.CreatedAt,
		UpdatedAt:         offer.UpdatedAt,
		CompanyName:       offer.CompanyName,
		JobTitle:          query.JobTitle,
		IsRemote:          offer.IsRemote,
		City:              offer.City,
		Country:           offer.Country,
		Department:        offer.Department,
		SalaryRangeMin:    offer.SalaryRangeMin,
		SalaryRangeMax:    offer.SalaryRangeMax,
		Description:       offer.Description,
		Benefits:          offer.Benefits,
		HowToApply:        offer.HowToApply,
		ApplyUrl:          offer.ApplyUrl,
		ApplyEmailAddress: offer.ApplyEmailAddress,
		ApplyPhoneNumber:  offer.ApplyPhoneNumber,
		Tags:              offer.Tags,
	})
}
