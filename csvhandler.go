package main

import (
	"encoding/csv"
	"os"
)

func readCSVFile(filename string) ([]string, [][]string, error) {
	// read in CSV data
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	headers := data[0]
	rows := data[1:]

	return headers, rows, err
}
