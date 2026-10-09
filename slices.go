package main

import "fmt"

func slices() {
	sliceSyntax()
}
func sliceSyntax() {
	mySlice := []int{}
	fmt.Printf("myslice: %v | len: %d | Cap: %d\n",
		mySlice, len(mySlice), cap(mySlice))
}
