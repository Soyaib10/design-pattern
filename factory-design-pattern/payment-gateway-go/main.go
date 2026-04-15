package main

import (
	"fmt"

	"github.com/Soyaib10/factory-design-pattern/payment-gateway-go/payment"
)

func main() {
	paymentTypes := []string{"creditcard", "paypal", "banktransfar", "bkash", "wth"}

	for _, paymentType := range paymentTypes {
		fmt.Printf("Payment Type: %s\n", paymentType)
		PaymentProcessor, err := payment.PaymentFactory(paymentType)
		fmt.Printf("Type: %T\n", PaymentProcessor)

		if err != nil {
			fmt.Println(err)
			continue
		}

		err = PaymentProcessor.Pay(343.323)
		if err != nil {
			fmt.Println("payment failed\n", err)
		}
	}
}
