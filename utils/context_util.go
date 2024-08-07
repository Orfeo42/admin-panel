package utils

import (
	"context"
	"strconv"

	"admin-panel/cmd/enum"

	"github.com/labstack/echo/v4"
)

type contextKey string

const (
	CustomerVisible      contextKey = "CustomerVisible"
	PageContextKey       contextKey = "pages"
	TitleContextKey      contextKey = "title"
	PageNumberContextKey contextKey = "pageNumber"
)

func getStringFromContext(ctx context.Context, key contextKey) string {
	if ctxVar, ok := ctx.Value(key).(string); ok {
		return ctxVar
	}
	return ""
}

func setInEchoContext(echoCtx echo.Context, key contextKey, value any) {
	ctx := context.WithValue(echoCtx.Request().Context(), key, value)
	echoCtx.SetRequest(echoCtx.Request().WithContext(ctx))
}

func GetPage(ctx context.Context) string {
	if page, ok := ctx.Value(PageContextKey).(enum.Page); ok {
		return string(page)
	}
	return ""
}

func SetPage(echoCtx echo.Context, value enum.Page) {
	setInEchoContext(echoCtx, PageContextKey, value)
}

func GetTitle(ctx context.Context) string {
	return getStringFromContext(ctx, TitleContextKey)
}

func SetTitle(echoCtx echo.Context, value string) {
	setInEchoContext(echoCtx, TitleContextKey, value)
}

func GetNextPageNumber(ctx context.Context) string {
	if page, ok := ctx.Value(PageNumberContextKey).(int); ok {
		page++
		return strconv.Itoa(page)
	}
	return strconv.Itoa(2)
}

func SetPageNumber(echoCtx echo.Context, value int) {
	setInEchoContext(echoCtx, PageNumberContextKey, value)
}
