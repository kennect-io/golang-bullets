package main

import (
	"encoding/json"
	"fmt"
)

// Config to store the config read from json string/file
type Config struct {
	// []Person is slice definition slice is a reference to the underlying array
	// [3]Person is array definition reference doesnt get passed but cant do this as we are reading json

	People []Person `json:"people"`
}

// Person to store age and name
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// AgeAfter50Years adds 50 years to all the person's ages
func AgeAfter50Years(people []Person) []Person {
	peopleCopy := people
	for index := range peopleCopy {
		peopleCopy[index].Age = peopleCopy[index].Age + 50
	}

	return peopleCopy
}

// AgeAfter50YearsWithCopy adds 50 years to all the person's ages but first copies the iterating slice first before mutating
func AgeAfter50YearsWithCopy(people []Person) []Person {
	peopleCopy := make([]Person, len(people))
	copy(peopleCopy, people)
	for index := range peopleCopy {
		peopleCopy[index].Age = peopleCopy[index].Age + 50
	}

	return peopleCopy
}

// PrintMemoryAddresses prints memory address pointers for the struct, and its elements
func PrintMemoryAddresses(p []Person, variableName string) {
	fmt.Printf("address of %s.people: %p\n", variableName, &p)
	for index := range p {
		fmt.Printf("address of elem %s: %p\n", p[index].Name, &p[index])
	}
}

// PrintAges prints ages for all people
func PrintAges(p []Person, suffix string) {
	for index := range p {
		fmt.Printf("%s: Age of %s is %d \n", suffix, p[index].Name, p[index].Age)
	}
}

func main() {

	var p1 Config
	json.Unmarshal([]byte(jsonStr), &p1)

	c1 := p1.People

	c2 := AgeAfter50Years(c1)
	// c2 := AgeAfter50YearsWithCopy(c1)

	PrintAges(c1, "At 0 years")
	fmt.Println("_____________________________________")
	PrintAges(c2, "After 50 years")

	// PrintMemoryAddresses(c1, "c1")
	// PrintMemoryAddresses(c2, "c2")
}

var jsonStr = `{
	"people": [
		{
			"name": "Berlin",
			"age": 35
		},
		{
			"name": "Moscow",
			"age": 32
		},
		{
			"name": "Nairobi",
			"age": 29
		}
	]
}`
