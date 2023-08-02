package main

import (
	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"github.com/osscameroon/jobsika/internal/handlers"
	_ "github.com/osscameroon/jobsika/swagger"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
	}))

	//Health check endpoint
	router.GET("/health", handlers.Health)

	//Salaries
	router.GET("/salaries", handlers.GetSalaries)
	router.GET("/salaries/:id", handlers.GetSalaryByID)

	//Companies
	router.GET("/companies", handlers.GetCompanies)
	router.GET("/companies/:id", handlers.GetCompanyByID)

	//Jobtitles
	router.GET("/jobtitles", handlers.GetJobTitles)

	//CompanyRatings
	router.GET("/company-ratings", handlers.GetCompanyRatings)
	router.GET("/company-ratings/:id", handlers.GetCompanyRatingsByID)

	//Ratings
	router.GET("/ratings", handlers.GetRatings)
	router.GET("/ratings/:id", handlers.GetRatingByID)
	router.GET("/average-rating", handlers.GetAverageRating)
	router.POST("/ratings", handlers.PostRatings)

	//Seniority
	router.GET("/seniority", handlers.GetSeniority)

	//Cities
	router.GET("/cities", handlers.GetCities)

	//JobOffers
	jobsRouter := router.Group("/jobs")
	{
		jobsRouter.Use(limits.RequestSizeLimiter(5242880)) // 5MB. if request body is larger than 5MB, it will return 413 error
		jobsRouter.GET("", handlers.GetJobOffers)
		jobsRouter.POST("", handlers.PostJobOffer)
		jobsRouter.GET("/:id/image", handlers.GetJobOfferImage)
	}

	//Subscribers
	router.POST("/subscribers", handlers.PostSubscribers)

	//Pay
	router.POST("/pay", handlers.PostPay)

	//GetOrderID
	router.GET("/getorder", handlers.GetOrderID)

	//OpenCollectionWebhook
	router.POST("/open-collective-webhook", handlers.OpenCollectiveWebhook)

	if err := router.Run(":7001"); err != nil {
		return
	}
}
