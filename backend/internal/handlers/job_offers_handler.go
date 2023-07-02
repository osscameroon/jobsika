package handlers

import (
	"errors"
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
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "could not post job offer"})
		return
	}

	err := query.Validate()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest,
			gin.H{"error": err.Error()})
		return
	}

	var uploadImageFunc func(int64) (string, error)
	// if we have an image, we need to upload it
	if len(query.CompanyImage) > 0 {
		// check if the image provided is a valid image and have the right extension/mime type
		extensionDeatils := mimetype.Detect([]byte(query.CompanyImage))
		if _, ok := allowedEtensions[extensionDeatils.Extension()]; !ok {
			log.Error(errors.New("wrong image extension"))
			c.JSON(http.StatusBadRequest,
				gin.H{"error": "provided image has wrong extension"})
			return
		}

		uploadImageFunc = func(jobOfferID int64) (string, error) {
			fs, err := server.GetDefaultFileStorage()
			if err != nil {
				return "", err
			}
			return fs.UploadJobOfferCompanyPicture(query.CompanyImage, jobOfferID, extensionDeatils.String(), extensionDeatils.Extension())
		}
	} else {
		uploadImageFunc = nil
	}

	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not post job offer"})
		return
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
		HasImage:                offer.CompanyImageLocation != "",
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

	jobOffer, err := db.GetJobOfferById(jobOfferID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound,
			gin.H{"error": "could not find job offer"})
		return
	}

	if jobOffer.CompanyImageLocation == "" {
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

	image, err := fs.DownloadJobOfferCompanyPicture(jobOffer.CompanyImageLocation)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "could not return job offer image"})
		return
	}

	c.DataFromReader(http.StatusOK, *image.ContentLength, *image.ContentType, image.Body, nil)
}
