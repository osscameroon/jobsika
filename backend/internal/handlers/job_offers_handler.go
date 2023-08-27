package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/osscameroon/jobsika/internal/server"
	"github.com/osscameroon/jobsika/pkg/models/v1beta"
	log "github.com/sirupsen/logrus"
)

var (
	allowedEtensions = map[string]struct{}{
		".jpg":  {},
		".jpeg": {},
		".png":  {},
		".gif":  {},
		".svg":  {},
		".webp": {},
	}
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
		// we check if the resuest body limiter middleware already responded with a 413 error. We do this avoid error when overrriding HTTP headers and body response
		if c.Writer.Status() == http.StatusRequestEntityTooLarge {
			return
		}
		c.JSON(http.StatusBadRequest,
			gin.H{"error": fmt.Sprintf("could not post job offer: %s", err.Error())})
		return
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

	// if we have an image, we need to upload it
	var hasImage bool
	if len(query.CompanyImage) > 0 {
		// check if the image provided is a valid image and have the right extension/mime type
		extensionDeatils := mimetype.Detect([]byte(query.CompanyImage))
		if _, ok := allowedEtensions[extensionDeatils.Extension()]; !ok {
			log.Error(errors.New("wrong image extension"))
			c.JSON(http.StatusBadRequest,
				gin.H{"error": "provided image has wrong extension"})
			return
		}

		fs, err := server.GetDefaultFileStorage()
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": "could not post job offer. Image upload failed"})
			return
		}
		location, err := fs.UploadJobOfferCompanyPicture(query.CompanyImage, offer.ID, extensionDeatils.String(), extensionDeatils.Extension())
		if err != nil {
			db.DeleteJobOffer(offer.ID)
			log.Error(err)
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": "could not post job offer. Image upload failed"})
			return
		}

		err = db.PostJobOfferImage(offer.ID, location)
		if err != nil {
			db.DeleteJobOffer(offer.ID)
			log.Error(err)
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": "could not post job offer image"})
			return
		}

		hasImage = true
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
		HasImage:                hasImage,
	})
}

//GetJobOfferImage handles /jobs/:id/image GET endpoint
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
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "failed to parse job offer id"})
		return
	}

	jobOfferImage, err := db.GetJobOfferImageByJobOfferId(jobOfferID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound,
			gin.H{"error": "could not find job offer image for this job offer"})
		return
	}

	fs, err := server.GetDefaultFileStorage()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not find job offer image"})
		return
	}

	image, err := fs.DownloadJobOfferCompanyPicture(jobOfferImage.ImageLocation)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not return job offer image"})
		return
	}

	c.DataFromReader(http.StatusOK, *image.ContentLength, *image.ContentType, image.Body, nil)
}
