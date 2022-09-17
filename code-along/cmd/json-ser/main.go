package main

import (
	"encoding/json"
	"fmt"
)

func filterValueGreaterThan18(arr []int) []int {
	// var filteredValues []int // default value is null
	var filteredValues = []int{} // default value is empty array i.e. []
	for _, v := range arr {
		if v > 18 {
			filteredValues = append(filteredValues, v)
		}
	}
	return filteredValues
}

func main() {
	a := []int{1, 2, 3, 4, 45, 4, 2}
	b := []int{1, 2, 3, 4, 2, 0, 0, 0}

	jsonStr, _ := json.Marshal(filterValueGreaterThan18((a)))
	fmt.Println(string(jsonStr))

	jsonStr2, _ := json.Marshal(filterValueGreaterThan18((b)))
	fmt.Println(string(jsonStr2))
}
