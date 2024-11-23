package routes

import (
	authHandler "camly-api/internal/auth/handler"

	"github.com/labstack/echo/v4"
)

// 全てのルートを定義
func RegisterRoutes(e *echo.Echo, authHandler *authHandler.AuthHandler) {
	v1 := e.Group("/api/v1")
	// 認証関連のルートを定義
	authRoutes := v1.Group("/auth")
	authRoutes.POST("/register", authHandler.SignUp)

}
