package storage

import (
	"os"

	"github.com/osscameroon/jobsika/pkg/models/v1beta"
)

// GetCities get cities
func (db DB) GetCities() ([]string, error) {
	var citiesNames []string
	cities := []v1beta.City{}

	//Check if the last city already exists
	defaultCity := v1beta.City{}
	ret := db.c.Table("cities").Where("name = ?", DefaultCities[len(DefaultCities)-1]).Find(&defaultCity)
	if ret.Error != nil {
		return citiesNames, ret.Error
	}

	//If we did not find the last jobTitle, we create every cities
	if defaultCity.Name == "" {
		tmpCities := DefaultCities
		//If we are running tests we use a short list of cities
		if os.Getenv("TEST") == "true" {
			tmpCities = []string{"Tester"}
		}

		for _, cityName := range tmpCities {
			tmpCity := v1beta.City{}

			//If the city exists we don't create it
			ret := db.c.Table("cities").Where("name = ?", cityName).Find(&tmpCity)
			if tmpCity.Name != "" || ret.Error != nil {
				continue
			}

			city := v1beta.City{
				Name: cityName,
			}

			db.c.Table("cities").Create(&city)
		}
	}

	//Get the list of cities again
	ret = db.c.Table("cities").Order("name").Find(&cities)
	if ret.Error != nil {
		return citiesNames, ret.Error
	}

	for _, t := range cities {
		citiesNames = append(citiesNames, t.Name)
	}

	return citiesNames, nil
}

// DefaultCities is a collection of default cities of Cameroon
var DefaultCities = []string{
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
