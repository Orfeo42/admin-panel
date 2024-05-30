package home

import (
	"admin-panel/cmd/enum"
	"admin-panel/cmd/web/components"
	"admin-panel/internal/database"
	"admin-panel/utils"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func RegisterHomeRoutes(application *echo.Echo) {

	homeGroup := application.Group("")

	controller := HomeControllerInstance()

	homeGroup.GET("", controller.HomeHandler)
}

var controllerInstance *homeController

type HomeController interface {
	HomeHandler(echoCtx echo.Context) error
}

type homeController struct {
	invRep database.InvoiceRepository
}

func HomeControllerInstance() HomeController {
	if controllerInstance != nil {
		return controllerInstance
	}
	controllerInstance = &homeController{
		invRep: database.InvoiceRepositoryInstance(),
	}
	return controllerInstance
}

func (c *homeController) HomeHandler(echoCtx echo.Context) error {
	utils.SetPage(echoCtx, enum.Home)
	utils.SetTitle(echoCtx, "Home Page")

	startOfYear := time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, time.Local)

	dateFrom := startOfYear
	dateTo := time.Now()

	labels := getMonthsBetweenDates(dateFrom, dateTo)

	salesList, err := c.invRep.SalesByMonth(dateFrom, dateTo)
	if err != nil {
		log.Errorf("Error in Sales retrive: %+v", err)
		return err
	}
	collectedList, err := c.invRep.CollectedByMonth(startOfYear, time.Now())
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

	dataSets := []components.AreaChartDataset{
		components.LoadBlueData("Fatturato", salesData),
		components.LoadGreenData("Riscosso", collectedData),
	}

	log.Infof("dataSets: %+v", dataSets)

	areaChartParams := components.AreaChartParams{
		XAxesLabels: labels,
		DataSets:    dataSets,
	}

	homePrams := HomeParameters{
		AreaChartParams: areaChartParams,
		SalesMonth:      salesMonth,
		SalesYear:       salesYear,
		CollectedMonth:  collectedMonth,
		CollectedYear:   collectedYear,
	}
	return utils.Render(HomeView(homePrams), echoCtx)
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
