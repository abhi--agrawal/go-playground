package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i float32 = 99.999
	fmt.Println(i)

	var j int
	// j = i  THIS IS AN ERROR
	j = int(i)
	fmt.Println(j)

	var k int = 45
	l := float32(k) //This will give float32, not float64
	m := 99.999
	fmt.Printf("%v %T\n",k,k)
	fmt.Printf("%v %T\n",l,l)
	fmt.Printf("%v %T\n",m,m)

	o := string(k) // o = '-'
	fmt.Printf("%v %T\n",o,o)
	p := strconv.Itoa(k)
	fmt.Printf("%v %T\n",p,p)
}
