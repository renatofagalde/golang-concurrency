package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("example 01")

	time.Sleep(3 * time.Second)

	printSomething("example 02")
}
