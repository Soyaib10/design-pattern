package main

import "fmt"

func main() {
	paymentType := "paypal"
	amount := 100.0

	processPayment(paymentType, amount)
}

func processPayment(paymentType string, amount float64) {
	if paymentType == "creditcard" {
		processCreditCardPayment(amount)
	} else if paymentType == "paypal" {
		processPaypalPayment(amount)
	} else {
		fmt.Println("unsupported")
	}
}

func processCreditCardPayment(amount float64) {
	fmt.Printf("credit card %f", amount)
}

func processPaypalPayment(amount float64) {
	fmt.Printf("Paypal %f", amount)
}
