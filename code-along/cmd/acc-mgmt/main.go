package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var wg sync.WaitGroup

func creditTransaction(accBal *int) {
	for i := 0; i < 1000; i++ {
		go func() {
			// lock.Lock()
			*accBal++
			// lock.Unlock()
			wg.Done()
		}()
	}
}

func debitTransaction(accBal *int) {
	for i := 0; i < 1000; i++ {
		// if i == 2 {
		// 	continue
		// }
		go func() {
			// lock.Lock()
			*accBal--
			// lock.Unlock()
			wg.Done()
		}()
	}
}

func simulateTransactions() int {
	accBal := 0

	wg.Add(2000)
	creditTransaction(&accBal)
	debitTransaction(&accBal)

	wg.Wait()
	return accBal
}

func main() {
	fmt.Println(simulateTransactions())
}
