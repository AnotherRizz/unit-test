package bank

import (
	"fmt"
)

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		 return fmt.Errorf("deposit amount must be greater than zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		 return fmt.Errorf("withdrawal amount must be greater than zero")
	}
	if amount > a.Balance {
		 return fmt.Errorf("insufficient funds")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}