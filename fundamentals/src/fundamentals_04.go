package main

import "fmt"

type newjeans int

var x newjeans

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Printf("%v", x)
}
