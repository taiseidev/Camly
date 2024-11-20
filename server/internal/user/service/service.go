package service

import (
	"camly-api/internal/user/model"
	"camly-api/internal/user/repository"
	"camly-api/internal/utils"
	"context"
	"fmt"
)

type UserService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
	// Validate input parameters
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if !utils.IsValidEmail(email) {
		return nil, fmt.Errorf("invalid email format")
	}

	// 新しいユーザーを構築
	user := &model.User{
		Name:  name,
		Email: email,
	}

	// データベースに保存
	if err := s.userRepo.SaveUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}
