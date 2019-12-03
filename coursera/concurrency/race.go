package main

import (
	"fmt"
	"sync"
)

func race(wg *sync.WaitGroup, x *int) {
	defer wg.Done()
	(*x)++
	fmt.Println(*x)
}

func race2(wg *sync.WaitGroup, x *int) {
	defer wg.Done()
	(*x) = (*x) * 20
	fmt.Println(*x)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	x := 10
	go race(&wg, &x)
	go race2(&wg, &x)
	wg.Wait()
}

//This prints
//ITAdmins-MacBook-Pro:concurrency abhishekagrawal$ ./race
//201
//200
//ITAdmins-MacBook-Pro:concurrency abhishekagrawal$ ./race
//220
//11
//As the value of x is dependent of the order of execution and any program can get executed first based on the go runtime.
//So the output is non deteminant.
