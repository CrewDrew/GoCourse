package main

import "math"

func SqrtValue(a, b, c float64) [2]float64 {

	var x1, x2 float64

	d := b*b - 4*a*c // Рассчитываем дискриминант

	if d < 0 { // Нет решений
		return [2]float64{}
	}

	if d > 0 { // Условие при дискриминанте больше нуля
		x1 = ((-b) + math.Sqrt(d)) / (2 * a)
		x2 = ((-b) - math.Sqrt(d)) / (2 * a)
	}

	if d == 0 { // Условие для дискриминанта равного нулю
		x1 = -(b / (2 * a))
		x2 = x1
	}

	return [2]float64{x1, x2}

}

func main() {

}
