// Package account has methods to solve bank exercise
package account

import "sync"

// Account type is concurrency safe thanks to sync.Mutex
type Account struct {
	sync.Mutex
	Amount int64
	Closed bool
}

// Open function opens an account if amount is > 0
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	account := &Account{Amount: amount}

	return account
}

// Balance function returns account status as it does not modify state we do not need to lock it
func (a *Account) Balance() (int64, bool) {
	if a.Closed {
		return 0, false
	}
	return a.Amount, a.Amount >= 0
}

// Deposit function adds or withdraws money from account, so it needs to be locked
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if a.Amount+amount < 0 {
		return a.Amount, false
	}
	a.Amount += amount

	return a.Balance()
}

// Close function closes account so it needs to handle concurrency
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	amount, balance := a.Balance()
	a.Closed = true

	return amount, balance
}
