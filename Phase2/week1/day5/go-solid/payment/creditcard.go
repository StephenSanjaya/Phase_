package payment

import "fmt"

type CreditCard struct{}

func (c CreditCard) Pay(amount float64) error {
	// 1. validasi card number
	// 2. validasi balance

	fmt.Printf("Paid %f using credit card\n", amount)
	return nil
}
