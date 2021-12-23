package main

import "csv2json/core"

func main() {
	records := core.ReadCsv("sample.csv")
	core.Convert2Json(records)
}
