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

func (csv *CSVData) GetCSVData(csvHeaders []string, csvData [][]string) {

	// Detect headings from no data
	for i, header := range csvHeaders {

		if csvData[i][0] == "" {
			csv.headings[i] = header
		}
	}

	csv.headers = csvHeaders[len(csv.headings):]

	for i, data := range csvData {
		data = data[len(csv.headings):]
		copy(csv.data[i], data)
	}

}
