package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsvFile(path string) [][]string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Unable to find current ", err)
	}
	absPath := pwd + "/" + path
	f, err := os.Open(absPath)
	if err != nil {
		log.Fatalln("Unable to read input file "+path, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse file as CSV for "+path, err)
	}

	return records
}
