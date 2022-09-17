package main

import "testing"

// run tests with race using command from root folder "go test -race ./cmd/bullet1-acct-mgmt"

// TestSimulateTransactionsWithRace tests the function having data race condition
func TestSimulateTransactionsWithRace(t *testing.T) {
	acctBal := SimulateTransactionsWithRace()
	if acctBal != 0 {
		t.Errorf("got %v$, expected to be 0$", acctBal)
	}
}

// TestSimulateTransactionsWithLock tests the functions without data race condition
func TestSimulateTransactionsWithLock(t *testing.T) {
	acctBal := SimulateTransactionsWithLock()
	if acctBal != 0 {
		t.Errorf("got %v$, expected to be 0$", acctBal)
	}
}

