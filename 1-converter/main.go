package main

import "fmt"

func main() {
	const USD_TO_EUR float64 = 0.8672
	const USD_TO_RUB float64 = 81.91
	const EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR
	fmt.Println("USD в EUR:", USD_TO_EUR)
	fmt.Println("USD в RUB:", USD_TO_RUB)
	fmt.Println("EUR в RUB:", EUR_TO_RUB)
}
