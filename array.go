package main

import (
	"fmt"
)

func Array() {
	// arrDefine()
	// arrInferred()
	// arrValueAccess()
	// arrValueChange()
	// arrInitialization()
	// arrKeyIndexedInitialization()
	arrLength()
}

// Note: Array declaration two way
// Define length
func arrDefine() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1, arr2)
}

// Inferred Length
func arrInferred() {
	var arr1 = [...]float32{1.2, 4.3, 5.4}
	arr2 := [...]int{3, 4, 5}
	fmt.Println(arr1, arr2)
}

// Note: Array element access
func arrValueAccess() {
	prices := [...]int{17, 15, 16}
	fmt.Println(prices[1])
	fmt.Println(prices[0])
	fmt.Println(prices[2])

}

// Note: Array value change
func arrValueChange() {
	prices := [...]int{17, 15, 16}
	fmt.Println(prices)
	// value change
	prices[2] = 20
	fmt.Println(prices)
}

// Note: Array Initialization- The default value for int is 0
func arrInitialization() {
	arr1 := [5]int{}
	arr2 := [5]int{2, 1}
	arr3 := [5]int{2, 1, 4, 5, 6}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

// Note: Array Key indexed Initialization
func arrKeyIndexedInitialization() {
	arr := [5]int{1: 10, 3: 34}
	fmt.Println(arr)
}

// Note: Find the length of an Array
func arrLength() {
	cars := [4]string{"Volvo", "Toyota", "BMW", "Ford"}
	number := [...]int{3, 4, 2, 5, 3, 4}

	fmt.Println(len(cars))
	fmt.Println(len(number))
}
