package pointers

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds means a wallet does not have enough Bitcoin to perform a withdraw.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin represents a number of Bitcoins.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoin someone owns.
type Wallet struct {
	balance Bitcoin
}

// Wallet stores the number of Bitcoin someone owns.
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of memory in Deposit is %p \n", &w.balance)
	w.balance += amount
}

// Balance returns the number of Bitcoin a wallet has.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Balance returns the number of Bitcoin a wallet has.
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
