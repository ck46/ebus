package user

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Login(db *gorm.DB, req *Request) (*Response, error) {
	user, err := FindUserByEmail(db, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, &PasswordMismatchError{}
	}
	return &Response{User: user}, nil
}

func Signup(db *gorm.DB, req *Request) (*Response, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &User{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PhoneNumber:  req.Phone,
		PasswordHash: string(passwordHash),
	}

	id, err := Create(db, newUser)
	if err != nil {
		return nil, err
	}
	return &Response{Id: id}, err
}
