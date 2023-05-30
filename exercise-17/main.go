package main

import "fmt"

func main() {
	fmt.Println("Hello Gophers!")
}

//GOOS=windows go build -o main_windows main.go
//GOOS=darwin go build -o main_macos main.go
//GOOS=linux go build -o main_linux main.go
