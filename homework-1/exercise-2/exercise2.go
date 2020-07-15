package main

import (
	"fmt"
	"math"
)

func main() {

	var legAB float64
	var legAC float64

	fmt.Print("Введите катет AB: ")
	fmt.Scanln(&legAB)

	fmt.Print("Введите катет AC: ")
	fmt.Scanln(&legAC)

	legBC := math.Sqrt(math.Pow(legAB, 2) + math.Pow(legAC, 2)) // Сторона BC это и есть гипотенуза

	fmt.Printf("Площадь треугольника: %.2f\n", (legAB*legAC)/2)
	fmt.Printf("Периметр треугольника: %.2f\n", legAB+legAC+legBC)
	fmt.Printf("Гипотенуза треугольника: %.2f\n", legBC)

}
