package main

import "fmt"

func main() {
	fmt.Println(greeting("Lukas"))
	fmt.Println(getSum(56, 102))
}

func greeting(name string) string { //func name(arg type) return-type
	return "Hello, " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}
