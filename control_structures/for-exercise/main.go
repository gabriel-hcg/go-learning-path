package main

import "fmt"

func main() {
	birthyear := 2000
	currentyear := 2026
	for {
		if birthyear > currentyear {
			break
		}
		fmt.Println(birthyear)
		birthyear++
	}
}

/* The exercise above demonstrates how a for loop works with a stopping condition.
The loop continues as long as the variable birthyear is less than or equal to the variable currentyear.
When birthyear becomes greater than currentyear, the loop is interrupted with the break statement.
In each iteration, the value of birthyear is incremented by 1 and printed to the screen. */
