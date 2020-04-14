package main

import "fmt"

func main() {
	//for loops
	/*	i := 1
		for i<=10 {
			fmt.Println(i)
			i++
		}

		//shorter way
		for x := 1; x <= 10; x++{
			fmt.Println(x)
		}
	*/
	//fun game:
	for x := 0; x <= 100; x++ {
		if x%3 == 0 && x%5 != 0 {
			fmt.Println("Fizz")
		} else if x%3 != 0 && x%5 == 0 {
			fmt.Println("Buzz")
		} else if x%3 == 0 && x%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(x)
		}
	}
}
