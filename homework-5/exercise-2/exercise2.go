package main

import (
	"fmt"
	"os"
)

const fileName = "exercise2.go"

func main() {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	defer file.Close()

	bs := make([]byte, 10)

	for {
		n, _ := file.Read(bs)
		if n == 0 { // если конец файла
			break // выходим из цикла
		}
		fmt.Println(string(bs[:n]))
	}
}
