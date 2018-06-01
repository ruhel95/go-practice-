package main

/*
func main() {

	fmt.Println("My favorite number is", rand.Seed(100))
}
*/
/*
func random(min, max int) int {
	rand.Seed(time.Now().unix())
	return rand.Intn(max-min) + min
}
*/

type genpass []string

func newPass() genpass {
	ran := genpass{}
	randomizeL := []string{"a", "b", "c", "d"}
	randomizeU := []string{"A", "B", "C", "D"}
	randomizeN := []string{"1", "2", "3", "4", "5"}

	for _, rl := range randomizeL {
		for _, rU := range randomizeU {
			for _, rN := range randomizeN {
				ran := append(ran, rl+rU+rN)
			}
		}
	}
	return ran
}

func main() {
	ran := newPass()
}
