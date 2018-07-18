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
