package util

import (
	"camly-api/internal/auth"
	"time"
)

type TokenResponse struct {
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	AccessTokenExpiration  string `json:"access_token_expiration"`
	RefreshTokenExpiration string `json:"refresh_token_expiration"`
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

	accessTokenExpiration := time.Now().Add(auth.AccessTokenExpiration).Format(time.RFC3339)
	refreshTokenExpiration := time.Now().Add(auth.RefreshTokenExpiration).Format(time.RFC3339)

	return TokenResponse{
		AccessToken:            accessToken,
		RefreshToken:           refreshToken,
		AccessTokenExpiration:  accessTokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}, nil
}
