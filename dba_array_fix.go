package main

import "fmt"

func main() {
	// Create an empty slice of slices.
	dba := [2][5]string{{"mysql", "oracle", "sqlserver", "mongo", "postgres"},
		{"ruhel", "steven", "amrutha", "vinod", "ajay"},
	}

	// Loop over slices in dba.
	for i := 1; i >= 0; i-- {
		for b := 4; b >= 0; b-- {
			fmt.Println(dba[i][b])
		}
	}

}
