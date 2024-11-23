package authHandler

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
		panic("userService cannot be nil")
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
		log.Fatalf("failed to migrate database: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "User creation failed",
			"details": "An error occurred while processing your request",
		})
	}

	// レスポンスとしてユーザーを返す
	return c.JSON(http.StatusOK, tokens)
}

func validateUserInput(user *model.User) error {
	if user.Name == "" {
		return fmt.Errorf("name is required")
	}
	if user.Email == "" {
		return fmt.Errorf("email is required")
	}
	if !utils.IsValidEmail(user.Email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}
