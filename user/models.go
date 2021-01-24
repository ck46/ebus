package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string `gorm:"type:varchar(100);not null;unique_index"`
	PasswordHash string
	PhoneNumber  string
}

type Request struct {
	Userid    string
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Phone     string
}

type Response struct {
	Id   uint
	User *User
}
