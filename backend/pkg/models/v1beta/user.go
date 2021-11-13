package v1beta

import "time"

//User defines the user structure
type User struct {
	ID          int64     `json:"id" gorm:"column:id;autoincrement"`
	Email       string    `json:"email" gorm:"column:email"`
	Password    string    `json:"password" gorm:"column:password"`
	FirstName   string    `json:"firstname" gorm:"column:firstname"`
	LastName    string    `json:"lastname" gorm:"column:lastname"`
	Gender      string    `json:"gender" gorm:"column:gender"`
	DateOfBirth time.Time `json:"dateofbirth" gorm:"column:dateofbirth"`
	CreatedAt   time.Time `json:"createdat" gorm:"column:createdat"`
	UpdatedAt   time.Time `json:"updatedat" gorm:"column:updatedat"`
}
