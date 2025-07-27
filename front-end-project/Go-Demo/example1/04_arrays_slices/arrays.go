package main

import "fmt"

func main() {
	//Arrays general way declaration
	// var myarr [3]string
	// myarr[0] = "appl"
	// myarr[1] = "grps"
	// myarr[2] = "napsti"
	// fmt.Println(myarr[2])

	//Arrays shorthand declaration
	// arr1 := [3]int{2, 5, 6}
	// fmt.Println(arr1[2])
	// fmt.Println(&arr1[2]) ---->>array location value will print

	// slices in GoLang just similar to array but in array size need to declare first and in slice no need to declar size like ArrayList in java
	// vegSlice := []string{"apple", "mango", "grps", "sitaphal"}
	// fmt.Println(vegSlice)
	// fmt.Println(len(vegSlice))

	// Range in Go

	vegSliceRange := []string{"apple", "mango", "grps", "sitaphal"}
	// fmt.Println(vegSliceRange[0:1])
	// fmt.Println(vegSliceRange)
	fmt.Println("before", len(vegSliceRange), cap(vegSliceRange))
	vegSliceRange = append(vegSliceRange, "ornge") //append method to add value at the end of slice
	fmt.Println("after", len(vegSliceRange), cap(vegSliceRange))
	vegSliceRange = append(vegSliceRange, "maalto")
	fmt.Printf("\n end the length is %v with capacity %v", len(vegSliceRange), cap(vegSliceRange)) // %v for value print

}
