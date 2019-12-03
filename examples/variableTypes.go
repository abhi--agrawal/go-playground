package main

import (
	"fmt"
)

func main() {

	var a bool
	fmt.Printf("%v %T\n",a,a)
	a = true
	fmt.Printf("%v %T\n",a,a)

	var b int64 = 10000000
	var c int8 = 677777

	fmt.Println(b)
	fmt.Println(c)

	var d uint32 = 1
	fmt.Println(d)
	d = -1
	fmt.Println(d)

	e := 100
	f := 3
	g := e / f
	fmt.Printf("%v %T\n",g,g)

	var h complex64 = 3 + 4i
	fmt.Printf("%v %T\n",h,h)
	fmt.Printf("%v %T\n",real(h),real(h))
	fmt.Printf("%v %T\n",imag(h),imag(h))

	var i complex128 = 3 + 4i
	fmt.Printf("%v %T\n",i,i)
	fmt.Printf("%v %T\n",real(i),real(i))
	fmt.Printf("%v %T\n",imag(i),imag(i))

	j := "String"
	fmt.Printf("%v %T\n",j,j)
	fmt.Printf("%v %T\n",j[1],j[1])

	var k []byte = []byte(j)
	fmt.Printf("%v %T\n",k,k)

	var l rune = 'a'
	fmt.Printf("%v %T\n",l,l)

}
