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

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${uri} (${remote_ip}) ${latency_human}\n",
		Output: app.Logger.Output(),
	}))

	app.Static("/", "web")

	controllers.CustomerController(app)

	controllers.InvoiceController(app)

	app.GET("/", homeController)

	err := app.Start(":8080")
	if err != nil {
		log.Fatalf("Error Stating application: %+v", err)
	}
}

func homeController(echoCtx echo.Context) error {
	echoCtx = utils.SetPage(echoCtx, enum.Home)
	echoCtx = utils.SetTitle(echoCtx, "Home Page")

	startOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Local)

	salesList, err := repositories.SalesByMonth(startOfYear, time.Now())
	if err != nil {
		log.Errorf("Error in Sales retrive: %+v", err)
		return err
	}
	collectedList, err := repositories.CollectedByMonth(startOfYear, time.Now())
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}
	salesData := earningsToAreaChartData(salesList)
	salesMonth := salesData.Data[len(salesData.Data)-1]
	salesYear := float64(0)
	for _, v := range salesData.Data {
		salesYear += v
	}
	collectedData := earningsToAreaChartData(collectedList)
	collectedMonth := salesData.Data[len(collectedData.Data)-1]
	collectedYear := float64(0)
	for _, v := range collectedData.Data {
		collectedYear += v
	}

	salesMonth = utils.ToFixed(salesMonth, 2)
	salesYear = utils.ToFixed(salesYear, 2)
	collectedMonth = utils.ToFixed(collectedMonth, 2)
	collectedYear = utils.ToFixed(collectedYear, 2)

	return utils.Render(pages.HomeView(salesData, collectedData, salesMonth, salesYear, collectedMonth, collectedYear), echoCtx)
}

func earningsToAreaChartData(earningList *[]repositories.MoneyByMonthResult) component.AreaChartData {
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

		amount := float64(earning.Amount) / 100

		amount = utils.ToFixed(amount, 2)
		data = append(data, amount)

	}
	return component.AreaChartData{
		Data:   data,
		Labels: labels,
	}
}
