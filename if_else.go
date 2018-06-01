package main

import (
	"fmt"
)

func pow(x, y int) int {
	option := 2

	if option < 1 {
		return x + y
	} else if option == 1 {
		return y - x
	} else if option == 2 {
		return x * y
	} else {
		return 0
	}
}

func main() {
	fmt.Println(pow(3, 2))
}
