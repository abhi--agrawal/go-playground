package main

import (
	"fmt"
	"sync"
)

func sort(ar []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("Sorting : ")
	fmt.Println(ar)
	l := len(ar)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if ar[j+1] < ar[j] {
				ar[j+1], ar[j] = ar[j], ar[j+1]
			}
		}
	}
	fmt.Print("Sorted : ")
	fmt.Println(ar)
}

func main() {

	ar := []int{2, 7, 4, 10, -3, 6, 1, 55, 12, 4, -8, 72, 0, 90, 38, 77}
	len := len(ar)
	slLen := len / 4
	var wg sync.WaitGroup
	wg.Add(1)
	go sort(ar[0:slLen], &wg)
	wg.Add(1)
	go sort(ar[slLen:2*slLen], &wg)
	wg.Add(1)
	go sort(ar[2*slLen:3*slLen], &wg)
	wg.Add(1)
	go sort(ar[3*slLen:], &wg)
	wg.Wait()

	wg.Add(1)
	sort(ar, &wg)
}
