package server

import (
	"github.com/osscameroon/jobsika/internal/config"
	"github.com/osscameroon/jobsika/internal/payment"
	"github.com/osscameroon/jobsika/internal/storage"
)

// Server defines the api server struct
type Server struct {
	DB            storage.DB
	Conf          config.Config
	PaymentClient payment.Client
}

var defaultServer *Server

// GetDefaultServer returns a default server configuration
func GetDefaultServer() (*Server, error) {
	if defaultServer == nil {
		conf := config.GetDefaultConfig()
		db, err := storage.NewDB(conf.DBOpts)
		if err != nil {
			return nil, err
		}

		paymentClient, err := payment.NewClient(conf.OCOpts)
		if err != nil {
			return nil, err
		}

		defaultServer = &Server{
			DB:            *db,
			Conf:          conf,
			PaymentClient: *paymentClient,
		}
	}

	return defaultServer, nil
}

// GetDefaultDBClient returns the database default client
func GetDefaultDBClient() (storage.DB, error) {
	s, err := GetDefaultServer()
	if err != nil {
		return storage.DB{}, err
	}

	return s.DB, nil
}

// GetDefaultConfig returns the server default config
func GetDefaultConfig() config.Config {
	s, err := GetDefaultServer()
	if err != nil {
		return config.Config{}
	}

	return s.Conf
}

// GetDefaultPaymentClient returns the default payment client
func GetDefaultPaymentClient() (payment.Client, error) {
	s, err := GetDefaultServer()
	if err != nil {
		return payment.Client{}, err
	}

	return s.PaymentClient, nil
}
