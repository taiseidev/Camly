package handler

import (
	"camly-api/internal/user/model"
	"camly-api/internal/user/service"
	"camly-api/internal/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	if userService == nil {
		panic("userService cannot be nil")
	}
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	// リクエストからデータをバインド
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
	user, err := h.userService.CreateUser(ctx, req.Name, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error":   "User creation failed",
			"details": "An error occurred while processing your request",
		})
	}

	// レスポンスとしてユーザーを返す
	return c.JSON(http.StatusOK, user)
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
