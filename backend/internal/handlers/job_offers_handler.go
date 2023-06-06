package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/osscameroon/jobsika/internal/server"
	"github.com/osscameroon/jobsika/pkg/models/v1beta"
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

	var uploadImageFunc func(int64) error
	// if we have an image, we need to upload it
	if len(query.CompanyImage) > 0 {
		uploadImageFunc = func(jobOfferID int64) error {
			fs, err := server.GetDefaultFileStorage()
			if err != nil {
				return err
			}
			_, err = fs.UploadJobOfferCompanyPicture(query.CompanyImage, jobOfferID)
			return err
		}
	} else {
		uploadImageFunc = nil
	}

	offer, err := db.PostJobOffer(query, uploadImageFunc)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not post job offer"})
		return
	}

	c.JSON(http.StatusCreated, v1beta.JobOfferPresenter{
		ID:                      offer.ID,
		CreatedAt:               offer.CreatedAt,
		UpdatedAt:               offer.UpdatedAt,
		CompanyName:             offer.CompanyName,
		JobTitle:                query.JobTitle,
		IsRemote:                offer.IsRemote,
		City:                    offer.City,
		Country:                 offer.Country,
		Department:              offer.Department,
		SalaryRangeMin:          offer.SalaryRangeMin,
		SalaryRangeMax:          offer.SalaryRangeMax,
		Description:             offer.Description,
		Benefits:                offer.Benefits,
		HowToApply:              offer.HowToApply,
		ApplicationUrl:          offer.ApplicationUrl,
		ApplicationEmailAddress: offer.ApplicationEmailAddress,
		ApplicationPhoneNumber:  offer.ApplicationPhoneNumber,
		Tags:                    offer.Tags,
	})
}

//GetJobOfferImage handles /offers/:id/image GET endpoint
func GetJobOfferImage(c *gin.Context) {
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find job offer"})
		return
	}

	jobOfferIDStr := c.Param("id")
	jobOfferID, err := strconv.ParseInt(jobOfferIDStr, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to parse job offer id"})
		return
	}

	jobOffer, err := db.GetJobOfferById(jobOfferID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound,
			gin.H{"error": "could not find job offer"})
		return
	}

	if !jobOffer.HasImage {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "job offer does not have an image"})
		return
	}

	fs, err := server.GetDefaultFileStorage()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find job offer image"})
		return
	}

	image, err := fs.DownloadJobOfferCompanyPicture(jobOffer.ID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not return job offer image"})
		return
	}

	c.DataFromReader(http.StatusOK, *image.ContentLength, *image.ContentType, image.Body, nil)
}
