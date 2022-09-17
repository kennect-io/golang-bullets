package main

import (
	"fmt"
	"sync"
)

/*

	given a array of numbers
	run 3 go routines
	and those 3 go routines should divide the array amongst themselves add it to the result variable

*/
var wg sync.WaitGroup
var lock sync.Mutex
var result = 0

func addToResult(ch chan int) {
	for { // infinite loop
		val, ok := <-ch //ok == false when all channels ends
		if !ok {
			wg.Done()
			return
		}
		lock.Lock()
		fmt.Println(val)
		result += val
		lock.Unlock()

	}
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	channel := make(chan int) // channels are in golang is blocking operation -> read/write (only for un-buffered channels [this is the example of un-buffered channel example])
	// channel := make(chan int, 10)

	wg.Add(3)

	go addToResult(channel)
	go addToResult(channel)
	go addToResult(channel)

	for _, v := range arr {
		channel <- v
	}

	close(channel)

	wg.Wait()

	fmt.Println("result: ", result)

}
