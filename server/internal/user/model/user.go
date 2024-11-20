package model

import "github.com/go-playground/validator/v10"

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255" json:"name" validate:"required,min=2,max=100"`
	Email string `gorm:"uniqueIndex;size:255" json:"email" validate:"required,email"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
