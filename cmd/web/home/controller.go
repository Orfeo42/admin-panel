package home

import (
	"admin-panel/cmd/enum"
	"admin-panel/internal/database"
	"admin-panel/utils"
	"errors"
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

	homeGroup.GET("/sales/graph", controller.salesGraph)
	homeGroup.GET("/sales/month", controller.salesMonth)
	homeGroup.GET("/sales/year", controller.salesYear)
	homeGroup.GET("/sales/all", controller.salesAll)

	homeGroup.GET("/collected/graph", controller.collectedGraph)
	homeGroup.GET("/collected/month", controller.collectedMonth)
	homeGroup.GET("/collected/year", controller.collectedYear)
	homeGroup.GET("/collected/all", controller.collectedAll)

	homeGroup.GET("/to-be-collected/graph", controller.toBeCollectedGraph)
	homeGroup.GET("/to-be-collected/month", controller.toBeCollectedMonth)
	homeGroup.GET("/to-be-collected/year", controller.toBeCollectedYear)
	homeGroup.GET("/to-be-collected/all", controller.toBeCollectedAll)

	homeGroup.GET("main-chart", controller.mainChart)
}

var controllerInstance *controller

type Controller interface {
	home(echoCtx echo.Context) error

	salesGraph(echoCtx echo.Context) error
	salesMonth(echoCtx echo.Context) error
	salesYear(echoCtx echo.Context) error
	salesAll(echoCtx echo.Context) error

	collectedGraph(echoCtx echo.Context) error
	collectedMonth(echoCtx echo.Context) error
	collectedYear(echoCtx echo.Context) error
	collectedAll(echoCtx echo.Context) error

	toBeCollectedGraph(echoCtx echo.Context) error
	toBeCollectedMonth(echoCtx echo.Context) error
	toBeCollectedYear(echoCtx echo.Context) error
	toBeCollectedAll(echoCtx echo.Context) error

	mainChart(echoCtx echo.Context) error
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

func (c *controller) salesGraph(echoCtx echo.Context) error {

	dateTo := time.Now()
	dateFrom := dateTo.AddDate(-1, 0, 0)

	salesList, err := c.invRep.SalesByMonth(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in Sales retrive: %+v", err)
		return err
	}
	chartData, err := earningsToAreaChartData(salesList)

	return echoCtx.JSON(http.StatusOK, chartData)
}
func (c *controller) collectedGraph(echoCtx echo.Context) error {

	dateTo := time.Now()
	dateFrom := dateTo.AddDate(-1, 0, 0)

	collectedList, err := c.invRep.CollectedByMonth(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in collected retrive: %+v", err)
		return err
	}
	chartData, err := earningsToAreaChartData(collectedList)

	return echoCtx.JSON(http.StatusOK, chartData)
}

func (c *controller) toBeCollectedGraph(echoCtx echo.Context) error {

	dateTo := time.Now()
	dateFrom := dateTo.AddDate(-1, 0, 0)

	earnings, err := c.invRep.ToBeCollectedByMonth(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in Sales retrive: %+v", err)
		return err
	}
	chartData, err := earningsToAreaChartData(earnings)

	return echoCtx.JSON(http.StatusOK, chartData)
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

func (c *controller) mainChart(echoCtx echo.Context) error {
	return echoCtx.String(http.StatusOK, "")
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

type ChartData struct {
	Labels []string
	Values []float64
}

func earningsToAreaChartData(earningList []database.MoneyByMonthResult) (ChartData, error) {
	var labels []string
	var data []float64
	for _, earning := range earningList {
		month, err := nomeMese(earning.Month)
		month = month + " - " + strconv.Itoa(earning.Year)
		if err != nil {
			log.Errorf("Error in collected retrive: %+v", err)
			return ChartData{}, err
		}
		labels = append(labels, month)

		amount := float64(earning.Amount) / 100
		amount = utils.ToFixed(amount, 2)
		data = append(data, amount)
	}
	return ChartData{
		labels,
		data,
	}, nil
}

var mesi = map[int]string{
	1:  "Gennaio",
	2:  "Febbraio",
	3:  "Marzo",
	4:  "Aprile",
	5:  "Maggio",
	6:  "Giugno",
	7:  "Luglio",
	8:  "Agosto",
	9:  "Settembre",
	10: "Ottobre",
	11: "Novembre",
	12: "Dicembre",
}

// Funzione per ottenere il nome del mese
func nomeMese(numero int) (string, error) {
	if mese, exists := mesi[numero]; exists {
		return mese, nil
	}
	return "", errors.New("numero del mese non valido")
}
