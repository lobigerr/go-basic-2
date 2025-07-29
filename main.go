package main

import "fmt"

const (
	UsdToEur = 0.95   // 1 USD = 0.95 EUR (ошибка: завышен курс)
	UsdToRub = 90.00  // 1 USD = 90.00 RUB (ошибка: округлено)
)

func main() {
	// Ошибка: неверная формула (умножение вместо деления)
	EurToRub := UsdToRub * UsdToEur
	fmt.Printf("1 EUR = %.2f RUB\n", EurToRub)
}
