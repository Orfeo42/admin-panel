package utils

import (
	"context"

	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/labstack/echo/v4"
)

type contextKey string

func getStringFromContext(ctx context.Context, key contextKey) string {
	if page, ok := ctx.Value(key).(string); ok {
		return page
	}
	return ""
}

func setInContext(echoCtx echo.Context, key contextKey, value any) echo.Context {
	ctx := context.WithValue(echoCtx.Request().Context(), key, value)
	echoCtx.SetRequest(echoCtx.Request().WithContext(ctx))
	return echoCtx
}

var pageContextKey contextKey = "page"

func GetPage(ctx context.Context) string {
	//fmt.Println(ctx.Value(pageContextKey))
	if page, ok := ctx.Value(pageContextKey).(pages.Page); ok {
		return string(page)
	}
	return ""
}

func SetPage(echoCtx echo.Context, value pages.Page) echo.Context {
	return setInContext(echoCtx, pageContextKey, value)
}

var titleContextKey contextKey = "title"

func GetTitle(ctx context.Context) string {
	return getStringFromContext(ctx, titleContextKey)
}

func SetTitle(echoCtx echo.Context, value string) echo.Context {
	return setInContext(echoCtx, titleContextKey, value)
}
