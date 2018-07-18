package raindrops

import "strconv"

func Convert(n int) string {

	var store string = ""

	if n%3 == 0 {
		store = store + "Pling"
	}
	if n%5 == 0 {
		store = store + "Plang"
	}
	if n%7 == 0 {
		store = store + "Plong"
	}
	if store == "" {
		return strconv.Itoa(n)
	}

	return store
}

func factor(x int) []int {
	fact := make([]int, 0)
	var save int = 0
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fact[save] = i
			save++
		}
	}
	return fact
}

//prints all the factors of a number
// stores it in a array
