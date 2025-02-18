package charts

import (
	"admin-panel/internal/database"
	"admin-panel/utils"
	"errors"
	"github.com/labstack/gommon/log"
	"strconv"
)

func moneyByMonthToChartData(earningList []database.MoneyByMonthResult) (ChartData, error) {
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
		amount := utils.AmountIntegerToFloat(earning.Amount)
		data = append(data, amount)
	}
	return ChartData{
		labels,
		data,
	}, nil
}

type ChartData struct {
	Labels []string  `json:"labels"`
	Values []float64 `json:"values"`
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

func nomeMese(numero int) (string, error) {
	if mese, exists := mesi[numero]; exists {
		return mese, nil
	}
	return "", errors.New("numero del mese non valido")
}

type MainChartData struct {
	Labels        []string  `json:"labels"`
	Sales         []float64 `json:"sales"`
	Collected     []float64 `json:"collected"`
	ToBeCollected []float64 `json:"to-be-collected"`
}

func earningsToAreaChartData(salesByDate []database.MoneyByDateResult, collectedByDate []database.MoneyByDateResult, toBeCollectedByDate []database.MoneyByDateResult) MainChartData {
	labels := make([]string, len(salesByDate))
	sales := make([]float64, len(salesByDate))
	collected := make([]float64, len(salesByDate))
	toBeCollected := make([]float64, len(salesByDate))
	for i := range salesByDate {
		labels[i] = salesByDate[i].Date.Format("02-01-2006")
		sales[i] = utils.AmountIntegerToFloat(salesByDate[i].Amount)
		collected[i] = utils.AmountIntegerToFloat(collectedByDate[i].Amount)
		toBeCollected[i] = utils.AmountIntegerToFloat(toBeCollectedByDate[i].Amount)
	}
	return MainChartData{
		Labels:        labels,
		Sales:         sales,
		Collected:     collected,
		ToBeCollected: toBeCollected,
	}
}

func moneyByDateToChartData(valByDate []database.MoneyByDateResult) ChartData {
	labels := make([]string, len(valByDate))
	values := make([]float64, len(valByDate))
	for i := range valByDate {
		labels[i] = valByDate[i].Date.Format("02-01-2006")
		values[i] = utils.AmountIntegerToFloat(valByDate[i].Amount)
	}
	return ChartData{
		Labels: labels,
		Values: values,
	}
}
