package main

import "fmt"

func main() {
	if x := 500; x > 100 {
		fmt.Println("x is more than 100")
	} else if x < 10 {
		fmt.Println("x is less than 10")
	} else {
		fmt.Println("x is neither less than 10 nor more than 100")
	}
}
