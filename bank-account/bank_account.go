// Package account has methods to solve bank exercise
package account

import "sync"

// Account type is concurrency safe thanks to sync.Mutex
type Account struct {
	sync.RWMutex
	amount int64
	closed bool
}

// handleConcurrency function lock and unlock the state of the account
func (a *Account) handleConcurrency(handler func() (int64, bool)) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	return handler()
}

// Open function opens an account if amount is > 0
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	account := &Account{amount: amount}

	return account
}

// Balance function returns account status as it does not modify state we do not need to lock it
func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()

	if a.closed {
		return 0, false
	}
	return a.amount, true
}

// deposit funciton handles the adding and withdrawal operations in account
func (a *Account) deposit(amount int64) func() (int64, bool) {
	return func() (int64, bool) {
		if !a.closed && a.amount+amount >= 0 {
			a.amount += amount
		} else {
			return a.amount, false
		}

		return a.amount, true
	}
}

// Deposit function handles deposit in a concurrency safe way
func (a *Account) Deposit(amount int64) (int64, bool) {
	return a.handleConcurrency(a.deposit(amount))
}

// close function handles closing account
func (a *Account) close() (int64, bool) {
	if a.closed {
		return 0, false
	}

	a.closed = true

	return a.amount, true
}

// Close function uses a concurrency safe version of close
func (a *Account) Close() (int64, bool) {
	return a.handleConcurrency(a.close)
}
