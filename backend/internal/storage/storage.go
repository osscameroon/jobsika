package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//IDB defines the DB interface
type IDB interface{}

//DB defines the DB storage data structure
type DB struct {
	//The database client
	c *gorm.DB
}

//DBOptions is a set of options used to initialise the database
type DBOptions struct {
	Password string
	DBName   string
	User     string
	Host     string
	Port     string
}

//NewDB returns a new DB client
func NewDB(opt DBOptions) (*DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		opt.Host, opt.User, opt.Password, opt.DBName, opt.Port)
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{
		c: client,
	}, nil
}
