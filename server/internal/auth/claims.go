// auth/claims.go
package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// アクセストークンの有効期間
const AccessTokenExpiration = time.Minute * 30

// リフレッシュトークンの有効期間
const RefreshTokenExpiration = time.Hour * 24 * 30
