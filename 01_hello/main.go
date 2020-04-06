/*
Intro to go!
Go is very sensitive with the workspace structure!
correct structure:
-bin
-src
	-github.com
		-githubusername
			-project1
				project1.go
			-project2
-pkg

you need to set the GOPATH System Environment Variable to your workspace!
*/
package main //important, set this to main every time!

import "fmt" //fmt allows us to write to the console

func main() {
	fmt.Println("Hello World")
}

//to make an .exe File, use go install, and the exe is in the bin folder
