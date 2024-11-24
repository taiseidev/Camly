package util

import (
	"camly-api/internal/auth"
	"time"
)

// TokenResponse represents the authentication tokens and their expiration times
type TokenResponse struct {
	AccessToken            string    `json:"access_token"`
	RefreshToken           string    `json:"refresh_token"`
	AccessTokenExpiration  time.Time `json:"access_token_expiration"`
	RefreshTokenExpiration time.Time `json:"refresh_token_expiration"`
}

func GenerateTokens(userID uint) (TokenResponse, error) {
	accessToken, err := auth.GenerateAccessToken(userID)
	if err != nil {
		return TokenResponse{}, err
	}

	refreshToken, err := auth.GenerateRefreshToken(userID)
	if err != nil {
		return TokenResponse{}, err
	}

	accessTokenExpiration := time.Now().Add(auth.AccessTokenExpiration)
	refreshTokenExpiration := time.Now().Add(auth.RefreshTokenExpiration)

	return TokenResponse{
		AccessToken:            accessToken,
		RefreshToken:           refreshToken,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}, nil
}
