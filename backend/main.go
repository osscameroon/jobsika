package main

import (
	"github.com/elhmn/camerdevs/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//Health check endpoint
	router.GET("/health", handlers.Health)

	//Salaries
	router.GET("/salaries", handlers.GetSalaries)
	router.GET("/salaries/:id", handlers.GetSalaryByID)

	//Companies
	router.GET("/companies", handlers.GetCompanies)

	//CompanyRatings
	router.GET("/company-ratings", handlers.GetCompanyRatings)
	router.GET("/company-ratings/:id", handlers.GetCompanyRatingsByID)

	//Seniority
	router.GET("/seniority", handlers.GetSeniority)

	if err := router.Run(":7000"); err != nil {
		return
	}
}
