package repository

import (
	authModel "camly-api/internal/auth/model"
	userModel "camly-api/internal/user/model"
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

// NOTE(onishi): multiple interfaces in the future
type IAuthRepository interface {
	SaveOrUpdateRefreshToken(ctx context.Context, tx *gorm.DB, model *authModel.RefreshToken) error
	DeleteRefreshToken(ctx context.Context, userID uint) error
	SaveUser(ctx context.Context, tx *gorm.DB, user *userModel.User) error
	GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (*userModel.User, error)
	BeginTransaction(ctx context.Context) *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	if db == nil {
		panic("db cannot be nil")
	}
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) SaveOrUpdateRefreshToken(ctx context.Context, tx *gorm.DB, model *authModel.RefreshToken) error {
	var refreshToken authModel.RefreshToken

	result := tx.WithContext(ctx).Where("user_id = ?", model.UserID).First(&refreshToken)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	if result.RowsAffected > 0 {
		// Update existing refresh token
		refreshToken.TokenHash = model.TokenHash
		refreshToken.ExpiresAt = model.ExpiresAt
		refreshToken.UpdatedAt = time.Now()
		// 更新処理
		if err := tx.WithContext(ctx).Save(&refreshToken).Error; err != nil {
			return err
		}
	} else {
		// Create new refresh token
		if err := tx.WithContext(ctx).Create(model).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *AuthRepository) DeleteRefreshToken(ctx context.Context, userID uint) error {
	// userID に基づいてリフレッシュトークンを削除
	result := r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&authModel.RefreshToken{})

	// エラーチェック
	if result.Error != nil {
		return fmt.Errorf("failed to delete refresh token: %w", result.Error)
	}

	return nil
}

func (r *AuthRepository) SaveUser(ctx context.Context, tx *gorm.DB, user *userModel.User) error {
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}

	// データの妥当性検証
	if err := user.Validate(); err != nil {
		return fmt.Errorf("invalid user data: %w", err)
	}

	// ユーザー情報の挿入
	if err := tx.WithContext(ctx).Create(user).Error; err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}

	return nil
}

func (r *AuthRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (*userModel.User, error) {
	var user userModel.User
	result := tx.WithContext(ctx).Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	} else if result.Error != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabase, result.Error)
	}

	return &user, nil
}

// トランザクション開始
func (r *AuthRepository) BeginTransaction(ctx context.Context) *gorm.DB {
	return r.db.Begin()
}
