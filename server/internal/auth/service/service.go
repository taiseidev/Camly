package service

import (
	"camly-api/internal/auth"
	authModel "camly-api/internal/auth/model"
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

	tokens, err := util.GenerateTokens(newUser.ID)
	if err != nil {
		return util.TokenResponse{}, err
	}

	refreshToken := authModel.RefreshToken{
		UserID:    newUser.ID,
		TokenHash: tokens.RefreshToken,
		ExpiresAt: tokens.RefreshTokenExpiration,
	}

	if err := s.authRepo.SaveOrUpdateRefreshToken(ctx, &refreshToken); err != nil {
		return util.TokenResponse{}, err
	}

	return tokens, nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (util.TokenResponse, error) {
	user, err := s.authRepo.GetUserByEmail(email)
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

	tokens, err := util.GenerateTokens(user.ID)
	if err != nil {
		return util.TokenResponse{}, err
	}

	refreshToken := authModel.RefreshToken{
		UserID:    user.ID,
		TokenHash: tokens.RefreshToken,
		ExpiresAt: tokens.RefreshTokenExpiration,
	}

	if err := s.authRepo.SaveOrUpdateRefreshToken(ctx, &refreshToken); err != nil {
		return util.TokenResponse{}, err
	}

	return tokens, nil
}

func (s *AuthService) Logout(ctx context.Context, accessToken string) error {
	// アクセストークンからuserIdを取得
	userID, err := auth.GetUserIDFromToken(accessToken)
	if err != nil {
		// トークンが無効な場合
		if err.Error() == "no token found" {
			return err
		}
		// トークンの解析に失敗した場合
		if err.Error() == "invalid token" {
			return err
		}
		// その他のエラー
		return err
	}

	if err := s.authRepo.DeleteRefreshToken(ctx, userID); err != nil {
		return err
	}
	
	return nil
}
