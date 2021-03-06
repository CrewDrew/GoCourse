package main

import (
	"fmt"

	"calculator"
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		if input == "help" {
			fmt.Println("Для расчета введите операнды разделенные операциями без пробелов")
			fmt.Println("Например, 1+2+3")
			fmt.Println("Результат будет выведен на экран")
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
