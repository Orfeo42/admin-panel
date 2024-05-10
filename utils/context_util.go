package utils

import (
	"context"

	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/labstack/echo/v4"
)

type contextKey string

const (
	CustomerVisible contextKey = "CustomerVisible"
	PageContextKey  contextKey = "page"
	TitleContextKey contextKey = "title"
)

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

func GetPage(ctx context.Context) string {
	if page, ok := ctx.Value(PageContextKey).(pages.Page); ok {
		return string(page)
	}
	return ""
}

func SetPage(echoCtx echo.Context, value pages.Page) echo.Context {
	return setInContext(echoCtx, PageContextKey, value)
}

func GetTitle(ctx context.Context) string {
	return getStringFromContext(ctx, TitleContextKey)
}

func SetTitle(echoCtx echo.Context, value string) echo.Context {
	return setInContext(echoCtx, TitleContextKey, value)
}

func GetCustomerVisible(ctx context.Context) bool {
	if page, ok := ctx.Value(CustomerVisible).(bool); ok {
		return page
	}
	return true
}

func SetCustomerVisible(ctx echo.Context, value bool) echo.Context {
	return setInContext(ctx, CustomerVisible, value)
}
