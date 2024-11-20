package repository

import (
	"camly-api/internal/user/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	SaveUser(ctx context.Context, user *model.User) error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	if db == nil {
		panic("db cannot be nil")
	}
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) SaveUser(ctx context.Context, user *model.User) error {
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}

	if err := user.Validate(); err != nil {
		return fmt.Errorf("invalid user data: %w", err)
	}

	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to save user: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
