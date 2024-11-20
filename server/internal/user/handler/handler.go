package handler

import (
	"camly-api/internal/user/model"
	"camly-api/internal/user/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	// リクエストからデータをバインド
	var req model.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Serviceを呼び出し
	user, err := h.userService.CreateUser(req.Name, req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	// レスポンスとしてユーザーを返す
	return c.JSON(http.StatusOK, user)
}
