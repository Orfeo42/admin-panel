package handler

import (
	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/home"
	"github.com/labstack/echo/v4"
)

func ChartsShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, enum.CHARTS)
	echoCtx = utils.SetTitle(echoCtx, "Charts")
	return utils.Render(home.HomeView(), echoCtx)
}
