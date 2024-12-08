package main

import (
	"encoding/csv"
	"os"
)

func (m *model) readCSVFile() {
	// read in CSV data
	f, err := os.Open("Input/data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvData, err := csvReader.ReadAll()

	if err != nil {
		panic(err)
	}

	headers := csvData[0] // []string
	rows := csvData[1:]   // [][]string

	m.csvData.headers = headers
	m.csvData.rows = make([]map[string]string, 0, len(rows))
	// rows := []map[string]string{}

	for _, data := range rows {
		csvData := map[string]string{}
		for j, header := range headers {
			csvData[header] = data[j]
		}
		m.csvData.rows = append(m.csvData.rows, csvData)
	}

	// m.csvData.headers = append(m.csvData.headers, csvData[0]...)
	// m.csvData.data = append(m.csvData.data, csvData[1:]...)

}
