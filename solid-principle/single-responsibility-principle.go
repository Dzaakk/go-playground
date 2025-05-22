package main

import "fmt"

type Invoice struct {
	Items []string
	Total float64
}

func (i *Invoice) CalculateTotal() float64 {
	return i.Total
}

type InvoicePrinter struct{}

func (p *InvoicePrinter) Print(i *Invoice) {
	fmt.Println("Total:", i.Total)
}

// ----------------------------
// wrong example
func (p *Invoice) Print(i *Invoice) {
	fmt.Println("Total:", i.Total)
}
