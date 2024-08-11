package pointer_error

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance int
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() (amount int) {
	return w.balance
}

func (w *Wallet) Withdraw(amount int) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
