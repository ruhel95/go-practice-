package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Wednesday?")
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days.")
	case today + 4:
		fmt.Println("In four days.")
	case today + 5:
		fmt.Println("In five days.")
	case today + 6:
		fmt.Println("In six days.")
	default:
		fmt.Println("Too far away.")
	}
}
