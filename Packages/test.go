package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Go programming")
	fmt.Println("The time right now is:", time.Now())
	fmt.Println("Random number:", GetRandom())
	printName()
	fmt.Println(returnName())
	fmt.Println(squareRoot32())
	fmt.Println(squareRoot64())
}
