package main

import (
	"fmt"
)

func main() {
	// var myMapp map[int]string
	// if myMapp == nil {
	// 	fmt.Println("yes")
	// } else {
	// 	fmt.Println("No")
	// }

	// dataMpp := map[int]string{
	// 	1: "CAT",
	// 	2: "BAT",
	// 	3: "MAT",
	// }
	// fmt.Println("Answer is :", dataMpp)

	// ------------creating slice using make function-------------

	mydt := make([]int, 5, 6)
	fmt.Printf("length is = %d\n", len(mydt))
	fmt.Printf("capcity is = %d\n", cap(mydt))
}
