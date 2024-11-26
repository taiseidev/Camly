package service

import (
	"camly-api/internal/auth"
	authModel "camly-api/internal/auth/model"
	authRepository "camly-api/internal/auth/repository"
	"camly-api/internal/auth/util"
	"camly-api/internal/user/model"
	"context"
	"errors"
	"fmt"

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
	tx := s.authRepo.BeginTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-panic after rollback
		}
	}()
	// パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return util.TokenResponse{}, err
	}

	newUser := model.User{Name: user.Name, Email: user.Email, Password: string(hash)}
	if err := s.authRepo.SaveUser(ctx, tx, &newUser); err != nil {
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

	if err := s.authRepo.SaveOrUpdateRefreshToken(ctx, tx, &refreshToken); err != nil {
		return util.TokenResponse{}, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return util.TokenResponse{}, fmt.Errorf("failed to commit transaction")
	}

	return tokens, nil
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (util.TokenResponse, error) {
	tx := s.authRepo.BeginTransaction(ctx)
	defer tx.Rollback()

	user, err := s.authRepo.GetUserByEmail(ctx, tx, email)
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

	if err := s.authRepo.SaveOrUpdateRefreshToken(ctx, tx, &refreshToken); err != nil {
		return util.TokenResponse{}, err
	}

	tx.Commit()

	return tokens, nil
}

// TODO(onishi): Ensure secure storage of refresh tokens.
// https://github.com/taiseidev/Camly/pull/16#discussion_r1859005820
func (s *AuthService) Logout(ctx context.Context, accessToken string) error {
	// アクセストークンからuserIdを取得
	userID, err := auth.GetUserIDFromToken(accessToken)
	if err != nil {
		return errors.New("invalid or expired access token")
	}

	tx := s.authRepo.BeginTransaction(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := s.authRepo.DeleteRefreshToken(ctx, tx, userID); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction")
	}

	return nil
}
