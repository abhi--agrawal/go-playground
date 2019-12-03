package main

import (
	"fmt"
)

func GenDisplaceFn(a,v0,s0 float64) func (float64) float64 {
	
	return func (t float64) float64  {
		return (a * t * t)/2 + (v0 * t) + s0;
	}

}

func main()  {
	
	var a,v,s,t float64
	fmt.Print("Enter accerlation: ")
	fmt.Scan(&a)
	fmt.Print("Enter inital velocity: ")
	fmt.Scan(&v)
	fmt.Print("Enter displacement: ")
	fmt.Scan(&s)

	time := GenDisplaceFn(a,v,s)
	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	fmt.Println(time(t))
}

