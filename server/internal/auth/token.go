package auth

import (
	"camly-api/internal/config"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getJwtSecret() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return []byte(config.GetConfig().JWTSecret)
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

func GetUserIDFromToken(accessToken string) (uint, error) {
	// JWT トークンを検証して userID を抽出
	userID, err := parseToken(accessToken)
	if err != nil {
		return 0, fmt.Errorf("invalid token: %w", err)
	}

	return userID, nil
}

// TODO(onishi):Several improvements needed in token parsing logic.
// https://github.com/taiseidev/Camly/pull/16#discussion_r1859005755
func parseToken(tokenStr string) (uint, error) {
	// JWT トークンを解析する
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// トークンが有効かどうかの検証を行う
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// config.JWTSecret は JWT の秘密鍵
		return []byte(config.GetConfig().JWTSecret), nil
	})
	if err != nil {
		return 0, err
	}

	// トークンが有効であれば、claims から userID を取得
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return 0, fmt.Errorf("invalid token: user_id not found")
		}
		return uint(userID), nil
	}
	return 0, fmt.Errorf("invalid token")
}
