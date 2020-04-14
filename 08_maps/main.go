package main

import "fmt"

func main() {
	//08_maps map a key to a value
	//define map
	emails := make(map[string]string) //make a map[key type]value type
	emails["lukas"] = "lukas.becker.lb@gmail.com"
	emails["franz"] = "franz.bauer@hirsche.at"
	emails["harold"] = "harold@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))

	//Delete from map

	delete(emails, "lukas")
	fmt.Println(emails)

	age := map[string]int{"lukas": 21, "franz": 88, "harold": 30}
	fmt.Println(age["franz"])
}
