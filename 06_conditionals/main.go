package main

import "fmt"

func main() {
	x := 15
	y := 15

	//if does not require ()
	if x < y {
		fmt.Printf("%d is less than %d \n", x, y)
	} else if x == y {
		fmt.Println("The numbers are equal")
	} else {
		fmt.Printf("%d is more than %d \n", x, y)
	}

	color := "red"

	//Switch statement
	switch color {
	case "red":
		fmt.Println("RED")
	case "blue":
		fmt.Println("Blue")
	default:
		fmt.Println("some other color")

	}
}
