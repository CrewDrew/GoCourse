package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	const rateUsd float64 = 70.80 // Это курс доллара
	var summRubStr string = "0"

	fmt.Print("Введите сумму в рублях: ")
	fmt.Scanln(&summRubStr)

	summRub, err := strconv.ParseFloat(summRubStr, 64)
	if err != nil {
		log.Fatalln("Вы ввели не число", err)
	}

	fmt.Printf("Сумма по курсу в долларах: %.2f", summRub/rateUsd)

}
