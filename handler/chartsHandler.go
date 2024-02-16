package handler

import (
	"github.com/Orfeo42/admin-panel/model"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/home"
	"github.com/labstack/echo/v4"
)

func ChartsShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, "charts")
	films := []model.FilmModel{
		{Title: "Titolo4", Director: "Director4"},
		{Title: "Titolo5", Director: "Director5"},
		{Title: "Titolo6", Director: "Director6"},
	}
	return utils.Render(home.HomeView(films), echoCtx)
}
