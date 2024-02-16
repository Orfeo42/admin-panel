package handler

import (
	"github.com/Orfeo42/admin-panel/model"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/home"
	"github.com/labstack/echo/v4"
)

func HomeShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, "HOME")
	films := []model.FilmModel{
		{Title: "Titolo1", Director: "Director1"},
		{Title: "Titolo2", Director: "Director2"},
		{Title: "Titolo3", Director: "Director3"},
	}
	return utils.Render(home.HomeView(films), echoCtx)
}
