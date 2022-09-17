package main

import (
	"fmt"
	"sync"
)

/*
	Problem statement
	given a slice of numbers
	run 3 go routines
	and those 3 go routines should add the numbers in the slice and print the result
*/

// wg used to wait for all the go routines to be processed
var wg sync.WaitGroup
// lock variable to gain an exclusive mutex lock
var lock sync.Mutex

// result to store the result of the sum
var result = 0

// AddToResultDeadlock results in deadlock as it only reads from the channel once and the function ends
func AddToResultDeadlock(ch chan int) {
	lock.Lock()
	val := <-ch
	fmt.Println("val", val)
	result += val
	lock.Unlock()
}

// AddToResultForeverDeadlock results in deadlock as it forever tries to read from the channel again resulting in a deadlock
func AddToResultForeverDeadlock(ch chan int) {
	for {
		lock.Lock()
		val := <-ch
		fmt.Println("val", val)
		result += val
		lock.Unlock()
	}
}

/* 
	AddToResultForeverDeadlock works perfectly 
	provided the channel "ch" is closed 
	as it tries to read from the channel until it can read no more
	then exits from the function
*/
func AddToResultWorkingAsExpected(ch chan int) {
	for {
		val, ok := <-ch
		if ok {
			lock.Lock()
			fmt.Println("val", val)
			result += val
			lock.Unlock()
		} else {
			return
		}
	}
}

func main() {
	numsToAdd := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	channel := make(chan int)

	wg.Add(3)

	go func() {
		AddToResultDeadlock(channel)
		wg.Done()
	}()
	go func() {
		AddToResultDeadlock(channel)
		wg.Done()
	}()
	go func() {
		AddToResultDeadlock(channel)
		wg.Done()
	}()

	for _, v := range numsToAdd {
		channel <- v
	}
	
	// not closing the channel also causes deadlock
	// close(channel)

	wg.Wait()

	fmt.Println(result)
}
