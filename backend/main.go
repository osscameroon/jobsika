package main

import (
	_ "github.com/elhmn/camerdevs/docs"
	"github.com/elhmn/camerdevs/internal/handlers"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposedHeaders:   []string{},
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

	//CompanyRatings
	router.GET("/company-ratings", handlers.GetCompanyRatings)
	router.GET("/company-ratings/:id", handlers.GetCompanyRatingsByID)

	//Seniority
	router.GET("/seniority", handlers.GetSeniority)

	if err := router.Run(":7000"); err != nil {
		return
	}
}
