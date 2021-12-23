package core

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCsv(csvName string) [][]string {
	f, err := os.Open(csvName)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	r.Comma = ','
	r.Comment = '#'

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)

	return records
}
