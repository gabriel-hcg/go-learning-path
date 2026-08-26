package main

import "fmt"

func main() {
	sega := 0
	if sega == 1994 {
		fmt.Println("セガサターン城！")
	} else if sega == 1998 {
		fmt.Println("１９９８年です。ドリームキャストがあります！")
	} else {
		fmt.Println("The end of SEGA...")
	}
}


/* the exercise above demonstrates the use of an if-else-if statement in Go. Depending on the value 
of the variable 'sega', it prints different messages related to SEGA's history. If 'sega' is equal to 1994, 
it prints "セガサターン城！". If it is equal to 1998, it prints "１９９８年です。ドリームキャストがあります！". 
Otherwise, it prints "The end of SEGA...". */
