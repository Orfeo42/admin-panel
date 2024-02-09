package handler

import (
	"github.com/Orfeo42/admin-panel/model"
	"github.com/labstack/echo/v4"

	"github.com/Orfeo42/admin-panel/view/page/home"
)

func HomeShow(context echo.Context) error {
	films := []model.FilmModel{
		{Title: "Titolo1", Director: "Director1"},
		{Title: "Titolo2", Director: "Director2"},
		{Title: "Titolo3", Director: "Director3"},
	}
	return render(home.HomeView(films), context)
}
