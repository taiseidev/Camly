package repository

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// TODO(onishi):Add method later.
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
