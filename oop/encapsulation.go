package main

type BankAccount struct {
	Owner   string
	balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) GetBalance() float64 {
	return b.balance
}
