package main

import (
	"fmt"
)

const n int = 100
const initN int = 550 // Столько нужно примерно перебрать целых чисел, чтобы набрать 100 простых

func main() {

	var primeNumbers [100]int
	var intNumbers [600]int

	for i := range intNumbers {
		intNumbers[i] = i
	}

	for p := 2; p < initN; p++ {
		if intNumbers[p] != 0 {
			for j := p * p; j < initN+1; j += p {
				intNumbers[j] = 0
			}
		}
	}

	j := 0
	for i := range intNumbers {
		if intNumbers[i] > 0 && j < 100 {
			primeNumbers[j] = intNumbers[i]
			j++
		}
	}

	for i := range primeNumbers {
		fmt.Printf("Простое число №%d : %d\n", i+1, primeNumbers[i])
	}

}
