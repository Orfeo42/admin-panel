package utils

import (
	"context"

	"github.com/labstack/echo/v4"
)

type contextKey string

func getFromContext(ctx context.Context, key contextKey) string {
	if page, ok := ctx.Value(pageContextKey).(string); ok {
		return page
	}
	return ""
}

func setInContext(echoCtx echo.Context, key contextKey, value string) echo.Context {
	ctx := context.WithValue(echoCtx.Request().Context(), key, value)
	echoCtx.SetRequest(echoCtx.Request().WithContext(ctx))
	return echoCtx
}

var pageContextKey contextKey = "page"

func GetPage(ctx context.Context) string {
	return getFromContext(ctx, pageContextKey)
}

func SetPage(echoCtx echo.Context, value string) echo.Context {
	return setInContext(echoCtx, pageContextKey, value)
}

var titleContextKey contextKey = "title"

func GetTitle(ctx context.Context) string {
	return getFromContext(ctx, titleContextKey)
}

func SetTitle(echoCtx echo.Context, value string) echo.Context {
	return setInContext(echoCtx, pageContextKey, value)
}
