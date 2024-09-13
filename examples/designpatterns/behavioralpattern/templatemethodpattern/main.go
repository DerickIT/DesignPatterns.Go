package main

import "fmt"

type PaymentProcessor interface {
	ValidateRequest()
	DebitMoney()
	CalculateFees()
	CreditMoney()
	// ProcessPayment()
}

type BasePaymentProcessor struct {
	processor PaymentProcessor
}

func (b *BasePaymentProcessor) ProcessPayment() {

	b.processor.ValidateRequest()
	b.processor.DebitMoney()
	b.processor.CalculateFees()
	b.processor.CreditMoney()
}

type FriendPaymentProcessor struct {
	BasePaymentProcessor
}

func NewFriendPaymentProcessor() *FriendPaymentProcessor {
	fpp := &FriendPaymentProcessor{}
	fpp.processor = fpp
	return fpp
}

func (f *FriendPaymentProcessor) ValidateRequest() {
	// no validation needed
	fmt.Println("Validating friend payment request")
}

func (f *FriendPaymentProcessor) DebitMoney() {
	// no debit needed
	fmt.Println("Debiting money for friend payment")
}

func (f *FriendPaymentProcessor) CalculateFees() {
	fmt.Println("Calculating fees for friend payment (0%)")
}

func (f *FriendPaymentProcessor) CreditMoney() {
	fmt.Println("Crediting money for friend payment")
}

type MerchanPaymentProcessor struct {
	BasePaymentProcessor
}

func NewMerchanPaymentProcessor() *MerchanPaymentProcessor {
	mpp := &MerchanPaymentProcessor{}
	mpp.processor = mpp
	return mpp
}
func (m *MerchanPaymentProcessor) ValidateRequest() {
	fmt.Println("Validating merchant payment request")
}

func (m *MerchanPaymentProcessor) DebitMoney() {
	fmt.Println("Debiting money for merchant payment")
}

func (m *MerchanPaymentProcessor) CalculateFees() {
	fmt.Println("Calculating fees for merchant payment (2%)")
}

func (m *MerchanPaymentProcessor) CreditMoney() {
	fmt.Println("Crediting money for merchant payment")
}

func main() {
	friendPayment := NewFriendPaymentProcessor()
	friendPayment.ProcessPayment()

	fmt.Println("-------------------")

	merchanPayment := NewMerchanPaymentProcessor()
	merchanPayment.ProcessPayment()
}
