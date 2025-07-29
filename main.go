package main

import "fmt"

const (
	UsdToEur = 0.92   // 1 USD = 0.92 EUR
	UsdToRub = 89.50  // 1 USD = 89.50 RUB
)

func main() {
	// EUR to RUB = (USD to RUB) / (USD to EUR)
	EurToRub := UsdToRub / UsdToEur
	fmt.Printf("1 EUR = %.2f RUB\n", EurToRub)
}
