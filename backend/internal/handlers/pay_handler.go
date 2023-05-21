package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/osscameroon/jobsika/internal/payment"
	"github.com/osscameroon/jobsika/internal/server"
	"github.com/osscameroon/jobsika/pkg/models/v1beta"
	log "github.com/sirupsen/logrus"
)

//PostPay handler for POST
func PostPay(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		log.Error(errors.New("Wrong contentType"))
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post rating"})
		return
	}

	query := v1beta.PayPostQuery{}
	if err := c.ShouldBind(&query); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post rating"})
		return
	}

	if err := query.Validate(); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	//Create a new opencollective tier
	//The deletion should happen on the webhook or a day after created
	paymentClient, err := server.GetDefaultPaymentClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not create payment"})
		return
	}
	response, err := paymentClient.CreateTier()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not create payment"})
		return
	}

	//Record customer payment request
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to get db client"})
		return
	}
	url := fmt.Sprintf("%s%s-%d", payment.OPEN_COLLECTIVE_CONTRIBUTE, response.CreateTier.Slug, response.CreateTier.LegacyID)
	err = db.CreatePaymentRecord(&v1beta.PaymentRecord{
		TierId:   response.CreateTier.ID,
		LegacyId: response.CreateTier.LegacyID,
		Slug:     response.CreateTier.Slug,
		Email:    query.Email,
		TierUrl:  url,
	})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find job titles"})
		return
	}

	//Send the link to the newly created opencollective tier to back to the client
	c.JSON(http.StatusOK, gin.H{"tier_url": url})
}
