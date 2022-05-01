package config

import (
	"os"

	"github.com/elhmn/jobsika/internal/storage"
	"github.com/joho/godotenv"
)

var (
	postgresEnvFile = ".postgres-env"
	//StageEnv stage environment
	StageEnv = "stage"
	//ProdEnv production environment
	ProdEnv = "production"
)

var defaultConfig *Config

//Config describes the server configuration
type Config struct {
	DBOpts storage.DBOptions
	//Environment can be set to stage or production
	Environment string
	JWTKey      string
}

//GetDefaultConfig returns a config with default values  and env variables
func GetDefaultConfig() Config {
	if defaultConfig == nil {
		//load postgres env variables
		godotenv.Load(postgresEnvFile)

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
		}
	}

	return *defaultConfig
}
