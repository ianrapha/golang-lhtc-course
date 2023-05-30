package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	fmt.Println(add(42, 3))
	fmt.Println(swap("IAN", "RAPHAEL"))
	fmt.Println(split(17))
}
