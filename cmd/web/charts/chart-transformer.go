package charts

import (
	"admin-panel/internal/database"
	"admin-panel/utils"
	"errors"
	"github.com/labstack/gommon/log"
	"strconv"
)

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

type ChartData struct {
	Labels []string
	Values []float64
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
