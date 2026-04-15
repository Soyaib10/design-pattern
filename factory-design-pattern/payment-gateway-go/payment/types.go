package payment

import "fmt"

type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) error {
	fmt.Printf("%.2f tk has been credited from credit card\n\n", amount)
	return nil
}

type PayPal struct{}

func (p *PayPal) Pay(amount float64) error {
	fmt.Printf("%.2f tk has been credited from paypal\n\n", amount)
	return nil
}

type BankTransfar struct{}

func (b *BankTransfar) Pay(amount float64) error {
	fmt.Printf("%.2f tk has been credited from bank transfar\n\n", amount)
	return nil
}

type BKash struct{}

func (b *BKash) Pay(amount float64) error {
	fmt.Printf("%.2f tk has been credited from bkash\n\n", amount)
	return nil
}
