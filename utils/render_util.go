package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(component templ.Component, context echo.Context) error {
	return component.Render(context.Request().Context(), context.Response())
}
