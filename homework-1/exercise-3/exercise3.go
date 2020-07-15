package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {

	var summRubStr string = "0"
	var percentByYearStr string = "0"

	fmt.Print("Введите сумму вклада: ")
	fmt.Scanln(&summRubStr)

	summRub, err := strconv.ParseFloat(summRubStr, 64)
	if err != nil {
		log.Fatalln("Сумма должна быть числом", err)
	}

	fmt.Print("Введите процент по вкладу в год: ")
	fmt.Scanln(&percentByYearStr)

	percentByYear, err := strconv.ParseFloat(percentByYearStr, 64)
	if err != nil {
		log.Fatalln("Процент должен быть числом", err)
	}

	fmt.Printf("Сумма по вкладу через 5 лет: %.2f", summRub*math.Pow(1+percentByYear/100, 5)) // Формула для расчета суммы вклада с ежегодной капитализацией

}
