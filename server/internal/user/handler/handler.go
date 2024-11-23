package handler

import (
	"camly-api/internal/user/service"
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
