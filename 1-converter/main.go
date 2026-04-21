package main

import "fmt"

const USD_TO_EUR float64 = 0.8672
const USD_TO_RUB float64 = 81.91
const EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR

func getUserImput() (float64, string, string) {
	var amount float64
	var currencyTo string
	var currencyFrom string

	fmt.Print("Введите сумму:")
	fmt.Scan(&amount)

	for {
		fmt.Print("Введите валюту(USD/EUR/RUB):")
		fmt.Scan(&currencyFrom)

		if currencyFrom != "USD" && currencyFrom != "RUB" && currencyFrom != "EUR" {
			fmt.Println("Ошибка ввода попробуй еще раз")
			continue
		} else {
			break
		}
	}

	if currencyFrom == "EUR" {
		for {
			fmt.Print("в какую валюту хотите перевести", "(USD/RUB):")
			fmt.Scan(&currencyTo)

			if currencyTo != "USD" && currencyTo != "RUB" {
				fmt.Println("Ошибка ввода попробуй еще раз")
				continue
			} else {
				break
			}
		}
	} else if currencyFrom == "USD" {
		for {
			fmt.Print("в какую валюту хотите перевести", "(EUR/RUB):")
			fmt.Scan(&currencyTo)

			if currencyTo != "EUR" && currencyTo != "RUB" {
				fmt.Println("Ошибка ввода попробуй еще раз")
				continue
			} else {
				break
			}
		}
	} else if currencyFrom == "RUB" {
		for {
			fmt.Print("в какую валюту хотите перевести", "(USD/EUR):")
			fmt.Scan(&currencyTo)

			if currencyTo != "USD" && currencyTo != "EUR" {
				fmt.Println("Ошибка ввода попробуй еще раз")
				continue
			} else {
				break
			}
		}
	}

	return amount, currencyFrom, currencyTo
}

func convert(amount float64, currencyTo string, currencyFrom string) float64 {
	if currencyFrom == "USD" && currencyTo == "EUR" {
		return amount * USD_TO_EUR
	} else if currencyFrom == "EUR" && currencyTo == "USD" {
		return amount / USD_TO_EUR
	} else if currencyFrom == "USD" && currencyTo == "RUB" {
		return amount * USD_TO_RUB
	} else if currencyFrom == "RUB" && currencyTo == "USD" {
		return amount / USD_TO_RUB
	} else if currencyFrom == "EUR" && currencyTo == "RUB" {
		return amount * EUR_TO_RUB
	} else if currencyFrom == "RUB" && currencyTo == "EUR" {
		return amount / EUR_TO_RUB
	} else if currencyFrom == currencyTo {
		return amount
	}

	return 0
}

func main() {
	amount, currencyFrom, currencyTo := getUserImput()
	result := convert(amount, currencyTo, currencyFrom)
	fmt.Printf("%.2f", result)
}
