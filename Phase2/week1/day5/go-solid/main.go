package main

import (
	"Phase2/week1/day5/go-solid/payment"
	"fmt"
)

func main() {
	// Create payment processor for credit card
	ccProcessor := payment.NewPaymentProcessor(payment.CreditCard{})
	err := ccProcessor.ProcessPayment(100.0)
	if err != nil {
		fmt.Println("PAYMENT WITH CC FAILED :", err)
	}

	ppProcessor := payment.NewPaymentProcessor(payment.PayPal{})
	err = ppProcessor.ProcessPayment(50.50)
	if err != nil {
		fmt.Println("PAYMENT WITH Paypal FAILED :", err)
	}
}
