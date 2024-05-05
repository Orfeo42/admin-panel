package utils

import (
	"encoding/csv"
	"github.com/labstack/gommon/log"
	"os"
)

func ReadCsvFile(path string) ([][]string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Error("Unable to find current", err)
		return [][]string{}, err
	}
	absPath := pwd + "/" + path
	f, err := os.Open(absPath)
	if err != nil {
		log.Errorf("Unable to read input file %s: %+v", path, err)
		return [][]string{}, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Errorf("Unable to close file %s: %+v", path, err)
		}
	}(f)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Errorf("Unable to parse file as CSV for %s: %+v", path, err)
		return [][]string{}, err
	}

	return records, nil
}
