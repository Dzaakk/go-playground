package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64)
}

type PayPal struct{}

func (p *PayPal) Pay(amount float64) {
	fmt.Println("Paid with Paypal: ", amount)
}

type Stripe struct{}

func (s *Stripe) Pay(amount float64) {
	fmt.Println("Paid with Stripe: ", amount)
}

type PaymentContext struct {
	Strategy PaymentStrategy
}

func (pc *PaymentContext) Execute(amount float64) {
	pc.Strategy.Pay(amount)
}

// func main() {
// 	paymentCtx := PaymentContext{Strategy: &Stripe{}}
// 	paymentCtx.Execute(250.0)
// }
