package main

import "fmt"

var x = 40

const y = 41
const keyword = "asana"

func main() {
	z := 42

	fmt.Printf("The value of z is %v and the type of z is %T\n", x, x)
	fmt.Printf("The value of z is %v and the type of z is %T\n", y, y)
	fmt.Printf("The value of z is %v and the type of z is %T\n", z, z)
}
