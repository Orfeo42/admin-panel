package handlers

import (
	"github.com/Orfeo42/admin-panel/data"
	"github.com/Orfeo42/admin-panel/enum/pages"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/page/film"
	"github.com/labstack/echo/v4"
)

func RefeshFilmList(context echo.Context) error {
	films := []data.FilmModel{
		{Title: "Titolo1", Director: "Director1"},
		{Title: "Titolo2", Director: "Director2"},
		{Title: "Titolo3", Director: "Director3"},
	}
	return utils.Render(film.FilmListShow(films), context)
}

func AddFilm(context echo.Context) error {
	//time.Sleep(1 * time.Second)
	req := context.Request()
	filmAggiunto := data.FilmModel{
		Title:    req.PostFormValue("title"),
		Director: req.PostFormValue("director"),
	}
	return utils.Render(film.FilmItemShow(filmAggiunto), context)
}

func FilmShow(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, pages.FilmList)
	echoCtx = utils.SetTitle(echoCtx, "Films")
	films := []data.FilmModel{
		{Title: "Titolo1", Director: "Director1"},
		{Title: "Titolo2", Director: "Director2"},
		{Title: "Titolo3", Director: "Director3"},
	}
	return utils.Render(film.FilmView(films), echoCtx)
}