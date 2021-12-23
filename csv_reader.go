package main

import (
    "encoding/csv"
    "encoding/json"
	"io/ioutil"
    "fmt"
    "log"
    "os"
)

func read_csv(csv_name string) {
    f, err := os.Open(csv_name)

    if err != nil {
        log.Fatal(err)
    }

    r := csv.NewReader(f)
    r.Comma = ';'
    r.Comment = '#'

    records, err := r.ReadAll()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(records)
}

func main() {
	read_csv("sample.csv")
}