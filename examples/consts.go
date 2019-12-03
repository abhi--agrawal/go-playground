package main

import (
	"fmt"
)

func main() {
	const a = 16
	// var c = 19
	var b int16 = 32
	fmt.Printf("%v %T\n",a ,a )
	fmt.Printf("%v %T\n",a + b,a + b)
	// fmt.Printf("%v %T\n",c + b,c + b)

	const (
		z = iota
		y = iota
		x = iota
	)
	fmt.Printf("%v %T\n",x ,x )
	fmt.Printf("%v %T\n",y,y)
	fmt.Printf("%v %T\n",z,z)

	const (
		w = iota
		v 
		u 
	)
	fmt.Printf("%v %T\n",u, u )
	fmt.Printf("%v %T\n",v,v)
	fmt.Printf("%v %T\n",w,w)

	const (
		t = 2 * iota
		s 
		r 
	)
	fmt.Printf("%v %T\n",t,t)
	fmt.Printf("%v %T\n",s,s)
	fmt.Printf("%v %T\n",r,r)

	const (
		p = 1 >> iota
		q  
	)
	fmt.Printf("%v %T\n",p,p)
	fmt.Printf("%v %T\n",q,q)
}
