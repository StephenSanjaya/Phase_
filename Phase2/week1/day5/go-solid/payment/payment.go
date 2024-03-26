package payment

type PaymentMethod interface {
	Pay(amount float64) error
}

type PaymentProcessor struct {
	paymentMethod PaymentMethod
}

func NewPaymentProcessor(method PaymentMethod) *PaymentProcessor {
	return &PaymentProcessor{
		paymentMethod: method,
	}
}

func (processor *PaymentProcessor) ProcessPayment(amount float64) error {
	return processor.paymentMethod.Pay(amount)
}
