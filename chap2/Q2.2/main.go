package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	var hoge = []string{"header1", "header2", "header3"}
	var fuga = []string{"body1", "body2", "body3"}
	writer.Write(hoge)
	writer.Write(fuga)
	writer.Flush()
}
