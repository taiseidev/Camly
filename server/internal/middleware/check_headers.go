package middleware

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckHeadersMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Authorization ヘッダーの確認
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Authorization header is missing")
		}

		// Bearer プレフィックスの確認
		if len(authHeader) < 7 || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, "Authorization header must be in the format 'Bearer <token>'")
		}

		// Content-Type ヘッダーの確認
		contentType := c.Request().Header.Get("Content-Type")
		if contentType != "application/json" {
			return c.JSON(http.StatusBadRequest, "Content-Type must be application/json")
		}

		// AppVersion ヘッダーの確認
		appVersion := c.Request().Header.Get("AppVersion")
		if appVersion == "" {
			return c.JSON(http.StatusBadRequest, "AppVersion header is missing")
		}

		// AppVersion バージョン形式の確認 (例: "1.0.0")
		versionRegex := `^(\d+\.\d+\.\d+)$`
		matched, err := regexp.MatchString(versionRegex, appVersion)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error in version validation: %v", err))
		}

		if !matched {
			return c.JSON(http.StatusBadRequest, "AppVersion header must follow the format 'X.Y.Z' (e.g., '1.0.0')")
		}

		return next(c)
	}
}
