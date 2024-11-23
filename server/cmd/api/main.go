package main

import (
	"camly-api/cmd/api/routes"
	authHandler "camly-api/internal/auth/handler"
	authRepository "camly-api/internal/auth/repository"
	authService "camly-api/internal/auth/service"
	"camly-api/internal/user/repository"
	"camly-api/internal/user/service"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	db          *gorm.DB
	userService *service.UserService
	authService *authService.AuthService
	cleanup     func()
}

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load("docker/db/.env")
		if err != nil {
			log.Fatalln(err)
		}
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	// データベースに接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	fmt.Println("🚀 DB connected!!")
	return db
}

func NewApp(db *gorm.DB) *App {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authRepo := authRepository.NewAuthRepository(db)
	authService := authService.NewAuthService(authRepo)

	return &App{
		db:          db,
		userService: userService,
		authService: authService,
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

	// DB接続
	db := NewDB()

	// App インスタンスを作成し、クリーンアップ関数を defer で登録
	app := NewApp(db)

	// appインスタンスのクリーンナップ
	defer app.cleanup()

	// ハンドラーの初期化
	authHandler := authHandler.NewAuthHandler(app.authService)

	// ルートを登録
	// TODO: 必要なミドルウェアの追加やグレースフルシャットダウンの実装を検討する
	routes.RegisterRoutes(e, authHandler)

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
