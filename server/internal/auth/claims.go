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

// Duration for access token validity
const AccessTokenExpiration = time.Minute * 7

// Duration for refresh token validity
const RefreshTokenExpiration = time.Hour * 24 * 7
