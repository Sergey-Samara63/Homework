package main

import "fmt"

func main() {
	const usdEur = 0.85
	const usdRub = 78.40
	const eurRub = usdRub / usdEur
	userMessage := getUserInput()
	fmt.Print (userMessage)
}

func getUserInput() (string) {
	var userInput string
	fmt.Print("Введите свое сообщение: ")
	fmt.Scan(&userInput)
	return userInput
}
//заготовка
func convert (number float64, usd float64, eur float64) float64 {



}