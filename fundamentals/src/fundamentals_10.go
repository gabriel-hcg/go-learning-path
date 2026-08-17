package main

import "fmt"

var x int = 10

func main() {
	x := `text
				with
	raw		string
					literal`
	fmt.Println(x)
}
