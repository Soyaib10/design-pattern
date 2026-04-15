package payment

import "fmt"

func PaymentFactory(paymentType string) (Payment, error) {
	switch paymentType {
	case "creditcard":
		return &CreditCard{}, nil
	case "paypal":
		return &PayPal{}, nil
	case "banktransfar":
		return &BankTransfar{}, nil
	case "bkash":
		return &BKash{}, nil
	default:
		return nil, fmt.Errorf("payment type not supported: %s\n", paymentType)
	}
}
