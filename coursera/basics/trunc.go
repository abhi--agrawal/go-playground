package main

import (
	"fmt"
)

/**
 * Write a program which prompts the user to enter a floating point number and prints the integer
 * which is a truncated version of the floating point number that was entered.
 * Truncation is the process of removing the digits to the right of the decimal place.
 */

func main() {

	var fN float64
	fmt.Print("Enter the floating point number: ")

	fmt.Scanf("%f", &fN)

	fmt.Println(int(fN))

}
