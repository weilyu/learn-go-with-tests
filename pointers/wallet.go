// Package pointers
package pointers

import "fmt"

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount BitCoin) {
	w.balance -= amount
}

type Stringer interface {
	String() string
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
