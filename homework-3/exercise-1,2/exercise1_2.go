package main

import (
	"fmt"
)

type AutomobileDescribe struct {
	CarBrand               string
	Year                   int
	TrunkVolume            float32
	IsEngineRunning        bool
	IsWindowsOpen          bool
	TrunkVolumeFreePercent float32
}

func main() {
	var auto1, auto2, auto4 AutomobileDescribe

	var auto3 = AutomobileDescribe{
		CarBrand:               "Optimus Prime",
		Year:                   1990,
		TrunkVolume:            100000,
		IsEngineRunning:        false,
		IsWindowsOpen:          false,
		TrunkVolumeFreePercent: 100,
	}

	auto1.CarBrand = "Batmobile"
	auto1.IsEngineRunning = false
	auto1.IsWindowsOpen = true
	auto1.TrunkVolume = 100
	auto1.TrunkVolumeFreePercent = 100
	auto1.Year = 1966

	auto2.CarBrand = "Ninja Turtles Van"
	auto2.IsEngineRunning = true
	auto2.IsWindowsOpen = true
	auto2.TrunkVolume = 1000
	auto2.TrunkVolumeFreePercent = 80
	auto2.Year = 1985

	fmt.Println("auto1", auto1)
	fmt.Println("auto2", auto2)
	fmt.Println("auto3", auto3)

	auto4 = auto1
	auto4.IsWindowsOpen = false
	fmt.Println("auto4", auto4)
	fmt.Println("auto1", auto1)
	fmt.Println("auto1 = auto4", auto1 == auto4)

	auto4 = auto1
	auto4.IsWindowsOpen = true
	fmt.Println("auto4", auto4)
	fmt.Println("auto1", auto1)
	fmt.Println("auto1 = auto4", auto1 == auto4)

}
