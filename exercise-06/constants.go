package main

import "fmt"

const Pi float64 = 3.14

func main() {
	const World = "Worlde"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go Rules?", Truth)
}
