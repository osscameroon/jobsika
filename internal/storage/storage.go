package storage

import (
	"fmt"

	"github.com/elhmn/camerdevs/pkg/models/v1beta"
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

//StoreLoginAndPassword store user's login and password
func (db DB) StoreLoginAndPassword(email string, password string) error {
	ret := db.c.Create(&v1beta.User{Email: email, Password: password})
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

//GetUserByEmail get user by email
func (db DB) GetUserByEmail(email string) (v1beta.User, error) {
	user := v1beta.User{}
	ret := db.c.Where("email = ?", email).First(&user)
	if ret.Error != nil {
		return user, ret.Error
	}

	return user, nil
}
