package charts

import (
	"admin-panel/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"time"
)

func RegisterRoutes(application *echo.Echo) {
	chartGroup := application.Group("/chart")
	controller := getControllerInstance()
	chartGroup.GET("/sales", controller.sales)
	chartGroup.GET("/collected", controller.collected)
	chartGroup.GET("/to-be-collected", controller.toBeCollected)
	chartGroup.GET("/main", controller.main)
}

func getControllerInstance() Controller {
	if controllerInstance != nil {
		return controllerInstance
	}
	controllerInstance = &chartController{
		invRep: database.InvoiceRepositoryInstance(),
	}
	return controllerInstance
}

var controllerInstance *chartController

type Controller interface {
	sales(echoCtx echo.Context) error
	collected(echoCtx echo.Context) error
	toBeCollected(echoCtx echo.Context) error
	main(echoCtx echo.Context) error
}

type chartController struct {
	invRep database.InvoiceRepository
}

func (c *chartController) sales(echoCtx echo.Context) error {

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
func (c *chartController) collected(echoCtx echo.Context) error {

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

func (c *chartController) toBeCollected(echoCtx echo.Context) error {

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

func (c *chartController) main(echoCtx echo.Context) error {

	//chartData, err := earningsToAreaChartData(earnings)
	//return echoCtx.JSON(http.StatusOK, chartData)
	//
	return echoCtx.JSON(http.StatusOK, map[string]string{
		"labels":        "",
		"sales":         "",
		"collected":     "",
		"toBeCollected": "",
	})
}
