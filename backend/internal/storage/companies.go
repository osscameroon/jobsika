package storage

import (
	"os"

	"github.com/elhmn/camerdevs/pkg/models/v1beta"
)

//GetCompanies get companies
func (db DB) GetCompanies() ([]v1beta.Company, error) {
	companies := []v1beta.Company{}
	ret := db.c.Table("companies").Order("name").Find(&companies)
	if ret.Error != nil {
		return companies, ret.Error
	}

	//Check if the first cameroonian company already exists
	cameroonianCompany := v1beta.Company{}
	ret = db.c.Table("companies").Where("name = ?", CameroonianCompanies[0]).Find(&cameroonianCompany)
	if ret.Error != nil {
		return companies, ret.Error
	}

	//If we did not find the first cameroonian company, we create every cameroonian companies
	if cameroonianCompany.Name == "" {
		tmpCameroonianCompanies := CameroonianCompanies
		//If we are running tests we use a short list of companies
		if os.Getenv("ENVIRONMENT") == "test" {
			tmpCameroonianCompanies = []string{"Tester"}
		}

		for _, c := range tmpCameroonianCompanies {
			company := v1beta.Company{
				Name: c,
			}
			ret := db.c.Create(&company)
			if ret.Error != nil {
				return companies, ret.Error
			}
		}

		//Get the list of companies again
		ret := db.c.Table("companies").Order("name").Find(&companies)
		if ret.Error != nil {
			return companies, ret.Error
		}
	}

	return companies, nil
}

//GetCompanyByID get company by `id`
func (db DB) GetCompanyByID(id int64) (v1beta.Company, error) {
	company := v1beta.Company{}
	ret := db.c.First(&company, "id = ?", id)
	if ret.Error != nil {
		return company, ret.Error
	}

	return company, nil
}

//CameroonianCompanies is collection of default cameroonian companies
var CameroonianCompanies = []string{
	"CCAA",
	"ADC",
	"ADER",
	"Alpicam Industries",
	"Alucam",
	"AMPCI",
	"Aberec Limited Machines électriques et matériels électroniques",
	"Afriland First Bank",
	"AFRILANE",
	"AFROLOGIX",
	"AFOUP",
	"ANAFOR",
	"ASMAP",
	"AT graphiline",
	"AFROSPHINX",
	"APN",
	"Brasseries du Cameroun",
	"Business Facilities Corporation S.A",
	"BT ROUTES S.A",
	"BICEC",
	"B.E.S Best Engineering System Sarl",
	"Boh Plantations Limited",
	"BUNS Sarl",
	"BYOOS sarl",
	"Bendo sarl",
	"BVS",
	"Camair",
	"Camair-Co",
	"Cambuild-BTP",
	"Camdev emploi et formation",
	"Camerinside Web Agency",
	"Camerounaise des Eaux",
	"Cameroon Water Utilities Corporation",
	"Cameroon Development Corporation",
	"Camercrack",
	"Camlait",
	"Campost",
	"Camtel",
	"Camrail",
	"Cameroon Chemical Company",
	"Cameroon Development Corporation",
	"Cameroon Tea Estates",
	"CamerPages",
	"CAWAD",
	"CCAA",
	"CCPC Finance",
	"Chanas assurances",
	"Chantier naval industriel du Cameroun",
	"CICAM",
	"Clarans Afrique",
	"Crédit lyonnais du Cameroun",
	"Cimenterie du Cameroun",
	"Crédit Foncier du Cameroun",
	"CNIC",
	"Commercial Bank of Cameroon",
	"Compagnie d'Opérations Pétrolières Schlumberger",
	"Conseil National des Chargeurs du Cameroun",
	"Coordonnerie du 3e Millénaire",
	"CRTV",
	"CONGELCAM",
	"Corlay Cameroun",
	"CCA",
	"CDS",
	"CA Integrated Systems & consulting",
	"CNPS",
	"Douala Stock Exchange",
	"DIPMAN",
	"DANAY EXPRESS",
	"DIGIT CAMEROON",
	"EBA GROUP",
	"E-Business Cameroon SARL",
	"Electricity Development Corporation",
	"Express Union",
	"EIF-CONSULTING AND SERVICES SARL",
	"Enéo Cameroon",
	"EVA Corporation SARL",
	"EWONGA SARL",
	"ETS KOT SUD DOUALA",
	"FRINT GROUP",
	"Fonds national de l'emploi",
	"Foongon Corporation",
	"FEICOM",
	"FB-Building",
	"GSEC",
	"Gaz du Cameroun",
	"Geovic Cameroon",
	"GETRACOM-INTER",
	"GEQUIPS",
	"GFBC",
	"Guinness Cameroun sa",
	"GNO Solutions",
	"GESSIIA SARL",
	"Haltech",
	"Hévécam",
	"Hope Management & Consulting",
	"Hope Music Group",
	"Hope Music Publishing",
	"HSC Cameroun",
	"IFATI",
	"ITGStore",
	"ICCSOFT",
	"IMS",
	"INSTRUMELEC",
	"INSURMIND TECH",
	"IELT CAMEROUN",
	"IFRIQUIYA Conseil Sarl",
	"Joy Bonapriso",
	"KEN Technology Sarl",
	"KPMG",
	"KIMPA SAS",
	"LANAVET",
	"LIGHTECH",
	"LAPIN237",
	"LOCAGRUES CAMEROUN",
	"LUMEN SARL",
	"MIDEPECAM",
	"MTN Cameroun",
	"MIRE WORLD",
	"Megasoft",
	"Maguysama Technologies Solaires",
	"MY WAY Sarl",
	"Mezadi SARL",
	"MEN TRAVEL",
	"NACYDATE",
	"Ndawara Tea Estates",
	"Nestlé Cameroun",
	"Nexttel cameroun",
	"NTARRA",
	"Nexa Industries",
	"Orange Cameroun",
	"Ok Plast Cam",
	"Opticam",
	"PERFITCOM",
	"Plantations Pamol du Cameroun",
	"PHP",
	"Philjohn Technologies",
	"PMUC",
	"Port autonome de Douala",
	"Port autonome de Kribi",
	"Port autonome de Limbé",
	"Plasticam",
	"Pro Services Emploi",
	"Projects Experts Consulting",
	"Prosygma Cameroun",
	"PwC Cameroun",
	"Quezil Language Services traduction",
	"ETS RIVER",
	"SABC",
	"SACONETS",
	"SARMETAL",
	"SBS Smart Business Solutions",
	"SCDP",
	"Semry",
	"SGBC",
	"SNH",
	"Sodecoton",
	"Socatral",
	"Socapalm",
	"SOCATUR",
	"SONARA",
	"Sopecam",
	"SANCOMS",
	"SOCOSER",
	"SODEPA",
	"SOTRACOM",
	"SUDCAM",
	"Secours auto",
	"Societé Camerounaise d'équipement",
	"SimaLap",
	"SYSCOM SARL",
	"S&P",
	"Touristiques Express",
	"Third",
	"Tradex",
	"The House of services",
	"Uccao",
	"UPTIMA.CM Sarl",
	"WAT&CO",
}
