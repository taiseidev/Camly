package authService

import (
	"camly-api/internal/auth"
	authRepository "camly-api/internal/auth/repository"
	"camly-api/internal/user/model"
	"context"
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
