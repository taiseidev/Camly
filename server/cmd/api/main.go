package main

import (
	"camly-api/cmd/api/routes"
	"camly-api/internal/user/handler"
	"camly-api/internal/user/repository"
	"camly-api/internal/user/service"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	db          *gorm.DB
	userService *service.UserService
	cleanup     func()
}

func NewApp(db *gorm.DB) *App {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	return &App{
		db:          db,
		userService: userService,
		cleanup: func() {
			sqlDB, err := db.DB()
			if err == nil {
				sqlDB.Close()
			}
		},
	}
}

// TODO: 下記PRの内容を確認して修正。
// https://github.com/taiseidev/Camly/pull/13#discussion_r1850223272
// https://github.com/taiseidev/Camly/pull/13#discussion_r1850223270

func main() {
	e := echo.New()

	dsn := "user:password@tcp(camly-db:3306)/camly-db?charset=utf8mb4&parseTime=True&loc=Local"

	// データベースに接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// App インスタンスを作成し、クリーンアップ関数を defer で登録
	app := NewApp(db)
	defer app.cleanup()

	// ハンドラーの初期化
	userHandler := handler.NewUserHandler(app.userService)

	// ルートを登録
	// TODO: 必要なミドルウェアの追加やグレースフルシャットダウンの実装を検討する
	routes.RegisterRoutes(e, userHandler)

	// 終了シグナルを受け取るチャネルを作成
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// サーバーを非同期で起動
	go func() {
		if err := e.Start(":8080"); err != nil && err != echo.ErrInternalServerError {
			e.Logger.Fatal("サーバーの起動に失敗しました")
		}
	}()

	// シグナルを待機
	<-quit

	// コンテキストを作成してサーバーをシャットダウン
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
