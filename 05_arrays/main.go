package main

import "fmt"

func main() {
	//Arrays in go always have a fixed size and you need to specify their type
	var fruitArray [2]string //Declare an Array
	fruitArray[0] = "apple"  //add something to the Array
	fruitArray[1] = "banana"

	nameArray := [2]string{"Lukas", "Michael"} //faster way of declaring an array
	fmt.Println(fruitArray)
	fmt.Println(nameArray)

	fruitSlice := []string{"orange", "pineapple"} //a slice doesn not need a fixed size
	fmt.Println(fruitSlice)
}
