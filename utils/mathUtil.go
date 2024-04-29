package utils

import "strconv"

func FormatAmount(amount int) string {
	result := float64(amount) / 100
	return strconv.FormatFloat(result, 'f', 2, 64)
}
