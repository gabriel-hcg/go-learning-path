// We can simulate a "while loop" in Go using this structure below, changing the init, condition and post structure:

package main

import "fmt"

func main() {

	x := 0

	for x < 10 {
		fmt.Println("x is less than 10")
		x++
	}
}

/* For an infinite loop, we can use a similar structure, just removing the post
or even removing the init and the condition along with it and keeping the for keyword:

import "fmt"

func main() {

	x := 0

	for x < 10 {
		fmt.Println("x is less than 10")
	}
}

We can use the break statement when we want to exit the current loop:

package main

import "fmt"

func main() {

	x := 0

	for {
		if x < 10 {
		fmt.Println("x is less than 10")
		x++
		} else {
		fmt.Println("x is 10 or more")
		break
		}
	}
	fmt.Println("We broke the loop!")
}
*/
