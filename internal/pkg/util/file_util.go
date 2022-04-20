package util

import (
	"encoding/csv"
	"log"
	"os"
)

func CreateFie(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Cannot create file %q: %s\n", filePath, err)
	}
	file.Close()
}

func WriteCsv(filePath string, record []string) {
	file, err := os.OpenFile("test.csv", os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("Cannot open file %q: %s\n", filePath, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(record)
}
