package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/elhmn/camerdevs/internal/server"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

//Crendentials is a user creds struct
type Crendentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claims is a JWT claims struct
type Claims struct {
	Userid int64 `json:"userid"`
	jwt.StandardClaims
}

//Signup handles user sign up
func Signup(c *gin.Context) {
	creds := Crendentials{}

	//Get creds from body
	if err := c.ShouldBind(&creds); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "body should contain an email and password"})
		return
	}

	//Check that email and password are set
	if creds.Email == "" || creds.Password == "" {
		log.Error("wrong body: should contain an email and password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong body"})
		return
	}

	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to signup, please try later"})
		return
	}

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to signup, please try later"})
		return
	}

	//encrypt the password
	salt := 8
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), salt)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to signup, please try again"})
		return
	}

	err = db.StoreLoginAndPassword(creds.Email, string(hashedPassword))
	if err != nil {
		log.Error(err)

		//If email already exists
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to signup, please try again"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfull signup"})
}

//Signin handles user sign in
func Signin(c *gin.Context) {
	creds := Crendentials{}

	//Get creds from body
	if err := c.ShouldBind(&creds); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "body should contain an email and password"})
		return
	}

	//Check that email and password are set
	if creds.Email == "" || creds.Password == "" {
		log.Error("wrong body: should contain an email and password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong body"})
		return
	}

	//Initialize db client
	db, err := server.GetDefaultDBClient()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to sign in, please try later"})
		return
	}

	user, err := db.GetUserByEmail(creds.Email)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized,
			gin.H{"error": "wrong email or password"})
		return
	}

	//Verify email and password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized,
			gin.H{"error": "wrong email or password"})
		return
	}

	//Generate a JWT key
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Userid: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(server.GetDefaultConfig().JWTKey))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "failed to sign in, please try later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": fmt.Sprintf("Barear %s", tokenString)})
}
