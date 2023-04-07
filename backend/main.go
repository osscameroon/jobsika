package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/osscameroon/jobsika/docs"
	"github.com/osscameroon/jobsika/internal/handlers"
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
	router.GET("/jobs", handlers.GetJobOffers)
	router.POST("/jobs", handlers.PostJobOffer)

	//Subscribers
	router.POST("/subscribers", handlers.PostSubscribers)

	if err := router.Run(":7000"); err != nil {
		return
	}
}
