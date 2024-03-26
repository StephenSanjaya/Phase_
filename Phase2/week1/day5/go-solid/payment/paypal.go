package payment

import "fmt"

type PayPal struct{}

func (pp PayPal) Pay(amount float64) error {
	// 1. validasi account
	// 2. validasi balance

	fmt.Printf("Paid %f using paypal\n", amount)
	return nil
}
