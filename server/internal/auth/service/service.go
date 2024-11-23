package service

import (
	"camly-api/internal/auth"
	authRepository "camly-api/internal/auth/repository"
	"camly-api/internal/user/model"
	"context"
	"errors"
	"time"

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

type TokenResponse struct {
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	AccessTokenExpiration  string `json:"access_token_expiration"`
	RefreshTokenExpiration string `json:"refresh_token_expiration"`
}

func (s *AuthService) SignUp(ctx context.Context, user model.User) (TokenResponse, error) {

	// パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		return TokenResponse{}, err
	}

	newUser := model.User{Name: user.Name, Email: user.Email, Password: string(hash)}

	if err := s.authRepo.SaveUser(ctx, &newUser); err != nil {
		return TokenResponse{}, err
	}

	// アクセストークン生成
	accessToken, err := auth.GenerateAccessToken(newUser.ID)
	if err != nil {
		return TokenResponse{}, err
	}

	// リフレッシュトークン生成
	refreshToken, err := auth.GenerateRefreshToken(newUser.ID)
	if err != nil {
		return TokenResponse{}, err
	}

	// 有効期限を取得
	accessTokenExpiration := time.Now().Add(auth.AccessTokenExpiration).Format(time.RFC3339)
	refreshTokenExpiration := time.Now().Add(auth.RefreshTokenExpiration).Format(time.RFC3339)

	tokenRes := TokenResponse{
		AccessToken:            accessToken,
		RefreshToken:           refreshToken,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}

	return tokenRes, nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (TokenResponse, error) {
	user, err := s.authRepo.GetUserById(ctx, email)
	if err != nil {
		return TokenResponse{}, errors.New("user not found")
	}
	// パスワードの検証
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return TokenResponse{}, errors.New("invalid credentials")
	}

	// トークン生成
	accessToken, err := auth.GenerateAccessToken(user.ID)
	if err != nil {
		return TokenResponse{}, err
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		return TokenResponse{}, err
	}

	// 有効期限を計算
	accessTokenExpiration := time.Now().Add(auth.AccessTokenExpiration).Format(time.RFC3339)
	refreshTokenExpiration := time.Now().Add(auth.RefreshTokenExpiration).Format(time.RFC3339)

	return TokenResponse{
		AccessToken:            accessToken,
		RefreshToken:           refreshToken,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}, nil
}
