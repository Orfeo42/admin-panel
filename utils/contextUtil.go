package utils

import (
	"context"

	"github.com/labstack/echo/v4"
)

type contextKey string

var pageContextKey contextKey = "page"

func GetPage(ctx context.Context) string {
	if page, ok := ctx.Value(pageContextKey).(string); ok {
		return page
	}
	return ""
}

func SetPage(echoCtx echo.Context, page string) echo.Context {
	ctx := context.WithValue(echoCtx.Request().Context(), pageContextKey, page)
	echoCtx.SetRequest(echoCtx.Request().WithContext(ctx))
	return echoCtx
}
