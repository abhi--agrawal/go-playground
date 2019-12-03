package main

import (
	"fmt"
)

func main() {

	ar1 := [3]int{1,2,3}
	fmt.Printf("%v %T\n",ar1 ,ar1 )

	ar2 := [...]int{1,2,3}
	fmt.Printf("%v %T\n",ar2 ,ar2 )

	ar3 := ar1
	ar3[0] = 10
	fmt.Printf("%v %T\n",ar1 ,ar1 )
	fmt.Printf("%v %T\n",ar3 ,ar3 )

	ar4 := &ar1
	ar4[0] = 10
	fmt.Printf("%v %T\n",ar1 ,ar1 )
	fmt.Printf("%v %T\n",ar4 ,ar4 )

	ar5 := make([]int, 3)
	fmt.Printf("%v %T\n",ar5 ,ar5 )
	fmt.Printf("%v %v\n",len(ar5), cap(ar5) )

	ar6 := make([]int, 3, 10)
	fmt.Printf("%v %T\n",ar6 ,ar6 )
	fmt.Printf("%v %v\n",len(ar6), cap(ar6) )

	a := []int{1,2,3,4,5}
	fmt.Println(a)

	b := a[:2]
	fmt.Println(b)
	fmt.Println(a)
	
	c := append(b,a...)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	
	d := append(a[:2],a[3:]...)
	fmt.Println(a)
	fmt.Println(d)

}