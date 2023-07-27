package main

import "fmt"

var myVal int

func main() {
	fmt.Println("Hello")
	myVal = calc(10, 20)

	// this comes from amazingGoRoutines, to able to use this you need to run this go file with  others either use go run main.go amazing.gorr or use go run *.go
	waitMe()

}

func calc(a, b int) int {

	return a * b

}

// you need to build all if you gonna build
