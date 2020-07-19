package main

import (
	"fmt"
	"math"
)

func fib(n int) float64 {

	sqrt5 := math.Sqrt(5)
	phi := (sqrt5 + 1) / 2
	fibNum := math.Round((math.Pow(phi, float64(n)) / sqrt5))
	return fibNum

}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("Число Фибонначи №%d = %0.f\n", i, fib(i))
	}
}
