package service

import (
	authRepository "camly-api/internal/auth/repository"
	"camly-api/internal/auth/util"
	"camly-api/internal/user/model"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo authRepository.IAuthRepository
}

func NewAuthService(authRepo authRepository.IAuthRepository) *AuthService {
	return &AuthService{
		authRepo: authRepo,
	}
}

func (s *AuthService) SignUp(ctx context.Context, user model.User) (util.TokenResponse, error) {

	// パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return util.TokenResponse{}, err
	}

	newUser := model.User{Name: user.Name, Email: user.Email, Password: string(hash)}

	if err := s.authRepo.SaveUser(ctx, &newUser); err != nil {
		return util.TokenResponse{}, err
	}

	return util.GenerateTokens(newUser.ID)
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (util.TokenResponse, error) {
	user, err := s.authRepo.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, authRepository.ErrUserNotFound) {
			return util.TokenResponse{}, errors.New("user not found")
		} else {
			return util.TokenResponse{}, errors.New("internal server error")
		}
	}
	// パスワードの検証
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return util.TokenResponse{}, errors.New("invalid credentials")
	}

	return util.GenerateTokens(user.ID)
}
