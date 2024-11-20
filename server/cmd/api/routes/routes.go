package routes

import (
	"camly-api/internal/user/handler"

	"github.com/labstack/echo/v4"
)

// 全てのルートを定義
func RegisterRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	v1 := e.Group("/api/v1")
	// ユーザー関連のルートを定義
	userRoutes := v1.Group("/users")
	// TODO: Add necessary middleware
	// userRoutes.Use(middleware.JWT([]byte("secret")))
	userRoutes.POST("", userHandler.CreateUser)

}
