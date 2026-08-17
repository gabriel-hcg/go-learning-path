package main

import "fmt"

var x int = 10

func main() {
	fmt.Printf("%d\t%b\t%#x\t\n", x, x, x)
	x := 10 << 1
	y := x
	fmt.Printf("%d\t%b\t%#x\t", y, y, y)
}
