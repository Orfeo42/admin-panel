package home

import (
	"admin-panel/cmd/enum"
	"admin-panel/cmd/web/components"
	"admin-panel/internal/database"
	"admin-panel/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func RegisterRoutes(application *echo.Echo) {

	homeGroup := application.Group("")

	controller := getControllerInstance()

	homeGroup.GET("", controller.home)
	homeGroup.GET("/graph", controller.graph)

	homeGroup.GET("/sales/month", controller.salesMonth)
	homeGroup.GET("/sales/year", controller.salesYear)
	homeGroup.GET("/sales/all", controller.salesAll)

	homeGroup.GET("/collected/month", controller.collectedMonth)
	homeGroup.GET("/collected/year", controller.collectedYear)
	homeGroup.GET("/collected/all", controller.collectedAll)

	homeGroup.GET("/to-be-collected/month", controller.toBeCollectedMonth)
	homeGroup.GET("/to-be-collected/year", controller.toBeCollectedYear)
	homeGroup.GET("/to-be-collected/all", controller.toBeCollectedAll)
}

var controllerInstance *controller

type Controller interface {
	home(echoCtx echo.Context) error
	graph(echoCtx echo.Context) error
	salesMonth(echoCtx echo.Context) error
	salesYear(echoCtx echo.Context) error
	salesAll(echoCtx echo.Context) error
	collectedMonth(echoCtx echo.Context) error
	collectedYear(echoCtx echo.Context) error
	collectedAll(echoCtx echo.Context) error

	toBeCollectedMonth(echoCtx echo.Context) error
	toBeCollectedYear(echoCtx echo.Context) error
	toBeCollectedAll(echoCtx echo.Context) error
}

type controller struct {
	invRep database.InvoiceRepository
}

func getControllerInstance() Controller {
	if controllerInstance != nil {
		return controllerInstance
	}
	controllerInstance = &controller{
		invRep: database.InvoiceRepositoryInstance(),
	}
	return controllerInstance
}

func (c *controller) home(echoCtx echo.Context) error {
	utils.SetPage(echoCtx, enum.Home)

	utils.SetTitle(echoCtx, "Home Page")

	return utils.Render(HomeView(), echoCtx)
}

func (c *controller) graph(echoCtx echo.Context) error {

	startOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Local)

	dateFrom := startOfYear
	dateTo := time.Now()

	salesList, err := c.invRep.SalesByMonth(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in Sales retrive: %+v", err)
		return err
	}
	salesData := earningsToAreaChartData(salesList)

	collectedList, err := c.invRep.CollectedByMonth(startOfYear, dateTo)
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}
	collectedData := earningsToAreaChartData(collectedList)

	dataSets := []components.AreaChartDataset{
		components.LoadBlueData("Fatturato", salesData),
		components.LoadGreenData("Riscosso", collectedData),
	}

	log.Infof("dataSets: %+v", dataSets)

	labels := getMonthsBetweenDates(dateFrom, dateTo)
	areaChartParams := components.AreaChartParams{
		XAxesLabels: labels,
		DataSets:    dataSets,
	}

	return utils.Render(components.AreaChart("salesChart", areaChartParams), echoCtx)
}

func (c *controller) salesMonth(echoCtx echo.Context) error {

	dateTo := time.Now()

	startOfMonth := time.Date(dateTo.Year(), dateTo.Month(), 1, 0, 0, 0, 0, dateTo.Location())

	dateFrom := startOfMonth

	amount, err := c.invRep.SalesTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in Sales retrive: %+v", err)
		return err
	}
	formattedAmount := utils.AmountIntegerToString(amount)
	return echoCtx.String(http.StatusOK, formattedAmount)
}

func (c *controller) salesYear(echoCtx echo.Context) error {
	dateTo := time.Now()

	startOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Local)

	dateFrom := startOfYear

	amount, err := c.invRep.SalesTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}

	formattedAmount := utils.AmountIntegerToString(amount)
	return echoCtx.String(http.StatusOK, formattedAmount)
}

func (c *controller) salesAll(echoCtx echo.Context) error {
	dateTo := time.Now()

	startOfYear := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)

	dateFrom := startOfYear

	amount, err := c.invRep.SalesTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}

	formattedAmount := utils.AmountIntegerToString(amount)
	return echoCtx.String(http.StatusOK, formattedAmount)
}

func (c *controller) collectedMonth(echoCtx echo.Context) error {
	dateTo := time.Now()

	startOfMonth := time.Date(dateTo.Year(), dateTo.Month(), 1, 0, 0, 0, 0, dateTo.Location())

	dateFrom := startOfMonth

	amount, err := c.invRep.CollectedTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}

	formattedAmount := utils.AmountIntegerToString(amount)
	return echoCtx.String(http.StatusOK, formattedAmount)
}

func (c *controller) collectedYear(echoCtx echo.Context) error {
	dateTo := time.Now()

	startOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Local)

	dateFrom := startOfYear

	amount, err := c.invRep.CollectedTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}

	formattedAmount := utils.AmountIntegerToString(amount)
	return echoCtx.String(http.StatusOK, formattedAmount)
}

func (c *controller) collectedAll(echoCtx echo.Context) error {
	dateTo := time.Now()

	startOfYear := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)

	dateFrom := startOfYear

	amount, err := c.invRep.CollectedTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}

	formattedAmount := utils.AmountIntegerToString(amount)
	return echoCtx.String(http.StatusOK, formattedAmount)
}

func (c *controller) toBeCollectedMonth(echoCtx echo.Context) error {
	dateTo := time.Now()

	startOfMonth := time.Date(dateTo.Year(), dateTo.Month(), 1, 0, 0, 0, 0, dateTo.Location())

	dateFrom := startOfMonth

	amount, err := c.invRep.ToBeCollectedTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}

	formattedAmount := utils.AmountIntegerToString(amount)
	return echoCtx.String(http.StatusOK, formattedAmount)
}

func (c *controller) toBeCollectedYear(echoCtx echo.Context) error {
	dateTo := time.Now()

	startOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Local)

	dateFrom := startOfYear

	amount, err := c.invRep.ToBeCollectedTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}

	formattedAmount := utils.AmountIntegerToString(amount)
	return echoCtx.String(http.StatusOK, formattedAmount)
}

func (c *controller) toBeCollectedAll(echoCtx echo.Context) error {
	dateTo := time.Now()

	startOfYear := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local)

	dateFrom := startOfYear

	amount, err := c.invRep.ToBeCollectedTotal(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}

	formattedAmount := utils.AmountIntegerToString(amount)
	return echoCtx.String(http.StatusOK, formattedAmount)
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

func earningsToAreaChartData(earningList []database.MoneyByMonthResult) []float64 {
	var data []float64
	for _, earning := range earningList {
		amount := float64(earning.Amount) / 100
		amount = utils.ToFixed(amount, 2)
		data = append(data, amount)
	}
	return data
}
