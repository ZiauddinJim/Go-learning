package main

import "fmt"

func main() {
	// Note: Many way to variable calling
	// New:  var name string = "zia"
	// New: name := "zia"
	/* New:
	var name string
	name = "Zia"
	*/

	// New: Group variable declaration
	var (
		name string = "zia"
		age  int    = 25
	)
	// New: Multiple value declaration
	var x, y int
	x = 10
	y = 20

	fmt.Println(x, y)
	fmt.Println(name)
	fmt.Println(age)
}
