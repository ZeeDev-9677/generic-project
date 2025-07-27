package main

import "fmt"

func main() {

	//-----------------------VARIABLES----------------
	// var name = "Zeshan"
	// var age = 29
	// const floorNo = 44 using const keyword declaration, if declare using "const" keyword will not be changed value
	var name string = "Zeeshan"
	myname := "ZeeAlam" //shorthand declaration
	name, email := "rock", "abcd@gmail.com"
	fmt.Println(myname)
	fmt.Printf("%T\n", myname) // %T= type of variable print
	//%v= value print
	//\n= for new line

	fmt.Println(name, email)
}
