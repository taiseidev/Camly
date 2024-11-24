package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"size:255" json:"name" validate:"min=2,max=100"`
	Email     string         `gorm:"uniqueIndex;size:100" json:"email" validate:"required,email"`
	Password  string         `gorm:"column:password_hash;size:255" json:"password" validate:"required,min=8"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
