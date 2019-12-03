package package1

import (
	"fmt"
)

var global_package_level float32 = 3.14
var Global_level float32 = 3.14

func main() {
	fmt.Println(global_package_level)
	fmt.Println(Global_level)
}
