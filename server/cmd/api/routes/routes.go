package routes

import (
	authHandler "camly-api/internal/auth/handler"

	"github.com/labstack/echo/v4"
)

// 全てのルートを定義
// TODO(onishi): 下記レビューの通り、ログイン試行に対するレートリミットを設定する
// https://github.com/taiseidev/Camly/pull/15#discussion_r1855351379
func RegisterRoutes(e *echo.Echo, authHandler *authHandler.AuthHandler) {
	// TODO(onishi): 認証ミドルウェアをauthグループまたは特にlogoutルートに追加する。
	v1 := e.Group("/api/v1")
	// 認証関連のルートを定義
	authRoutes := v1.Group("/auth")
	authRoutes.POST("/signup", authHandler.SignUp)
	authRoutes.POST("/login", authHandler.Login)
	authRoutes.POST("/logout", authHandler.Logout)

}
