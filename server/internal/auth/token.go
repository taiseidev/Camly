package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getJwtSecret() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	return []byte(jwtSecret)
}

// アクセストークンを生成
func GenerateAccessToken(userID uint) (string, error) {
	if userID == 0 {
		return "", errors.New("invalid user ID")
	}
	jwtSecret := getJwtSecret()
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpiration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// リフレッシュトークンを生成
func GenerateRefreshToken(userID uint) (string, error) {
	if userID == 0 {
		return "", errors.New("invalid user ID")
	}
	jwtSecret := getJwtSecret()
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenExpiration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// トークンを検証
func ValidateToken(tokenString string) (*Claims, error) {
	if tokenString == "" {
		return nil, errors.New("empty token")
	}
	jwtSecret := getJwtSecret()
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("could not parse claims")
	}

	// Verify token hasn't expired
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
