package main

import "fmt"

var name string = "Lukas" //it is DuckTyped, so we dont need to tell it this is a string
//name :="Lukas" // This doesnt work here!

func main() {

	//var unusedVar ="useless"//Variables have to be used in Go! Uncomment to see error
	secondName := "Becker" //this is a shorter Version of creating Variables, but this only works inside Functions
	var age = 21
	const isCool = true
	fmt.Println(name, secondName, "Age:", age, "isCool:", isCool)
	fmt.Printf("%T\n", age) //Printf takes so called verbs. Google go ftm to see which verbs exist
}
