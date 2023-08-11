// Namn
// Datum
// Implementera de 5 funktionerna och testa att det blir korrekt utskrift

package main

import "fmt"

func calculateAverage(numbers []int) float64 {

}

func isPalindrome(input string) bool {

}

func calculate(op string, a, b float64) float64 {

}

func celsiusToFahrenheit(celsius float64) float64 {

}

func fahrenheitToCelsius(fahrenheit float64) float64 {

}

func main() {
	// Funktion 1 - Printsatsen ska skriva ut talet 30
	numbers := []int{10, 20, 30, 40, 50}
	average := calculateAverage(numbers)
	fmt.Printf("Medelvärde: %.2f\n", average)

	// Funktion 2 - Första printsatsen ska skriva att radar är ett palindrom
	// Andra printsatsen ska skriva ut att monkey inte är ett palindrom
	word := "radar"
	if isPalindrome(word) {
		fmt.Println(word, " är ett palindrom!")
	} else {
		fmt.Println(word, " är inte ett palindrom.")
	}

	word2 := "monkey"
	if isPalindrome(word2) {
		fmt.Println(word2, " är ett palindrom!")
	} else {
		fmt.Println(word2, " är inte ett palindrom.")
	}

	var (
		operator   string
		num1, num2 float64
	)

	// Funktion 3 - Ska skriva ut korrekt resultat utifrån input
	fmt.Print("Ange en operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Ange det första talet: ")
	fmt.Scan(&num1)

	fmt.Print("Ange det andra talet: ")
	fmt.Scan(&num2)

	result := calculate(operator, num1, num2)
	fmt.Println("Resultat:", result)

	// Funktion 4
	// Ska skriva ut att 25 celsius är 77 Fahrenheit
	celsius := 25.0
	fahrenheit := celsiusToFahrenheit(celsius)
	fmt.Println(celsius, " Celsius är", fahrenheit, " Fahrenheit")

	// Funktion 5
	// Ska skriva ut att 68 Fahrenheit är 20 celsius
	fahrenheit = 68.0
	celsius = fahrenheitToCelsius(fahrenheit)
	fmt.Println(fahrenheit, " Fahrenheit är", celsius, " Celsius")

}
