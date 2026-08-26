package main

import "fmt"

func main() {
	favoriteSport := "Ryoma"
	switch favoriteSport {
	case "Ryoma":
		fmt.Println("My favorite sport is tennis!")
	case "Sakuragi":
		fmt.Println("My favorite sport is basketball!")
	case "Goenji":
		fmt.Println("My favorite sport is soccer!")
	}
}

/*
The exercise above is an example of how to use the switch statement in Go. This exercise 
checks the value of the variable favoriteSport and prints a message corresponding to the favorite sport. 
In this case, since favoriteSport is "Ryoma", it prints "My favorite sport is tennis!"

*/
