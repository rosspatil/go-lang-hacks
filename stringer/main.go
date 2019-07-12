package main

import "fmt"

func main() {
	a1 := A{10, false}
	fmt.Println(a1)
	a2 := A{20, true}
	fmt.Println(a2)
}

type A struct {
	Count int
	Valid bool
}

func (a A) String() string {
	return fmt.Sprintf("%d (%v)", a.Count, a.Valid)
}

/* output

10 (false)
20 (true)

*/
