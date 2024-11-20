package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:255" json:"name" validate:"required,min=2,max=100"`
	Email     string         `gorm:"uniqueIndex;size:100" json:"email" validate:"required,email"`
	Password  string         `gorm:"size:255" json:"-" validate:"required,min=8"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
