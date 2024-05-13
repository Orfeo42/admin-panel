package main

import (
	"strconv"
	"time"

	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/view/pages"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Orfeo42/admin-panel/controllers"
	"github.com/Orfeo42/admin-panel/repositories"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/component"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	app := echo.New()

	app.Use(middleware.Logger())

	app.Static("/", "web")

	controllers.CustomerController(app)

	controllers.InvoiceController(app)

	controllers.OrderController(app)

	app.GET("/", func(echoCtx echo.Context) error {
		echoCtx = utils.SetPage(echoCtx, enum.Home)
		echoCtx = utils.SetTitle(echoCtx, "Home Page")

		startOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Local)

		earningList, err := repositories.EarningsByMonth(startOfYear, time.Now())
		if err != nil {
			log.Errorf("Error in earnings retrive: %+v", err)
			return err
		}
		data := earningsToAreaChartData(earningList)
		lastMonthEarning := data.Data[len(data.Data)-1]
		yearEarning := float64(0)
		for _, v := range data.Data {
			yearEarning += v
		}
		return utils.Render(pages.HomeView(data, lastMonthEarning, yearEarning), echoCtx)
	})

	err := app.Start(":8080")
	if err != nil {
		log.Fatalf("Error Stating application: %+v", err)
	}
}

func earningsToAreaChartData(earningList *[]repositories.EarningsByMonthResult) component.AreaChartData {
	var labels []string
	var data []float64
	addYear := false
	if len(*earningList) > 12 {
		addYear = true
	}
	for _, earning := range *earningList {
		label := time.Month(earning.Month).String()
		if addYear {
			label = strconv.Itoa(earning.Year) + "_" + label
		}
		labels = append(labels, label)

		data = append(data, float64(earning.Amount)/100)

	}
	return component.AreaChartData{
		Data:   data,
		Labels: labels,
	}
}
