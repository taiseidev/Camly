package service

import (
	"camly-api/internal/user/model"
	"camly-api/internal/user/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(name string, email string) (*model.User, error) {

	// 新しいユーザーを構築
	user := &model.User{
		Name:  name,
		Email: email,
	}

	// データベースに保存
	if err := s.userRepo.SaveUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
