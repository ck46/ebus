package user

import (
	"github.com/jinzhu/gorm"
)

func FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	if email != "" {
		var user User
		res := db.Find(&user, &User{Email: email})
		if res.RecordNotFound() {
			return nil, &EmailNotExistsError{}
		}
		return &user, nil
	} else {
		return nil, &EmailNotExistsError{}
	}
}

func FindByUserid(db *gorm.DB, uid string) (*User, error) {
	user, err := FindUserByEmail(db, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
