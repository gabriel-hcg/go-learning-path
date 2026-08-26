package main

import "fmt"

func main() {
	birthyear := 2000
	currentyear := 2026
	for birthyear <= currentyear {
		fmt.Println(birthyear)
		birthyear++
	}
}

/* The exercise above serves to illustrate the use of a for loop 
with a condition. The variable birthyear is initialized with the 
value 2000 and the variable currentyear is initialized with the 
value 2026. The for loop continues as long as birthyear is less 
than or equal to currentyear, printing the value of birthyear on 
each iteration and incrementing it by 1. When birthyear exceeds 
currentyear, the loop ends. */
