package controllers

import (
	"strconv"
	"time"

	"github.com/Orfeo42/admin-panel/enum"
	"github.com/Orfeo42/admin-panel/mvc/repositories"
	"github.com/Orfeo42/admin-panel/utils"
	"github.com/Orfeo42/admin-panel/view/component"
	"github.com/Orfeo42/admin-panel/view/pages/viewhome"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func HomeController(application *echo.Echo) {

	homeGroup := application.Group("")

	homeGroup.GET("", homeController)
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

	homePrams := viewhome.HomeParameters{
		AreaChartParams: areaChartParams,
		SalesMonth:      salesMonth,
		SalesYear:       salesYear,
		CollectedMonth:  collectedMonth,
		CollectedYear:   collectedYear,
	}
	return utils.Render(viewhome.HomeView(homePrams), echoCtx)
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
