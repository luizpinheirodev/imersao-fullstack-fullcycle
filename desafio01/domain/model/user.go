package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID   string `json:"id" valid:"required"`
	Name string `json:"name" valid:"required"`
	Email string `json:"email" valid:"required"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func newUser (name string, email string) (*User, error){
	user := User{
		Name:  name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
