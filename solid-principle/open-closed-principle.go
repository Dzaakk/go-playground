package main

type Discount interface {
	Apply(price float64) float64
}

type NoDiscount struct{}

func (d NoDiscount) Apply(price float64) float64 {
	return price
}

type PercentageDiscount struct{}

func (d PercentageDiscount) Apply(price float64) float64 {
	return price * 0.9
}

func Checkout(price float64, d Discount) float64 {
	return d.Apply(price)
}
