package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/osscameroon/jobsika/internal/server"
	log "github.com/sirupsen/logrus"
)

type WebhookPayload struct {
	CreatedAt    time.Time `json:"createdAt"`
	ID           int       `json:"id"`
	CollectiveID int       `json:"CollectiveId"`
	Type         string    `json:"type"`
	Data         struct {
		FromCollective struct {
			ID            int    `json:"id"`
			Type          string `json:"type"`
			Slug          string `json:"slug"`
			Name          string `json:"name"`
			TwitterHandle string `json:"twitterHandle"`
			GithubHandle  string `json:"githubHandle"`
			RepositoryURL string `json:"repositoryUrl"`
			Image         string `json:"image"`
		} `json:"fromCollective"`
		Collective struct {
			ID            int         `json:"id"`
			Type          string      `json:"type"`
			Slug          string      `json:"slug"`
			Name          string      `json:"name"`
			TwitterHandle interface{} `json:"twitterHandle"`
			GithubHandle  string      `json:"githubHandle"`
			RepositoryURL string      `json:"repositoryUrl"`
			Image         string      `json:"image"`
		} `json:"collective"`
		Transaction struct {
			ID                                int         `json:"id"`
			Kind                              string      `json:"kind"`
			Type                              string      `json:"type"`
			UUID                              string      `json:"uuid"`
			Group                             string      `json:"group"`
			Amount                            int         `json:"amount"`
			IsDebt                            bool        `json:"isDebt"`
			OrderID                           int64       `json:"OrderId"`
			Currency                          string      `json:"currency"`
			IsRefund                          bool        `json:"isRefund"`
			ExpenseID                         interface{} `json:"ExpenseId"`
			CreatedAt                         time.Time   `json:"createdAt"`
			TaxAmount                         interface{} `json:"taxAmount"`
			Description                       string      `json:"description"`
			CollectiveID                      int         `json:"CollectiveId"`
			HostCurrency                      string      `json:"hostCurrency"`
			CreatedByUserID                   int         `json:"CreatedByUserId"`
			FromCollectiveID                  int         `json:"FromCollectiveId"`
			AmountInHostCurrency              int         `json:"amountInHostCurrency"`
			HostFeeInHostCurrency             int         `json:"hostFeeInHostCurrency"`
			NetAmountInHostCurrency           int         `json:"netAmountInHostCurrency"`
			PlatformFeeInHostCurrency         int         `json:"platformFeeInHostCurrency"`
			UsingGiftCardFromCollectiveID     interface{} `json:"UsingGiftCardFromCollectiveId"`
			NetAmountInCollectiveCurrency     int         `json:"netAmountInCollectiveCurrency"`
			AmountSentToHostInHostCurrency    int         `json:"amountSentToHostInHostCurrency"`
			PaymentProcessorFeeInHostCurrency int         `json:"paymentProcessorFeeInHostCurrency"`
			FormattedAmount                   string      `json:"formattedAmount"`
			FormattedAmountWithInterval       string      `json:"formattedAmountWithInterval"`
		} `json:"transaction"`
	} `json:"data"`
}

func OpenCollectiveWebhook(c *gin.Context) {
	var payload WebhookPayload
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	paymentClient, err := server.GetDefaultPaymentClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not create payment client"})
		return
	}
	response, err := paymentClient.GetOrder(payload.Data.Transaction.OrderID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not create payment client"})
		return
	}

	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to get db client"})
		return
	}

	tierID := response.Order.Tier.LegacyID
	_, err = db.GetPaymentRecordByID(tierID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to get payment record"})
		return
	}

	if err := paymentClient.DeleteTier(tierID); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to delete tier"})
		return
	}

	log.Info("OpenCollectiveWebhook was triggered!")
}

// GetOrderID retrieves an open collective order id
// This endpoint was created for testing purposes
func GetOrderID(c *gin.Context) {
	orderID, err := strconv.ParseInt(c.Query("orderID"), 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to parse order id"})
		return
	}

	paymentClient, err := server.GetDefaultPaymentClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not create payment client"})
		return
	}
	response, err := paymentClient.GetOrder(orderID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not create payment client"})
		return
	}

	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to get db client"})
		return
	}

	tierID := response.Order.Tier.LegacyID
	paymentRecord, err := db.GetPaymentRecordByID(tierID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to get payment record"})
		return
	}
	fmt.Printf("paymentRecord: %+v\n", paymentRecord)

	fmt.Printf("legacyID: %d\n", response.Order.Tier.LegacyID)
	log.Info("GetOrders was triggered!")
}
