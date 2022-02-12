package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetCities returns a list of cities
func GetCities(c *gin.Context) {
	c.JSON(http.StatusOK, cities)
}

//Cities of Cameroon
//for now we get them from that list but
//later they could be fetch from the DB or a different api
var cities = []string{
	"Douala",
	"Yaoundé",
	"Garoua",
	"Bamenda",
	"Maroua",
	"Nkongsamba",
	"Bafoussam",
	"Ngaoundéré",
	"Bertoua",
	"Loum",
	"Kumba",
	"Edéa",
	"Kumbo",
	"Foumban",
	"Mbouda",
	"Dschang",
	"Limbé",
	"Ebolowa",
	"Kousséri",
	"Guider",
	"Meiganga",
	"Yagoua",
	"Mbalmayo",
	"Bafang",
	"Tiko",
	"Bafia",
	"Wum",
	"Kribi",
	"Buea",
	"Sangmélima",
	"Foumbot",
	"Bangangté",
	"Batouri",
	"Banyo",
	"Nkambé",
	"Bali",
	"Mbanga",
	"Mokolo",
	"Melong",
	"Manjo",
	"Garoua-Boulaï",
	"Mora",
	"Kaélé",
	"Tibati",
	"Ndop",
	"Akonolinga",
	"Eséka",
	"Mamfé",
	"Obala",
	"Muyuka",
	"Nanga-Eboko",
	"Abong-Mbang",
	"Fundong",
	"Nkoteng",
	"Fontem",
	"Mbandjock",
	"Touboro",
	"Ngaoundal",
	"Yokadouma",
	"Pitoa",
	"Tombel",
	"Kékem",
	"Magba",
	"Bélabo",
	"Tonga",
	"Maga",
	"Koutaba",
	"Blangoua",
	"Guidiguis",
	"Bogo",
	"Batibo",
	"Yabassi",
	"Figuil",
	"Makénéné",
	"Gazawa",
	"Tcholliré",
}
