package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/osscameroon/jobsika/internal/payment"
	"github.com/osscameroon/jobsika/internal/storage"
)

var (
	postgresEnvFile       = ".postgres-env"
	openCollectiveEnvFile = ".opencollective-env"
	//StageEnv stage environment
	StageEnv = "stage"
	//ProdEnv production environment
	ProdEnv = "production"
)

var defaultConfig *Config

// Config describes the server configuration
type Config struct {
	DBOpts storage.DBOptions
	//Environment can be set to stage or production
	Environment string
	JWTKey      string

	//OCOpts contains the open collective options
	OCOpts payment.OpenCollectiveOptions
}

// GetDefaultConfig returns a config with default values  and env variables
func GetDefaultConfig() Config {
	if defaultConfig == nil {
		//load postgres env variables
		godotenv.Load(postgresEnvFile)
		godotenv.Load(openCollectiveEnvFile)

		defaultConfig = &Config{
			DBOpts: storage.DBOptions{
				DBName:   os.Getenv("POSTGRES_DB"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				User:     os.Getenv("POSTGRES_USER"),
				Host:     os.Getenv("POSTGRES_HOST"),
				Port:     os.Getenv("POSTGRES_PORT"),
			},
			Environment: os.Getenv("ENVIRONMENT"),
			JWTKey:      os.Getenv("JWT_KEY"),
			OCOpts: payment.OpenCollectiveOptions{
				URL:     os.Getenv("OPEN_COLLECTIVE_API_URL"),
				KEY:     os.Getenv("OPEN_COLLECTIVE_API_KEY"),
				OrgSlug: os.Getenv("OPEN_COLLECTIVE_ORG_SLUG"),
			},
		}
	}

	return *defaultConfig
}
