package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/
func main() {

	var n, a string
	fmt.Print("Enter the name: ")
	fmt.Scan(&n)
	fmt.Print("Enter the address: ")
	fmt.Scan(&a)

	j := map[string]string{
		"name":    n,
		"address": a,
	}

	js, _ := json.Marshal(j)

	os.Stdout.Write(js)

}
