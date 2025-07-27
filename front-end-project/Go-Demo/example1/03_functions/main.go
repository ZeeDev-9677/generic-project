package main

import "fmt"

func welcome() string {
	return "Hi There, please respond!!!"
}

// passing arguments a, b into sum function
func sum(a int, b int) int {

	return a + b
}

func sub(c int, d int) int {
	return c - d
}

func main() {
	fmt.Println(welcome()) //calling "welcome" method
	fmt.Println(sum(2, 3)) //calling "sum" method
	fmt.Println(sub(5, 4))
}
