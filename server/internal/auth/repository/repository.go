package repository

import (
	"camly-api/internal/user/model"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

// NOTE(onishi): multiple interfaces in the future
type IAuthRepository interface {
	SaveUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, email string) (*model.User, error)
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	if db == nil {
		panic("db cannot be nil")
	}
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) SaveUser(ctx context.Context, user *model.User) error {
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}

	// データの妥当性検証
	if err := user.Validate(); err != nil {
		return fmt.Errorf("invalid user data: %w", err)
	}

	// ユーザー情報の挿入
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

func (r *AuthRepository) GetUserById(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
