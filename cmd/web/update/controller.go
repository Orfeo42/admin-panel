package update

import (
	"admin-panel/cmd/web/update/validation"
	"admin-panel/internal/database"
	"admin-panel/preload/dbupdate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"mime/multipart"
	"net/http"
)

func RegisterRoutes(application *echo.Echo) {
	group := application.Group("/update-data")
	controller := getControllerInstance()
	group.POST("", controller.update)
}

func getControllerInstance() Controller {
	if controllerInstance != nil {
		return controllerInstance
	}
	controllerInstance = &preloadController{
		invRep: database.InvoiceRepositoryInstance(),
	}
	return controllerInstance
}

var controllerInstance *preloadController

type Controller interface {
	update(echoCtx echo.Context) error
}

type preloadController struct {
	invRep database.InvoiceRepository
}

func (p *preloadController) update(c echo.Context) error {

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "File non trovato"})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Impossibile aprire il file"})
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Error("Error closing file")
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Errore nella lettura del file"})
	}

	customerList, invoiceList, err := validation.ValidateExcel(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Errore validazione excel"})
	}

	err = dbupdate.EmptyData()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Errore svuotamento vecchi"})
	}

	err = dbupdate.LoadData(customerList, invoiceList)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Errore caricamento dati"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":        "Upload effettuato con successo",
		"customer-count": len(customerList),
		"invoice-count":  len(invoiceList),
	})
}
