package main

import (
	"fmt"
	"os"
)

func BubbleSort(ar []int) {

	for i:=0; i < len(ar) -1; i++ {
		for j:=0; j < len(ar) - i -1; j++ {
			if ar[j+1] < ar[j] {
				Swap(ar, j)
			}
		}
	}
}

func Swap(slice []int, i int )  {

	if len(slice) == i+1 {
		return
	}

	slice[i] += slice[i+1]
	slice[i+1] = slice[i] - slice[i+1]
	slice[i] -= slice[i+1]
}

func main ()  {

	var n int
	fmt.Print("Enter the number of ints you want to input: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		os.Exit(0)
	}

	fmt.Print(n)
	ar := make([]int, 0, 0)
	for i:=0;i<n;i++ {
		var num int
		fmt.Print("Enter the number : ")
		fmt.Scan(&num)
		ar = append(ar, num)
	}

	BubbleSort(ar)
	fmt.Print(ar)
}