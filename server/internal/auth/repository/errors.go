package repository

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrDatabase     = errors.New("database error")
)
