package main

import (
	"fmt"
)

func main() {
	//if-else conditional statements
	// a := 10
	// b := 50
	// if a >= b {
	// 	fmt.Println("ready")
	// } else {
	// 	fmt.Println("not ready")
	// }
	//else-if conditional statements
	time := 20
	if time < 10 {
		fmt.Println("Morning")
	} else if time <= 20 {
		fmt.Println("Noon")
	} else {
		fmt.Println("Evening")
	}

}
