package routes

import (
	"camly-api/internal/user/handler"

	"github.com/labstack/echo/v4"
)

// 全てのルートを定義
func RegisterRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	// ユーザー関連のルートを定義
	userRoutes := e.Group("/users")
	userRoutes.POST("", userHandler.CreateUser)

}
