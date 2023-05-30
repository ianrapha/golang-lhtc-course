package main

import "fmt"

func main() {
	a, b, c := "Ian", 42, 42.42

	fmt.Printf("%v is of type %T\n", a, a)
	fmt.Printf("%v is of type %T\n", b, b)
	fmt.Printf("%v is of type %T\n", c, c)
}
