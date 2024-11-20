package main

import (
	"camly-api/cmd/api/routes"
	"camly-api/internal/user/handler"
	"camly-api/internal/user/repository"
	"camly-api/internal/user/service"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	dsn := "user:password@tcp(camly-db:3306)/camly-db?charset=utf8mb4&parseTime=True&loc=Local"

	// データベースに接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// リポジトリ、サービス、ハンドラーの初期化
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// ルートを登録
	routes.RegisterRoutes(e, userHandler)

	// サーバーを起動
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
