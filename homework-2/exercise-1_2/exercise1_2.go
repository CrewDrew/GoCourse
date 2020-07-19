package main

import (
	"fmt"
	"log"
)

func parityOfNumber(value int64) string {
	if value%2 == 0 {
		return "четное"
	} else {
		return "нечетное"
	}
}

func multipleOfThree(value int64) string {
	if value%3 == 0 {
		return "делится на три"
	} else {
		return "не делится на три"
	}
}

func main() {

	var someNumber int64

	fmt.Println("Программа определения четности числа")
	fmt.Print("Введите число: ")
	_, err := fmt.Scanf("%d\n", &someNumber)

	if err != nil {
		log.Fatalln("Вы ввели не число")
	}

	fmt.Printf("Вы ввели %s число\n", parityOfNumber(someNumber))
	fmt.Printf("Число %d %s\n", someNumber, multipleOfThree(someNumber))

}
