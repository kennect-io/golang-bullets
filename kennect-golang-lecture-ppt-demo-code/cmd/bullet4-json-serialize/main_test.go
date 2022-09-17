package main

import (
	"reflect"
	"testing"
)

// Equal just iterates over one slices and checks if all elements are the same in both slices
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// TestFilter tests the filter slice function
func TestFilter(t *testing.T) {
	specs := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 2, 3}, output: []int{}},
		{input: []int{1, 23, 3}, output: []int{23}},
		{input: []int{1, 23, 3, 33, 44, 55}, output: []int{23, 33, 44, 55}},
	}

	for _, spec := range specs {
		filteredSlice := FilterLessThan18(spec.input)

		if !Equal(spec.output, filteredSlice) {
			t.Errorf("Output %v doesnt match Expected %v", filteredSlice, spec.output)
		}
	}
}

// TestFilterWithDeepCheck tests the filter slice function using the reflect.DeepEqual checking for slice
func TestFilterWithDeepCheck(t *testing.T) {
	specs := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 2, 3}, output: []int{}},
		{input: []int{1, 23, 3}, output: []int{23}},
		{input: []int{1, 23, 3, 33, 44, 55}, output: []int{23, 33, 44, 55}},
	}

	for _, spec := range specs {
		filteredSlice := FilterLessThan18(spec.input)

		// to actually Deep Check things which catches the nil slice
		if !reflect.DeepEqual(spec.output, filteredSlice) {
			t.Errorf("Output %v doesnt match Expected %v", filteredSlice, spec.output)
		}

	}
}
