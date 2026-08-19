package main

import "fmt"

func main() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("%d\t%#x\t%#U\n", x, x, x)
		fmt.Println()
	}
}

/* 
Exercise showing just the numerical code and converting to string to show the character
package main

import "fmt"

func main() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("%d - %v", x, string(x))
		fmt.Println()
	}
}
*/
