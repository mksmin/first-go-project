package main

import (
	"fmt"
	"math"
)

func calculateIMT(userWeight float64, userHeight float64) float64 {
	IMT := userWeight / math.Pow(userHeight, 2)
	return IMT
}

func main() {
	fmt.Println("___ Калькулятор индекса массы тела ___")
	weight, height := userInput()

	imt := calculateIMT(weight, height/100)
	fmt.Printf("Твой ИМТ: %.2f", imt)
}

func userInput() (float64, float64) {
	fmt.Print("Введи свой вес в кг: ")
	var weight float64
	fmt.Scan(&weight)

	fmt.Print("Введи свой рост в сантиметрах: ")
	var height float64
	fmt.Scan(&height)

	return weight, height
}
