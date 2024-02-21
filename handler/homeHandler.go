package handler

import (
	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/home"
	"github.com/labstack/echo/v4"
)

func HomeShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, enum.HOME)
	echoCtx = utils.SetTitle(echoCtx, "Home Page")
	return utils.Render(home.HomeView(), echoCtx)
}
