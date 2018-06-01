package main
import "fmt"

type T struct {
    name string
}
func main() {
    t1 := T{"foo"}
    t2 := T{"bar"}
    p1 := &t1
    p2 := &t1
    p3 := &t2
    fmt.Println(p1 == p2)   // true
    fmt.Println(p2 == p3)   // false
    fmt.Println(p3 == nil)  // false
}
