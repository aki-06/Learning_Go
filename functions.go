package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func minus(x int, y int) int {
	return x - y
}

func main() {
	fmt.Println(add(100, 200))
	fmt.Println(minus(1000, 3000))
}