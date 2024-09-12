package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paying %.2f using Credit Card", amount)
}

type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paying %.2f using PayPal", amount)
}

type BankTransferPayment struct{}

func (b *BankTransferPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paying %.2f using Bank Transfer", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) ExecutePayment(amount float64) string {
	return pc.strategy.Pay(amount)
}

func main() {
	context := &PaymentContext{}

	context.SetStrategy(&CreditCardPayment{})
	fmt.Println(context.ExecutePayment(100.54))

	context.SetStrategy(&PayPalPayment{})
	fmt.Println(context.ExecutePayment(200.54))

	context.SetStrategy(&BankTransferPayment{})
	fmt.Println(context.ExecutePayment(300.54))
}
