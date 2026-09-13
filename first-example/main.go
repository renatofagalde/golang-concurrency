package main

import "fmt"

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("example 01")

	printSomething("example 02")
}
