package repository

import (
	authModel "camly-api/internal/auth/model"
	userModel "camly-api/internal/user/model"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	return tx.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"token_hash", "expires_at", "updated_at"}),
		}).
		Create(model).Error
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
	return r.db.WithContext(ctx).Begin()
}
