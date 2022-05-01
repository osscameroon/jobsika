package server

import (
	"github.com/elhmn/jobsika/internal/config"
	"github.com/elhmn/jobsika/internal/storage"
)

//Server defines the api server struct
type Server struct {
	DB   storage.DB
	Conf config.Config
}

var defaultServer *Server

//GetDefaultServer returns a default server configuration
func GetDefaultServer() (*Server, error) {
	if defaultServer == nil {
		conf := config.GetDefaultConfig()
		db, err := storage.NewDB(conf.DBOpts)
		if err != nil {
			return nil, err
		}

		defaultServer = &Server{
			DB:   *db,
			Conf: conf,
		}
	}

	return defaultServer, nil
}

//GetDefaultDBClient returns the database default client
func GetDefaultDBClient() (storage.DB, error) {
	s, err := GetDefaultServer()
	if err != nil {
		return storage.DB{}, err
	}

	return s.DB, nil
}

//GetDefaultConfig returns the server default config
func GetDefaultConfig() config.Config {
	s, err := GetDefaultServer()
	if err != nil {
		return config.Config{}
	}

	return s.Conf
}
