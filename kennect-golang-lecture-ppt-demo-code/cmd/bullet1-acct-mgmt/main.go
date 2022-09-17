package main

import (
	"fmt"
	"sync"
)

// lock variable to gain an exclusive mutex lock
var lock sync.Mutex

// wg used to wait for all the go routines to be processed
var wg sync.WaitGroup

// CreditTransactionWhichExit exits as no waitgroup
func CreditTransactionWhichExit(balance *int) {
	// credit 1000 dollars
	for i := 0; i < 1000; i++ {
		go func() {
			*balance++
		}()
	}
	
}

// DebitTransactionWithRace exits as no waitgroup
func DebitTransactionWhichExit(balance *int) {
	// debit 1000 dollars
	for i := 0; i < 1000; i++ {
		go func() {
			*balance--
		}()
	}
}

// CreditTransactionWithRace has a race condition as the balance variable is incremented without exclusive mutex lock
func CreditTransactionWithRace(balance *int) {
	// credit 1000 dollars
	for i := 0; i < 1000; i++ {
		go func() {
			*balance++
			wg.Done()
		}()
	}
	
}

// DebitTransactionWithRace has a race condition as the balance variable is decremented without exclusive mutex lock
func DebitTransactionWithRace(balance *int) {
	// debit 1000 dollars
	for i := 0; i < 1000; i++ {
		go func() {
			*balance--
			wg.Done()
		}()
	}
}

// CreditTransactionWithLock very safely increments balance variable as it has a exclusive mutex lock
func CreditTransactionWithLock(balance *int) {
	// credit 1000 dollars
	for i := 0; i < 1000; i++ {
		go func() {
			lock.Lock()
			*balance++
			lock.Unlock()
			wg.Done()
		}()
	}
	
}

// DebitTransactionWithLock very safely decrements balance variable as it has a exclusive mutex lock
func DebitTransactionWithLock(balance *int) {
	// debit 1000 dollars
	for i := 0; i < 1000; i++ {
		go func() {
			lock.Lock()
			*balance--
			lock.Unlock()
			wg.Done()
		}()
	}
}

// SimulateTransactionsWithoutWaitGroup to simulate credit and debit transactions without waiting for go routines
func SimulateTransactionsWithoutWaitGroup() int {
	accountBalance := 0
	// we are going to be running 2000 goroutines
	CreditTransactionWhichExit(&accountBalance)
	DebitTransactionWhichExit(&accountBalance)

	return accountBalance
}

// SimulateTransactionsWithRace to simulate credit and debit transactions with race
func SimulateTransactionsWithRace() int {
	accountBalance := 0
	// we are going to be running 2000 goroutines
	wg.Add(2000)
	CreditTransactionWithRace(&accountBalance)
	DebitTransactionWithRace(&accountBalance)
	wg.Wait()

	return accountBalance
}

// SimulateTransactionsWithLock to simulate credit and debit transactions safely and correctly without any unexpected output
func SimulateTransactionsWithLock() int {
	accountBalance := 0
	// we are going to be running 2000 goroutines
	wg.Add(2000)
	CreditTransactionWithLock(&accountBalance)
	DebitTransactionWithLock(&accountBalance)
	wg.Wait()

	return accountBalance
}

func main() {
	acctBal := SimulateTransactionsWithLock()
	fmt.Printf("account balance: %v$\n", acctBal)
}
