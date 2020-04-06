package main

import (
	"fmt"
	"github.com/lukasbeckercode/HelloGo/03_packages/strutil" //my own package
	"math"
) //more than 1 import NO COMMA

func main() {
	fmt.Println(math.Floor(3.7))
	fmt.Println(math.Ceil(3.4))
	fmt.Println(math.Round(3.7))
	fmt.Println(strutil.Reverse("Lukas"))
}
