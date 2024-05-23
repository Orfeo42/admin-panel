package main

import (
	"strconv"
	"time"

	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/view/component"
	"github.com/Orfeo42/admin-panel/view/pages"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Orfeo42/admin-panel/mvc/controllers"
	"github.com/Orfeo42/admin-panel/mvc/repositories"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	app := echo.New()

	app.HTTPErrorHandler = customHTTPErrorHandler
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
	utils.SetPage(echoCtx, enum.Home)
	utils.SetTitle(echoCtx, "Home Page")

	startOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Local)

	dateFrom := startOfYear
	dateTo := time.Now()

	labels := getMonthsBetweenDates(dateFrom, dateTo)

	salesList, err := repositories.SalesByMonth(dateFrom, dateTo)
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
	collectedData := earningsToAreaChartData(collectedList)

	salesMonth := salesData[len(salesData)-1]
	salesMonth = utils.ToFixed(salesMonth, 2)

	salesYear := float64(0)
	for _, v := range salesData {
		salesYear += v
	}
	salesYear = utils.ToFixed(salesYear, 2)

	collectedMonth := collectedData[len(collectedData)-1]
	collectedMonth = utils.ToFixed(collectedMonth, 2)

	collectedYear := float64(0)
	for _, v := range collectedData {
		collectedYear += v
	}
	collectedYear = utils.ToFixed(collectedYear, 2)

	dataSets := []component.AreaChartDataset{
		component.LoadBlueData("Fatturato", salesData),
		component.LoadGreenData("Riscosso", collectedData),
	}

	log.Infof("dataSets: %+v", dataSets)

	areaChartParams := component.AreaChartParams{
		XAxesLabels: labels,
		DataSets:    dataSets,
	}

	return utils.Render(pages.HomeView(areaChartParams, salesMonth, salesYear, collectedMonth, collectedYear), echoCtx)
}

func getMonthsBetweenDates(dateFrom, dateTo time.Time) []string {
	addYear := false
	if monthsBetween(dateFrom, dateTo) > 12 {
		addYear = true
	}
	var months []string
	for !dateFrom.After(dateTo) {
		month := dateFrom.Month().String()
		if addYear {
			month = strconv.Itoa(dateFrom.Year()) + "_" + month
		}
		months = append(months, month)
		dateFrom = dateFrom.AddDate(0, 1, 0)
	}
	return months
}

func monthsBetween(dateFrom, dateTo time.Time) int {
	if dateFrom.After(dateTo) {
		dateFrom, dateTo = dateTo, dateFrom
	}

	years := dateTo.Year() - dateFrom.Year()
	months := int(dateTo.Month()) - int(dateFrom.Month())

	totalMonths := years*12 + months

	if dateFrom.Day() > dateTo.Day() {
		totalMonths--
	}

	return totalMonths
}

func earningsToAreaChartData(earningList *[]repositories.MoneyByMonthResult) []float64 {
	var data []float64
	for _, earning := range *earningList {
		amount := float64(earning.Amount) / 100
		amount = utils.ToFixed(amount, 2)
		data = append(data, amount)
	}
	return data
}

// TODO DA TESTARE E FARE PER BENE
func customHTTPErrorHandler(err error, echoCtx echo.Context) {
	/*code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	echoCtx.Logger().Error(err)
	errorPage := fmt.Sprintf("%d.html", code)
	if err := echoCtx.File(errorPage); err != nil {
		echoCtx.Logger().Error(err)
	}*/
}
