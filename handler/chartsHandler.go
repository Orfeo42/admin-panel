package handler

import (
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/home"
	"github.com/labstack/echo/v4"
)

func ChartsShow(echoCtx echo.Context) error {
	return utils.Render(home.HomeView(), echoCtx)
}
