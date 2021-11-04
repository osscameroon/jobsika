package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//GetUser get user
func GetUser(c *gin.Context) {
	fmt.Printf("GetUser")
}

//GetUsers get users
func GetUsers(c *gin.Context) {
	fmt.Printf("GetUsers")
}

//PostUser post users data
func PostUser(c *gin.Context) {
	fmt.Printf("PostUser")
}
