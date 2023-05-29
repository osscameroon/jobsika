package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

	//We can't proceed with a payment on test mode as this will create unecessary
	//payment on opencollective
	if os.Getenv("TEST") == "true" {
		//Send the link to the newly created opencollective tier to back to the client
		c.JSON(http.StatusOK, gin.H{"tier_url": "https://opencollective.com/osscameroon/something-something"})
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

	jobOfferID, err := strconv.ParseInt(query.JobOfferID, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to parse job offer id"})
		return
	}

	url := fmt.Sprintf("%s%s-%d", paymentClient.GetContributionURL(), response.CreateTier.Slug, response.CreateTier.LegacyID)
	err = db.CreatePaymentRecord(&v1beta.PaymentRecord{
		TierId:     response.CreateTier.ID,
		JobOfferID: jobOfferID,
		LegacyId:   response.CreateTier.LegacyID,
		Slug:       response.CreateTier.Slug,
		Email:      query.Email,
		TierUrl:    url,
	})
	if err != nil {
		retry := 0
		for retry < 3 {
			if err := paymentClient.DeleteTier(response.CreateTier.LegacyID); err != nil {
				retry++
				time.Sleep(time.Duration(retry*100) * time.Millisecond)
			} else {
				break
			}
		}

		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to create payment record"})
		return
	}

	//Send the link to the newly created opencollective tier to back to the client
	c.JSON(http.StatusOK, gin.H{"tier_url": url})
}
