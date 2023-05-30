package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(1000))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(9))
	fmt.Println(math.Pi)
}
