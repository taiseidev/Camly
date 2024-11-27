package handler

import (
	authService "camly-api/internal/auth/service"
	"camly-api/internal/user/model"
	"camly-api/internal/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *authService.AuthService
}

func NewAuthHandler(authService *authService.AuthService) *AuthHandler {
	if authService == nil {
		panic("authService cannot be nil")
	}
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	var req model.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
	}

	ctx := c.Request().Context()

	// Validate input
	if err := validateUserInput(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	// Serviceを呼び出し
	tokens, err := h.authService.SignUp(ctx, req)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "User creation failed",
			"details": "An error occurred while processing your request",
			"text":    err.Error(),
		})
	}

	// レスポンスとしてユーザーを返す
	return c.JSON(http.StatusOK, tokens)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req model.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
	}

	ctx := c.Request().Context()

	// Validate input
	if err := validateUserInput(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Validation failed",
			"details": err.Error(),
		})
	}

	// Serviceを呼び出し
	tokens, err := h.authService.Login(ctx, req.Email, req.Password)
	if err != nil {
		log.Printf("login failed: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Login failed",
			"details": "An error occurred while processing your request",
		})
	}
	// レスポンスとしてユーザーを返す
	return c.JSON(http.StatusOK, tokens)
}

type LogoutRequest struct {
	AccessToken string `json:"access_token"`
}

func (h *AuthHandler) Logout(c echo.Context) error {
	// リクエストボディからaccessTokenを取得
	var req LogoutRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
	}

	// accessTokenがリクエストボディに含まれているかを確認
	if req.AccessToken == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Validation failed",
			"details": "accessToken is required",
		})
	}

	ctx := c.Request().Context()

	err := h.authService.Logout(ctx, req.AccessToken)
	if err != nil {
		log.Printf("logout failed: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "Logout failed",
			"details": "An error occurred while processing your request",
		})
	}

	// 成功レスポンスを返す
	return c.JSON(http.StatusOK, nil)
}

func validateUserInput(user *model.User) error {
	if user.Email == "" {
		return fmt.Errorf("email is required")
	}
	if !utils.IsValidEmail(user.Email) {
		return fmt.Errorf("invalid email format")
	}
	if user.Password == "" {
		return fmt.Errorf("password is required")
	}
	if len(user.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	// Additional password complexity checks can be added here
	return nil
}
