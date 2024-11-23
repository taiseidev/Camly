package repository

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	if db == nil {
		panic("db cannot be nil")
	}
	return &UserRepository{
		db: db,
	}
}
