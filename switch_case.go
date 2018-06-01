package main

import "fmt"

func abc(x, y int) int {

	option := 0
	switch option {
	case 0:
		return x + y
	case 1:
		return y - x
	case 2:
		return x * x
	default:
		return 0
	}
}

func main() {

	fmt.Println(abc(3, 2))
}
