package main

import (
	"fmt"
)

var (
	i_am_block_var_1 int = 2
	i_am_block_var_2 string = "block"
	i_am_block_var_3 = 99.1
)

var global_var float32 = 3.14

func main() {
	fmt.Println(i_am_block_var_1)
	fmt.Println(i_am_block_var_2)
	fmt.Println(global_var)	
}
