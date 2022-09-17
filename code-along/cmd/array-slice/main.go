package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	MoneyHeistCharacters []Person `json:"people"`
}

type Person struct {
	Name string
	Age  int
}

func AgeAfter50Years(arr []Person) []Person {
	arrCopy := make([]Person, len(arr))
	copy(arrCopy, arr)
	for i := range arrCopy {
		arrCopy[i].Age += 50
	}
	return arrCopy
}

func printMemAdd(p []Person) {
	fmt.Printf("add of p: %p\n", &p)
	for index := range p {
		fmt.Printf("addr of p %s %p\n", p[index].Name, &p[index])
	}
}

func main() {
	var p1 Config

	json.Unmarshal([]byte(jsonStr), &p1)

	arr := AgeAfter50Years(p1.MoneyHeistCharacters)
	fmt.Println(arr)
	printMemAdd(arr)

	fmt.Println(p1.MoneyHeistCharacters)
	printMemAdd(p1.MoneyHeistCharacters)
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
