package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

//In Go if a symbol (so variables, types, functions et al) starts with a lowercase symbol
//then it is private outside the package it's defined in.
type Wallet struct {
	balance Bitcoin
}

//In Go, when you call a function or a method the arguments are copied.
//a pointer to a wallet
//struct pointers
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
