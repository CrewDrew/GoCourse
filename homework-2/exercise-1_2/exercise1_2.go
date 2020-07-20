package main

import (
	"fmt"
	"log"
)

func parityOfNumber(value int64) (isEven bool, parityOfNumber string) {
	if value%2 == 0 {
		return true, "четное"
	}
	return false, "нечетное"

}

func multipleOfThree(value int64) (isMultipleOfThree bool, multipleOfThree string) {
	if value%3 == 0 {
		return true, "делится на три"
	}
	return false, "не делится на три"
}

func main() {

	var someNumber int64

	fmt.Println("Программа определения четности числа")
	fmt.Print("Введите число: ")
	_, err := fmt.Scanf("%d\n", &someNumber)

	if err != nil {
		log.Fatalln("Вы ввели не число")
	}

	_, isEvenStr := parityOfNumber(someNumber)
	_, multipleOfThreeStr := multipleOfThree(someNumber)

	fmt.Printf("Вы ввели %s число\n", isEvenStr)
	fmt.Printf("Число %d %s\n", someNumber, multipleOfThreeStr)

}
