package handler

import (
	"time"

	"github.com/Orfeo42/admin-panel/model"
	"github.com/Orfeo42/admin-panel/utils"
	component "github.com/Orfeo42/admin-panel/view/component/filmItem"
	"github.com/labstack/echo/v4"
)

func RefeshFilmList(context echo.Context) error {
	films := []model.FilmModel{
		{Title: "Titolo1", Director: "Director1"},
		{Title: "Titolo2", Director: "Director2"},
		{Title: "Titolo3", Director: "Director3"},
	}
	return utils.Render(component.FilmListShow(films), context)
}

func AddFilm(context echo.Context) error {
	time.Sleep(1 * time.Second)
	req := context.Request()
	film := model.FilmModel{Title: req.PostFormValue("title"),
		Director: req.PostFormValue("director"),
	}
	return utils.Render(component.FilmItemShow(film), context)
}
