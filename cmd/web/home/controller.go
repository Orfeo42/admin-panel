package home

import (
	"admin-panel/cmd/enum"
	"admin-panel/internal/database"
	"admin-panel/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func RegisterRoutes(application *echo.Echo) {

	homeGroup := application.Group("")

	controller := getControllerInstance()

	homeGroup.GET("", controller.home)

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

var controllerInstance *homeController

type Controller interface {
	home(echoCtx echo.Context) error

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

type homeController struct {
	invRep database.InvoiceRepository
}

func getControllerInstance() Controller {
	if controllerInstance != nil {
		return controllerInstance
	}
	controllerInstance = &homeController{
		invRep: database.InvoiceRepositoryInstance(),
	}
	return controllerInstance
}

func (c *homeController) home(echoCtx echo.Context) error {
	utils.SetPage(echoCtx, enum.Home)

	utils.SetTitle(echoCtx, "Home Page")

	return utils.Render(HomeView(), echoCtx)
}

func (c *homeController) salesMonth(echoCtx echo.Context) error {

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

func (c *homeController) salesYear(echoCtx echo.Context) error {
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

func (c *homeController) salesAll(echoCtx echo.Context) error {
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

func (c *homeController) collectedMonth(echoCtx echo.Context) error {
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

func (c *homeController) collectedYear(echoCtx echo.Context) error {
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

func (c *homeController) collectedAll(echoCtx echo.Context) error {
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

func (c *homeController) toBeCollectedMonth(echoCtx echo.Context) error {
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

func (c *homeController) toBeCollectedYear(echoCtx echo.Context) error {
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

func (c *homeController) toBeCollectedAll(echoCtx echo.Context) error {
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
