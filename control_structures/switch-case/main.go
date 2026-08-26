package main

import "fmt"

func main() {
	sega := 1994
	switch {
	case sega < 1983:
		fmt.Println("セガ本体がありません")
	case sega == 1983:
		fmt.Println("1983年です。SG-1000があります！")
	case sega == 1984:
		fmt.Println("1984年です。SG-1000IIがあります！")
	case sega == 1985:
		fmt.Println("1985年です。セガマークIIIがあります！")
	case sega == 1987:
		fmt.Println("1987年です。セガマスターシステムがあります！")
	case sega == 1988:
		fmt.Println("1988年です。セガメガドライブがあります！")
	case sega == 1990:
		fmt.Println("1990年です。セガゲームギアがあります！")
	case sega == 1994:
		fmt.Println("1994年です。セガサターンがあります！")
	case sega == 1998:
			fmt.Println("1998年です。ドリームキャストがあります！")
	}
}

/* The exercise above is an example of how to use the switch statement in Go. This exercise checks 
the value of the variable 'sega' and prints a corresponding message for the year in which a specific Sega 
console was released. If the value of 'sega' is less than 1983, it prints that there is no Sega console. 
For specific values from 1983 to 1998, it prints the name of the console released in that year. */
