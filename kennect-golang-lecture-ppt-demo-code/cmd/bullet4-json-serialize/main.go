package main

import (
	"encoding/json"
	"fmt"
)

// FilterLessThan18 filters slice elements where values are less than 18 nothing special just a random filter function
func FilterLessThan18(a []int) []int {
	// one way to avoid nil initializations of slices
	// filteredSlice := []int{}

	var filteredSlice []int
	for _, v := range a {
		if v > 18 {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice
}

func main() {
	a := []int{1, 2, 3, 4, 54, 6, 7}
	b := []int{1, 2, 3, 4, 5, 6, 7}

	jsonStr, _ := json.Marshal(FilterLessThan18(a))
	fmt.Println("Json string:", string(jsonStr[:]))
	jsonStr, _ = json.Marshal(FilterLessThan18(b))
	fmt.Println("Json string:", string(jsonStr[:]))
}
