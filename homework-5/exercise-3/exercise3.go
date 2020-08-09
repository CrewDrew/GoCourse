package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("test.csv")
	if err != nil {
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comment = '#'
	reader.Comma = ';'

	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(record)
	}
}
