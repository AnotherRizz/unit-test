package bank

import (
	
	"errors"
)

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		  return errors.New("withdrawal amount must be greater than zero")
	}
	if amount > a.Balance {
		  return errors.New("insufficient funds")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}