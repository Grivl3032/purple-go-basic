package main

import "fmt"

func main() {
	const usd2Eur float64 = 0.8672
	const usd2Rub float64 = 81.91

	fmt.Println("USD в EUR:", usdEur)
	fmt.Println("USD в RUB:", usdRub)
	fmt.Println("EUR в RUB:", usdRub/usdEur)

}
