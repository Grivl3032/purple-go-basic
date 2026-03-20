package main

import "fmt"
//возвращает то что пользователь ввел 
func getUserImput() float64, string, string {
	var emount float64
	var currencyTo string
	var currencyFrom string
	fmt.Print("Введите сумму:")
	fmt.Scan(&emount)
	ftm.Print("Введите валюту(USD/EUR/RUB):")
	fmt.Scan(&currencyFrom)
	fmt.Print("в какую валюту хотите перевести(USD/EUR/RUB)?")
	fmt.Scan(&currencyTo)

	return 
}

func convert(emount float64, currencyTo string, currencyFrom string) float64 {
}

func main() {
	const USD_TO_EUR float64 = 0.8672
	const USD_TO_RUB float64 = 81.91
	const EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR
}
